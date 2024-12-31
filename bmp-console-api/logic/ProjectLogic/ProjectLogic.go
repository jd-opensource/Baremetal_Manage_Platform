package ProjectLogic

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/BaseLogic"
	UserLoigc "coding.jd.com/aidc-bmp/bmp-console-api/logic/UserLogic"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/excel"
	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/exception"

	sdkProjectParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/project"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
)

func GetProjectById(logger *log.Logger, ProjectId string) (*sdkModels.Project, error) {

	// language := logger.GetPoint("language").(string)
	logid := logger.GetPoint("logid").(string)
	request := sdkProjectParameters.NewDescribeUserProjectParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithProjectID(ProjectId)
	logger.Info("GetProjectById openapi request:", util.InterfaceToJson(request))
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	responseApi, err := openApi.SdkClient.Project.DescribeUserProject(request, nil)
	logger.Info("GetProjectById openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		fmt.Println("DescribeUserProject openapi error:", err.Error())
		logger.Info("DescribeUserProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	project := responseApi.Payload.Result.Project
	if project == nil {
		return nil, errors.New("openapi returns empty project")
	}
	return project, nil
}

func CreateProject(logger *log.Logger, param *sdkModels.CreateProjectRequest) (string, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkProjectParameters.NewCreateUserProjectParams()
	req.WithTraceID(logid)
	req.WithBody(param)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("CreateProject openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.CreateUserProject(req, nil)
	logger.Info("CreateProject openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		fmt.Println("CreateProject openapi error:", err.Error())
		logger.Info("CreateProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return "", err
		}
		return "", errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.ProjectID, nil
}

func ModifyProject(logger *log.Logger, param *sdkModels.ModifyProjectRequest, ProjectId string) error {
	logid := logger.GetPoint("logid").(string)
	req := sdkProjectParameters.NewModifyUserProjectParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithProjectID(ProjectId)
	req.WithBody(param)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("ModifyProject openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.ModifyUserProject(req, nil)
	logger.Info("ModifyProject openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		fmt.Println("ModifyProject openapi error:", err.Error())
		logger.Info("ModifyProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return err
		}
		return errors.New(respErr.Message)
	}
	return nil
}

func ModifyProjectDescription(logger *log.Logger, param *sdkModels.ModifyProjectDescriptionRequest, ProjectId string) error {
	logid := logger.GetPoint("logid").(string)
	req := sdkProjectParameters.NewModifyUserProjectDescriptionParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithProjectID(ProjectId)
	req.WithBody(param)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("ModifyProjectDescription openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.ModifyUserProjectDescription(req, nil)
	logger.Info("ModifyProjectDescription openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		fmt.Println("ModifyProjectDescription openapi error:", err.Error())
		logger.Info("ModifyProjectDescription openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return err
		}
		return errors.New(respErr.Message)
	}
	return nil
}

func DeleteProject(logger *log.Logger, ProjectId string) error {
	logid := logger.GetPoint("logid").(string)
	req := sdkProjectParameters.NewDeleteUserProjectParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithProjectID(ProjectId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DeleteProject openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.DeleteUserProject(req, nil)
	logger.Info("DeleteProject openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DeleteProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return err
		}
		return errors.New(respErr.Message)
	}
	return nil
}

func GetProjectList(logger *log.Logger, param requestTypes.QueryProjectsRequest) (*sdkModels.ProjectList, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkProjectParameters.NewDescribeUserProjectsParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithPageNumber(&param.PageNumber)
	req.WithPageSize(&param.PageSize)
	if param.ExportType != "" || param.IsAll == "1" {
		isAll := "1"
		req.WithIsAll(&isAll)
	}
	var owned int64 = 0
	if param.Owned == "1" {
		owned = 1
	} else if param.Owned == "2" {
		owned = 2
	} else {
		owned = 0
	}
	req.Owned = &owned
	if param.OwnerName != "" {
		req.OwnerName = &param.OwnerName
	}
	if param.SharerName != "" {
		req.SharerName = &param.SharerName
	}

	req.WithOrderByCreatetime(&param.OrderByCreatetime)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("GetProjectList openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.DescribeUserProjects(req, nil)
	logger.Info("GetProjectList openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DeleteProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result, nil
}

func ExportProject(logger *log.Logger, result []*sdkModels.Project) (fileName, downLoadFileName string) {

	language := logger.GetPoint("language").(string)

	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, data := range result {
		var shareStatus string
		ownerText := "拥有者"
		shareText := "共享者"
		if logger.GetPoint("language").(string) == BaseLogic.EN_US {
			ownerText = "Owner"
			shareText = "collaborator"
		}
		if data.Owned == 1 {
			shareStatus = ownerText
		} else {
			shareStatus = shareText
		}
		sheetData = append(sheetData, []string{
			data.ProjectName,
			data.ProjectID,
			shareStatus,
			fmt.Sprintf("%d", data.InstanceCount),
			data.CreatedTime,
			data.UpdatedTime,
		})
	}
	exportProjectHeader := BaseLogic.ExportProjectHeader
	if language == "en_US" {
		exportProjectHeader = BaseLogic.ExportProjectHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportProjectHeader, sheetData); err != nil {
		panic(exception.Exception{Msg: "新建excel失败", Status: BaseLogic.ERROR_CODE})
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "project_list_" + time.Now().Format(BaseLogic.FormatDate) + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Info("保存excel到本地失败" + err.Error())
		panic(exception.Exception{Msg: "保存excel到本地失败", Status: BaseLogic.ERROR_CODE})
	}
	return fileName, downloadFileName
}

func ShareProject(logger *log.Logger, ProjectId string, param requestTypes.ShareProjectRequest) error {

	logid := logger.GetPoint("logid").(string)

	ownUser, _ := UserLoigc.GetUserByName(logger, param.OwnerName)
	if ownUser == nil {
		return fmt.Errorf("param.OwnerName invalid")
	}
	shareUser, _ := UserLoigc.GetUserByName(logger, param.SharerName)
	if shareUser == nil {
		return fmt.Errorf("param.SharerName invalid")
	}

	p := &sdkModels.ShareProjectRequest{
		OwnerID:     &ownUser.UserID,
		SharerID:    &shareUser.UserID,
		InstanceIDs: &param.InstanceIDs,
	}

	req := sdkProjectParameters.NewShareUserProjectParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithProjectID(ProjectId)

	req.WithBody(p)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("ShareUserProject openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.ShareUserProject(req, nil)
	logger.Info("ShareUserProject openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ShareUserProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return err
		}
		return errors.New(respErr.Message)
	}
	return nil
}

func CancelShareProject(logger *log.Logger, ProjectId string, param requestTypes.ShareProjectRequest) error {
	logid := logger.GetPoint("logid").(string)

	ownUser, _ := UserLoigc.GetUserByName(logger, param.OwnerName)
	if ownUser == nil {
		return fmt.Errorf("param.OwnerName invalid")
	}
	shareUser, _ := UserLoigc.GetUserByName(logger, param.SharerName)
	if shareUser == nil {
		return fmt.Errorf("param.SharerName invalid")
	}

	p := &sdkModels.CalcelShareProjectRequest{
		OwnerID:  &ownUser.UserID,
		SharerID: &shareUser.UserID,
	}

	req := sdkProjectParameters.NewCancelShareUserProjectParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithProjectID(ProjectId)
	req.WithBody(p)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("ShareUserProject openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.CancelShareUserProject(req, nil)
	logger.Info("ShareUserProject openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ShareUserProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return err
		}
		return errors.New(respErr.Message)
	}
	return nil
}

func GetShareProjectInfo(logger *log.Logger, ProjectId string) (*sdkModels.ShareProjectInfo, error) {
	logid := logger.GetPoint("logid").(string)

	req := sdkProjectParameters.NewDescribeShareProjectParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithProjectID(ProjectId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DescribeShareProject openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.DescribeShareProject(req, nil)
	logger.Info("DescribeShareProject openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeShareProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result, nil
}

func MoveProject(logger *log.Logger, ProjectId string, param requestTypes.MoveProjectRequest) error {

	logid := logger.GetPoint("logid").(string)

	ownUser, _ := UserLoigc.GetUserByName(logger, param.OwnerName)
	if ownUser == nil {
		return fmt.Errorf("param.OwnerName invalid")
	}
	moverUser, _ := UserLoigc.GetUserByName(logger, param.MoverName)
	if moverUser == nil {
		return fmt.Errorf("param.MoverName invalid")
	}

	p := &sdkModels.MoveProjectRequest{
		OwnerID: &ownUser.UserID,
		MoverID: &moverUser.UserID,
	}

	req := sdkProjectParameters.NewMoveUserProjectParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithProjectID(ProjectId)
	req.WithBody(p)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("MoveUserProject openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.MoveUserProject(req, nil)
	logger.Info("MoveUserProject openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("MoveUserProject openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return err
		}
		return errors.New(respErr.Message)
	}
	return nil
}

func MoveInstances(logger *log.Logger, param *requestTypes.MoveInstancesRequest) error {

	logid := logger.GetPoint("logid").(string)

	ownUser, _ := UserLoigc.GetUserByName(logger, param.OwnerName)
	if ownUser == nil {
		return fmt.Errorf("param.OwnerName invalid")
	}
	moverUser, _ := UserLoigc.GetUserByName(logger, param.MoverName)
	if moverUser == nil {
		return fmt.Errorf("param.MoverName invalid")
	}
	ownerProject, _ := GetProjectById(logger, param.OwnerProjectID)
	if ownerProject == nil {
		return fmt.Errorf("param.OwnerProjectID invalid")
	}
	moverProject, _ := GetProjectById(logger, param.MoverProjectID)
	if moverProject == nil {
		return fmt.Errorf("param.MoverProjectID invalid")
	}

	p := &sdkModels.MoveInstancesRequest{
		OwnerID:        &ownUser.UserID,
		MoverID:        &moverUser.UserID,
		OwnerProjectID: &param.OwnerProjectID,
		MoverProjectID: &param.MoverProjectID,
		InstanceIDs:    &param.InstanceIDs,
	}

	req := sdkProjectParameters.NewMoveUserInstancesParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithBody(p)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("MoveUserInstance openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Project.MoveUserInstances(req, nil)
	logger.Info("MoveUserInstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("MoveUserInstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return err
		}
		return errors.New(respErr.Message)
	}
	return nil
}
