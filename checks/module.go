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

//Handle unmarshalling inputs configuration
func (c *check) unmarshalKernelModuleCheck(unmarshal func(interface{}) error) error {
	var km kernelModuleCheck
	if err := unmarshal(&km); err != nil {
		return err
	}
	c.kernelModuleCheck = &km
	c.auditType = core.KernelModuleAudit{Name: km.Params.Name}
	c.remediateType = core.KernelModuleRemediate{Name: km.Params.Name}
	return nil
}
