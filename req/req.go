package req

import (
	"fmt"
	"net/http"
	"os/exec"
)

func Systeminfo() []byte {
	var result []byte
	var err error
	cmd := exec.Command("osqueryi", "--json", "SELECT * FROM time")
	if result, err = cmd.Output(); err != nil {
		fmt.Println(err)
	}

	return result
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	result := Systeminfo()
	_, err := w.Write(result)
	if err != nil {
		fmt.Println(err)
	}

}
