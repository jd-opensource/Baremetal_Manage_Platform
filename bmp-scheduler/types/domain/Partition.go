package domain

type Partition struct {
	Size       int    `json:"size"`
	FsType     string `json:"fs_type"`
	MountPoint string `json:"mountpoint"`
	Label      string `json:"label"`
}
