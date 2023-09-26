package service

import (
	"product/domain/model"
	"product/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

func NewProductDataService(productRepository repository.IProductRepository) IProductDataService {
	return &ProductDataService{ProductRepository: productRepository}
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

func (c ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return c.ProductRepository.AddProduct(product)
}

func (c ProductDataService) DeleteProduct(i int64) error {
	return c.ProductRepository.DeleteProduct(i)
}

func (c ProductDataService) UpdateProduct(category *model.Product) error {
	return c.ProductRepository.UpdateProduct(category)
}

func (c ProductDataService) FindProductByID(i int64) (*model.Product, error) {
	return c.ProductRepository.FindProductByID(i)
}

func (c ProductDataService) FindAllProduct() ([]model.Product, error) {
	return c.ProductRepository.FindAll()
}
