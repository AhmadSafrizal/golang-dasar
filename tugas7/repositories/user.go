package repository

import (
	"log"

	model "github.com/AhmadSafrizal/golang-dasar/tugas7/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) Migrate() {
	err := u.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
}

func (u *UserRepository) Create(user *model.User) error {
	err := u.DB.Debug().Model(&model.User{}).Create(user).Error
	return err
}

func (u *UserRepository) Get() ([]*model.User, error) {
	users := []*model.User{}
	err := u.DB.Debug().Model(&model.User{}).Find(&users).Error
	return users, err
}

func (u *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := u.DB.Debug().Model(&model.User{}).Where("email = ?", email).First(&user).Error
	return user, err
}
