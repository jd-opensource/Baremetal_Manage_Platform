package mailLogic

const BaseSubjest string = "[BMP消息中心]"

const MessageInBondAlertTemplate string = `



<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>报警消息</title>
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

        .count .desc .id {
            color: #108ee9;
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
            <p class="tip">尊敬的用户 {{.UserName}}，您好！</p>
            <p class="desc">
                【监控报警】：默认机房 的实例<span class="id">{{.InstanceName}}/{{.InstanceID}}</span>当前<span class="line">{{.MessageSubType}}</span> - 告警值 {{.AlertValue}}，持续时间：{{.AlertPeriod}}分钟，所属规则：{{.RuleName}}，请您及时查看。
            </p>
        </div>
        <div class="footer">
            <p>BMP平台</p>
            <p>本邮件由系统自动发出，请勿直接回复!</p>
        </div>
    </div>
</body>
</html>
`

type InbondAlertMailMessage struct {
	MessageID      string
	SN             string
	MessageType    string
	MessageSubType string
	ResourceType   string
	AlertTime      string
	AlertPeriod    int
	AlertValue     string
	Trigger        string
	UserID         string
	UserName       string
	InstanceID     string
	InstanceName   string
	Metric         string
	MetricName     string
	IsRecover      int
	RuleName       string
	RuleID         string
}
