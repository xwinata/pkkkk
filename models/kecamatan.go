package models

import (
	"github.com/xwinata/basemodel"
)

// Kecamatan main type
type Kecamatan struct {
	basemodel.BaseModel
	KotaID uint64 `json:"kota_id" gorm:"not null"`
	Name   string `json:"name" gorm:"column:name;type:varchar(255)"`
}

// Create function
func (model *Kecamatan) Create() error {
	return basemodel.Create(&model)
}

// Save function
func (model *Kecamatan) Save() error {
	return basemodel.Save(&model)
}

// Delete function
func (model *Kecamatan) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID function
func (model *Kecamatan) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// FindFilter function
func (model *Kecamatan) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) ([]Kecamatan, error) {
	slices := []Kecamatan{}
	err := basemodel.FindFilter(&slices, order, sort, limit, offset, filter)
	return slices, err
}

// SingleFindFilter function
func (model *Kecamatan) SingleFindFilter(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}
