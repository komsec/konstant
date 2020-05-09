package checks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/komsec/konstant/core"
	yaml "gopkg.in/yaml.v2"
)

type general struct {
	ID     string // check id
	Desc   string // description of the check
	Scored bool   //scored or not
}

type check struct {
	checkType core.Runner // audit type
	general               //general params
	// Params may be in different types based on actual check params,
	//  which will be defined check specific definitions
	Params interface{}
}

var checkTypes = make(map[string]func(func(interface{}) error) (interface{}, core.Runner, error))

func (c *check) UnmarshalYAML(unmarshal func(interface{}) error) error {

	t := struct {
		Type string `yaml:"type"`
	}{}
	if err := unmarshal(&t); err != nil {
		return err
	}

	d := general{}
	if err := unmarshal(&d); err != nil {
		return err
	}
	c.ID = d.ID
	c.Desc = d.Desc
	c.Scored = d.Scored
	var err error
	if ct := checkTypes[t.Type]; ct != nil {
		c.Params, c.checkType, err = ct(unmarshal)
		if err != nil {
			return err
		}
	} else {
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
	files, err := filepath.Glob(fmt.Sprintf("%s/*.yml", abs))
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
func RunAudit(dir string) (res string, ok bool, err error) {
	//Detect Operating system and fail in case of unsupported operating system or version
	_, _, err = core.GetOSNameVersion()
	if err != nil {
		return res, ok, errors.New("Unsupported Operating system or version")
	}
	checkList, err := getChecks(dir)
	if err != nil {
		return res, ok, err
	}
	var o response
	var r []result
	var failed bool
	for i := range checkList {
		for c := range checkList[i].Checks {
			t := time.Now().Format(time.RFC1123)
			s, msg, err := checkList[i].Checks[c].checkType.Audit()
			var e string
			if err != nil {
				e = err.Error()
				failed = true
			}
			r = append(r, result{
				Parents:     checkList[i].Parents,
				ID:          checkList[i].Checks[c].ID,
				Description: checkList[i].Checks[c].Desc,
				Scored:      checkList[i].Checks[c].Scored,
				Time:        t,
				Status:      s,
				Message:     msg,
				Error:       e,
			})
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
		return res, ok, fmt.Errorf("Failed jsonify the result: %s", err)
	}

	return string(j), o.Success, nil
}
