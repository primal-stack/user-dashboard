package userdashboard

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

var (
	// JWTAuth is the authentication middleware
	JWTAuth *jwtauth.JWTAuth
)

func init() {
	// Initialize JWTAuth to validate JSON Web Tokens
	JWTAuth = jwtauth.New(&jwtauth.JWTAuthenticator{
		// Set the secret key for signing and verifying tokens
		// Replace with your own secret key
		SecretKey: []byte("your-secret-key"),
	})
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	// Set a JSON error response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": "Unauthorized",
	})
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	// Set a JSON error response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return middleware.Handler(next)
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return JWTAuth.HandlerFunc(next)
}

func DBConnect(dsn string) (*sql.DB, error) {
	// Connect to the database
	return sql.Open("postgres", dsn)
}

func DBClose(db *sql.DB) error {
	// Close the database connection
	return db.Close()
}