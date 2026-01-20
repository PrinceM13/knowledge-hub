package health

import (
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
