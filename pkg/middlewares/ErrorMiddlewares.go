package middlewares

import (
	. "davinci/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//context.JSON(500, gin.H{"error": err})
				ret := ResultWrapper(context)(nil, err.(string))(Error)
				context.JSON(500, ret)
			}
		}()
		context.Next()
	}
}
