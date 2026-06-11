package middleware

import (
	"github.com/go-chi/cors"
)

func GetCorsConfig(origins []string, methods []string, allowedHeaders []string, exposedHeaders []string) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   methods,
		AllowedHeaders:   allowedHeaders,
		ExposedHeaders:   exposedHeaders,
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
