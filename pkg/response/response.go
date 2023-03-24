package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResponseWithPage struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
}
type MsgResponse struct {
	Result  bool     `json:"result"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, MsgResponse{Result: true, Message: "success"})
}

func Error(c *gin.Context, err string) {
	c.JSON(http.StatusInternalServerError, MsgResponse{Result: false, Message: err})
}

func Data(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, Response{Result: true, Message: "success", Data: v})
}
func DataWtihPage(c *gin.Context, v interface{}, total int) {
	c.JSON(http.StatusOK, ResponseWithPage{Result: true, Message: "success", Data: v, Total: total})
}

func BadRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, MsgResponse{Result: false, Message: err})
}
