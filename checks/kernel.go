package checks

import (
	"github.com/komsec/konstant/core/kernel"
)

type kernelModuleParams struct {
	Name string `yaml:"name"` // module name
}

type kernelModuleCheck struct {
	ID     string             `yaml:"id"`     // check id
	Desc   string             `yaml:"desc"`   // description of the check
	Scored bool               `yaml:"scored"` //scored or not
	Type   string             `yaml:"type"`   // check type
	Params kernelModuleParams `yaml:"params"` // check parameters
}

//Handle unmarshalling inputs configuration
func (c *check) unmarshalKernelModuleCheck(unmarshal func(interface{}) error) error {
	var km kernelModuleCheck
	if err := unmarshal(&km); err != nil {
		return err
	}
	c.kernelModuleCheck = &km
	c.auditType = kernel.ModuleAudit{Name: c.Params.Name}
	c.remediateType = kernel.ModuleRemediate{Name: c.Params.Name}
	return nil
}
