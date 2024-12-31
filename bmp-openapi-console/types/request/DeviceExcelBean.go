package request

type DeviceExcelBean struct {
	Id         int    `xlsx:"0"`  //序号
	Region     string `xlsx:"1"`  //region
	Az         string `xlsx:"2"`  //az
	Sn         string `xlsx:"3"`  //SN号
	SystemIp   string `xlsx:"4"`  //系统IP
	IloIp      string `xlsx:"5"`  //管理IP
	DeviceType string `xlsx:"6"`  //型号
	Cabinet    string `xlsx:"7"`  //机柜
	UPosition  string `xlsx:"8"`  //U位
	Mac1       string `xlsx:"9"`  //mac1
	Mac2       string `xlsx:"10"` //mac2
	SwitchIp   string `xlsx:"11"` //switchIp
}

var DeviceExcelColumn []string = []string{"序号", "region", "az", "SN号", "系统IP", "管理IP", "型号", "机柜", "U位", "Mac1", "Mac2", "switchIP"}
