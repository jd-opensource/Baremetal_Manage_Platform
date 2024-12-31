package emailTemplate

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
)

type CheckTimeoutCommandTpl struct {
	CurrentDate string
	Commands    []*commandDao.Command
	Device      deviceDao.Device
	Instance    instanceDao.Instance
}

const CheckTimeoutCommandMailContent string = `
<html>
<meta charset="UTF-8"/>
<head>
    <style>
        table {
            font-size: 14px;
            border-collapse: collapse;
            border-spacing: 0;
            width: 100%;
        }

        td, th {
            border: 1px solid #ddd;
            padding: 8px;
            text-align: left;
        }

        th {
            padding-top: 12px;
            padding-bottom: 12px;
            background-color: rgb(240, 240, 240);
        }
    </style>
</head>
<body>
<p>hi all:</p>
<p>统计时间：<span th:text={{.CurrentDate}}></span></p>
<table>
    <thead>
    <tr>
        <th>idcId</th>
        <th>device_type_id</th>
        <th>sn</th>
        <th>ilo_ip</th>
        <th>instance_id</th>
        <th>name</th>
        <th>status</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <td th:text={{.Device.IDcID}}></td>
        <td th:text={{.Device.DeviceTypeID}}></td>
        <td th:text={{.Device.Sn}}></td>
        <td th:text={{.Device.IloIP}}></td>
        <td th:text={{.Instance.InstanceID}}></td>
        <td th:text={{.Instance.InstanceName}}></td>
        <td th:text={{.Instance.Status}}></td>
    </tr>
    </tbody>
</table>
<br/>
<table>
    <thead>
    <tr>
        <th>id</th>
        <th>request_id</th>
        <th>sn</th>
        <th>instance_id</th>
        <th>action</th>
        <th>type</th>
        <th>status</th>
        <th>parent_command_id</th>
        <th>execute_count</th>
        <th>create_time</th>
        <th>update_time</th>
    </tr>
    </thead>
    <tbody>
    {{range .Commands}}
        <tr>
            <td th:text={{.ID}}></td>
            <td th:text={{.RequestID}}></td>
            <td th:text={{.Sn}}></td>
            <td th:text={{.InstanceID}}></td>
            <td th:text={{.Action}}></td>
            <td th:text={{.Type}}></td>
            <td th:text={{.Status}}></td>
            <td th:text={{.ParentCommandID}}></td>
            <td th:text={{.ExecuteCount}}></td>
            <td th:text={{.CreatedTime}}></td>
            <td th:text={{.UpdatedTime}}></td>
        </tr>
    {{end}}
    </tbody>
</table>
</body>
</html>`
