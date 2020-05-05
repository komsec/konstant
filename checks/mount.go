package checks

import (
	"github.com/komsec/konstant/core"
)

type mountPointParams struct {
	Path string `yaml:"path"` // module name
}

type mountPointCheck struct {
	Params mountPointParams `yaml:"params"` // check parameters
}

//Handle unmarshalling inputs configuration
func (c *check) unmarshalMountPointCheck(unmarshal func(interface{}) error) error {
	var m mountPointCheck
	if err := unmarshal(&m); err != nil {
		return err
	}
	c.mountPointCheck = &m
	c.auditType = core.MountPointAudit{Path: m.Params.Path}
	c.remediateType = core.MountPointRemediate{Path: m.Params.Path}
	return nil
}
