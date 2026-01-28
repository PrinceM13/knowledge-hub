package middleware

import (
	"log"

	"github.com/PrinceM13/knowledge-hub-api/internal/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// process request first and collect errors
		c.Next()

		// check if there were any errors during request processing
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// return just the last error but log all errors for debugging
			for _, e := range c.Errors {
				log.Printf("[ERROR] %s | Path: %s | Method: %s",
					e.Err.Error(),
					c.Request.URL.Path,
					c.Request.Method)
			}

			// try to convert to AppError
			if appErr, ok := errors.AsAppError(err); ok {
				c.JSON(appErr.HTTPStatus, gin.H{
					"error":   appErr.Code,
					"message": appErr.Message,
				})
				return
			}

			// default to internal server error for unknown errors
			c.JSON(errors.ErrInternal.HTTPStatus, gin.H{
				"error":   errors.ErrInternal.Code,
				"message": errors.ErrInternal.Message,
			})
		}
	}
}
