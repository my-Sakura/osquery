package req

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Result struct {
}

func New() *Result {
	return &Result{}
}
func (a *Result) Method(r gin.IRouter) {
	r.GET("/", a.OsQueryi)
}
func (a *Result) OsQueryi(c *gin.Context) {
	var result []byte
	var err error
	cmd := exec.Command("osqueryi", "--json", "SELECT * FROM time")
	if result, err = cmd.Output(); err != nil {
		fmt.Println(err)
	}
	var data interface{}
	json.Unmarshal(result, &data)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}
