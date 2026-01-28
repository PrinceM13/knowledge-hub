package userhttp

import (
	"net/http"
	"strconv"

	"github.com/PrinceM13/knowledge-hub-api/internal/app"
	"github.com/PrinceM13/knowledge-hub-api/internal/errors"
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

		// bind JSON body to request struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(errors.InvalidRequest(err))
			return
		}

		u, err := a.RegisterUser(c.Request.Context(), req.Email, req.Name)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusCreated, toUserDetail(u))
	}
}

func listUsers(a *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse query parameters with validation
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
		if err != nil {
			c.Error(errors.InvalidQueryParam(err, "limit"))
			return
		}

		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		if err != nil {
			c.Error(errors.InvalidQueryParam(err, "offset"))
			return
		}

		users, err := a.ListUsers(c.Request.Context(), limit, offset)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, toUserListItems(users))
	}
}

func getUserByID(a *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse user ID from path parameter
		userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Error(errors.InvalidPathParam(err, "id"))
			return
		}

		u, err := a.GetUserByID(c.Request.Context(), userID)

		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, toUserDetail(u))
	}
}
