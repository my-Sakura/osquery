package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/my-sakura/osquery/utils"
)

type OsqueryController struct{}

func New() *OsqueryController {
	return &OsqueryController{}
}

func (oc *OsqueryController) Register(r gin.IRouter) {
	r.GET("/mounts", oc.mounts)
	r.GET("/system_info", oc.systemInfo)
}

func (oc *OsqueryController) mounts(c *gin.Context) {
	result, err := utils.Query("\"SELECT * FROM mounts\"")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	var resp struct {
		Data []struct {
			Blocks          string `json:"blocks"`
			BlocksAvailable string `json:"blocks_available"`
			BlocksFree      string `json:"blocks_free"`
			BlocksSize      string `json:"blocks_size"`
			Device          string `json:"device"`
			DeviceAlias     string `json:"device_alias"`
			Flags           string `json:"flags"`
			Inodes          string `json:"inodes"`
			InodesFree      string `json:"inodes_free"`
			Path            string `json:"path"`
			Type            string `json:"type"`
		} `json:"data"`
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(result, &resp.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp.Data})
}

func (oc *OsqueryController) systemInfo(c *gin.Context) {
	result, err := utils.Query("\"SELECT * FROM system_info\"")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	var resp struct {
		Data []struct {
			BoardModel       string `json:"board_model"`
			BoardSerial      string `json:"board_serial"`
			BoardVendor      string `json:"board_vendor"`
			BoardVersion     string `json:"board_version"`
			ComputerName     string `json:"computer_name"`
			CPUBrand         string `json:"cpu_brand"`
			CPULogicalCores  string `json:"cpu_logical_cores"`
			CPUMicrocode     string `json:"cpu_microcode"`
			CPUPhysicalCores string `json:"cpu_physical_cores"`
			CPUSubtype       string `json:"cpu_subtype"`
			CPUType          string `json:"cpu_type"`
			HardwareModel    string `json:"hardware_model"`
			HardwareSerial   string `json:"hardware_serial"`
			HardwareVendor   string `json:"hardware_vendor"`
			HardwareVersion  string `json:"hardware_version"`
			Hostname         string `json:"hostname"`
			LocalHostname    string `json:"local_hostname"`
			PhysicalMemory   string `json:"physical_memory"`
			UUID             string `json:"uuid"`
		} `json:"data"`
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(result, &resp.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp.Data})
}
