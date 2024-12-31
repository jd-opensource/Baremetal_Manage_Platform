package swaggermodels

// swagger:parameters describeRaid
type DescribeRaidRequest struct {
	ReadRequestHeader

	// in: path
	RaidID string `json:"raid_id"`
}
