package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tjhooper1/bookstore_users-api/domain/users"
	"github.com/tjhooper1/bookstore_users-api/services"
	"github.com/tjhooper1/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.BindJSON(&user); err != nil {
		//TODO HANDLE JSON ERROR
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		//TODO HANDLE USER CREATION ERROR
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
