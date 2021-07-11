package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"github.com/gin-gonic/gin"
)

func info() []byte {
	var result []byte
	var err error
	gin.Default()
	cmd := exec.Command("osqueryi", "--json", "SELECT * FROM system_info")
	if result, err = cmd.Output(); err != nil {
		fmt.Println(err)
	}
	return result
}

func text(w http.ResponseWriter, r *http.Request) {
	var err error
	w.WriteHeader(http.StatusOK)
	result := info()
	_, err = w.Write(result)
	if err != nil{
		fmt.Println(err)
	}
}

func main()  {
	http.HandleFunc("/", text)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Println(err)
	}
}
