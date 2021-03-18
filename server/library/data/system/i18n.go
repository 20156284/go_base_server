package data

import (
	"go_base_server/library/global"
)

var I18nHash map[string]string

func Init() {
	if global.Config.System.Language == "en" {
		I18nHash = GetEnI18n()
	} else {
		I18nHash = GetChI18n()
	}
}

func GetEnI18n() map[string]string {
	return map[string]string{
		"SuperAdmin": "SuperAdmin",
		"OtherUser":  "OtherUser",

		"UserLogin":                     "UserLogin",
		"UserRegister":                  "UserRegister",
		"CreateApi":                     "CreateAPI",
		"GetApiList":                    "GetAPIList",
		"GetApiDetail":                  "GetAPIDetail",
		"DeleteApi":                     "DeleteAPI",
		"UpdateApi":                     "UpdateAPI",
		"GetAllApis":                    "GetAllAPIs",
		"CreateAuthority":               "CreateAuthority",
		"DeleteAuthority":               "DeleteAuthority",
		"GetAuthorityList":              "GetAuthorityList",
		"GetMenu":                       "GetMenu",
		"GetMenuList":                   "GetMenuList(Pagination)",
		"AddBaseMenu":                   "AddBaseMenu",
		"GetBaseMenuTree":               "GetUserDynamicMenu",
		"AddMenuAuthority":              "AddMenuAuthorityConnection",
		"GetMenuAuthority":              "GetMenuAuthorityConnection",
		"DeleteBaseMenu":                "DeleteBaseMenu",
		"UpdateBaseMenu":                "UpdateBaseMenu",
		"GetBaseMenuById":               "GetBaseMenuById",
		"ChangePassword":                "ChangePassword",
		"GetUserList":                   "GetUserList",
		"SetUserAuthority":              "SetUserAuthority",
		"UploadFile":                    "FileUploadDemo",
		"GetFileList":                   "GetUploadFileList",
		"UpdateCasbin":                  "UpdateUserAPIAuthority",
		"GetPolicyPathByAuthorityId":    "GetPolicyPathByAuthorityId",
		"DeleteFile":                    "DeleteFile",
		"JsonInBlacklist":               "JWTBlock(LogOut)",
		"SetDataAuthority":              "SetUserDataAuthority",
		"GetSystemConfig":               "GetSystemConfig",
		"SetSystemConfig":               "SetSystemConfig",
		"CreateCustomer":                "CreateCustomer",
		"UpdateCustomer":                "UpdateCustomer",
		"DeleteCustomer":                "DeleteCustomer",
		"GetCustomer":                   "GetCustomer",
		"GetCustomerList":               "GetCustomerList",
		"RESTFULTest":                   "RESTFUL Model",
		"CreateTemp":                    "AutomationCode",
		"UpdateAuthority":               "UpdateAuthority",
		"CopyAuthority":                 "CopyAuthority",
		"DeleteUser":                    "DeleteUser",
		"CreateSysDictionaryDetail":     "CreateDictionaryDetail",
		"DeleteSysDictionaryDetail":     "DeleteDictionaryDetail",
		"UpdateSysDictionaryDetail":     "UpdateDictionaryDetail",
		"FindSysDictionaryDetail":       "FindDictionaryDetail",
		"GetSysDictionaryDetailList":    "GetDictionaryDetailList",
		"CreateSysDictionary":           "CreateDictionary",
		"DeleteSysDictionary":           "DeleteDictionary",
		"UpdateSysDictionary":           "UpdateDictionary",
		"FindSysDictionary":             "FindDictionary",
		"GetSysDictionaryList":          "GetDictionaryList",
		"CreateSysOperationRecord":      "CreateOperationRecord",
		"DeleteSysOperationRecord":      "DeleteOperationRecord",
		"FindSysOperationRecord":        "FindOperationRecord",
		"GetSysOperationRecordList":     "GetOperationRecordList",
		"GetTables":                     "GetTables",
		"GetDB":                         "GetDB",
		"GetColumn":                     "GetColumn",
		"DeleteSysOperationRecordByIds": "DeleteOperationRecordByIds",
		"SubsectionUpload":              "SubsectionUpload",
		"CheckFileMd5":                  "CheckFile",
		"MergeFileMd5":                  "MergeFile",
		"SetUserInfo":                   "SetUserInfo",
		"GetServerInfo":                 "GetServerInfo",
		"EmailTest":                     "SendEmailTest",
		"CreateWorkflowProcess":         "CreateWorkflowProcess",
		"DeleteWorkflowProcess":         "DeleteWorkflowProcess",
		"DeleteWorkflowProcessByIds":    "BatchDeleteWorkflowProcess",
		"UpdateWorkflowProcess":         "UpdateWorkflowProcess",
		"FindWorkflowProcess":           "FindWorkflowProcessByID",
		"GetWorkflowProcessList":        "GetWorkflowProcessList",
		"FindWorkflowStep":              "FindWorkflowStep",
		"StartWorkflow":                 "StartWorkflow",
		"GetMyStated":                   "GetMyFlowStated",
		"GetMyNeed":                     "GetMyFlowNeed",
		"GetWorkflowMoveByID":           "GetWorkflowDetailAndHistoryByID",
		"CompleteWorkflowMove":          "CompleteWorkflowMove",
		"Preview":                       "PreviewAutoCode",
		"ImportExcel":                   "ImportExcel",
		"LoadExcel":                     "LoadExcel",
		"ExportExcel":                   "ExportExcel",
		"DownloadTemplate":              "DownloadTemplate",

		"OrdinaryUser":      "OrdinaryUser",
		"NormalUserSubRole": "NormalUserSubRole",
		"TestRole":          "TestRole",

		"Sex":            "Sex",
		"SexDictionary":  "SexDictionary",
		"DBTypeInt":      "DBTypeInt",
		"DBTypeDateTime": "DBTypeDateTime",
		"DBTypeFloat":    "DBTypeFloat",
		"DBTypeString":   "DBTypeString",
		"DBTypeBool":     "DBTypeBool",

		"InstrumentPanel": "InstrumentPanel",
		"AboutUs":         "AboutUs",
		// SuperAdmin
		"RoleManagement":             "RoleManagement",
		"MenuManagement":             "MenuManagement",
		"APIManagement":              "APIManagement",
		"UserManagement":             "UserManagement",
		"PersonalInformation":        "PersonalInformation",
		"SampleFile":                 "SampleFile",
		"ExcelImportExport":          "Excel(ImportAndExport)",
		"MediaLibrary":               "MediaLibrary(UploadAndDownload)",
		"BreakpointContinuation":     "BreakpointContinuation",
		"CustomerList":               "CustomerList(ResourceExample)",
		"SystemTool":                 "SystemTool",
		"CodeGenerator":              "CodeGenerator",
		"FormGenerator":              "FormGenerator",
		"SystemConfiguration":        "SystemConfiguration",
		"DictionaryManagement":       "DictionaryManagement",
		"DictionaryDetails":          "DictionaryDetails",
		"OperationHistory":           "OperationHistory",
		"BreakpointContinuationPlug": "BreakpointContinuation(Plug-in)",
		"OfficialWebsite":            "OfficialWebsite",
		"ServerStatus":               "ServerStatus",
		"WorkflowFunction":           "WorkflowFunction",
		"WorkflowDrawing":            "WorkflowDrawing",
		"WorkflowList":               "WorkflowList",
		"UseWorkflow":                "UseWorkflow",
		"IInitiatedIt":               "IInitiatedIt",
		"MyToDoList":                 "MyToDoList",

		"LeaveProcess":             "LeaveProcess(Demo)",
		"LeaveFail":                "LeaveFail",
		"LeaveSuccess":             "LeaveSuccess",
		"LeaveSuccessDesc":         "The request for leave is completed, and the specific results are waiting to be submitted.",
		"InitiateRequestLeave":     "InitiateRequestLeave",
		"InitiateRequestLeaveDesc": "Initiate a leave process",
		"ExaminationApproval":      "ExaminationApproval",
		"ExaminationApprovalDesc":  "Examination and approval countersignature",

		"Agree":    "Agree",
		"Disagree": "Disagree",
	}
}

