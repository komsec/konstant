package core

import (
	"errors"
	"os/exec"
	"regexp"
	"runtime"
)

func hasRPMCmd() bool {
	_, err := exec.LookPath("rpm")
	if err != nil {
		return false
	}
	return true
}

func getCentOS() (bool, string, error) {
	out, err := exec.Command("rpm", "-q", "centos-release").Output()
	if err != nil {
		return false, "", nil
	}
	re := regexp.MustCompile(`centos-release-(\d+)`)
	m := re.FindStringSubmatch(string(out))
	if m[1] == "" {
		return true, "", errors.New("Could not find centos major version")
	}
	return true, m[1], nil
}

//GetOSNameVersion get operating system name and version
func GetOSNameVersion() (os string, version string, err error) {
	var ok bool
	if runtime.GOOS == "linux" {
		//CentOS
		ok, version, err = getCentOS()
		if err != nil {
			return os, version, err
		}
		if ok {
			return "CentOS", version, err
		}

	}
	return "", "", errors.New("Unknown Operating system or version")
}
