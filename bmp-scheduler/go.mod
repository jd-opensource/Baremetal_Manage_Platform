module coding.jd.com/aidc-bmp/bmp-scheduler

go 1.15

require (
	git.jd.com/cps-golang/ironic-common v0.0.0-00010101000000-000000000000
	git.jd.com/cps-golang/log v0.0.0-00010101000000-000000000000
	github.com/beego/beego/v2 v2.0.1
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
	github.com/streadway/amqp v1.0.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/goyy/goyy.v0 v0.0.0-20190218140538-82e7740e526e
)

replace git.jd.com/cps-golang/log => ./bmp_vendor/log

replace git.jd.com/cps-golang/ironic-common => ./bmp_vendor/ironic-common
