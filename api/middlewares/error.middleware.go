package middlewares

import (
	"carveo/utils"

	"github.com/gin-gonic/gin"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.ErrorResponse(ctx, 500, utils.ErrInternalServer, err)
			}
		}()
		ctx.Next()
	}
}
