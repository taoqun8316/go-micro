package repository

import (
	"cart/domain/model"
	"errors"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)
	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDB: db}
}

type CartRepository struct {
	mysqlDB *gorm.DB
}

func (u *CartRepository) FindCartByID(i int64) (cart *model.Cart, err error) {
	return cart, u.mysqlDB.First(cart, i).Error
}

func (u *CartRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.Cart{}).Error
}

func (u *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := u.mysqlDB.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}

func (u *CartRepository) DeleteCartByID(i int64) error {
	return u.mysqlDB.Where("id = ?", i).Delete(&model.Cart{}).Error
}

func (u *CartRepository) UpdateCart(cart *model.Cart) error {
	return u.mysqlDB.Model(cart).Update(cart).Error
}

func (u *CartRepository) FindAll(uid int64) (all []model.Cart, err error) {
	return all, u.mysqlDB.Where("user_id = ?", uid).Find(&all).Error
}

func (u *CartRepository) CleanCart(uid int64) error {
	return u.mysqlDB.Where("user_id = ?", uid).Delete(&model.Cart{}).Error
}

func (u *CartRepository) IncrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	return u.mysqlDB.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (u *CartRepository) DecrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	db := u.mysqlDB.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
