module coding.jd.com/aidc-bmp/bmp-openapi

go 1.15

require github.com/beego/beego/v2 v2.0.1

require (
	coding.jd.com/aidc-bmp/bmp_log v0.0.0-00010101000000-000000000000
	git.jd.com/cps-golang/ironic-common v0.0.0-00010101000000-000000000000
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis v6.15.1+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/streadway/amqp v1.0.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/goyy/goyy.v0 v0.0.0-20190218140538-82e7740e526e
)

replace github.com/astaxie/beego => github.com/beego/beego v1.12.3

replace git.jd.com/cps-golang/ironic-common => ./bmp_vendor/ironic-common

replace coding.jd.com/aidc-bmp/bmp_log => ./bmp_vendor/bmp_log
