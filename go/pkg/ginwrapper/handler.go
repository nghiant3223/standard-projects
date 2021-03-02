package ginwrapper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/standard-project/internal/todo/apperror"
)

type Response struct {
	Data  interface{}
	Error error
}

type HandlerFunc func(ctx *gin.Context) *Response

func Wrap(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := handler(ctx)
		if resp.Error != nil {
			var appErr apperror.Error
			if !errors.As(resp.Error, &appErr) {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(appErr.StatusCode, resp)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
