package system

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct{}

func DatabaseHandler() *Database {
	return &Database{}
}

func (h *Database) Connection() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func (h *Database) GetOne(where interface{}) (interface{}, error) {

	db := h.Connection()
	result := map[string]interface{}{}
	db.Find(where).Take(&result)
	return result, nil

}

func (h *Database) CreateOne(model interface{}) (interface{}, error) {

	db := h.Connection()
	result := db.Create(&model)
	return model, result.Error

}
