package models

import (
	"github.com/jinzhu/gorm"
)

// Provinsi indonesa
type Provinsi struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
}
