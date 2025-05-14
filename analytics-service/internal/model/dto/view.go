package dto

import "time"

type ViewResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Timestamp time.Time `json:"timestamp"`
}

type CreateViewRequest struct {
	UserID    int `json:"user_id" binding:"required"`
	ProductID int `json:"product_id" binding:"required"`
}

type CreateViewResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Timestamp time.Time `json:"timestamp"`
}

type UserViewsResponse struct {
	Views []ViewResponse `json:"views"`
}

type ProductViewsResponse struct {
	Views []ViewResponse `json:"views"`
}

type UserProductViewsResponse struct {
	Views []ViewResponse `json:"views"`
}

type RecentViewsRequest struct {
	Limit int `form:"limit"`
}

type RecentViewsResponse struct {
	Views []ViewResponse `json:"views"`
}

type MostViewedProductsRequest struct {
	Start string `form:"start" time_format:"2006-01-02"`
	End   string `form:"end" time_format:"2006-01-02"`
}

type ProductViewCount struct {
	ProductID int `json:"product_id"`
	ViewCount int `json:"view_count"`
}

type MostViewedProductsResponse struct {
	Data []ProductViewCount `json:"data"`
}

type UserTopProductsResponse struct {
	UserID int                `json:"user_id"`
	Data   []ProductViewCount `json:"data"`
}

type ProductViewCountResponse struct {
	ProductID int `json:"product_id"`
	ViewCount int `json:"view_count"`
}

type UserViewCountResponse struct {
	UserID    int `json:"user_id"`
	ViewCount int `json:"view_count"`
}

type DailyViewsRequest struct {
	Start string `form:"start"`
	End   string `form:"end"`
}

type DailyViewStat struct {
	Date      string `json:"date"`
	ViewCount int    `json:"view_count"`
}

type DailyViewsResponse struct {
	Data []DailyViewStat `json:"data"`
}

type HourlyViewStat struct {
	Hour      int `json:"hour"`
	ViewCount int `json:"view_count"`
}

type HourlyViewsResponse struct {
	ProductID int              `json:"product_id"`
	Data      []HourlyViewStat `json:"data"`
}

type DeleteOldViewsRequest struct {
	Before string `form:"before"`
}

type DeleteOldViewsResponse struct {
	DeletedCount int `json:"deleted_count"`
}
