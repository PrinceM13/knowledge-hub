package server

import (
	"github.com/PrinceM13/knowledge-hub-api/internal/health"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// routes
	r.GET("/health", health.Handler)

	return r
}
