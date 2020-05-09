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

func init() {
	//Add unmarshal functions to checkTypes map, so Unmarshal can call it
	checkTypes["mountPoint"] = unmarshalMountPointCheck
	checkTypes["mountOption"] = unmarshalMountOptionCheck
}

//Handle unmarshalling inputs configuration
func unmarshalMountPointCheck(unmarshal func(interface{}) error) (param interface{}, chk core.Runner, err error) {
	var m mountPointCheck
	if err := unmarshal(&m); err != nil {
		return param, chk, err
	}
	return m.Params, core.MountPoint{Path: m.Params.Path}, nil
}

type mountOptionParams struct {
	Path        string `yaml:"path"` // module name
	MountOption string `yaml:"mountOption"`
}

type mountOptionCheck struct {
	Params mountOptionParams `yaml:"params"` // check parameters
}

//Handle unmarshalling inputs configuration
func unmarshalMountOptionCheck(unmarshal func(interface{}) error) (param interface{}, chk core.Runner, err error) {
	var m mountOptionCheck
	if err := unmarshal(&m); err != nil {
		return param, chk, err
	}
	ct := core.MountOption{
		Path:        m.Params.Path,
		MountOption: m.Params.MountOption,
	}
	return m.Params, ct, nil
}
