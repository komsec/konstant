package core

import (
	"fmt"
	"os/exec"
	"strings"
)

// KernelModule define kernel module struct
//It must implement both core.Runner and core.checker interfaces
type KernelModule struct {
	Name string // module name
}

//Audit implements core.Runner interface
func (m KernelModule) Audit() (status CheckStatus, msg string, err error) {
	return runCheck(m, "audit")
}

//Remediate implements core.Runner interface
func (m KernelModule) Remediate() (status CheckStatus, msg string, err error) {
	return runCheck(m, "remediate")
}

//auditCentos7 implements core.checker interface
func (m KernelModule) auditCentos7() (status CheckStatus, msg string, err error) {
	out, err := exec.Command("modprobe", "-n", "-v", m.Name).Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(err.Stderr), fmt.Sprintf("Module %s not found", m.Name)) {
				return StatusNA, "Module does not found", nil
			}
			return StatusFail, msg, fmt.Errorf("Failed executing modprobe. Error: %s", err.Error())
		}
		if err, ok := err.(*exec.Error); ok {
			return StatusFail, msg, fmt.Errorf("Failed executing modprobe. Error: %s", err.Error())
		}
	}
	if strings.Contains(string(out), "install /bin/true") {
		// checking lsmod
		out, err := exec.Command("lsmod").Output()
		if strings.Contains(string(out), fmt.Sprintf("%s ", m.Name)) {
			return StatusFail, msg, fmt.Errorf("Module is loaded")
		}
		if err != nil {
			return StatusFail, msg, err
		}
	} else {
		return StatusFail, msg, fmt.Errorf("Output: %s, Error: Unexpected output", strings.Trim(string(out), "\n"))
	}
	return StatusPass, msg, nil
}

//remediateCentos7 implements core.checker interface
func (m KernelModule) remediateCentos7() (status CheckStatus, msg string, err error) {
	fmt.Println("Executing centos7 Remediate for", m.Name)
	return StatusPass, msg, err
}
