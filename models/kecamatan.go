package models

import (
	"github.com/jinzhu/gorm"
)

// Kecamatan main type
type Kecamatan struct {
	gorm.Model
	KotaID uint64 `json:"kota_id" gorm:"not null"`
	Name   string `json:"name" gorm:"column:name;type:varchar(255)"`
	Kota   Kota   `gorm:"association_foreignkey:ID"`
}
