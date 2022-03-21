package database

import (
	"gorm.io/gorm"
)

func ExecuteMigrations(db *gorm.DB) error {
	return db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate()

}
