package repository

import (
	"log"

	model "github.com/AhmadSafrizal/golang-dasar/tugas7/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (prod *ProductRepository) Migrate() {
	err := prod.DB.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatal(err)
	}
}

func (prod *ProductRepository) Create(product *model.Product) error {
	err := prod.DB.Debug().Model(&model.Product{}).Create(product).Error
	return err
}

func (prod *ProductRepository) Get() ([]*model.Product, error) {
	products := []*model.Product{}
	err := prod.DB.Debug().Model(&model.Product{}).Preload("UserDetail").Find(&products).Error
	return products, err
}

func (prod *ProductRepository) Update(id uint, product *model.Product) error {
	err := prod.DB.Debug().Model(&model.Product{}).Preload("UserDetail").Where("id = ?", id).Save(product).Error
	return err
}

func (prod *ProductRepository) Delete(id uint) error {
	err := prod.DB.Debug().Model(&model.Product{}).Where("id = ?", id).Delete(&model.Product{}).Error
	return err
}
