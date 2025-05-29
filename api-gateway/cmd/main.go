package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	analyticspb "github.com/RakhatLukum/CodeMart/api-gateway/proto/analytics"
	productpb "github.com/RakhatLukum/CodeMart/api-gateway/proto/product"
)

func main() {
	r := gin.Default()
	gin.SetMode(getEnv("GIN_MODE", "release"))

	productConn, err := grpc.NewClient(getEnv("PRODUCT_SERVICE", "localhost:50052"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()
	productClient := productpb.NewProductServiceClient(productConn)

	analyticsConn, err := grpc.NewClient(getEnv("ANALYTICS_SERVICE", "localhost:50054"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to analytics service: %v", err)
	}
	defer analyticsConn.Close()
	analyticsClient := analyticspb.NewViewServiceClient(analyticsConn)

	r.POST("/api/v1/products", func(c *gin.Context) {
		var req productpb.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := productClient.CreateProduct(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		res, err := productClient.GetProduct(context.Background(), &productpb.ProductIDRequest{Id: int32(id)})
		handleResponse(c, res, err)
	})

	r.PUT("/api/v1/products/:id", func(c *gin.Context) {
		var req productpb.UpdateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		req.Id = int32(id)
		res, err := productClient.UpdateProduct(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.DELETE("/api/v1/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		res, err := productClient.DeleteProduct(context.Background(), &productpb.ProductIDRequest{Id: int32(id)})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/products", func(c *gin.Context) {
		res, err := productClient.ListProducts(context.Background(), &productpb.Empty{})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/products/search", func(c *gin.Context) {
		res, err := productClient.SearchProducts(context.Background(), &productpb.SearchProductsRequest{
			Query: c.Query("query"),
			Tags:  c.Query("tags"),
		})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/products/by-tag", func(c *gin.Context) {
		res, err := productClient.GetProductsByTag(context.Background(), &productpb.TagRequest{Tag: c.Query("tag")})
		handleResponse(c, res, err)
	})

	r.POST("/api/v1/products/cache", func(c *gin.Context) {
		var req productpb.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := productClient.SetProductCache(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.DELETE("/api/v1/products/cache/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		res, err := productClient.InvalidateProductCache(context.Background(), &productpb.ProductIDRequest{Id: int32(id)})
		handleResponse(c, res, err)
	})

	r.POST("/api/v1/products/email", func(c *gin.Context) {
		var req productpb.SendProductEmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := productClient.SendProductEmail(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/products/cache/redis", func(c *gin.Context) {
		res, err := productClient.GetAllFromRedis(context.Background(), &productpb.Empty{})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/products/cache/memory", func(c *gin.Context) {
		res, err := productClient.GetAllFromCache(context.Background(), &productpb.Empty{})
		handleResponse(c, res, err)
	})

	r.POST("/api/v1/products/bulk", func(c *gin.Context) {
		var req productpb.BulkCreateProductsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := productClient.BulkCreateProducts(context.Background(), &req)
		handleResponse(c, res, err)
	})

	// Analytics Service Endpoints
	r.POST("/api/v1/views", func(c *gin.Context) {
		var req analyticspb.CreateViewRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := analyticsClient.CreateView(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/user/:user_id", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
			return
		}
		res, err := analyticsClient.GetViewsByUser(context.Background(), &analyticspb.UserRequest{UserId: int32(userID)})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/product/:product_id", func(c *gin.Context) {
		productID, err := strconv.Atoi(c.Param("product_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id"})
			return
		}
		res, err := analyticsClient.GetViewsByProduct(context.Background(), &analyticspb.ProductRequest{ProductId: int32(productID)})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/user-product", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
			return
		}
		productID, err := strconv.Atoi(c.Query("product_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id"})
			return
		}
		res, err := analyticsClient.GetViewsByUserAndProduct(context.Background(), &analyticspb.UserProductRequest{
			UserId:    int32(userID),
			ProductId: int32(productID),
		})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/recent", func(c *gin.Context) {
		limit := queryInt(c, "limit", 10)
		res, err := analyticsClient.GetRecentViews(context.Background(), &analyticspb.RecentViewsRequest{Limit: int32(limit)})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/most-viewed", func(c *gin.Context) {
		res, err := analyticsClient.GetMostViewedProducts(context.Background(), &analyticspb.Empty{})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/user-top-products/:user_id", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
			return
		}
		limit := queryInt(c, "limit", 5)
		res, err := analyticsClient.GetUserTopProducts(context.Background(), &analyticspb.UserTopProductsRequest{
			UserId: int32(userID),
			Limit:  int32(limit),
		})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/product-count/:product_id", func(c *gin.Context) {
		productID, err := strconv.Atoi(c.Param("product_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id"})
			return
		}
		res, err := analyticsClient.GetProductViewCount(context.Background(), &analyticspb.ProductRequest{ProductId: int32(productID)})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/user-count/:user_id", func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
			return
		}
		res, err := analyticsClient.GetUserViewCount(context.Background(), &analyticspb.UserRequest{UserId: int32(userID)})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/daily", func(c *gin.Context) {
		res, err := analyticsClient.GetDailyViews(context.Background(), &analyticspb.Empty{})
		handleResponse(c, res, err)
	})

	r.POST("/api/v1/views/report-email", func(c *gin.Context) {
		var req analyticspb.ReportEmailRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := analyticsClient.GenerateDailyViewReportEmail(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/hourly/:product_id", func(c *gin.Context) {
		productID, err := strconv.Atoi(c.Param("product_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id"})
			return
		}
		res, err := analyticsClient.GetHourlyViews(context.Background(), &analyticspb.ProductRequest{ProductId: int32(productID)})
		handleResponse(c, res, err)
	})

	r.DELETE("/api/v1/views/old", func(c *gin.Context) {
		var req analyticspb.DeleteOldViewsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := analyticsClient.DeleteOldViews(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/cache/:product_id", func(c *gin.Context) {
		productID, err := strconv.Atoi(c.Param("product_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id"})
			return
		}
		res, err := analyticsClient.GetCachedView(context.Background(), &analyticspb.ProductRequest{ProductId: int32(productID)})
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/views/memory-cache/:product_id", func(c *gin.Context) {
		productID, err := strconv.Atoi(c.Param("product_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id"})
			return
		}
		res, err := analyticsClient.GetMemoryCachedView(context.Background(), &analyticspb.ProductRequest{ProductId: int32(productID)})
		handleResponse(c, res, err)
	})

	server := &http.Server{
		Addr:    "0.0.0.0:" + getEnv("HTTP_PORT", "50050"),
		Handler: r,
	}

	go func() {
		log.Println("API Gateway running on", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	waitForShutdown(server)
}

func handleResponse(c *gin.Context, res any, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func waitForShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown failed: %v", err)
	}
	log.Println("Server exited gracefully")
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func queryInt(c *gin.Context, key string, defaultVal int) int {
	val, err := strconv.Atoi(c.Query(key))
	if err != nil {
		return defaultVal
	}
	return val
}
