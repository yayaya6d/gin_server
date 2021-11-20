package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/yayaya6d/gin_server/api/fib"
	"github.com/yayaya6d/gin_server/middleware"
)

func main() {
	r := gin.New()

	r.Use(middleware.TraceIDMiddleware())

	r.GET("/ping", func(ctx *gin.Context) {
		traceID := ctx.Request.Header.Get("trace_id")
		ctx.JSON(200, gin.H{
			"message":  "pong...",
			"trace id": traceID,
		})
	})

	r.GET("/fib/:num", fib.Fib)

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
