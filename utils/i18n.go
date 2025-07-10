package utils

import (
	"fmt"
	"strings"
)

// MessageType định nghĩa loại message
type MessageType string

const (
	// Error messages
	ErrorInvalidEmail           MessageType = "error.invalid_email"
	ErrorInvalidPassword        MessageType = "error.invalid_password"
	ErrorUserNotFound           MessageType = "error.user_not_found"
	ErrorProductNotFound        MessageType = "error.product_not_found"
	ErrorCartItemNotFound       MessageType = "error.cart_item_not_found"
	ErrorOrderNotFound          MessageType = "error.order_not_found"
	ErrorInvalidQuantity        MessageType = "error.invalid_quantity"
	ErrorInvalidUserID          MessageType = "error.invalid_user_id"
	ErrorInvalidProductID       MessageType = "error.invalid_product_id"
	ErrorInvalidToken           MessageType = "error.invalid_token"
	ErrorExpiredToken           MessageType = "error.expired_token"
	ErrorMissingAuthHeader      MessageType = "error.missing_auth_header"
	ErrorFailedToHashPassword   MessageType = "error.failed_to_hash_password"
	ErrorIncorrectOldPassword   MessageType = "error.incorrect_old_password"
	ErrorFailedToUpdateProfile  MessageType = "error.failed_to_update_profile"
	ErrorEmailAlreadyExists     MessageType = "error.email_already_exists"
	ErrorAccountDisabled        MessageType = "error.account_disabled"
	ErrorInvalidCredentials     MessageType = "error.invalid_credentials"
	ErrorNoItemsToOrder         MessageType = "error.no_items_to_order"
	ErrorInvalidRequestBody     MessageType = "error.invalid_request_body"
	ErrorFailedToCreateOrder    MessageType = "error.failed_to_create_order"
	ErrorFailedToUpdateCart     MessageType = "error.failed_to_update_cart"
	ErrorFailedToAddToCart      MessageType = "error.failed_to_add_to_cart"
	ErrorFailedToRemoveFromCart MessageType = "error.failed_to_remove_from_cart"
	ErrorFailedToClearCart      MessageType = "error.failed_to_clear_cart"
	ErrorFailedToRestoreCart    MessageType = "error.failed_to_restore_cart"

	// Success messages
	SuccessPasswordChanged        MessageType = "success.password_changed"
	SuccessProfileUpdated         MessageType = "success.profile_updated"
	SuccessOrderPlaced            MessageType = "success.order_placed"
	SuccessProductAddedToCart     MessageType = "success.product_added_to_cart"
	SuccessCartItemUpdated        MessageType = "success.cart_item_updated"
	SuccessProductRemovedFromCart MessageType = "success.product_removed_from_cart"
	SuccessCartCleared            MessageType = "success.cart_cleared"
	SuccessProductRestored        MessageType = "success.product_restored"
	SuccessCheckEmailForReset     MessageType = "success.check_email_for_reset"
	SuccessPasswordReset          MessageType = "success.password_reset"

	// Info messages
	InfoCartEmpty      MessageType = "info.cart_empty"
	InfoNoDeletedItems MessageType = "info.no_deleted_items"
)

// Language định nghĩa ngôn ngữ
type Language string

const (
	Vietnamese Language = "vi"
	English    Language = "en"
)

