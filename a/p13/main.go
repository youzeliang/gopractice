package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

type A struct {
	m map[string]string
}

func (a *A) Init() {

	a.m = make(map[string]string, 0)
}

func (a *A) get(s string) string {
	return a.m[s]
}

func (a *A) set(k, v string) {
	a.m[k] = v
}
