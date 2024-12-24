package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{code, msg, data})
}

var (
	CODE       = 20000
	ERROR_CODE = 40001
	MSG        = "success"
)

func OkSuccess(c *gin.Context) {
	Result(CODE, MSG, nil, c)
}

func Ok(data interface{}, c *gin.Context) {
	Result(CODE, MSG, data, c)
}

func Fail(code int, msg string, c *gin.Context) {
	Result(code, msg, map[string]any{}, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR_CODE, msg, map[string]any{}, c)
}

func FailWithError(err error, c *gin.Context) {
	Result(ERROR_CODE, err.Error(), map[string]any{}, c)
}

func FailWithData(code int, msg string, data any, c *gin.Context) {
	Result(code, msg, data, c)
}
