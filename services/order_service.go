package services

import (
	"errors"
	"foods-drinks-app/models"
	"foods-drinks-app/repository"

	"time"

	"github.com/google/uuid"
)

type OrderHistoryItem struct {
	OrderID     uuid.UUID             `json:"order_id"`
	Status      string                `json:"status"`
	TotalAmount float64               `json:"total_amount"`
	CreatedAt   string                `json:"created_at"`
	Items       []OrderHistoryProduct `json:"items"`
}

type OrderHistoryProduct struct {
	ProductID uuid.UUID `json:"product_id"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	Quantity  int64     `json:"quantity"`
	Price     int64     `json:"price"`
}

type OrderService interface {
	GetOrderHistory(userID uuid.UUID) ([]OrderHistoryItem, error)
	PlaceOrder(userID uuid.UUID, req PlaceOrderRequest) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo}
}

func (s *orderService) GetOrderHistory(userID uuid.UUID) ([]OrderHistoryItem, error) {
	orders, err := s.repo.GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}
	var history []OrderHistoryItem
	for _, order := range orders {
		products, err := s.repo.GetOrderItemsWithProduct(order.ID)
		if err != nil {
			return nil, err
		}
		var items []OrderHistoryProduct
		for _, p := range products {
			items = append(items, OrderHistoryProduct{
				ProductID: p.Product.ID,
				Name:      p.Product.Name,
				ImageURL:  p.Product.ImageURL,
				Quantity:  p.OrderItem.Quantity,
				Price:     p.OrderItem.Price,
			})
		}
		history = append(history, OrderHistoryItem{
			OrderID:     order.ID,
			Status:      order.Status,
			TotalAmount: order.TotalAmount,
			CreatedAt:   order.CreatedAt.Format("2006-01-02 15:04:05"),
			Items:       items,
		})
	}
	return history, nil
}

type PlaceOrderRequest struct {
	Items           []OrderItemRequest `json:"items"`
	CustomerName    string             `json:"customer_name"`
	CustomerPhone   string             `json:"customer_phone"`
	DeliveryAddress string             `json:"delivery_address"`
}

type OrderItemRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int64     `json:"quantity"`
}

func (s *orderService) PlaceOrder(userID uuid.UUID, req PlaceOrderRequest) error {
	if len(req.Items) == 0 {
		return errors.New("no items to order")
	}
	var orderItems []models.OrderItem
	total := 0.0
	for _, item := range req.Items {
		product, err := s.repo.GetProductByID(item.ProductID)
		if err != nil {
			return errors.New("product not found")
		}
		if item.Quantity <= 0 {
			return errors.New("invalid quantity")
		}
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     int64(product.Price),
		})
		total += product.Price * float64(item.Quantity)
	}
	order := &models.Order{
		ID:              uuid.New(),
		UserID:          userID,
		Status:          models.OrderStatusPending,
		TotalAmount:     total,
		CustomerName:    req.CustomerName,
		CustomerPhone:   req.CustomerPhone,
		DeliveryAddress: req.DeliveryAddress,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return s.repo.CreateOrder(order, orderItems)
}
