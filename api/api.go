package api

import (
	"context"
	"log"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/NickDubelman/fantasy-bball/auth"
	"github.com/NickDubelman/fantasy-bball/config"
	"github.com/NickDubelman/fantasy-bball/db"
)

// Load initializes the API routes, middlewares, context, etc...
func Load() (*gin.Engine, error) {
	router := gin.Default()

	// Connect to app database
	driver, err := sql.Open("mysql", config.Get().Database.DSN())
	if err != nil {
		return nil, err
	}

	appDB := driver.DB()
	appDB.SetConnMaxLifetime(3 * time.Minute)
	appDB.SetMaxIdleConns(100)
	appDB.SetMaxOpenConns(100)

	client := db.NewClient(db.Driver(driver))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Middleware to make db client accessible via request context
	router.Use(func(c *gin.Context) {
		ctx := db.NewContext(c.Request.Context(), client)
		c.Request = c.Request.WithContext(ctx)
	})

	router.Use(auth.GoogleAuthFromConfig()) // Auth handlers

	return router, nil
}
