package repository

import (
	"category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAll() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDB: db}
}

type CategoryRepository struct {
	mysqlDB *gorm.DB
}

func (u *CategoryRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.Category{}).Error
}

func (u *CategoryRepository) AddCategory(category *model.Category) (int64, error) {
	return category.ID, u.mysqlDB.Create(category).Create(category).Error
}

func (u *CategoryRepository) DeleteCategory(i int64) error {
	return u.mysqlDB.Delete(i).Error
}

func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	return u.mysqlDB.Model(category).Update(category).Error
}

func (u *CategoryRepository) FindCategoryByID(i int64) (category *model.Category, err error) {
	return category, u.mysqlDB.Where("id = ?", i).Find(&category).Error
}

func (u *CategoryRepository) FindAll() (categoryAll []model.Category, err error) {
	return categoryAll, u.mysqlDB.Find(&categoryAll).Error
}

func (u *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	return category, u.mysqlDB.Where("category_name = ?", categoryName).Find(&category).Error
}

func (u *CategoryRepository) FindCategoryByLevel(categoryLevel uint32) (category []model.Category, err error) {
	return category, u.mysqlDB.Where("category_level = ?", categoryLevel).Find(&category).Error
}

func (u *CategoryRepository) FindCategoryByParent(categoryParent int64) (category []model.Category, err error) {
	return category, u.mysqlDB.Where("category_level = ?", categoryParent).Find(&category).Error
}
