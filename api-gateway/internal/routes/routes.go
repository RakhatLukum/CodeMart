package routes

import (
	"api-gateway/internal/handler"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	handler.RegisterHandlers(mux)
}
