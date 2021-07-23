package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/my-sakura/osquery/osquery/controller"
)

var (
	osQueryRouterGroup = "api/v1/osquery"
)

func main() {
	r := gin.Default()

	controller.New().Register(r.Group(osQueryRouterGroup))

	log.Println(r.Run("0.0.0.0:8081"))
}
