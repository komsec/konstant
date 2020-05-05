package core

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

//mountCheck is to perform mount related checks
type mountCheck struct {
	path   string // Mount Path
	search string //Search string in mount output
}

func (c mountCheck) nix() (status CheckStatus, msg string, err error) {
	out, err := exec.Command("mount").Output()
	if err != nil {
		return StatusFail, msg, fmt.Errorf("Failed executing mount. Error: %s", err.Error())
	}
	s := string(out)
	if strings.Contains(s, fmt.Sprintf(" %s ", c.path)) {
		if c.search != "" {
			m, err := regexp.MatchString(c.search, s)
			if err != nil {
				return StatusFail, msg, errors.New("Invalid regex format in Search string")
			}
			//Search pattern not matched
			if !m {
				return StatusFail, msg, fmt.Errorf("Search pattern does not matched")
			}
		}
	} else {
		return StatusFail, msg, fmt.Errorf("Could not find matching mountpoint path %s", c.path)
	}
	return StatusPass, msg, nil
}

//MountPointAudit define kernel module struct
type MountPointAudit struct {
	Path   string // Mount Path
	Runner        //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//MountPointRemediate define kernel module struct
type MountPointRemediate struct {
	Path   string // Mount Path
	Runner        //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//Centos7 implement audit for any kernel module based checks
func (a MountPointAudit) Centos7() (status CheckStatus, msg string, err error) {
	m := mountCheck{
		path: a.Path,
	}
	return m.nix()
}

//MountOptionAudit define kernel module struct
type MountOptionAudit struct {
	Path        string // Mount Path
	MountOption string //Mount option
	Runner             //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//MountOptionRemediate define kernel module struct
type MountOptionRemediate struct {
	Path        string // Mount Path
	MountOption string //Mount option
	Runner             //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//Centos7 implement audit for any kernel module based checks
func (a MountOptionAudit) Centos7() (status CheckStatus, msg string, err error) {
	m := mountCheck{
		a.Path,
		a.MountOption,
	}
	status, msg, err = m.nix()
	if err != nil {
		if err.Error() == "Search pattern does not matched" {
			return status, msg, fmt.Errorf("Mount Option %s is not set for path %s", a.MountOption, a.Path)
		}
	}
	return status, msg, err
}
