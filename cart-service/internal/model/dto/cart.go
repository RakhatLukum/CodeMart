package dto

import "github.com/RakhatLukum/CodeMart/cart-service/internal/model"

type CartResponse struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

type CreateCartRequest struct {
	UserID    int `json:"user_id" binding:"required"`
	ProductID int `json:"product_id" binding:"required"`
}

type CreateCartResponse struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

type DeleteCartItemRequest struct {
	UserID    int `json:"user_id" binding:"required"`
	ProductID int `json:"product_id" binding:"required"`
}

type DeleteCartItemResponse struct {
	Deleted bool `json:"deleted"`
}

type UserCartResponse struct {
	UserID int             `json:"user_id"`
	Items  []model.Product `json:"items"`
}
