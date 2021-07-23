package req

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

type Table struct {
	Datetime      time.Time `json:"datetime"`
	Day           string    `json:"day"`
	Hour          string    `json:"hour"`
	Iso8601       time.Time `json:"iso_8601"`
	LocalTime     string    `json:"local_time"`
	LocalTimezone string    `json:"local_timezone"`
	Minutes       string    `json:"minutes"`
	Month         string    `json:"month"`
	Seconds       string    `json:"seconds"`
	Timestamp     string    `json:"timestamp"`
	Timezone      string    `json:"timezone"`
	UnixTime      string    `json:"unix_time"`
	Weekday       string    `json:"weekday"`
	WinTimestamp  string    `json:"win_timestamp"`
	Year          string    `json:"year"`
}

type Result struct {
}

func New() *Result {
	return &Result{}
}
func (a *Result) Method(r gin.IRouter) {
	r.GET("/", a.OsQueryi)
}
func (a *Result) OsQueryi(c *gin.Context) {
	cmd := exec.Command("osqueryi", "--json",
		"SELECT weekday, year, month, day, hour, minutes, seconds, timezone, local_time, timestamp, datetime, isl_8601 FROM time")
	result, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	var data Table
	json.Unmarshal(result, &data)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}
