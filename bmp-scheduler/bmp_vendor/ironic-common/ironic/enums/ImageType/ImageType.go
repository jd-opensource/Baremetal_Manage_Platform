package ImageType

import "strings"

const (
	CENTOS = "centos"

	UBUNTU = "ubuntu"

	WINDOWS = "windows"
)

func IsWindows(imageType string) bool {
	return strings.HasPrefix(strings.ToUpper(imageType), WINDOWS)
}
