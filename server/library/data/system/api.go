package data

import (
	"github.com/gookit/color"
	model "go_base_server/app/model/system"
	"go_base_server/library/global"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

//@description: apis 表数据初始化
func (a *api) Init() error {
	apis := []model.Api{
		{Model: global.Model{ID: 1}, Path: "/base/login", Description: I18nHash["UserLogin"], ApiGroup: "base", Method: "POST"},
		{Model: global.Model{ID: 2}, Path: "/user/register", Description: I18nHash["UserRegister"], ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 3}, Path: "/api/createApi", Description: I18nHash["CreateApi"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 4}, Path: "/api/getApiList", Description: I18nHash["GetApiList"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 5}, Path: "/api/getApiById", Description: I18nHash["GetApiDetail"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 6}, Path: "/api/deleteApi", Description: I18nHash["DeleteApi"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 7}, Path: "/api/updateApi", Description: I18nHash["UpdateApi"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 8}, Path: "/api/getAllApis", Description: I18nHash["GetAllApis"], ApiGroup: "api", Method: "POST"},
		{Model: global.Model{ID: 9}, Path: "/authority/createAuthority", Description: I18nHash["CreateAuthority"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 10}, Path: "/authority/deleteAuthority", Description: I18nHash["DeleteAuthority"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 11}, Path: "/authority/getAuthorityList", Description: I18nHash["GetAuthorityList"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 12}, Path: "/menu/getMenu", Description: I18nHash["GetMenu"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 13}, Path: "/menu/getMenuList", Description: I18nHash["GetMenuList"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 14}, Path: "/menu/addBaseMenu", Description: I18nHash["AddBaseMenu"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 15}, Path: "/menu/getBaseMenuTree", Description: I18nHash["GetBaseMenuTree"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 16}, Path: "/menu/addMenuAuthority", Description: I18nHash["AddMenuAuthority"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 17}, Path: "/menu/getMenuAuthority", Description: I18nHash["GetMenuAuthority"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 18}, Path: "/menu/deleteBaseMenu", Description: I18nHash["DeleteBaseMenu"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 19}, Path: "/menu/updateBaseMenu", Description: I18nHash["UpdateBaseMenu"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 20}, Path: "/menu/getBaseMenuById", Description: I18nHash["GetBaseMenuById"], ApiGroup: "menu", Method: "POST"},
		{Model: global.Model{ID: 21}, Path: "/user/changePassword", Description: I18nHash["ChangePassword"], ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 23}, Path: "/user/getUserList", Description: I18nHash["GetUserList"], ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 24}, Path: "/user/setUserAuthority", Description: I18nHash["SetUserAuthority"], ApiGroup: "user", Method: "POST"},
		{Model: global.Model{ID: 25}, Path: "/fileUploadAndDownload/upload", Description: I18nHash["UploadFile"], ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 26}, Path: "/fileUploadAndDownload/getFileList", Description: I18nHash["GetFileList"], ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 27}, Path: "/casbin/updateCasbin", Description: I18nHash["UpdateCasbin"], ApiGroup: "casbin", Method: "POST"},
		{Model: global.Model{ID: 28}, Path: "/casbin/getPolicyPathByAuthorityId", Description: I18nHash["GetPolicyPathByAuthorityId"], ApiGroup: "casbin", Method: "POST"},
		{Model: global.Model{ID: 29}, Path: "/fileUploadAndDownload/deleteFile", Description: I18nHash["DeleteFile"], ApiGroup: "fileUploadAndDownload", Method: "POST"},
		{Model: global.Model{ID: 30}, Path: "/jwt/jsonInBlacklist", Description: I18nHash["JsonInBlacklist"], ApiGroup: "jwt", Method: "POST"},
		{Model: global.Model{ID: 31}, Path: "/authority/setDataAuthority", Description: I18nHash["SetDataAuthority"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 32}, Path: "/system/getSystemConfig", Description: I18nHash["GetSystemConfig"], ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 33}, Path: "/system/setSystemConfig", Description: I18nHash["SetSystemConfig"], ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 34}, Path: "/customer/customer", Description: I18nHash["CreateCustomer"], ApiGroup: "customer", Method: "POST"},
		{Model: global.Model{ID: 35}, Path: "/customer/customer", Description: I18nHash["UpdateCustomer"], ApiGroup: "customer", Method: "PUT"},
		{Model: global.Model{ID: 36}, Path: "/customer/customer", Description: I18nHash["DeleteCustomer"], ApiGroup: "customer", Method: "DELETE"},
		{Model: global.Model{ID: 37}, Path: "/customer/customer", Description: I18nHash["GetCustomer"], ApiGroup: "customer", Method: "GET"},
		{Model: global.Model{ID: 38}, Path: "/customer/customerList", Description: I18nHash["GetCustomerList"], ApiGroup: "customer", Method: "GET"},
		{Model: global.Model{ID: 39}, Path: "/casbin/casbinTest/:pathParam", Description: I18nHash["RESTFULTest"], ApiGroup: "casbin", Method: "GET"},
		{Model: global.Model{ID: 40}, Path: "/autoCode/createTemp", Description: I18nHash["CreateTemp"], ApiGroup: "autoCode", Method: "POST"},
		{Model: global.Model{ID: 41}, Path: "/authority/updateAuthority", Description: I18nHash["UpdateAuthority"], ApiGroup: "authority", Method: "PUT"},
		{Model: global.Model{ID: 42}, Path: "/authority/copyAuthority", Description: I18nHash["CopyAuthority"], ApiGroup: "authority", Method: "POST"},
		{Model: global.Model{ID: 43}, Path: "/user/deleteUser", Description: I18nHash["DeleteUser"], ApiGroup: "user", Method: "DELETE"},
		{Model: global.Model{ID: 44}, Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: I18nHash["CreateSysDictionaryDetail"], ApiGroup: "sysDictionaryDetail", Method: "POST"},
		{Model: global.Model{ID: 45}, Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: I18nHash["DeleteSysDictionaryDetail"], ApiGroup: "sysDictionaryDetail", Method: "DELETE"},
		{Model: global.Model{ID: 46}, Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: I18nHash["UpdateSysDictionaryDetail"], ApiGroup: "sysDictionaryDetail", Method: "PUT"},
		{Model: global.Model{ID: 47}, Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: I18nHash["FindSysDictionaryDetail"], ApiGroup: "sysDictionaryDetail", Method: "GET"},
		{Model: global.Model{ID: 48}, Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: I18nHash["GetSysDictionaryDetailList"], ApiGroup: "sysDictionaryDetail", Method: "GET"},
		{Model: global.Model{ID: 49}, Path: "/sysDictionary/createSysDictionary", Description: I18nHash["CreateSysDictionary"], ApiGroup: "sysDictionary", Method: "POST"},
		{Model: global.Model{ID: 50}, Path: "/sysDictionary/deleteSysDictionary", Description: I18nHash["DeleteSysDictionary"], ApiGroup: "sysDictionary", Method: "DELETE"},
		{Model: global.Model{ID: 51}, Path: "/sysDictionary/updateSysDictionary", Description: I18nHash["UpdateSysDictionary"], ApiGroup: "sysDictionary", Method: "PUT"},
		{Model: global.Model{ID: 52}, Path: "/sysDictionary/findSysDictionary", Description: I18nHash["FindSysDictionary"], ApiGroup: "sysDictionary", Method: "GET"},
		{Model: global.Model{ID: 53}, Path: "/sysDictionary/getSysDictionaryList", Description: I18nHash["GetSysDictionaryList"], ApiGroup: "sysDictionary", Method: "GET"},
		{Model: global.Model{ID: 54}, Path: "/sysOperationRecord/createSysOperationRecord", Description: I18nHash["CreateSysOperationRecord"], ApiGroup: "sysOperationRecord", Method: "POST"},
		{Model: global.Model{ID: 55}, Path: "/sysOperationRecord/deleteSysOperationRecord", Description: I18nHash["DeleteSysOperationRecord"], ApiGroup: "sysOperationRecord", Method: "DELETE"},
		{Model: global.Model{ID: 56}, Path: "/sysOperationRecord/findSysOperationRecord", Description: I18nHash["FindSysOperationRecord"], ApiGroup: "sysOperationRecord", Method: "GET"},
		{Model: global.Model{ID: 57}, Path: "/sysOperationRecord/getSysOperationRecordList", Description: I18nHash["GetSysOperationRecordList"], ApiGroup: "sysOperationRecord", Method: "GET"},
		{Model: global.Model{ID: 58}, Path: "/autoCode/getTables", Description: I18nHash["GetTables"], ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 59}, Path: "/autoCode/getDB", Description: I18nHash["GetDB"], ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 60}, Path: "/autoCode/getColumn", Description: I18nHash["GetColumn"], ApiGroup: "autoCode", Method: "GET"},
		{Model: global.Model{ID: 61}, Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: I18nHash["DeleteSysOperationRecordByIds"], ApiGroup: "sysOperationRecord", Method: "DELETE"},
		{Model: global.Model{ID: 62}, Path: "/simpleUploader/upload", Description: I18nHash["SubsectionUpload"], ApiGroup: "simpleUploader", Method: "POST"},
		{Model: global.Model{ID: 63}, Path: "/simpleUploader/checkFileMd5", Description: I18nHash["CheckFileMd5"], ApiGroup: "simpleUploader", Method: "GET"},
		{Model: global.Model{ID: 64}, Path: "/simpleUploader/mergeFileMd5", Description: I18nHash["MergeFileMd5"], ApiGroup: "simpleUploader", Method: "GET"},
		{Model: global.Model{ID: 65}, Path: "/user/setUserInfo", Description: I18nHash["SetUserInfo"], ApiGroup: "user", Method: "PUT"},
		{Model: global.Model{ID: 66}, Path: "/system/getServerInfo", Description: I18nHash["GetServerInfo"], ApiGroup: "system", Method: "POST"},
		{Model: global.Model{ID: 67}, Path: "/email/emailTest", Description: I18nHash["EmailTest"], ApiGroup: "email", Method: "POST"},
		{Model: global.Model{ID: 68}, Path: "/workflowProcess/createWorkflowProcess", Description: I18nHash["CreateWorkflowProcess"], ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 69}, Path: "/workflowProcess/deleteWorkflowProcess", Description: I18nHash["DeleteWorkflowProcess"], ApiGroup: "workflowProcess", Method: "DELETE"},
		{Model: global.Model{ID: 70}, Path: "/workflowProcess/deleteWorkflowProcessByIds", Description: I18nHash["DeleteWorkflowProcessByIds"], ApiGroup: "workflowProcess", Method: "DELETE"},
		{Model: global.Model{ID: 71}, Path: "/workflowProcess/updateWorkflowProcess", Description: I18nHash["UpdateWorkflowProcess"], ApiGroup: "workflowProcess", Method: "PUT"},
		{Model: global.Model{ID: 72}, Path: "/workflowProcess/findWorkflowProcess", Description: I18nHash["FindWorkflowProcess"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 73}, Path: "/workflowProcess/getWorkflowProcessList", Description: I18nHash["GetWorkflowProcessList"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 74}, Path: "/workflowProcess/findWorkflowStep", Description: I18nHash["FindWorkflowStep"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 75}, Path: "/workflowProcess/startWorkflow", Description: I18nHash["StartWorkflow"], ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 76}, Path: "/workflowProcess/getMyStated", Description: I18nHash["GetMyStated"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 77}, Path: "/workflowProcess/getMyNeed", Description: I18nHash["GetMyNeed"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 78}, Path: "/workflowProcess/getWorkflowMoveByID", Description: I18nHash["GetWorkflowMoveByID"], ApiGroup: "workflowProcess", Method: "GET"},
		{Model: global.Model{ID: 79}, Path: "/workflowProcess/completeWorkflowMove", Description: I18nHash["CompleteWorkflowMove"], ApiGroup: "workflowProcess", Method: "POST"},
		{Model: global.Model{ID: 80}, Path: "/autoCode/preview", Description: I18nHash["Preview"], ApiGroup: "autoCode", Method: "POST"},
		{Model: global.Model{ID: 81}, Path: "/excel/importExcel", Description: I18nHash["ImportExcel"], ApiGroup: "excel", Method: "POST"},
		{Model: global.Model{ID: 82}, Path: "/excel/loadExcel", Description: I18nHash["LoadExcel"], ApiGroup: "excel", Method: "GET"},
		{Model: global.Model{ID: 83}, Path: "/excel/exportExcel", Description: I18nHash["ExportExcel"], ApiGroup: "excel", Method: "POST"},
		{Model: global.Model{ID: 84}, Path: "/excel/downloadTemplate", Description: I18nHash["DownloadTemplate"], ApiGroup: "excel", Method: "GET"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 84}).Find(&[]model.Api{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 定义表名
func (a *api) TableName() string {
	return "apis"
}
