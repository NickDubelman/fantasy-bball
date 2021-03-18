package api

import (
	"github.com/gin-gonic/gin"

	"github.com/NickDubelman/fantasy-bball/auth"
)

// Load initializes the API routes, middlewares, context, etc...
func Load() (*gin.Engine, error) {
	router := gin.Default()
	router.Use(auth.GoogleAuthFromConfig()) // Auth handlers
	return router, nil
}
