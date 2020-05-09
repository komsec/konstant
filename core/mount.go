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

//MountPoint define kernel module struct
//It must implement both core.Runner and core.checker interfaces
type MountPoint struct {
	Path   string // Mount Path
	Runner        //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//Audit perform audit checks
func (m MountPoint) Audit() (status CheckStatus, msg string, err error) {
	return runCheck(m, checkAudit)
}

//Remediate perform audit checks
func (m MountPoint) Remediate() (status CheckStatus, msg string, err error) {
	return runCheck(m, checkRemdiate)
}

func (m MountPoint) auditCentos7() (status CheckStatus, msg string, err error) {
	mc := mountCheck{
		path: m.Path,
	}
	return mc.nix()
}

func (m MountPoint) remediateCentos7() (status CheckStatus, msg string, err error) {
	return status, msg, errors.New("Not implemented")
}

//MountOption define kernel module struct
//It must implement both core.Runner and core.checker interfaces
type MountOption struct {
	Path        string // Mount Path
	MountOption string //Mount option
}

//Audit perform audit checks
func (mo MountOption) Audit() (status CheckStatus, msg string, err error) {
	return runCheck(mo, checkAudit)
}

//Remediate perform audit checks
func (mo MountOption) Remediate() (status CheckStatus, msg string, err error) {
	return runCheck(mo, checkRemdiate)
}

//Centos7 implement audit for any kernel module based checks
func (mo MountOption) auditCentos7() (status CheckStatus, msg string, err error) {
	m := mountCheck{
		mo.Path,
		mo.MountOption,
	}
	status, msg, err = m.nix()
	if err != nil {
		if err.Error() == "Search pattern does not matched" {
			return status, msg, fmt.Errorf("Mount Option %s is not set for path %s", mo.MountOption, mo.Path)
		}
	}
	return status, msg, err
}

func (mo MountOption) remediateCentos7() (status CheckStatus, msg string, err error) {
	return status, msg, errors.New("Not implemented")
}
