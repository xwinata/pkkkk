package models

import "github.com/xwinata/basemodel"

// Provinsi indonesa
type Provinsi struct {
	basemodel.BaseModel
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
}

// Create function
func (model *Provinsi) Create() error {
	return basemodel.Create(&model)
}

// Save function
func (model *Provinsi) Save() error {
	return basemodel.Save(&model)
}

// Delete function
func (model *Provinsi) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID function
func (model *Provinsi) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// FindFilter function
func (model *Provinsi) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) ([]Provinsi, error) {
	slices := []Provinsi{}
	err := basemodel.FindFilter(&slices, order, sort, limit, offset, filter)
	return slices, err
}

// SingleFindFilter function
func (model *Provinsi) SingleFindFilter(filter interface{}) error {
	err := basemodel.SingleFindFilter(&model, filter)
	return err
}
