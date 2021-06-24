package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("trace_id") == "" {
			traceID := uuid.New()
			ctx.Request.Header.Set("trace_id", traceID.String())
			fmt.Println("new trace id =", traceID)
		}

		ctx.Next()
	}
}
