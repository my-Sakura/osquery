package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

var (
	errOSNotFound = func(OS string) error { return fmt.Errorf("unsupported operating system: %s", OS) }
)

const (
	cmdDownloadHomeBrew      = "/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
	cmdDownloadChoco         = "@\"%SystemRoot%\\System32\\WindowsPowerShell\v1.0\\powershell.exe\" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command \"[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))\" && SET \"PATH=%PATH%;%ALLUSERSPROFILE%\\chocolatey\bin\""
	cmdInstallOSQueryByBrew  = "brew install --cask osquery"
	cmdInstallOSQueryByYum   = "yum install osquery"
	cmdInstallOSQueryByChoco = "choco install osquery"
)

func init() {
	if !checkCMDIsExist("osqueryi") {
		err := downloadOSQuery()
		if err != nil {
			panic(err)
		}
	}
}

func Query(SQL string) (output []byte, err error) {
	return runCommand("osqueryi --json " + SQL)
}

func downloadOSQuery() error {
	switch runtime.GOOS {
	case "darwin":
		if !checkCMDIsExist("brew") {
			_, err := runCommand(cmdDownloadHomeBrew)
			if err != nil {
				return err
			}
		}

		output, err := runCommand(cmdInstallOSQueryByBrew)
		fmt.Println(string(output))

		return err

	case "linux":
		// TODO check yum exist
		// if !checkCMDIsExist("yum") {
		// }

		output, err := runCommand(cmdInstallOSQueryByYum)
		fmt.Println(string(output))

		return err

	case "windows":
		if !checkCMDIsExist("choco") {
			_, err := runCommand(cmdDownloadChoco)
			return err
		}

		output, err := runCommand(cmdInstallOSQueryByChoco)
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
