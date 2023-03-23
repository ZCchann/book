package view

import (
	"book/initalize/database/mysql/user"
	"github.com/gin-gonic/gin"
	"log"
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
		log.Println("2  ", err)

		c.JSON(http.StatusNotFound, "Username does not exist")
		return
	}

	if u.Password != request.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	//t, err := token.CreateToken(u.Username)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}

	c.JSON(http.StatusOK, map[string]string{"jwt": "token123456789"}) //test
}
