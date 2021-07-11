package main

import (
	"github.com/gin-gonic/gin"
	osController "github.com/my-sakura/osquery/pkgs/osquery/controller"
)

var (
	osQueryRouterGroup = "/api/v1/osquery"
)

func main() {
	r := gin.Default()

	osCon := osController.New()
	osCon.Register(r.Group(osQueryRouterGroup))

	r.Run("0.0.0.0:8081")
}
