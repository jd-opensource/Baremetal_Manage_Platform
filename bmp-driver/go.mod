module coding.jd.com/aidc-bmp/bmp-driver

go 1.15

require (
	coding.jd.com/aidc-bmp/bmp-scheduler v0.0.0-00010101000000-000000000000
	git.jd.com/cps-golang/ironic-common v0.0.0-00010101000000-000000000000
	git.jd.com/cps-golang/log v0.0.0-00010101000000-000000000000
	github.com/beego/beego/v2 v2.0.1
	github.com/streadway/amqp v1.0.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace coding.jd.com/aidc-bmp/bmp-scheduler => ./bmp_vendor/bmp-scheduler

replace git.jd.com/cps-golang/ironic-common => ./bmp_vendor/ironic-common

replace git.jd.com/cps-golang/log => ./bmp_vendor/log
