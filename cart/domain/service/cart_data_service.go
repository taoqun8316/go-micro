package service

import (
	"cart/domain/model"
	"cart/domain/repository"
)

type ICartDataService interface {
	FindCartByID(int64) (*model.Cart, error)
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindAllCart(int64) ([]model.Cart, error)
	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{CartRepository: cartRepository}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func (c CartDataService) FindCartByID(i int64) (*model.Cart, error) {
	return c.CartRepository.FindCartByID(i)
}

func (c CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return c.CartRepository.CreateCart(cart)
}

func (c CartDataService) DeleteCart(i int64) error {
	return c.CartRepository.DeleteCartByID(i)
}

func (c CartDataService) UpdateCart(cart *model.Cart) error {
	return c.CartRepository.UpdateCart(cart)
}

func (c CartDataService) FindAllCart(i int64) ([]model.Cart, error) {
	return c.CartRepository.FindAll(i)
}

func (c CartDataService) CleanCart(i int64) error {
	return c.CartRepository.CleanCart(i)
}

func (c CartDataService) IncrNum(cartID int64, num int64) error {
	return c.CartRepository.IncrNum(cartID, num)
}

func (c CartDataService) DecrNum(cartID int64, num int64) error {
	return c.CartRepository.DecrNum(cartID, num)
}
