package models

import "gorm.io/gorm"

type RegistryImage struct {
	gorm.Model
	Image     string `json:"image"`
	Tag       string `json:"tag"`
	Timestamp string `json:"timestamp"`
}
