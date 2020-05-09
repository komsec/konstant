package core

import "errors"

// Runner module interface that define how a konstant module look like
type Runner interface {
	Audit() (CheckStatus, string, error)
	Remediate() (CheckStatus, string, error)
}

type checker interface {
	auditCentos7() (CheckStatus, string, error)
	remediateCentos7() (CheckStatus, string, error)
}

// CheckStatus check status
type CheckStatus string

const (
	//StatusPass pass status
	StatusPass CheckStatus = "Pass"

	//StatusFail pass status
	StatusFail CheckStatus = "Fail"

	//StatusNA not applicable status
	StatusNA CheckStatus = "NA"
)

type checkOp string

const (
	checkAudit    checkOp = "audit"
	checkRemdiate checkOp = "remediate"
)

//Audit implement audit method
func runCheck(c checker, op checkOp) (status CheckStatus, msg string, err error) {
	//error is ignore because it is evaluated as pre-check
	os, ver, _ := GetOSNameVersion()

	//Add platforms/OS versions specific methods here
	checkFuncs := map[string]map[checkOp]func() (CheckStatus, string, error) {
		"CentOS7": {
			"audit":     c.auditCentos7,
			"remediate": c.remediateCentos7,
		},
	}

	if ct := checkFuncs[os+ver]; ct[checkAudit] != nil {
		return ct[checkAudit]()
	}

	return status, msg, errors.New("Unsupported Operating system")
}
