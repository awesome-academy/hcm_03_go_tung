package repository

import (
	"errors"
	"foods-drinks-app/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetCartItemsWithProduct(userID uuid.UUID) ([]CartItemWithProduct, error)
	AddToCart(userID, productID uuid.UUID, quantity int64) error
	UpdateCartItem(userID, productID uuid.UUID, quantity int64) error
	RemoveFromCart(userID, productID uuid.UUID) error
	ClearCart(userID uuid.UUID) error
	GetCartItem(userID, productID uuid.UUID) (*models.Cart, error)
	GetProductByID(productID uuid.UUID) (*models.Product, error)
	GetDeletedCartItems(userID uuid.UUID) ([]CartItemWithProduct, error)
	RestoreCartItem(userID, productID uuid.UUID) error
}

type cartRepository struct {
	db *gorm.DB
}

type CartItemWithProduct struct {
	Cart    models.Cart    `json:"cart"`
	Product models.Product `json:"product"`
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

// GetCartItemsWithProduct lấy tất cả items trong giỏ hàng với thông tin sản phẩm (chỉ những item chưa bị xóa)
func (r *cartRepository) GetCartItemsWithProduct(userID uuid.UUID) ([]CartItemWithProduct, error) {
	var results []CartItemWithProduct
	var cartItems []models.Cart

	if err := r.db.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}

	for _, cartItem := range cartItems {
		var product models.Product
		if err := r.db.First(&product, "id = ?", cartItem.ProductID).Error; err != nil {
			return nil, err
		}
		results = append(results, CartItemWithProduct{
			Cart:    cartItem,
			Product: product,
		})
	}

	return results, nil
}

// AddToCart thêm sản phẩm vào giỏ hàng
func (r *cartRepository) AddToCart(userID, productID uuid.UUID, quantity int64) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// Kiểm tra sản phẩm đã có trong giỏ hàng chưa
	var existingCart models.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingCart).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Thêm mới nếu chưa có
			cart := models.Cart{
				ID:        uuid.New(),
				UserID:    userID,
				ProductID: productID,
				Quantity:  quantity,
			}
			return r.db.Create(&cart).Error
		}
		return err
	}

	// Cập nhật số lượng nếu đã có
	existingCart.Quantity += quantity
	return r.db.Save(&existingCart).Error
}

// UpdateCartItem cập nhật số lượng sản phẩm trong giỏ hàng
func (r *cartRepository) UpdateCartItem(userID, productID uuid.UUID, quantity int64) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	var cart models.Cart
	if err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error; err != nil {
		return errors.New("cart item not found")
	}

	cart.Quantity = quantity
	return r.db.Save(&cart).Error
}

// RemoveFromCart xóa sản phẩm khỏi giỏ hàng (soft delete)
func (r *cartRepository) RemoveFromCart(userID, productID uuid.UUID) error {
	now := time.Now()
	result := r.db.Model(&models.Cart{}).
		Where("user_id = ? AND product_id = ? AND deleted_at IS NULL", userID, productID).
		Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("cart item not found")
	}
	return nil
}

// ClearCart xóa toàn bộ giỏ hàng của user (soft delete)
func (r *cartRepository) ClearCart(userID uuid.UUID) error {
	now := time.Now()
	result := r.db.Model(&models.Cart{}).
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetCartItem lấy thông tin một item trong giỏ hàng
func (r *cartRepository) GetCartItem(userID, productID uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepository) GetProductByID(productID uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, "id = ?", productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetDeletedCartItems lấy danh sách các sản phẩm đã bị xóa
func (r *cartRepository) GetDeletedCartItems(userID uuid.UUID) ([]CartItemWithProduct, error) {
	var results []CartItemWithProduct
	var cartItems []models.Cart

	if err := r.db.Where("user_id = ? AND deleted_at IS NOT NULL", userID).Find(&cartItems).Error; err != nil {
		return nil, err
	}

	for _, cartItem := range cartItems {
		var product models.Product
		if err := r.db.First(&product, "id = ?", cartItem.ProductID).Error; err != nil {
			return nil, err
		}
		results = append(results, CartItemWithProduct{
			Cart:    cartItem,
			Product: product,
		})
	}

	return results, nil
}

// RestoreCartItem khôi phục sản phẩm đã bị xóa
func (r *cartRepository) RestoreCartItem(userID, productID uuid.UUID) error {
	result := r.db.Model(&models.Cart{}).
		Where("user_id = ? AND product_id = ? AND deleted_at IS NOT NULL", userID, productID).
		Update("deleted_at", nil)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("deleted cart item not found")
	}
	return nil
}
