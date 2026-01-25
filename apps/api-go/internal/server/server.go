package server

import (
	"github.com/PrinceM13/knowledge-hub-api/internal/app"
	"github.com/PrinceM13/knowledge-hub-api/internal/health"
	v1 "github.com/PrinceM13/knowledge-hub-api/internal/http/v1"
	"github.com/gin-gonic/gin"
)

func New(app *app.App) *gin.Engine {
	r := gin.New()

	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// routes
	health.Register(r)
	v1.Register(r, app)

	return r
}