func GetChI18n() map[string]string {
	return map[string]string{
		"SuperAdmin": "超级管理员",
		"OtherUser":  "QMPlusUser",

		"UserLogin":                     "用户登录",
		"UserRegister":                  "用户注册",
		"CreateApi":                     "创建API",
		"GetApiList":                    "获取API列表",
		"GetApiDetail":                  "获取API详细信息",
		"DeleteApi":                     "删除API",
		"UpdateApi":                     "更新API",
		"GetAllApis":                    "获取所有API",
		"CreateAuthority":               "创建角色",
		"DeleteAuthority":               "删除角色",
		"GetAuthorityList":              "获取角色列表",
		"GetMenu":                       "获取菜单树",
		"GetMenuList":                   "分页获取菜单树列表",
		"AddBaseMenu":                   "新增菜单",
		"GetBaseMenuTree":               "获取用户动态路由",
		"AddMenuAuthority":              "增加菜单和角色关联关系",
		"GetMenuAuthority":              "获取指定角色菜单",
		"DeleteBaseMenu":                "删除菜单",
		"UpdateBaseMenu":                "更新菜单",
		"GetBaseMenuById":               "根据ID获取菜单",
		"ChangePassword":                "修改密码",
		"GetUserList":                   "获取用户列表",
		"SetUserAuthority":              "修改用户角色",
		"UploadFile":                    "文件上传示例",
		"GetFileList":                   "获取上传文件列表",
		"UpdateCasbin":                  "更改角色API权限",
		"GetPolicyPathByAuthorityId":    "获取权限列表",
		"DeleteFile":                    "删除文件",
		"JsonInBlacklist":               "JWT加入黑名单(退出)",
		"SetDataAuthority":              "设置角色资源权限",
		"GetSystemConfig":               "获取配置文件内容",
		"SetSystemConfig":               "设置配置文件内容",
		"CreateCustomer":                "创建客户",
		"UpdateCustomer":                "更新客户",
		"DeleteCustomer":                "删除客户",
		"GetCustomer":                   "获取单一客户",
		"GetCustomerList":               "获取客户列表",
		"RESTFULTest":                   "RESTFUL模式测试",
		"CreateTemp":                    "自动化代码",
		"UpdateAuthority":               "更新角色信息",
		"CopyAuthority":                 "拷贝角色",
		"DeleteUser":                    "删除用户",
		"CreateSysDictionaryDetail":     "新增字典内容",
		"DeleteSysDictionaryDetail":     "删除字典内容",
		"UpdateSysDictionaryDetail":     "更新字典内容",
		"FindSysDictionaryDetail":       "根据ID获取字典内容",
		"GetSysDictionaryDetailList":    "获取字典内容列表",
		"CreateSysDictionary":           "新增字典",
		"DeleteSysDictionary":           "删除字典",
		"UpdateSysDictionary":           "更新字典",
		"FindSysDictionary":             "根据ID获取字典",
		"GetSysDictionaryList":          "获取字典列表",
		"CreateSysOperationRecord":      "新增操作记录",
		"DeleteSysOperationRecord":      "删除操作记录",
		"FindSysOperationRecord":        "根据ID获取操作记录",
		"GetSysOperationRecordList":     "获取操作记录列表",
		"GetTables":                     "获取数据库表",
		"GetDB":                         "获取所有数据库",
		"GetColumn":                     "获取所选table的所有字段",
		"DeleteSysOperationRecordByIds": "批量删除操作历史",
		"SubsectionUpload":              "插件版分片上传",
		"CheckFileMd5":                  "文件完整度验证",
		"MergeFileMd5":                  "上传完成合并文件",
		"SetUserInfo":                   "设置用户信息",
		"GetServerInfo":                 "获取服务器信息",
		"EmailTest":                     "发送测试邮件",
		"CreateWorkflowProcess":         "新建工作流",
		"DeleteWorkflowProcess":         "删除工作流",
		"DeleteWorkflowProcessByIds":    "批量删除工作流",
		"UpdateWorkflowProcess":         "更新工作流",
		"FindWorkflowProcess":           "根据ID获取工作流",
		"GetWorkflowProcessList":        "获取工作流",
		"FindWorkflowStep":              "获取工作流步骤",
		"StartWorkflow":                 "启动工作流",
		"GetMyStated":                   "获取我发起的工作流",
		"GetMyNeed":                     "获取我的待办",
		"GetWorkflowMoveByID":           "根据ID获取当前节点详情和历史",
		"CompleteWorkflowMove":          "提交工作流",
		"Preview":                       "预览自动化代码",
		"ImportExcel":                   "导入Excel",
		"LoadExcel":                     "下载Excel",
		"ExportExcel":                   "导出Excel",
		"DownloadTemplate":              "下载Excel模板",

		"OrdinaryUser":      "普通用户",
		"NormalUserSubRole": "普通用户子角色",
		"TestRole":          "测试角色",

		"Sex":            "性别",
		"SexDictionary":  "性别字典",
		"DBTypeInt":      "数据库int类型",
		"DBTypeDateTime": "数据库时间日期类型",
		"DBTypeFloat":    "数据库浮点型",
		"DBTypeString":   "数据库字符串",
		"DBTypeBool":     "数据库bool类型",

		"InstrumentPanel": "仪表盘",
		"AboutUs":         "关于我们",
		// superAdmin
		"RoleManagement":             "角色管理",
		"MenuManagement":             "菜单管理",
		"APIManagement":              "API管理",
		"UserManagement":             "用户管理",
		"PersonalInformation":        "个人信息",
		"SampleFile":                 "示例文件",
		"ExcelImportExport":          "Excel(导入导出)",
		"MediaLibrary":               "媒体库(上传下载)",
		"BreakpointContinuation":     "断点续传",
		"CustomerList":               "客户列表(资源示例)",
		"SystemTool":                 "系统工具",
		"CodeGenerator":              "代码生成器",
		"FormGenerator":              "表单生成器",
		"SystemConfiguration":        "系统配置",
		"DictionaryManagement":       "字典管理",
		"DictionaryDetails":          "字典详情",
		"OperationHistory":           "操作历史",
		"BreakpointContinuationPlug": "断点续传(插件版)",
		"OfficialWebsite":            "官方网站",
		"ServerStatus":               "服务器状态",
		"WorkflowFunction":           "工作流功能",
		"WorkflowDrawing":            "工作流绘制",
		"WorkflowList":               "工作流列表",
		"UseWorkflow":                "使用工作流",
		"IInitiatedIt":               "我发起的",
		"MyToDoList":                 "我的待办",

		"LeaveProcess":             "请假流程(演示)",
		"LeaveFail":                "请假失败",
		"LeaveSuccess":             "请假成功",
		"LeaveSuccessDesc":         "请假完成，具体结果等待提交",
		"InitiateRequestLeave":     "发起请假",
		"InitiateRequestLeaveDesc": "发起一个请假流程",
		"ExaminationApproval":      "审批",
		"ExaminationApprovalDesc":  "审批会签",

		"Agree":    "同意",
		"Disagree": "不同意",
	}
}
