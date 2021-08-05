package internal

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
	cmdDownloadHomeBrew     = "/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
	cmdInstallOSQueryByBrew = "brew install --cask osquery"

	cmdIsUbuntuOrCentos     = "cat /etc/*-release | grep NAME"
	cmdInstallOSQueryByYum1 = "curl -L https://pkg.osquery.io/rpm/GPG | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-osquery"
	cmdInstallOSQueryByYum2 = "yum-config-manager --add-repo https://pkg.osquery.io/rpm/osquery-s3-rpm.repo"
	cmdInstallOSQueryByYum3 = "yum-config-manager --enable osquery-s3-rpm-repo"
	cmdInstallOSQueryByYum  = "yum install -y osquery"

	cmdInstallOSQueryByAptGet1 = "apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 1484120AC4E9F8A1A577AEEE97A80C63C9D8B80B"
	cmdInstallOSQueryByAptGet2 = "add-apt-repository 'deb [arch=amd64] https://pkg.osquery.io/deb deb main'"
	cmdInstallOSQueryByAptGet3 = "apt-get update"
	cmdInstallOSQueryByAptGet  = "apt-get install osquery"

	cmdDownloadChoco         = "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"
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

func Tables() (output []byte, err error) {
	return runCommand("osqueryi .table")
}

func Table(tableName string) (output []byte, err error) {
	cmd := fmt.Sprintf("osqueryi \"SELECT * FROM %s\"", tableName)
	return runCommand(cmd)
}

func downloadOSQuery() error {
	switch runtime.GOOS {
	case "darwin":
		output, err := installOSQueryInMacDarwin()
		if err != nil {
			return fmt.Errorf("Error: installOSQueryInMacDarwin error: %s", err)
		}
		fmt.Print(string(output))

	case "linux":
		var (
			output []byte
			err    error
		)

		// check is centos or ubuntu
		output, err = runCommand(cmdIsUbuntuOrCentos)
		if err != nil {
			fmt.Printf("Error: The Command [%s] is err: %s", cmdIsUbuntuOrCentos, err)
			return nil
		}

		if strings.Contains(string(output), "Ubuntu") {
			output, err = installOSQueryInDebianLinux()
			if err != nil {
				return fmt.Errorf("Error: installOSQueryInDebianLinux error: %s", err)
			}
		} else if strings.Contains(string(output), "CentOS") {
			output, err = installOSQueryInRPMLinux()
			if err != nil {
				return fmt.Errorf("Error: installOSQueryInRPMLinux error: %s", err)
			}
		}

		fmt.Print(string(output))

	case "windows":
		output, err := installOSQueryInWindows()
		if err != nil {
			return fmt.Errorf("Error: installOSQueryInWindows error: %s", err)
		}
		fmt.Print(string(output))
	}

	return errOSNotFound(runtime.GOOS)
}

func installOSQueryInDebianLinux() (output []byte, err error) {
	// TODO check apt-get exist
	output, err = runCommand(cmdInstallOSQueryByAptGet1)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(output))

	output, err = runCommand(cmdInstallOSQueryByAptGet2)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(output))

	output, err = runCommand(cmdInstallOSQueryByAptGet3)
	if err != nil {
		return nil, err
	}
	fmt.Print(string(output))

	return runCommand(cmdInstallOSQueryByAptGet)
}

func installOSQueryInRPMLinux() (output []byte, err error) {
	// TODO check yum exist
	output, err = runCommand(cmdInstallOSQueryByYum1)
	if err != nil {
		return output, err
	}
	fmt.Print(string(output))
	output, err = runCommand(cmdInstallOSQueryByYum2)
	if err != nil {
		return output, err
	}
	fmt.Print(string(output))
	output, err = runCommand(cmdInstallOSQueryByYum3)
	if err != nil {
		return output, err
	}
	fmt.Print(string(output))

	return runCommand(cmdInstallOSQueryByYum)

}

func installOSQueryInMacDarwin() (output []byte, err error) {
	if !checkCMDIsExist("brew") {
		_, err := runCommand(cmdDownloadHomeBrew)
		if err != nil {
			return nil, err
		}
	}

	return runCommand(cmdInstallOSQueryByBrew)
}

func installOSQueryInWindows() (output []byte, err error) {
	if !checkCMDIsExist("choco") {
		_, err := runCommand(cmdDownloadChoco)
		return nil, err
	}

	return runCommand(cmdInstallOSQueryByChoco)
}

func runCommand(cmd string) (output []byte, err error) {
	switch runtime.GOOS {
	case "darwin":
		// log.Println("Running Mac cmd:", cmd)
		return exec.Command("/bin/sh", "-c", cmd).Output()

	case "linux":
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
