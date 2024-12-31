package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/exception"
	"github.com/go-playground/validator/v10"
)

type CreateProjectRequest struct {
	ProjectID   string `json:"projectId"`   // 项目id
	ProjectName string `json:"projectName"` //项目名称
	IsDefault   int8   `json:"isDefault"`
	IsSystem    int8   `json:"isSystem"`
	UserID      string `json:"userId"`
}
type ModifyProjectRequest struct {
	ProjectName string `json:"projectName"` // 项目名称
}

func (req *CreateProjectRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(req); err != nil {
		logger.Warn("CreateProjectRequest.Validate error:", err.Error())
		panic(exception.CommonParamValid)
	}
}

type QueryProjectsRequest struct {
	ProjectName string `json:"projectName"`
	PagingRequest
	ExportType        string `json:"exportType"`
	IsAll             string `json:"isAll"`
	OrderByCreatetime string `json:"orderByCreatetime"`
	Owned             string `json:"owned"`
	//新增，可选，根据用户名获取项目列表
	OwnerName string `json:"ownerName"`
	//新增，可选，根据被转移用户的用户名获取它拥有的项目列表
	SharerName string `json:"sharerName"`
}

func (req *QueryProjectsRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(req); err != nil {
		logger.Warn("QueryProjectsRequest.Validate error:", err.Error())
		panic(exception.CommonParamValid)
	}
}

type ShareProjectRequest struct {
	//
	OwnerName  string `json:"ownerName"`
	SharerName string `json:"sharerName"`
	//新增， 如果部分分享，instance_id逗号分隔； 如果全部分享，传all
	InstanceIDs string `json:"instanceIDs" validate:"required"` //all表示项目下的全部资源,否则instance_id逗号分隔传过来
}

type MoveProjectRequest struct {
	//
	OwnerName string `json:"ownerName"`
	MoverName string `json:"moverName"`
}

type MoveInstancesRequest struct {
	//
	OwnerName      string `json:"ownerName"`
	OwnerProjectID string `json:"ownerProjectID"`
	MoverName      string `json:"moverName"`
	MoverProjectID string `json:"moverProjectID"`
	InstanceIDs    string `json:"instanceIDs"`
}
