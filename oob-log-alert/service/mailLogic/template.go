package mailLogic

const MessageOobTemplate string = `

<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>故障消息</title>
    <style>
        * {
            padding: 0;
            margin: 0;
        }

        .box {
            width: 700px;
            height: 495px;
            position: relative;
        }

        .header {
            background: #108ee9;
            height: 48px;
            font-size: 16px;
            color: #fff;
            font-weight: 500;
            display: flex;
            align-items: center;
        }

        .header span {
            margin-left: 24px;
        }

        .count {
            padding: 0 24px;
            margin-top: 30px;
        }

        .count .tip {
            font-size: 14px;
            color: #333;
            font-weight: 500;
        }

        .count .desc {
            font-size: 14px;
            color: #333;
            text-align: justify;
            line-height: 20px;
            font-weight: 400;
            margin-top: 41px;
            margin-left: 20px;
        }

        .count .desc .line {
            text-decoration: underline;
        }

        .table {
            margin-top: 32px;
            margin-left: 20px;
            width: 624px;
            height: 87px;
            border: .6px solid #979797;
        }

        .table .title {
            font-weight: 400;
            font-size: 12px;
            padding: 5px 0;
            border: .6px solid #979797;
        }

        .table .count {
            font-weight: 400;
            font-size: 12px;
            border: .6px solid #979797;
        }

        .footer {
            position: absolute;
            right: 24px;
            bottom: 59px;
        }

        .footer p:first-child {
            font-size: 14px;
            color: #333;
            font-weight: 500;
            text-align: right;
        }

        .footer p:nth-child(2) {
            font-size: 14px;
            color: #333;
            font-weight: 400;
            text-align: right;
        }
    </style>
</head>
<body>
    <div class="box">
        <div class="header">
            <span>BMP裸金属管理平台</span>
        </div>
        <div class="count">
            <p class="tip">尊敬的用户 admin，您好！</p>
            <p class="desc">
                您帐号中的服务器资源（SN号： <span class="line">{{.SN}}</span> ）于（{{.AlertTime}}）发生【<span class="line">{{.MessageSubType}}</span>】。详情如下：
            </p>
            <table
                class="table"
                border="1"
                cellpadding="0"
                cellspacing="0"
            >
                <thead>
                    <tr>
                        <th class="title">LogID</th>
                        <th class="title">故障类型</th>
                        <th class="title">故障描述</th>
                        <th class="title">故障报警时间</th>
                        <th class="title">报警次数</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                        <th class="count">{{.LogID}}</th>
                        <th class="count">{{.MessageSubType}}</th>
                        <th class="count">{{.Detail}}</th>
                        <th class="count">{{.AlertTime}}</th>
                        <th class="count">1</th>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="footer">
            <p>BMP平台</p>
            <p>本邮件由系统自动发出，请勿直接回复!</p>
        </div>
    </div>
</body>
</html>
`

type MailMessage struct {
	MessageID      string
	ResourceType   string
	FinishTime     string
	MessageType    string
	MessageSubType string
	SN             string
	FaultType      string
	Content        string
	Detail         string
	LogID          string
	AlertTime      string
	AlertCount     string

	InstanceID   string
	InstanceName string
	UserID       string
	UserName     string
}
