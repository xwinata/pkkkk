package migration

import (
	"pkkkk/models"

	"github.com/jinzhu/gorm"
)

// AutoMigrate migrates tables using gorm automigrate
func AutoMigrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&models.Provinsi{}, &models.Kota{}, &models.Kelurahan{}, &models.Kecamatan{}).Error
	if err != nil {
		return err
	}

	db.Model(&models.Kota{}).AddForeignKey("provinsi_id", "provinsis(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Kecamatan{}).AddForeignKey("kota_id", "kota(id)", "CASCADE", "RESTRICT")
	db.Model(&models.Kelurahan{}).AddForeignKey("kecamatan_id", "kecamatans(id)", "CASCADE", "RESTRICT")

	return err
}
