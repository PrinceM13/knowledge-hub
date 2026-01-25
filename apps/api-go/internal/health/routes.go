package health

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	r.GET("/health", handler)
}
