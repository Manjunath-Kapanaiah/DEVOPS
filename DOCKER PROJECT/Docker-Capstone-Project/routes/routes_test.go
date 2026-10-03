package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LikithKumar0112/container-registry-tracker/config"
	"github.com/LikithKumar0112/container-registry-tracker/models"
	"github.com/LikithKumar0112/container-registry-tracker/routes"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// setupRouter wires the real routes against a fresh in-memory SQLite database,
// so each test runs end-to-end (routing + handlers + persistence) in isolation.
func setupRouter(t *testing.T) *gin.Engine {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}

	if err := db.AutoMigrate(&models.Registry{}, &models.Image{}, &models.Deployment{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	config.DB = db

	gin.SetMode(gin.TestMode)
	router := gin.New()
	routes.SetupRoutes(router)
	return router
}

func TestHealthEndpoint(t *testing.T) {
	router := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d (body: %s)", w.Code, w.Body.String())
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if body["status"] != "ok" || body["database"] != "up" {
		t.Fatalf("unexpected health body: %v", body)
	}
}

func TestCreateAndListRegistries(t *testing.T) {
	router := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/registries",
		bytes.NewBufferString(`{"name":"Docker Hub","type":"public"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 on create, got %d (body: %s)", w.Code, w.Body.String())
	}

	// The list endpoint should now return exactly the one we created.
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest(http.MethodGet, "/registries", nil)
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("expected 200 on list, got %d", w2.Code)
	}

	var registries []models.Registry
	if err := json.Unmarshal(w2.Body.Bytes(), &registries); err != nil {
		t.Fatalf("invalid JSON list: %v", err)
	}
	if len(registries) != 1 || registries[0].Name != "Docker Hub" {
		t.Fatalf("unexpected registries: %+v", registries)
	}
}

func TestGetImageByID_NotFound(t *testing.T) {
	router := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/images/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for missing image, got %d (body: %s)", w.Code, w.Body.String())
	}
}

func TestCreateRegistry_InvalidJSON(t *testing.T) {
	router := setupRouter(t)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/registries",
		bytes.NewBufferString(`{"name":`)) // malformed
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for malformed JSON, got %d", w.Code)
	}
}
