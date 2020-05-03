package core

// Runner module interface that define how a konstant module look like
type Runner interface {
	Centos7() (string, error)
	Centos8() (string, error)
}
