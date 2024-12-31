package response

type RegionsResponse struct {
	Regions []Region `json:"regions"`
}
type Region struct {
	Region     string `json:"region"`
	RegionName string `json:"regionName"`
	EnableIpv6 string `json:"enableIpv6"`
	Azs        []Az   `json:"azs"`
}
type Az struct {
	Region string `json:"region"`
	Az     string `json:"az"`
	AzName string `json:"azName"`
}
