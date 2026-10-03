package models

import "gorm.io/gorm"

type Registry struct {
	gorm.Model

	Name string `json:"name"`
	Type string `json:"type"`
}
