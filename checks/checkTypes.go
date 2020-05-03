package checks

import "github.com/komsec/konstant/core/kernel"

func getKernelModuleCheck(id, moduleName, desc, group string, scored bool) check {
	return check{
		id:          id,
		description: desc,
		scored:      scored,
		types: checkType{
			group:         group,
			auditType:     kernel.ModuleAudit{Name: moduleName},
			remediateType: kernel.ModuleRemediate{Name: moduleName},
		},
	}
}
