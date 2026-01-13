package handlers

import (
	"fmt"
	"net/http"
)

// GetRootHandler handles GET requests to root path
func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! HTTP server is running on port %s", r.URL.Query().Get("port"))
}