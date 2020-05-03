package kernel

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/komsec/konstant/core"
)

//ModuleAudit define kernel module struct
type ModuleAudit struct {
	Name        string // module name
	core.Runner        //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//ModuleRemediate define kernel module struct
type ModuleRemediate struct {
	Name        string // module name
	core.Runner        //Embed Runner interface to avoid error on any Methods not being implemented by this
}

//Centos7 implement audit for any kernel module based checks
func (a ModuleAudit) Centos7() (status core.CheckStatus, msg string, err error) {
	out, err := exec.Command("modprobe", "-n", "-v", a.Name).Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(err.Stderr), fmt.Sprintf("Module %s not found", a.Name)) {
				return core.StatusNA, "Module does not found", nil
			}
			return core.StatusFail, msg, fmt.Errorf("Failed executing modprobe. Error: %s", err.Error())
		}
		if err, ok := err.(*exec.Error); ok {
			return core.StatusFail, msg, fmt.Errorf("Failed executing modprobe. Error: %s", err.Error())
		}
	}
	if strings.Contains(string(out), "install /bin/true") {
		// checking lsmod
		out, err := exec.Command("lsmod").Output()
		if strings.Contains(string(out), fmt.Sprintf("%s ", a.Name)) {
			return core.StatusFail, msg, fmt.Errorf("Module is loaded")
		}
		if err != nil {
			return core.StatusFail, msg, err
		}
	} else {
		return core.StatusFail, msg, fmt.Errorf("Output: %s, Error: Unexpected output", strings.Trim(string(out), "\n"))
	}
	return core.StatusPass, msg, nil
}

//Centos7 implement audit for any kernel module based checks
func (r ModuleRemediate) Centos7() (status core.CheckStatus, msg string, err error) {
	fmt.Println("Executing centos7 Remediate for", r.Name)
	return core.StatusPass, msg, err
}
