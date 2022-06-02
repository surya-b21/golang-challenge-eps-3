package services

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB ....
var DB *gorm.DB

// InitDB function
func InitDB() {
	if DB == nil {
		dir, _ := os.Getwd()
		connection := sqlite.Open(dir + "/app/services/eps3.db")
		config := &gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		}
		if db, err := gorm.Open(connection, config); err == nil {
			DB = db.Debug()
		}
	}
}
