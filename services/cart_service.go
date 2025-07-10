package services

import (
	"errors"
	"foods-drinks-app/repository"

	"github.com/google/uuid"
)

type CartService interface {
	GetCart(userID uuid.UUID) ([]CartItemResponse, error)
	AddToCart(userID, productID uuid.UUID, quantity int64) error
	UpdateCartItem(userID, productID uuid.UUID, quantity int64) error
	RemoveFromCart(userID, productID uuid.UUID) error
	ClearCart(userID uuid.UUID) error
	GetDeletedCartItems(userID uuid.UUID) ([]CartItemResponse, error)
	RestoreCartItem(userID, productID uuid.UUID) error
}

type cartService struct {
	repo repository.CartRepository
}

type CartItemResponse struct {
	CartID      uuid.UUID `json:"cart_id"`
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	ImageURL    string    `json:"image_url"`
	Price       float64   `json:"price"`
	Quantity    int64     `json:"quantity"`
	Subtotal    float64   `json:"subtotal"`
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo: repo}
}

// GetCart lấy thông tin giỏ hàng với sản phẩm
func (s *cartService) GetCart(userID uuid.UUID) ([]CartItemResponse, error) {
	cartItems, err := s.repo.GetCartItemsWithProduct(userID)
	if err != nil {
		return nil, err
	}

	var responses []CartItemResponse
	for _, item := range cartItems {
		subtotal := item.Product.Price * float64(item.Cart.Quantity)
		responses = append(responses, CartItemResponse{
			CartID:      item.Cart.ID,
			ProductID:   item.Cart.ProductID,
			ProductName: item.Product.Name,
			ImageURL:    item.Product.ImageURL,
			Price:       item.Product.Price,
			Quantity:    item.Cart.Quantity,
			Subtotal:    subtotal,
		})
	}

	return responses, nil
}

// AddToCart thêm sản phẩm vào giỏ hàng
func (s *cartService) AddToCart(userID, productID uuid.UUID, quantity int64) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// Kiểm tra sản phẩm có tồn tại không
	_, err := s.repo.GetProductByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	return s.repo.AddToCart(userID, productID, quantity)
}

// UpdateCartItem cập nhật số lượng sản phẩm trong giỏ hàng
func (s *cartService) UpdateCartItem(userID, productID uuid.UUID, quantity int64) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	return s.repo.UpdateCartItem(userID, productID, quantity)
}

// RemoveFromCart xóa sản phẩm khỏi giỏ hàng
func (s *cartService) RemoveFromCart(userID, productID uuid.UUID) error {
	return s.repo.RemoveFromCart(userID, productID)
}

// ClearCart xóa toàn bộ giỏ hàng
func (s *cartService) ClearCart(userID uuid.UUID) error {
	return s.repo.ClearCart(userID)
}

// GetDeletedCartItems lấy danh sách sản phẩm đã bị xóa
func (s *cartService) GetDeletedCartItems(userID uuid.UUID) ([]CartItemResponse, error) {
	cartItems, err := s.repo.GetDeletedCartItems(userID)
	if err != nil {
		return nil, err
	}

	var responses []CartItemResponse
	for _, item := range cartItems {
		subtotal := item.Product.Price * float64(item.Cart.Quantity)
		responses = append(responses, CartItemResponse{
			CartID:      item.Cart.ID,
			ProductID:   item.Cart.ProductID,
			ProductName: item.Product.Name,
			ImageURL:    item.Product.ImageURL,
			Price:       item.Product.Price,
			Quantity:    item.Cart.Quantity,
			Subtotal:    subtotal,
		})
	}

	return responses, nil
}

// RestoreCartItem khôi phục sản phẩm đã bị xóa
func (s *cartService) RestoreCartItem(userID, productID uuid.UUID) error {
	return s.repo.RestoreCartItem(userID, productID)
}
