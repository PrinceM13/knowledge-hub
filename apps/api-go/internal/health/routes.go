package health

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/health", handler)
}
