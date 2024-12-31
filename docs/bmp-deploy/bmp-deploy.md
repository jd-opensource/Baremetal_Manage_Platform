# File and Directory Explanation for `bmp-deploy`
## List of Files and Directories
~~~
bmp-deploy directory
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

## Explanation of Files and Directories
* `bmp-deploy.sh`: Main deployment script.
* `cache`: Cache directory for GuestOS images and LiveOS images.
* `config`: Configuration directory.
* `data`: Cache directory for other binary files, including PXE boot programs and device import templates.
* `docker-compose.yml`: Docker Compose service orchestration file.
* `env`: Environment variable directory for each component.
* `.env`: Main configuration file for Docker Compose environment variables.
* `ReleaseNotes`: Version release information.
* `script`: Directory for initialization scripts of each component.
* `sql`: Directory for database initialization scripts.
* `template`: Template file directory.
* `tool`: Tool directory.
