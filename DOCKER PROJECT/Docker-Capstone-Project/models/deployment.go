package models

import "gorm.io/gorm"

type Deployment struct {
	gorm.Model

	ImageID     uint   `json:"image_id"`
	Environment string `json:"environment"`
}

