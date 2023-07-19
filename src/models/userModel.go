package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255);unique"`
	Username  string    `json:"username" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
}

type UserModelService struct {
	db *gorm.DB
}

func UserModelHandler(db *gorm.DB) *UserModelService {
	return &UserModelService{
		db: db,
	}
}

func (h *UserModelService) GetOne(model *UserModel) (*UserModel, error) {

	if err := h.db.Where(&model).First(&model); err != nil {
		return model, err.Error
	}
	return model, nil

}

func (h *UserModelService) CreateOne(model *UserModel) (*UserModel, error) {

	result := h.db.Create(&model)
	h.db.Commit()
	return model, result.Error

}
