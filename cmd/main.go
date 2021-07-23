package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

var (
	// osQueryRouterGroup = "/api/v1/osquery"

	errOSNotFound = func(OS string) error { return fmt.Errorf("unsupported operating system: %s", OS) }
)

const (
	SQL = "\"SELECT * FROM system_info\""

	CMDDownloadHomeBrew      = "/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
	CMDDownloadChoco         = "@\"%SystemRoot%\\System32\\WindowsPowerShell\v1.0\\powershell.exe\" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command \"[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))\" && SET \"PATH=%PATH%;%ALLUSERSPROFILE%\\chocolatey\bin\""
	CMDInstallOSQueryByBrew  = "brew install --cask osquery"
	CMDInstallOSQueryByYum   = "yum install osquery"
	CMDInstallOSQueryByChoco = "choco install osquery"
)

func main() {
	// r := gin.Default()
	output, err := Query(SQL)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(output))
	// controller.New().Register(r.Group(osQueryRouterGroup))

	// log.Println(r.Run("0.0.0.0:8081"))
}

func Query(SQL string) (output []byte, err error) {
	if !checkCMDIsExist("osqueryi") {
		err = downloadOSQuery()
		if err != nil {
			return nil, err
		}
	}

	output, err = runCommand("osqueryi --json " + SQL)

	return
}

func downloadOSQuery() error {
	switch runtime.GOOS {
	case "darwin":
		if !checkCMDIsExist("brew") {
			_, err := runCommand(CMDDownloadHomeBrew)
			if err != nil {
				return err
			}
		}

		output, err := runCommand(CMDInstallOSQueryByBrew)
		fmt.Println(string(output))

		return err

	case "linux":
		// TODO check yum exist
		// if !checkCMDIsExist("yum") {
		// }

		output, err := runCommand(CMDInstallOSQueryByYum)
		fmt.Println(string(output))

		return err

	case "windows":
		if !checkCMDIsExist("choco") {
			_, err := runCommand(CMDDownloadChoco)
			return err
		}

		output, err := runCommand(CMDInstallOSQueryByChoco)
		fmt.Println(string(output))

		return err
	}

	return errOSNotFound(runtime.GOOS)
}

func runCommand(cmd string) (output []byte, err error) {
	switch runtime.GOOS {
	case "darwin":
		log.Println("Running Mac cmd:", cmd)
		return exec.Command("/bin/sh", "-c", cmd).Output()

	case "linux":
		log.Println("Running Linux cmd:", cmd)
		return exec.Command("/bin/sh", "-c", cmd).Output()

	case "windows":
		log.Println("Running Windows cmd:", cmd)
		return exec.Command("cmd", "/c", cmd).Output()
	}

	return
}

func checkCMDIsExist(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		if strings.Contains(err.Error(), exec.ErrNotFound.Error()) {
			return false
		} else {
			log.Fatalf("[checkCMDIsExist] unknown error: %s", err)
		}
	}

	return true
}
