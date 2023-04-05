package authority

import (
	"book/pkg/response"
	"github.com/gin-gonic/gin"
)

func GetRoute(c *gin.Context) {

	result := Admin()

	response.Data(c, result)

}
