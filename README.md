# Konstant

Konstant is a tool to help on security scanning and benchmarking.

## Add new check types
1. Under core, add file for new check type ([sample check type](core/mount.go))
    1. Define type that describe new check type and implement core.Runner and
    core.checker interfaces
2. Under checks, add new checktype file with ([sample check type](checks/mount.go)):
    1. Define types to define check specific parameters
    2. add an unmarshal function to unmarshal the yaml input for that specific check type
    3. Add it to checkTypes using checkTypes.add() from within init() function.

## Add new checks
1. Create or update yaml files in your config directory ([sample cfg](cfg/))
    defining check type and parameters.
