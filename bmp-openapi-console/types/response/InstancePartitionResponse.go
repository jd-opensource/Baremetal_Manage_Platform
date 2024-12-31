package response

type QueryInstancePartitionResponse struct {
	PartitionResult struct {
		SystemPartition []Partition `json:"systemPartition"`
		DataPartition   []Partition `json:"dataPartition"`
	} `json:"partition_result"`
}
