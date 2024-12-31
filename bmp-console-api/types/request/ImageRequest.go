package request

type QueryImagesByDeviceTypeRequest struct {
	// Pin        string `json:"pin"`
	IdcID        string `json:"idc_id"`
	DeviceTypeID string `json:"deviceTypeId"`
	// OsType     string   `json:"os_type"`
	// OsVersion  string   `json:"os_version"`
	// OsID       string   `json:"os_id"`
	// Source     string   `json:"source"`
	// ImageIds   []string `json:"image_ids"`
}
