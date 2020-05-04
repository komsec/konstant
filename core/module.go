package core

import (
	"fmt"
	"os/exec"
	"strings"
)

//KernelModuleAudit define kernel module struct
type KernelModuleAudit struct {
	Name        string // module name
	Runner        //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//KernelModuleRemediate define kernel module struct
type KernelModuleRemediate struct {
	Name        string // module name
	Runner        //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//Centos7 implement audit for any kernel module based checks
func (a KernelModuleAudit) Centos7() (status CheckStatus, msg string, err error) {
	out, err := exec.Command("modprobe", "-n", "-v", a.Name).Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(err.Stderr), fmt.Sprintf("Module %s not found", a.Name)) {
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
		if strings.Contains(string(out), fmt.Sprintf("%s ", a.Name)) {
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

//Centos7 implement audit for any kernel module based checks
func (r KernelModuleRemediate) Centos7() (status CheckStatus, msg string, err error) {
	fmt.Println("Executing centos7 Remediate for", r.Name)
	return StatusPass, msg, err
}
