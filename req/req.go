package req

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func Systeminfo() []byte {
	var result []byte
	var err error
	cmd := exec.Command("osqueryi", "--json", "SELECT * FROM system_info")
	if result, err = cmd.Output(); err != nil {
		fmt.Println(err)
	}

	return result
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	result := Systeminfo()
	var a interface{}
	json.Unmarshal(result, &a)
	b, err := json.Marshal(&a)
	if err != nil {
		log.Println(nil)
	}
	w.Write(b)

}
