package models

import "github.com/jinzhu/gorm"

// Kota main type
type Kota struct {
	gorm.Model
	ProvinsiID uint64   `json:"provinsi_id" gorm:"not null"`
	Name       string   `json:"name" gorm:"column:name;type:varchar(255)"`
	Type       string   `json:"type" gorm:"column:type;type:varchar(255)"`
	Provinsi   Provinsi `gorm:"association_foreignkey:ID"`
}
