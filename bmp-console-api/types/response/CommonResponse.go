package response

//新增加或者删除返给前端的true或者false
type CommonResponse struct {
	Success bool `json:"success"`
}
