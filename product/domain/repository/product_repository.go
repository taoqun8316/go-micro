package repository

import (
	"github.com/jinzhu/gorm"
	"product/domain/model"
)

type IProductRepository interface {
	InitTable() error
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAll() ([]model.Product, error)
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDB: db}
}

type ProductRepository struct {
	mysqlDB *gorm.DB
}

func (u *ProductRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.Product{}, &model.ProductImage{}, &model.ProductSeo{}, &model.ProductSize{}).Error
}

func (u *ProductRepository) AddProduct(product *model.Product) (int64, error) {
	return product.ID, u.mysqlDB.Create(product).Create(product).Error
}

func (u *ProductRepository) DeleteProduct(productID int64) error {
	tx := u.mysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("image_product_id = ?", productID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (u *ProductRepository) UpdateProduct(product *model.Product) error {
	return u.mysqlDB.Model(product).Update(product).Error
}

func (u *ProductRepository) FindProductByID(i int64) (product *model.Product, err error) {
	return product, u.mysqlDB.Where("id = ?", i).Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&product).Error
}

func (u *ProductRepository) FindAll() (productAll []model.Product, err error) {
	return productAll, u.mysqlDB.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&productAll).Error
}
