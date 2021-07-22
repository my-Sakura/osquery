package main

import (
	"github.com/shuaiqidechuan/req"

	"github.com/gin-gonic/gin"
)

func main() {
	a := "/"
	r := gin.Default()
	Re := handler.New()
	Re.Method(r.Group(a))
	r.Run(":8080") //监听 localhost:8080
}
