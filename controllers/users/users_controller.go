package users

import (
	"net/http"
	"strconv"

	"github.com/galileoluna/bookstore_users-api/domain/users"
	"github.com/galileoluna/bookstore_users-api/services"
	"github.com/galileoluna/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		//TODO : return bad request
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		// TODO: handdle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
