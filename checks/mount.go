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
	c.checkType = core.MountPoint{Path: m.Params.Path}
	return nil
}

type mountOptionParams struct {
	Path        string `yaml:"path"` // module name
	MountOption string `yaml:"mountOption"`
}

type mountOptionCheck struct {
	Params mountOptionParams `yaml:"params"` // check parameters
}

//Handle unmarshalling inputs configuration
func (c *check) unmarshalMountOptionCheck(unmarshal func(interface{}) error) error {
	var m mountOptionCheck
	if err := unmarshal(&m); err != nil {
		return err
	}
	c.mountOptionCheck = &m
	c.checkType = core.MountOption{
		Path:        m.Params.Path,
		MountOption: m.Params.MountOption,
	}
	return nil
}
