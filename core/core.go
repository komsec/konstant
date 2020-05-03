package core

// Runner module interface that define how a konstant module look like
type Runner interface {
	Centos7() (CheckStatus, string, error)
	Centos8() (CheckStatus, string, error)
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
