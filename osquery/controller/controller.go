package controller

import (
	"fmt"
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
	r.GET("/sql", oc.sql)
	r.GET("/mounts", oc.mounts)
	r.GET("/system_info", oc.systemInfo)
}

func (oc *OsqueryController) sql(c *gin.Context) {
	var req struct {
		SQL string `json:"sql"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}
	fmt.Println(req.SQL)
	SQL := fmt.Sprintf("\"%s\"", req.SQL)
	result, err := utils.Query(SQL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	var resp interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(result, &resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp})
}

func (oc *OsqueryController) mounts(c *gin.Context) {
	result, err := utils.Query("\"SELECT blocks,blocks_available,blocks_free,blocks_size,device,device_alias,flags,inodes,inodes_free,path,type FROM mounts\"")
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
	result, err := utils.Query("\"SELECT board_model,board_serial,board_vendor,board_version,computer_name,cpu_brand,cpu_logical_cores,cpu_microcode,cpu_physical_cores,cpu_subtype,cpu_type,hardware_model,hardware_serial,hardware_vendor,hardware_version,hostname,local_hostname,physical_memory,uuid FROM system_info\"")
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
