package view

import (
	"book/api/user/internal/token"
	"book/initalize/database/mysql/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	u, err := user.GetUser(request.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, "Username does not exist")
		return
	}

	if u.Password != request.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	t, err := token.CreateToken(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	uuid := u.UUID
	c.JSON(http.StatusOK, map[string]string{"jwt": t, "uuid": uuid}) //test
}
