package service

import (
	"time"

	"coding.jd.com/aidc-bmp/oob-log-alert/object"
)

type logMailTpl struct {
	CPS  object.CPSRecord
	Logs []MailAlertLog
}

type MailAlertLog struct {
	EventTime   time.Time
	FaultTypeZh string
	FaultNameZh string
}

type alertOOBMailTpl struct {
	CPS object.CPSRecord
	Err string
}

type cpsResult struct {
	CPS  object.CPSRecord
	Code int
	Msg  string
	// LastMonitorTimeField string
}

type dailyReportMailTpl struct {
	IdcId           string
	TotalCount      int
	SuccessCount    int
	OOBErrorCount   int
	OtherErrorCount int
	CPSResult       []cpsResult
}

const logMailSubject string = "带外日志告警"
const alertOOBMailSubject string = "带外连接失败"
const dailyReportMailSubject string = "监控日报"
const powerStatusMailSubject string = "实例状态监控日报"

const powerStatusTemplate string = `
<html>
    <head>
    </head>
</html>
<body>
    <div>{{.Az}}可用区实例状态被带外巡检自动更新详情:</div>
    <table  border="1" cellspacing="0">
        <tr>
            <th>实例sn</th>
            <th>实例id</th>
            <th>az</th>
            <th>用户pin</th>
            <th>ilo_ip</th>
            <th>更改前状态</th>
            <th>更改后状态</th>
        </tr>
        {{range .DiffStatus}}
            <tr>
                <td>{{.SN}}</td>
                <td>{{.InstanceID}}</td>
                <td>{{.Az}}</td>
                <td>{{.Pin}}</td>
                <td>{{.IloIp}}</td>
                <td>{{.Status}}</td>
                <td>{{.IloStatus}}</td>
            </tr>
        {{end}}
    </table>

    <div>{{.Az}}可用区带外不通的实例如下:</div>
    <table  border="1" cellspacing="0">
        <tr>
            <th>实例sn</th>
            <th>实例id</th>
            <th>az</th>
            <th>用户pin</th>
            <th>ilo_ip</th>
            <th>控制台状态</th>
        </tr>
        {{range .Unreachable}}
            <tr>
                <td>{{.SN}}</td>
                <td>{{.InstanceID}}</td>
                <td>{{.Az}}</td>
                <td>{{.Pin}}</td>
                <td>{{.IloIp}}</td>
                <td>{{.Status}}</td>
            </tr>
        {{end}}
    </table>
</body>
`

const logMailTemplate string = `
<html>
    <head>
    </head>
</html>
<body>
    <div>CPS信息</div>
    <table  border="1" cellspacing="0">
        <tr>
            <th>SN</th>
            <th>IdcId</th>
            <th>用户pin</th>
            <th>品牌</th>
            <th>型号</th>
            <th>套餐</th>
            <th>机柜</th>
            <th>已售</th>
            <th>iLO IP</th>
        </tr>
        <tr>
            <td>{{.CPS.SN}}</td>
            <td>{{.CPS.IdcId}}</td>
            <td>{{.CPS.Pin}}</td>
            <td>{{.CPS.Brand}}</td>
            <td>{{.CPS.Model}}</td>
            <td>{{.CPS.Package}}</td>
            <td>{{.CPS.Cabinet}}</td>
            <td>{{.CPS.IsSale}}</td>
            <td>{{.CPS.IloIP}}</td>
        </tr>
    </table>
    <br><br>
    <div>告警日志</div>
    <table border="1" cellspacing="0">
        <tr>
            <th>故障带外时间</th>
            <th>故障类型</th>
            <th>故障名称</th>
		</tr>
    
    {{range .Logs}}
        <tr>
            <td>{{.EventTime}}</td>
            <td>{{.FaultTypeZh}}</td>
            <td>{{.FaultNameZh}}</td>
        </tr>
	{{end}}
            
        </table>
</body>
`

const alertOOBMailTemplate string = `
<html>
    <head>
    </head>
</html>
<body>
    <div>CPS信息</div>
    <table  border="1" cellspacing="0">
        <tr>
            <th>SN</th>
            <th>IdcId</th>
            <th>品牌</th>
            <th>型号</th>
            <th>套餐</th>
            <th>机柜</th>
            <th>已售</th>
            <th>iLO IP</th>
        </tr>
        <tr>
            <td>{{.CPS.SN}}</td>
            <td>{{.CPS.IdcId}}</td>
            <td>{{.CPS.Brand}}</td>
            <td>{{.CPS.Model}}</td>
            <td>{{.CPS.Package}}</td>
            <td>{{.CPS.Cabinet}}</td>
            <td>{{.CPS.IsSale}}</td>
            <td>{{.CPS.IloIP}}</td>
        </tr>
    </table>
    <br><br>
    <div>带外连接错误</div>
    <table border="1" cellspacing="0">
        <tr>
            <th>错误信息</th>      
		</tr>
			
		<tr>
            <td>{{.Err}}</td>
        </tr>
    </table>
</body>
`

const dailyReportMailTemplate string = `
<html>
    <head>
    </head>
</html>
<body>
    <div>CPS监控统计</div>
    <table  border="1" cellspacing="0">
        <tr>
            <th>IdcId</th>
            <th>总数</th>
            <th>成功</th>
            <th>带外访问错误</th>
            <th>其他错误</th>
        </tr>
        <tr>
            <td>{{.IdcId}}</td>
            <td>{{.TotalCount}}</td>
            <td>{{.SuccessCount}}</td>
            <td><font color="red">{{.OOBErrorCount}}</font></td>
            <td><font color="red">{{.OtherErrorCount}}</font></td>
        </tr>
    </table>
    <br><br><br><br>
    <div>CPS监控详情</div>
    <table border="1" cellspacing="0">
        <tr>
            <th>SN</th>
            <th>品牌</th>
            <th>型号</th>
            <th>套餐</th>
            <th>机柜</th>
            <th>已售</th>
            <th>iLO IP</th>
            <th>Code</th>
            <th>Msg</th>
        </tr>
        {{range .CPSResult}}
        {{if eq .Code 0}}
        <tr>
        {{else}}
        <tr class="color:red;">
        {{end}}
            <td>{{.CPS.SN}}</td>
            <td>{{.CPS.Brand}}</td>
            <td>{{.CPS.Model}}</td>
            <td>{{.CPS.Package}}</td>
            <td>{{.CPS.Cabinet}}</td>
            <td>{{.CPS.IsSale}}</td>
            <td>{{.CPS.IloIP}}</td>
            <td>{{.Code}}</td>
            <td>{{.Msg}}</td>
        </tr>
		{{end}}
            
        </table>
</body>
`
