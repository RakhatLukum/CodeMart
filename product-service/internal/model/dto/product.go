package dto

type ProductResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Tags  string  `json:"tags"`
}

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Tags  string  `json:"tags"`
}

type CreateProductResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Tags  string  `json:"tags"`
}

type UpdateProductRequest struct {
	ID    int     `json:"id" binding:"required"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Tags  string  `json:"tags"`
}

type DeleteProductResponse struct {
	Deleted bool `json:"deleted"`
}

type ListProductsResponse struct {
	Products []ProductResponse `json:"products"`
}

type SearchProductsRequest struct {
	Query string `form:"query"`
	Tags  string `form:"tags"`
}

type TagRequest struct {
	Tag string `json:"tag" binding:"required"`
}

type ProductIDRequest struct {
	ID int `json:"id" binding:"required"`
}

type CacheResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type SendProductEmailRequest struct {
	ProductID int    `json:"product_id" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type EmailStatusResponse struct {
	Sent    bool   `json:"sent"`
	Message string `json:"message,omitempty"`
}

type BulkCreateProductsRequest struct {
	Products []CreateProductRequest `json:"products" binding:"required"`
}

type BulkCreateProductsResponse struct {
	CreatedCount int `json:"created_count"`
}
