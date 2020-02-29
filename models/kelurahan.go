package models

import "github.com/jinzhu/gorm"

// Kelurahan main type
type Kelurahan struct {
	gorm.Model
	KecamatanID uint64    `json:"kecamatan_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"column:name;type:varchar(255)"`
	PostalCode  string    `json:"postal_code" gorm:"column:postal_code;type:varchar(255)"`
	AreaCode    string    `json:"area_code" gorm:"column:area_code;type:varchar(255)"`
	Kecamatan   Kecamatan `gorm:"association_foreignkey:ID"`
}
