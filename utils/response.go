package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 封装返回的代码
func Response(ctx *gin.Context, httpStatus int, code int, msg string, data gin.H) {
	ctx.JSON(httpStatus, gin.H{
		"code":    code,
		"data":    data,
		"message": msg,
	})
}

// Success 成功时返回的代码
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, msg, data)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, msg, data)
}
