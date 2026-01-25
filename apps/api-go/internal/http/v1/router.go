package v1

import (
	"github.com/PrinceM13/knowledge-hub-api/internal/app"
	userhttp "github.com/PrinceM13/knowledge-hub-api/internal/http/v1/user"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, app *app.App) {
	api := r.Group("/api/v1")

	userhttp.Register(api, app)
}
