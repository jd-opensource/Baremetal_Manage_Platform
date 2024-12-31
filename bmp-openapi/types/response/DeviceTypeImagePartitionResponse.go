package response

type QueryDeviceTypeImagePartitionResponse struct {
	//系统盘分区列表
	SystemPartition []Partition `json:"systemPartition"`
	//数据盘分区列表
	DataPartition []Partition `json:"dataPartition"`
}

type Partition struct {
	//format,如[swap, xfs]
	FsType string `json:"format"`
	//point,如[swap, /, /var]
	MountPoint string `json:"point"`
	// 分区大小, MB为单位
	Size int `json:"size"`
}
