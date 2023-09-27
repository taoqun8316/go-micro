package model

type Cart struct {
	ID        int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ProductID int64 `gorm:"not_null" json:"product_id"`
	num       int64 `gorm:"not_null" json:"num"`
	SizeID    int64 `json:"size_id"`
	UserID    int64 `json:"user_id"`
}
