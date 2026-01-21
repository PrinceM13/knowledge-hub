package health

import (
	"net/http"

	"github.com/PrinceM13/knowledge-hub-api/internal/db"
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	if err := db.DB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "db down",
		})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