// messages chứa tất cả message theo ngôn ngữ
var messages = map[Language]map[MessageType]string{
	Vietnamese: {
		// Error messages
		ErrorInvalidEmail:           "Email không hợp lệ",
		ErrorInvalidPassword:        "Mật khẩu phải có ít nhất 6 ký tự",
		ErrorUserNotFound:           "Không tìm thấy người dùng",
		ErrorProductNotFound:        "Không tìm thấy sản phẩm",
		ErrorCartItemNotFound:       "Không tìm thấy sản phẩm trong giỏ hàng",
		ErrorOrderNotFound:          "Không tìm thấy đơn hàng",
		ErrorInvalidQuantity:        "Số lượng không hợp lệ",
		ErrorInvalidUserID:          "ID người dùng không hợp lệ",
		ErrorInvalidProductID:       "ID sản phẩm không hợp lệ",
		ErrorInvalidToken:           "Token không hợp lệ",
		ErrorExpiredToken:           "Token đã hết hạn",
		ErrorMissingAuthHeader:      "Thiếu hoặc không hợp lệ header Authorization",
		ErrorFailedToHashPassword:   "Không thể mã hóa mật khẩu",
		ErrorIncorrectOldPassword:   "Mật khẩu cũ không đúng",
		ErrorFailedToUpdateProfile:  "Không thể cập nhật hồ sơ",
		ErrorEmailAlreadyExists:     "Email đã tồn tại",
		ErrorAccountDisabled:        "Tài khoản đã bị vô hiệu hóa",
		ErrorInvalidCredentials:     "Thông tin đăng nhập không đúng",
		ErrorNoItemsToOrder:         "Không có sản phẩm để đặt hàng",
		ErrorInvalidRequestBody:     "Dữ liệu yêu cầu không hợp lệ",
		ErrorFailedToCreateOrder:    "Không thể tạo đơn hàng",
		ErrorFailedToUpdateCart:     "Không thể cập nhật giỏ hàng",
		ErrorFailedToAddToCart:      "Không thể thêm vào giỏ hàng",
		ErrorFailedToRemoveFromCart: "Không thể xóa khỏi giỏ hàng",
		ErrorFailedToClearCart:      "Không thể xóa giỏ hàng",
		ErrorFailedToRestoreCart:    "Không thể khôi phục giỏ hàng",

		// Success messages
		SuccessPasswordChanged:        "Đổi mật khẩu thành công",
		SuccessProfileUpdated:         "Cập nhật hồ sơ thành công",
		SuccessOrderPlaced:            "Đặt hàng thành công",
		SuccessProductAddedToCart:     "Thêm sản phẩm vào giỏ hàng thành công",
		SuccessCartItemUpdated:        "Cập nhật giỏ hàng thành công",
		SuccessProductRemovedFromCart: "Xóa sản phẩm khỏi giỏ hàng thành công",
		SuccessCartCleared:            "Xóa giỏ hàng thành công",
		SuccessProductRestored:        "Khôi phục sản phẩm thành công",
		SuccessCheckEmailForReset:     "Kiểm tra email để đặt lại mật khẩu",
		SuccessPasswordReset:          "Đặt lại mật khẩu thành công",

		// Info messages
		InfoCartEmpty:      "Giỏ hàng trống",
		InfoNoDeletedItems: "Không có sản phẩm đã xóa",
	},
	English: {
		// Error messages
		ErrorInvalidEmail:           "Invalid email format",
		ErrorInvalidPassword:        "Password must be at least 6 characters",
		ErrorUserNotFound:           "User not found",
		ErrorProductNotFound:        "Product not found",
		ErrorCartItemNotFound:       "Cart item not found",
		ErrorOrderNotFound:          "Order not found",
		ErrorInvalidQuantity:        "Invalid quantity",
		ErrorInvalidUserID:          "Invalid user ID",
		ErrorInvalidProductID:       "Invalid product ID",
		ErrorInvalidToken:           "Invalid token",
		ErrorExpiredToken:           "Token expired",
		ErrorMissingAuthHeader:      "Missing or invalid Authorization header",
		ErrorFailedToHashPassword:   "Failed to hash password",
		ErrorIncorrectOldPassword:   "Incorrect old password",
		ErrorFailedToUpdateProfile:  "Failed to update profile",
		ErrorEmailAlreadyExists:     "Email already exists",
		ErrorAccountDisabled:        "Account is disabled",
		ErrorInvalidCredentials:     "Invalid credentials",
		ErrorNoItemsToOrder:         "No items to order",
		ErrorInvalidRequestBody:     "Invalid request body",
		ErrorFailedToCreateOrder:    "Failed to create order",
		ErrorFailedToUpdateCart:     "Failed to update cart",
		ErrorFailedToAddToCart:      "Failed to add to cart",
		ErrorFailedToRemoveFromCart: "Failed to remove from cart",
		ErrorFailedToClearCart:      "Failed to clear cart",
		ErrorFailedToRestoreCart:    "Failed to restore cart",

		// Success messages
		SuccessPasswordChanged:        "Password changed successfully",
		SuccessProfileUpdated:         "Profile updated successfully",
		SuccessOrderPlaced:            "Order placed successfully",
		SuccessProductAddedToCart:     "Product added to cart successfully",
		SuccessCartItemUpdated:        "Cart item updated successfully",
		SuccessProductRemovedFromCart: "Product removed from cart successfully",
		SuccessCartCleared:            "Cart cleared successfully",
		SuccessProductRestored:        "Product restored successfully",
		SuccessCheckEmailForReset:     "Check your email to reset password",
		SuccessPasswordReset:          "Password reset successfully",

		// Info messages
		InfoCartEmpty:      "Cart is empty",
		InfoNoDeletedItems: "No deleted items",
	},
}

// GetMessage lấy message theo ngôn ngữ và loại
func GetMessage(lang Language, msgType MessageType) string {
	if langMessages, exists := messages[lang]; exists {
		if message, exists := langMessages[msgType]; exists {
			return message
		}
	}
	// Fallback to English if language or message not found
	if enMessages, exists := messages[English]; exists {
		if message, exists := enMessages[msgType]; exists {
			return message
		}
	}
	return string(msgType)
}

// GetMessageWithParams lấy message với tham số
func GetMessageWithParams(lang Language, msgType MessageType, params map[string]interface{}) string {
	message := GetMessage(lang, msgType)

	for key, value := range params {
		placeholder := fmt.Sprintf("{%s}", key)
		message = strings.ReplaceAll(message, placeholder, fmt.Sprintf("%v", value))
	}

	return message
}

// GetLanguageFromHeader lấy ngôn ngữ từ Accept-Language header
func GetLanguageFromHeader(acceptLanguage string) Language {
	if strings.Contains(acceptLanguage, "vi") {
		return Vietnamese
	}
	return English
}
