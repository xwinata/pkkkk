package models

import (
	"github.com/xwinata/basemodel"
)

// Kelurahan main type
type Kelurahan struct {
	basemodel.BaseModel
	KecamatanID uint64 `json:"kecamatan_id" gorm:"not null"`
	Name        string `json:"name" gorm:"column:name;type:varchar(255)"`
	PostalCode  string `json:"postal_code" gorm:"column:postal_code;type:varchar(255)"`
	AreaCode    string `json:"area_code" gorm:"column:area_code;type:varchar(255)"`
}

// Create function
func (model *Kelurahan) Create() error {
	return basemodel.Create(&model)
}

// Save function
func (model *Kelurahan) Save() error {
	return basemodel.Save(&model)
}

// Delete function
func (model *Kelurahan) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID function
func (model *Kelurahan) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// FindFilter function
func (model *Kelurahan) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) ([]Kelurahan, error) {
	slices := []Kelurahan{}
	err := basemodel.FindFilter(&slices, order, sort, limit, offset, filter)
	return slices, err
}

// SingleFindFilter function
func (model *Kelurahan) SingleFindFilter(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}
