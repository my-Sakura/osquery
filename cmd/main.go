package main

import (
	"fmt"
	"log"

	"github.com/my-sakura/osquery/utils"
)

const (
	SQL = "\"SELECT * FROM mounts\""
)

func main() {
	// r := gin.Default()
	output, err := utils.Query(SQL)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(output))
	// controller.New().Register(r.Group(osQueryRouterGroup))

	// log.Println(r.Run("0.0.0.0:8081"))
}
