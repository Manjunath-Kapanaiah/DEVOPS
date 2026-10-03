# Container Registry Tracker

A lightweight **DevOps capstone project** — a Go REST API that tracks container **registries**, the **images** inside them, and **which environment each image is deployed to** (a `prod` / `dev` / `qa` label stored on every deployment record). It is backed by PostgreSQL, containerized with a multi-stage Docker build, orchestrated with Docker Compose, and shipped through a Jenkins CI/CD pipeline that publishes the image to **both Docker Hub and AWS ECR**.

> **The pitch:** Organisations running containers at scale lose track of which image version is deployed where. This app solves that with a simple REST API backed by PostgreSQL — and demonstrates the full build → push → deploy lifecycle.

---

## Table of Contents

- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Highlights](#highlights)
- [Project Structure](#project-structure)
- [Data Model](#data-model)
- [API Reference](#api-reference)
- [Getting Started](#getting-started)
- [CI/CD Pipeline](#cicd-pipeline)
- [AWS ECR Integration](#aws-ecr-integration)
- [Docker Image Optimization](#docker-image-optimization)
- [Screenshots](#screenshots)
- [Author](#author)

---

## Architecture

```
                                       ┌──►  Docker Hub  ─┐
GitHub  →  Jenkins  →  Docker Build  ──┤                  ├──►  Docker Compose  →  Running App + PostgreSQL
                                       └──►  AWS ECR     ─┘
```

The same image is pushed to **two registries** in the same pipeline run: Docker Hub (`likith0129/registry-tracker`) and AWS ECR (`container-registry-tracker`).

```
            ┌────────────────────────────────────────────┐
            │              Docker Compose                 │
            │                                             │
   :8081 ───┼──►  registry-tracker-app  ───►  postgres    │
            │     (Go + Gin + GORM)         (postgres:16) │
            │            host=postgres        :5432       │
            └────────────────────────────────────────────┘
```

Inside Docker Compose, the app reaches the database using the **service name** `postgres` as the hostname (not `localhost`) — each container is its own network namespace.

---

## Tech Stack

| Layer            | Technology                                  |
| ---------------- | ------------------------------------------- |
| Language         | Go 1.25                                     |
| Web framework    | [Gin](https://github.com/gin-gonic/gin)     |
| ORM              | [GORM](https://gorm.io)                      |
| Database         | PostgreSQL 16                               |
| Containerization | Docker (multi-stage build)                  |
| Orchestration    | Docker Compose                              |
| Registry         | Docker Hub (`likith0129/registry-tracker`)  |
| Cloud Registry   | AWS ECR (`container-registry-tracker`, `ap-south-1`) |
| CI/CD            | Jenkins (Declarative Pipeline)              |
| SCM              | Git / GitHub                                |

---

## Highlights

- **~96% smaller images** — multi-stage Docker build shrinks the image from **1.33 GB → 46 MB** by shipping only the compiled static binary on Alpine.
- **Resilient startup** — the app retries the database connection up to 10 times (5s apart), so it survives PostgreSQL not being ready yet inside Compose.
- **Auto-migration** — GORM creates/updates tables from the Go structs on boot; no manual SQL.
- **Fully automated delivery** — a Jenkins pipeline builds, logs in, and pushes the image to Docker Hub on every change.
- **Dual-registry publishing** — the same pipeline also authenticates to **AWS ECR** (via a dedicated `jenkins-ecr-user` IAM user) and pushes the image there, demonstrating cloud-native registry integration.
- **Tested CI gate** — `go vet` + `go test` run before any image is built, so broken code never ships.
- **Immutable, versioned tags** — every build is tagged `<build#>-<git-sha>` (alongside `latest`), giving full traceability and a real rollback target.
- **Verified deploys** — a `/health` endpoint (which checks the database) is polled after each deploy; the pipeline only passes if the app is genuinely healthy.
- **No secrets in git** — credentials come from a gitignored `.env`; the compose file requires `POSTGRES_PASSWORD` rather than hardcoding it.
- **Twelve-factor config** — configuration is injected via environment variables.

---

## Project Structure

```
container-registry-tracker/
├── cmd/
│   └── main.go            # Entry point: connect DB, auto-migrate, register routes, start server (:8081)
├── config/
│   └── database.go        # PostgreSQL connection via GORM, with 10x retry loop
├── models/
│   ├── registry.go        # Registry struct  → registries table
│   ├── image.go           # Image struct     → images table (FK: RegistryID)
│   └── deployment.go      # Deployment struct → deployments table (FK: ImageID)
├── handlers/
│   ├── registry.go        # Create / list registries
│   ├── image.go           # Create / list / get-by-id images
│   └── deployment.go      # Create / list / filter-by-environment deployments
├── routes/
│   └── routes.go          # Maps HTTP routes → handler functions
├── Dockerfile             # Multi-stage build (golang:1.25 → alpine)
├── docker-compose.yml     # App + PostgreSQL services
├── Jenkinsfile            # CI/CD pipeline
├── go.mod / go.sum        # Module definition & dependency checksums
└── README.md
```

---

## Data Model

All models embed `gorm.Model`, so every record also carries `ID`, `CreatedAt`, `UpdatedAt`, and `DeletedAt`.

```
Registry (1) ──< Image (1) ──< Deployment
```

| Model          | Fields                                            | Notes                                  |
| -------------- | ------------------------------------------------- | -------------------------------------- |
| **Registry**   | `name`, `type`                                    | e.g. `Docker Hub`, type `public`       |
| **Image**      | `name`, `version`, `registry_id`                  | belongs to a Registry                  |
| **Deployment** | `image_id`, `environment`                         | environment = `prod` / `dev` / `qa`    |

The three tables are auto-created by GORM on first boot (verified with `\dt` inside the PostgreSQL container):

![PostgreSQL auto-migrated tables](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/postgres-tables.png)

---

## API Reference

Base URL: `http://localhost:8081`

### System

| Method | Endpoint   | Description                                                       |
| ------ | ---------- | ---------------------------------------------------------------- |
| GET    | `/health`  | Liveness + DB check. Returns `200 {"status":"ok","database":"up"}`, or `503` if the database is unreachable. Used by the pipeline to verify a deployment. |

### Registries

| Method | Endpoint       | Description            |
| ------ | -------------- | ---------------------- |
| POST   | `/registries`  | Create a registry      |
| GET    | `/registries`  | List all registries    |

### Images

| Method | Endpoint        | Description            |
| ------ | --------------- | ---------------------- |
| POST   | `/images`       | Create an image        |
| GET    | `/images`       | List all images        |
| GET    | `/images/:id`   | Get one image by ID    |

### Deployments

| Method | Endpoint                                  | Description                          |
| ------ | ----------------------------------------- | ------------------------------------ |
| POST   | `/deployments`                            | Create a deployment                  |
| GET    | `/deployments`                            | List all deployments                 |
| GET    | `/deployments/environment/:environment`   | List deployments in an environment   |

### Example requests

```bash
# 1. Create a registry
curl -X POST http://localhost:8081/registries \
  -H "Content-Type: application/json" \
  -d '{"name": "Docker Hub", "type": "public"}'

# 2. Create an image in registry 1
curl -X POST http://localhost:8081/images \
  -H "Content-Type: application/json" \
  -d '{"name": "registry-tracker", "version": "v1.0", "registry_id": 1}'

# 3. Deploy image 1 to production
curl -X POST http://localhost:8081/deployments \
  -H "Content-Type: application/json" \
  -d '{"image_id": 1, "environment": "prod"}'

# 4. List everything deployed to prod
curl http://localhost:8081/deployments/environment/prod
```

---

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) & Docker Compose
- (Optional, for local non-Docker runs) Go 1.25+ and a local PostgreSQL

### Run with Docker Compose (recommended)

```bash
git clone https://github.com/LikithKumar0112/Docker-Capstone-Project.git
cd Docker-Capstone-Project

docker compose up -d
```

This starts:
- **postgres** — PostgreSQL 16 on port `5432` (data persisted in the `postgres_data` volume)
- **registry-tracker-app** — the API on [http://localhost:8081](http://localhost:8081)

Both containers running, with the app connected to the database and all routes registered:

![docker ps and application logs](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/docker-ps-and-app-logs.png)

Stop and clean up:

```bash
docker compose down
```

### Run locally without Docker

```bash
# Start a local PostgreSQL with database 'registrydb' (user/pass: registryuser/registrypass)
go mod download
go run ./cmd
```

If `DATABASE_URL` is not set, the app falls back to a local DSN:
`host=localhost user=registryuser password=registrypass dbname=registrydb port=5432 sslmode=disable`

### Configuration

Configuration is supplied via environment variables. Docker Compose reads them from a `.env` file (copy [`.env.example`](./.env.example) to `.env`); the `.env` file is gitignored so secrets are never committed.

| Variable             | Description                                   | Required? / Default                        |
| -------------------- | --------------------------------------------- | ------------------------------------------ |
| `POSTGRES_PASSWORD`  | Database password                             | **Required** — Compose won't start without it |
| `POSTGRES_USER`      | Database user                                 | `registryuser`                             |
| `POSTGRES_DB`        | Database name                                 | `registrydb`                               |
| `IMAGE_TAG`          | Image tag to deploy (set by the pipeline)     | `latest`                                   |
| `DATABASE_URL`       | PostgreSQL connection string (DSN)            | `host=localhost ... dbname=registrydb ...` |

---

## CI/CD Pipeline

The [`Jenkinsfile`](./Jenkinsfile) defines a declarative pipeline:

| Stage                  | What it does                                                      |
| ---------------------- | ---------------------------------------------------------------- |
| **Checkout**           | Pulls the latest code from SCM and computes an immutable version tag `<build#>-<git-sha>` |
| **Verify Workspace**   | Prints the working directory and file listing to confirm the checkout |
| **Test**               | **CI gate** — runs `go vet ./...` and `go test ./...` in a `golang:1.25` container; the pipeline stops here if anything fails, so broken code is never built or shipped |
| **Build Docker Image** | Builds and tags the image as `<build#>-<git-sha>` (immutable) **and** `latest` |
| **Docker Login**       | Authenticates to Docker Hub using `dockerhub-creds` (stored in Jenkins Credentials, piped via `--password-stdin`) |
| **Push Docker Image**  | Pushes both the versioned tag and `latest` to Docker Hub |
| **AWS ECR Login**      | Authenticates to AWS ECR using the `aws-access-key-id` / `aws-secret-access-key` Jenkins credentials, then `aws ecr get-login-password \| docker login` against `790505843920.dkr.ecr.ap-south-1.amazonaws.com` |
| **Tag Image For ECR**  | Re-tags the image (versioned + `latest`) for the `container-registry-tracker` ECR repository |
| **Push To ECR**        | Pushes both tags to AWS ECR |
| **Deploy**             | Pins the freshly built version (`IMAGE_TAG`), then redeploys with `docker compose pull` + `docker compose up -d` |
| **Verify Deployment**  | Polls `GET /health` with retries; the build only succeeds if the app comes up healthy and connected to the database |

Jenkins **Stage View** showing the pipeline running through every stage:

![Jenkins pipeline stage view](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/jenkins-stage-view.png)

The resulting image published to Docker Hub:

![Docker Hub repository](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/dockerhub-repository.png)

**Server prerequisite:** Jenkins must be able to talk to the Docker daemon:

```bash
sudo usermod -aG docker jenkins
sudo systemctl restart jenkins docker
```

> When deploying via Compose in the pipeline, run `docker compose down || true` before `docker compose up -d` to avoid the *"container name already in use"* conflict from a previous deployment.

---

## AWS ECR Integration

Beyond Docker Hub, the pipeline also publishes the image to **Amazon Elastic Container Registry (ECR)** — a private, cloud-native registry. This mirrors a real-world setup where images live in the same cloud as the workloads that run them.

**How it works:**

1. A dedicated IAM user, **`jenkins-ecr-user`**, is granted the `AmazonEC2ContainerRegistryFullAccess` managed policy. Its access key / secret are stored in Jenkins as the `aws-access-key-id` and `aws-secret-access-key` credentials.
2. The **AWS ECR Login** stage runs `aws ecr get-login-password --region ap-south-1 | docker login …`, authenticating Docker to the registry `790505843920.dkr.ecr.ap-south-1.amazonaws.com`.
3. The **Tag Image For ECR** stage re-tags the freshly built image for the `container-registry-tracker` ECR repository.
4. The **Push To ECR** stage pushes it — the layers upload (or are skipped if already present) and the image becomes available in ECR.

The dedicated IAM user **`jenkins-ecr-user`** with the ECR full-access policy attached:

![IAM jenkins-ecr-user with ECR access](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/iam-jenkins-ecr-user.png)

Jenkins authenticating to ECR — **`Login Succeeded`**:

![Jenkins AWS ECR Login stage log](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/jenkins-ecr-login-stage.png)

Jenkins pushing the image to ECR:

![Jenkins Push To ECR stage log](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/jenkins-ecr-push-stage.png)

The **`container-registry-tracker`** repository in the ECR console, showing the repository URI and configuration:

![ECR repository details](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/ecr-repository-details.png)

The pushed images listed in the repository:

![ECR images list](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/ecr-images-list.png)

Image details — tag, digest, size, and scan status:

![ECR image details](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/ecr-image-details.png)

---

## Docker Image Optimization

The [`Dockerfile`](./Dockerfile) uses a **multi-stage build**:

| Approach                       | Image Size | What it contains                        |
| ------------------------------ | ---------- | --------------------------------------- |
| Single stage (`golang:1.25`)   | **1.33 GB** | Go SDK + compiler + source + binary     |
| Multi-stage (Alpine runtime)   | **46 MB**   | Alpine OS + compiled static binary only |

The builder stage compiles a static binary (`CGO_ENABLED=0`), and the runtime stage copies only that binary onto a minimal `alpine:latest` base — roughly a **96% size reduction**.

The single-stage build (`registry-tracker:v1`) weighing in at **1.33 GB**:

![docker images showing v1 at 1.33 GB](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/docker-images-v1-1.33gb.png)

After the multi-stage rebuild — `v1` (1.33 GB) next to `v2` (46 MB):

![docker images size comparison v1 vs v2](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/docker-image-size-comparison.png)

---

## Screenshots

All screenshots live in [`docs/screenshots/`](https://github.com/LikithKumar0112/Docker-Capstone-Project/tree/Images/docs/screenshots) and are referenced throughout this README. The full lifecycle at a glance:

| What | Screenshot |
| ---- | ---------- |
| App + database running, routes registered | [`docker-ps-and-app-logs.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/docker-ps-and-app-logs.png) |
| Image size: single-stage build at 1.33 GB | [`docker-images-v1-1.33gb.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/docker-images-v1-1.33gb.png) |
| Image size: v1 (1.33 GB) vs v2 (46 MB) | [`docker-image-size-comparison.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/docker-image-size-comparison.png) |
| Image published to Docker Hub | [`dockerhub-repository.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/dockerhub-repository.png) |
| Jenkins pipeline stage view | [`jenkins-stage-view.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/jenkins-stage-view.png) |
| PostgreSQL auto-migrated tables | [`postgres-tables.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/postgres-tables.png) |
| IAM `jenkins-ecr-user` with ECR access | [`iam-jenkins-ecr-user.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/iam-jenkins-ecr-user.png) |
| Jenkins AWS ECR Login stage log | [`jenkins-ecr-login-stage.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/jenkins-ecr-login-stage.png) |
| Jenkins Push To ECR stage log | [`jenkins-ecr-push-stage.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/jenkins-ecr-push-stage.png) |
| ECR repository details | [`ecr-repository-details.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/ecr-repository-details.png) |
| ECR images list | [`ecr-images-list.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/ecr-images-list.png) |
| ECR image details | [`ecr-image-details.png`](https://raw.githubusercontent.com/LikithKumar0112/Docker-Capstone-Project/Images/docs/screenshots/ecr-image-details.png) |

---

DevOps Capstone Project — *Container Registry Tracker*

Docker Hub image: [`likith0129/registry-tracker`](https://hub.docker.com/r/likith0129/registry-tracker)
