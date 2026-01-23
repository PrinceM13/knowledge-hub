package server

import (
	"github.com/PrinceM13/knowledge-hub-api/internal/health"
	"github.com/gin-gonic/gin"
)

func New(app *App) *gin.Engine {
	r := gin.New()

	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// versioned API group
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// routes
	health.Register(v1)

	return r
}
