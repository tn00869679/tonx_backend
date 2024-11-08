package response

import (
	"net/http"
	"tonx_backend/internal/structure/error_code"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "ok",
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, code int, message string, _error error_code.ErrorCode) {
	c.JSON(code, gin.H{
		"code":    code,
		"status":  "error",
		"message": message,
		"error":   _error,
	})
}

func ServerError(c *gin.Context) {
	Error(c, http.StatusServiceUnavailable, "server error", error_code.Server)
}

func InputError(c *gin.Context) {
	Error(c, http.StatusBadRequest, "input error", error_code.Input)
}

func NotAcceptableError(c *gin.Context) {
	Error(c, http.StatusNotAcceptable, "not acceptable error", error_code.Acceptable)
}
