package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := uuid.New()

		ctx.Header("trace_id", traceID.String())
		ctx.Next()
	}
}
