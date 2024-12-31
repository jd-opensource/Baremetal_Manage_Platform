module coding.jd.com/bmp/cps-agent

go 1.17

require (
	git.jd.com/jcloud-api-gateway/jcloud-sdk-go v0.0.0-00010101000000-000000000000
	github.com/StackExchange/wmi v1.2.1
	github.com/VividCortex/godaemon v1.0.0
	github.com/astaxie/beego v1.12.3
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/gofrs/flock v0.8.1
	github.com/shirou/gopsutil v3.21.11+incompatible
)

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.3.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/sys v0.0.0-20210816074244-15123e1e1f71 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace git.jd.com/jcloud-api-gateway/jcloud-sdk-go => ./bmp_vendor/jcloud-sdk-go
