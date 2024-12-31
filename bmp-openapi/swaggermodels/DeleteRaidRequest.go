package swaggermodels

// swagger:parameters deleteRaid
type DeleteRaidRequest struct {
	WriteRequestHeader

	// in: path
	RaidID string `json:"raid_id"`
}
