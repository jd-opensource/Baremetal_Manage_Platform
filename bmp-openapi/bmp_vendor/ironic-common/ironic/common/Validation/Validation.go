package Validation

const (
	REGEX_ID               string = `^[a-zA-Z0-9-_]{2,36}$`
	REGEX_RAID_VOLUME_TYPE string = `^system$|^data$`
	REGEX_NETWORK_TYPE     string = `^basic$|^vpc$|^retail$`
	INTERFACE_MODE         string = `^bond$|^dual$`
	REGEX_IP_ADDRESS       string = `(^127\.0\.0\.1)|(^192\.168)|(^10\.)|(^172\.1[6-9])|(^172\.2[0-9])|(^172\.3[0-1])`
	REGEX_CIDR             string = `(10|11|172|192)\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})/(\\d{1,3})`
	REGEX_PASSWORD         string = "^(?=.*[a-z])(?=.*[A-Z])(?=.*[\\d\\* \\(\\)`~!@#\\$%&_\\-+=\\|\\{\\}\\[\\]:\";\\'<>,.\\?\\/\\）])[a-zA-Z\\d\\* \\(\\)`~!@#\\$%&_\\-+=\\|\\{\\}\\[\\]:\";\\'<>,.\\?\\/\\）]{8,30}$"
	REGEX_DEVICE_STATUS    string = "^putawaying$|^creating$|^starting$|^running$|^stopping$|^stopped$|^restarting$|^error$|^upgrading$|^reinstalling$|^destroying$|^resetting_password$"
	REGEX_RSA_PUBLIC_KEY   string = "^ssh-rsa AAAAB3NzaC1yc2.*"
	REGEX_MAIL_TYPE        string = "^api_error$|^command_warning$"
	REGEX_EMAIL            string = "[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,6}"
	REGEX_SOURCE_TYPE      string = "^common$|^customize$|^user_defined$"
)
