package models

import "github.com/jinzhu/gorm"

// Transaction gorm helper for create, update, and delete
func Transaction(db *gorm.DB, fn string, i interface{}) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	switch fn {
	case "create":
		err = db.Create(i).Error
	}
	if err != nil {
		return err
	}

	return tx.Commit().Error
}
