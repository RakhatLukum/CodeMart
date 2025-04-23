package routes

import (
	"net/http"
	"api-gateway/internal/handler"
)

func RegisterRoutes(mux *http.ServeMux) {
	handler.RegisterHandlers(mux)
}
