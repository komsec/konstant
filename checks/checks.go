package checks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/komsec/konstant/core"
	yaml "gopkg.in/yaml.v2"
)

type check struct {
	auditType     core.Runner // audit type
	remediateType core.Runner // remediate type

	// Embed Various type of checks - there will be different parameters for different type of checks
	*kernelModuleCheck //Kernel module specific params
}

func (c *check) UnmarshalYAML(unmarshal func(interface{}) error) error {
	t := struct {
		Type string `yaml:"type"`
	}{}
	if err := unmarshal(&t); err != nil {
		return err
	}
	switch {
	case t.Type == "kernelModule":
		c.unmarshalKernelModuleCheck(unmarshal)
	default:
		return fmt.Errorf("Invalid check type - %s", t.Type)
	}
	return nil
}

type checksInput struct {
	Parents []map[string]string `yaml:"parents"`
	Checks  []check             `yaml:"checks"`
}

type checksInputs []checksInput

//TODO: There should be OS version (in case of devices what to do?)
type result struct {
	Parents     []map[string]string `json:"parents"`           //check parent groups
	ID          string              `json:"id"`                // check name
	Description string              `json:"description"`       //check description
	Scored      bool                `json:"scored"`            //scored or not
	Status      core.CheckStatus    `json:"status"`            //check status
	Error       string              `json:"error,omitempty"`   //Error message
	Message     string              `json:"message,omitempty"` // optional message
	Time        string              `json:"datetime"`          // check run time
}

type response struct {
	Success bool     `json:"success"` //success or not - if any checks are failed, it return false
	Error   string   `json:"error"`   //Error message
	Results []result `json:"results"` //Results
}

//getChecks get checks
func getChecks(dir string) ([]checksInput, error) {
	var in []checksInput
	abs, _ := filepath.Abs(dir)
	files, err := filepath.Glob(fmt.Sprintf("%s/*.yaml", abs))
	if err != nil {
		return nil, err
	}
	for i := range files {
		y, err := ioutil.ReadFile(files[i])
		if err != nil {
			return nil, fmt.Errorf("Error reading yaml file %s ", err)
		}
		c := []checksInput{}
		err = yaml.Unmarshal(y, &c)
		if err != nil {
			return nil, fmt.Errorf("Unmarshal: %v", err)
		}
		in = append(in, c...)
	}
	return in, nil
}

//RunAudit run checks
func RunAudit(dir string) (string, error) {
	checkList, err := getChecks(dir)
	if err != nil {
		return "", err
	}
	var o response
	var r []result
	var failed bool
	for i := range checkList {
		for c := range checkList[i].Checks {
			r = append(r, result{
				Parents:     checkList[i].Parents,
				ID:          checkList[i].Checks[c].ID,
				Description: checkList[i].Checks[c].Desc,
				Scored:      checkList[i].Checks[c].Scored,
			})
			// TODO: Detect OS/Device and call appropriate method
			r[c].Time = time.Now().Format(time.RFC1123)
			r[c].Status, r[c].Message, err = checkList[i].Checks[c].auditType.Centos7()
			if err != nil {
				r[c].Error = err.Error()
				failed = true
			}
		}
	}
	o.Results = r
	if failed {
		o.Success = false
		o.Error = "Some checks have been failed"
	} else {
		o.Success = true
	}
	j, err := json.Marshal(o)
	if err != nil {
		return "", fmt.Errorf("Failed jsonify the result: %s", err)
	}
	return string(j), nil
}
