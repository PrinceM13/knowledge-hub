package userhttp

import (
	"net/http"
	"strconv"

	"github.com/PrinceM13/knowledge-hub-api/internal/app"
	"github.com/gin-gonic/gin"
)

func Register(rg *gin.RouterGroup, a *app.App) {
	r := rg.Group("/users")

	r.POST("", registerUser(a))
	r.GET("", listUsers(a))
	r.GET("/:id", getUserByID(a))
}

func registerUser(a *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest

		// validate input
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		u, err := a.RegisterUser(c.Request.Context(), req.Email, req.Name)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, toUserDetail(u))
	}
}

func listUsers(a *app.App) gin.HandlerFunc {
	{
		return func(c *gin.Context) {
			// parse query parameters
			limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
			offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

			users, err := a.ListUsers(c.Request.Context(), limit, offset)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
				return
			}

			c.JSON(http.StatusOK, toUserListItems(users))
		}
	}
}

func getUserByID(a *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse user ID from path parameter
		userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			return
		}

		u, err := a.GetUserByID(c.Request.Context(), userID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
			return
		}

		c.JSON(http.StatusOK, toUserDetail(u))
	}
}
