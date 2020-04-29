package kernel

import (
	"fmt"
	"os/exec"
	"strings"
)

//ModuleAudit define kernel module struct
type ModuleAudit struct {
	Name string // module name
}

//ModuleRemediate define kernel module struct
type ModuleRemediate struct {
	Name string // module name
}

//Centos7 implement audit for any kernel module based checks
func (a ModuleAudit) Centos7() (msg string, err error) {
	out, err := exec.Command("modprobe", "-n", "-v", a.Name).Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(err.Stderr), fmt.Sprintf("Module %s not found", a.Name)) {
				return "Module does not found", nil
			}
			return msg, fmt.Errorf("Failed executing modprobe. Error: %s", err.Error())
		}
		if err, ok := err.(*exec.Error); ok {
			return msg, fmt.Errorf("Failed executing modprobe. Error: %s", err.Error())
		}
	}
	if strings.Contains(string(out), "install /bin/true") {
		// checking lsmod
		out, err := exec.Command("lsmod").Output()
		if strings.Contains(string(out), fmt.Sprintf("%s ", a.Name)) {
			return msg, fmt.Errorf("Module is loaded")
		}
		if err != nil {
			return msg, err
		}
	} else {
		return msg, fmt.Errorf("Output: %s, Error: Unexpected output", strings.Trim(string(out), "\n"))
	}
	return msg, nil
}

//Centos7 implement audit for any kernel module based checks
func (r ModuleRemediate) Centos7() (msg string, err error) {
	fmt.Println("Executing centos7 Remediate for", r.Name)
	return msg, err
}
