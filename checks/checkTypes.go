package checks

import "github.com/komsec/konstant/core/kernel"

func getKernelModuleCheck(name, desc, group string, scored bool) check {
	return check{
		name:        name,
		description: desc,
		scored:      scored,
		types: checkType{
			group:         group,
			auditType:     kernel.ModuleAudit{Name: name},
			remediateType: kernel.ModuleRemediate{Name: name},
		},
	}
}
