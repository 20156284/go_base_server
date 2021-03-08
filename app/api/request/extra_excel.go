package request

import system "go_base_server/app/model/system"

type ExcelInfo struct {
	FileName string        `json:"fileName"`
	InfoList []system.Menu `json:"infoList"`
}
