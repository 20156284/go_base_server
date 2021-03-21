package request

import model "go_base_server/app/model/workflow"

type SearchWorkflowProcess struct {
	model.WorkflowProcess
	PageInfo
}
