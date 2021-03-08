package request

import model "go_base_server/app/model/extra"

type SearchWorkflowProcess struct {
	model.WorkflowProcess
	PageInfo
}
