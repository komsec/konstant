package checks

import (
	"encoding/json"
	"fmt"

	"github.com/komsec/konstant/core"
)

type checkType struct {
	group         string // check group like kernel
	auditType     core.Runner
	remediateType core.Runner
}

type check struct {
	id          string            // check ID
	params      map[string]string // check parameters
	description string            // description
	scored      bool              // scored or not
	types       checkType         //check type
}

//TODO: name should be the check number, there should OS version (in case of devices what to do?)
type result struct {
	Group       string `json:"group"`             //check group e.g filesystem
	ID          string `json:"id"`                // check name
	Description string `json:"description"`       //check description
	Scored      bool   `json:"scored"`            //scored or not
	Success     bool   `json:"success"`           //success or not
	Error       string `json:"error,omitempty"`   //Error message
	Message     string `json:"message,omitempty"` // optional message
}

type results []result

// Set a list of checks
var checkList []check = fsCheckList

//RunAudit run checks
func RunAudit() (string, error) {
	var r results
	var err error
	var failed bool
	for i := range checkList {
		r = append(r, result{
			Group:       checkList[i].types.group,
			ID:          checkList[i].id,
			Description: checkList[i].description,
			Scored:      checkList[i].scored,
			Success:     true,
		})
		//Detect OS/Device
		//Call appropriate method
		r[i].Message, err = checkList[i].types.auditType.Centos7()
		if err != nil {
			r[i].Success = false
			r[i].Error = err.Error()
			failed = true
		}
	}
	j, err := json.Marshal(r)
	if err != nil {
		return "", fmt.Errorf("Failed jsonify the result: %s", err)
	}
	if failed {
		return string(j), fmt.Errorf("Some checks have been failed")
	}
	return string(j), nil
}
