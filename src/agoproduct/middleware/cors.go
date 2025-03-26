package middleware

import (
	"fmt"
	"net/http"
)

type CorsHandler struct {
}

func (h *CorsHandler) addCorsHeaders(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("addCorsHeaders")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
		return
	}
}

func (h *CorsHandler) PreflightHandler() http.Handler {
	return http.HandlerFunc(h.addCorsHeaders)
}
