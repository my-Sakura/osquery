package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type OsqueryController struct{}

func New() *OsqueryController {
	return &OsqueryController{}
}

func (oc *OsqueryController) Register(r gin.IRouter) {
	r.GET("/mounts/get", oc.mounts)
}

func (oc *OsqueryController) mounts(c *gin.Context) {
	var result []byte
	var err error
	cmd := exec.Command("osqueryi", "--json", "SELECT * FROM mounts;")
	if result, err = cmd.Output(); err != nil {
		fmt.Println(err)
	}

	var data interface{}
	json.Unmarshal(result, &data)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}
