package repository

import (
	"foods-drinks-app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrdersByUserID(userID uuid.UUID) ([]models.Order, error)
	GetOrderItemsWithProduct(orderID uuid.UUID) ([]OrderItemWithProduct, error)
	CreateOrder(order *models.Order, items []models.OrderItem) error
	GetProductByID(productID uuid.UUID) (*models.Product, error)
}

type orderRepository struct {
	db *gorm.DB
}

type OrderItemWithProduct struct {
	OrderItem models.OrderItem
	Product   models.Product
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetOrdersByUserID(userID uuid.UUID) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetOrderItemsWithProduct(orderID uuid.UUID) ([]OrderItemWithProduct, error) {
	var results []OrderItemWithProduct
	var items []models.OrderItem
	if err := r.db.Where("order_id = ?", orderID).Find(&items).Error; err != nil {
		return nil, err
	}
	for _, item := range items {
		var product models.Product
		if err := r.db.First(&product, "id = ?", item.ProductID).Error; err != nil {
			return nil, err
		}
		results = append(results, OrderItemWithProduct{OrderItem: item, Product: product})
	}
	return results, nil
}

func (r *orderRepository) CreateOrder(order *models.Order, items []models.OrderItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for _, item := range items {
			item.OrderID = order.ID
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *orderRepository) GetProductByID(productID uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, "id = ?", productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
