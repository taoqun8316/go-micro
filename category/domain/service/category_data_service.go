package service

import (
	"category/domain/model"
	"category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

func NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{CategoryRepository: categoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

func (c CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return c.CategoryRepository.AddCategory(category)
}

func (c CategoryDataService) DeleteCategory(i int64) error {
	return c.CategoryRepository.DeleteCategory(i)
}

func (c CategoryDataService) UpdateCategory(category *model.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}

func (c CategoryDataService) FindCategoryByID(i int64) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByID(i)
}

func (c CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return c.CategoryRepository.FindAll()
}

func (c CategoryDataService) FindCategoryByName(name string) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByName(name)
}
func (c CategoryDataService) FindCategoryByLevel(i uint32) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByLevel(i)
}
func (c CategoryDataService) FindCategoryByParent(i int64) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByParent(i)
}
