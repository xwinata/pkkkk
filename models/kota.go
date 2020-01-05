package models

import (
	"github.com/xwinata/basemodel"
)

// Kota main type
type Kota struct {
	basemodel.BaseModel
	ProvinsiID uint64 `json:"provinsi_id" gorm:"not null"`
	Name       string `json:"name" gorm:"column:name;type:varchar(255)"`
	Type       string `json:"type" gorm:"column:type;type:varchar(255)"`
}

// Create function
func (model *Kota) Create() error {
	return basemodel.Create(&model)
}

// Save function
func (model *Kota) Save() error {
	return basemodel.Save(&model)
}

// Delete function
func (model *Kota) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID function
func (model *Kota) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// FindFilter function
func (model *Kota) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) ([]Kota, error) {
	slices := []Kota{}
	err := basemodel.FindFilter(&slices, order, sort, limit, offset, filter)
	return slices, err
}

// SingleFindFilter function
func (model *Kota) SingleFindFilter(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}
