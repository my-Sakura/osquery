package controller

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type OsqueryController struct{}

func New() *OsqueryController {
	return &OsqueryController{}
}

func (oc *OsqueryController) Register(r gin.IRouter) {
	// r.GET("/mounts", oc.mounts)
	// r.GET("/system_info", oc.systemInfo)
	r.GET("/time", oc.time)
}

func (oc *OsqueryController) mounts(c *gin.Context) {
	cmd := exec.Command("osqueryi", "--json", "SELECT * FROM mounts;")
	result, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": string(result)})
}

func (oc *OsqueryController) systemInfo(c *gin.Context) {
	cmd := exec.Command("osqueryi", "--json", "SELECT * FROM system_info")
	result, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": string(result)})
}
func (oc *OsqueryController) time(c *gin.Context) {
	cmd := exec.Command("osqueryi", "--json",
		"SELECT weekday, year, month, day, hour, minutes, seconds, timezone, local_time, timestamp, datetime, isl_8601 FROM time")
	result, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": string(result)})
}
