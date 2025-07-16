package models

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	OrderID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"order_id"`
	ProductID uuid.UUID `gorm:"type:uuid;primaryKey" json:"product_id"`
	Quantity  int64     `gorm:"type:bigint" json:"quantity"`
	Price     int64     `gorm:"type:bigint" json:"price"`
}
