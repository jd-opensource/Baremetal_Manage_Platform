module coding.jd.com/aidc-bmp/bmp-operation-api

go 1.17

require github.com/beego/beego/v2 v2.0.1

require (
	coding.jd.com/aidc-bmp/bmp-openapi-sdk v0.0.0-00010101000000-000000000000
	github.com/gorilla/websocket v1.5.1
)

require (
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/go-openapi/analysis v0.21.2 // indirect
	github.com/go-openapi/errors v0.20.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/loads v0.21.1 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-openapi/validate v0.21.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-redis/redis/v7 v7.4.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.7.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.1.3 // indirect
	github.com/richardlehane/mscfb v1.0.3 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/xuri/efp v0.0.0-20210322160811-ab561f5b45e3 // indirect
	go.mongodb.org/mongo-driver v1.10.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	google.golang.org/protobuf v1.23.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	git.jd.com/cps-golang/ironic-common v0.0.0-00010101000000-000000000000
	git.jd.com/cps-golang/log v0.0.0-00010101000000-000000000000
	github.com/go-openapi/runtime v0.24.2
	github.com/go-openapi/strfmt v0.21.3
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v4.2.0+incompatible
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/xuri/excelize/v2 v2.5.0
	golang.org/x/crypto v0.14.0
)

replace git.jd.com/cps-golang/log => ./bmp_vendor/log

replace git.jd.com/cps-golang/ironic-common => ./bmp_vendor/ironic-common

replace coding.jd.com/aidc-bmp/bmp-openapi-sdk => ./bmp_vendor/bmp-openapi-sdk
