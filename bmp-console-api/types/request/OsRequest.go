package request

type QueryOssRequest struct {
	DeviceType string `json:"deviceType"`
	Region     string `json:"region"`
	ImageType  string `json:"imageType"`
}
type QueryOsPartitionsRequest struct {
	DeviceType string `json:"deviceType"`
	Region     string `json:"region"`
	ImageId    string `json:"imageId"`
}
type QueryOsMountPointsRequest struct {
	VolumeType string `json:"volumeType"`
	Region     string `json:"region"`
}
type QueryOOSFileSystemFormatsRequest struct {
	Region string `json:"region"`
}
