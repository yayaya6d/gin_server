package fib

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Fib(ctx *gin.Context) {
	n, err := strconv.Atoi(ctx.Param("num"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "wrong param",
		})
	}

	mem := fib(n)

	ctx.JSON(200, gin.H{
		"ans": mem,
	})
}

func fib(n int) []int {
	ret := []int{}

	for i := 0; i <= n; i++ {
		if i < 2 {
			ret = append(ret, 1)
		} else {
			ret = append(ret, ret[i-1]+ret[i-2])
		}
	}

	return ret
}
