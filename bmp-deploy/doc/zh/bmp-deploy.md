# bmp-deploy文件和目录说明
## 文件和目录列表
~~~
bmp-deploy目录
├── bmp-deploy.sh
├── cache
├── config
├── data
├── docker-compose.yml
├── env
├── .env
├── ReleaseNotes
├── script
├── sql
├── template
└── tool
~~~
## 文件和目录说明
* bmp-deploy.sh 主部署脚本
* cache GuestOS镜像和LiveOS镜像缓存目录
* config 配置目录
* data 其它二进制文件缓存目录，包含pxe引导程序，设备导入模板
* docker-compose.yml docker-compose服务编排文件
* env 各组件环境变量目录
* .env docker-compose环境变量主配置文件
* ReleaseNotes 版本发布信息
* script 各组件初始化脚本目录
* sql 数据库初始化脚本目录
* template 模板文件目录
* tool 工具目录
