package checks

import (
	"github.com/komsec/konstant/core"
)

type kernelModuleParams struct {
	Name string `yaml:"name"` // module name
}

type kernelModuleCheck struct {
	Params kernelModuleParams `yaml:"params"` // check parameters
}

func init() {
	//Add unmarshal function to checkTypes map, so Unmarshal can call it
	checkTypes["kernelModule"] = unmarshalKernelModuleCheck
}

//Unmarshal params
func unmarshalKernelModuleCheck(unmarshal func(interface{}) error) (param interface{}, chk core.Runner, err error) {
	var km kernelModuleCheck
	if err := unmarshal(&km); err != nil {
		return param, chk, err
	}
	return km.Params, core.KernelModule{Name: km.Params.Name},  nil
}
