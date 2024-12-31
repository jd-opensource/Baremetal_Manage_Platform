package object

import (
	"encoding/json"
	"fmt"
)

type CMDBObjectJD struct {
	Asset             string   `json:"asset,omitempty"`
	Cabinet           string   `json:"cabinet,omitempty"`
	Configure         string   `json:"configure,omitempty"`
	Cost              string   `json:"cost,omitempty"`
	CreatedAt         string   `json:"created_at,omitempty"`
	Department        string   `json:"department,omitempty"`
	DeviceName        string   `json:"device_name,omitempty"`
	DeviceSn          string   `json:"device_sn,omitempty"`
	EthGateway        string   `json:"eth_gateway,omitempty"`
	EthIP             string   `json:"eth_ip,omitempty"`
	EthMac            string   `json:"eth_mac,omitempty"`
	EthMask           string   `json:"eth_mask,omitempty"`
	EthSwitchIP       string   `json:"eth_switch_ip,omitempty"`
	EthSwitchPort     string   `json:"eth_switch_port,omitempty"`
	Idc               string   `json:"idc,omitempty"`
	IdcID             int      `json:"idc_id,omitempty"`
	IloGateway        string   `json:"ilo_gateway"`
	IloIP             string   `json:"ilo_ip"`
	IloMac            string   `json:"ilo_mac"`
	IloMask           string   `json:"ilo_mask"`
	IloSwitchIP       string   `json:"ilo_switch_ip"`
	IloSwitchPort     string   `json:"ilo_switch_port"`
	MaintenanceEnd    string   `json:"maintenance_end"`
	MaintenanceStart  string   `json:"maintenance_start"`
	MaintenanceStatus int      `json:"maintenance_status"`
	Manufacturer      string   `json:"manufacturer"`
	Model             string   `json:"model"`
	Monitor           string   `json:"monitor"`
	Nodeid            int      `json:"nodeid"`
	OpAssetNumber     string   `json:"op_asset_number"`
	OpGroup           []string `json:"op_group"`
	Operation         string   `json:"operation"`
	OrganizationCode  string   `json:"organization_code"`
	OsKernel          string   `json:"os_kernel"`
	OsVersion         string   `json:"os_version"`
	Package           string   `json:"package"`
	Sn                string   `json:"sn"`
	Type              string   `json:"type"`
	UPosition         string   `json:"u_position"`
	UpDown            string   `json:"up_down"`
}

type CMDBObjectJDCloud struct {
	Cabinet struct {
		Code  string `json:"code,omitempty"`
		ID    int    `json:"id,omitempty"`
		IdcID int    `json:"idc_id,omitempty"`
	} `json:"cabinet"`
	DeviceSn string `json:"device_sn,omitempty"`
	Eth      struct {
		DeviceSn   string `json:"device_sn,omitempty"`
		Gateway    string `json:"gateway,omitempty"`
		IP         string `json:"ip,omitempty"`
		Mac        string `json:"mac,omitempty"`
		Mask       string `json:"mask,omitempty"`
		SwitchIP   string `json:"switch_ip,omitempty"`
		SwitchPort string `json:"switch_port,omitempty"`
		Type       string `json:"type,omitempty"`
	} `json:"eth,omitempty"`
	Idc struct {
		CnName string `json:"cn_name,omitempty"`
		ID     int    `json:"id,omitempty"`
	} `json:"idc,omitempty"`
	Ilo struct {
		DeviceSn   string `json:"device_sn,omitempty"`
		Gateway    string `json:"gateway,omitempty"`
		IP         string `json:"ip,omitempty"`
		Mac        string `json:"mac,omitempty"`
		Mask       string `json:"mask,omitempty"`
		SwitchIP   string `json:"switch_ip,omitempty"`
		SwitchPort string `json:"switch_port,omitempty"`
		Type       string `json:"type,omitempty"`
	} `json:"ilo,omitempty"`
	IloGateway    string `json:"ilo_gateway,omitempty"`
	IloIP         string `json:"ilo_ip,omitempty"`
	IloMac        string `json:"ilo_mac,omitempty"`
	IloMask       string `json:"ilo_mask,omitempty"`
	IloSwitchIP   string `json:"ilo_switch_ip,omitempty"`
	IloSwitchPort string `json:"ilo_switch_port,omitempty"`
	Manufacturer  string `json:"manufacturer,omitempty"`
	Model         string `json:"model,omitempty"`
	OpGroup       []struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"op_group,omitempty"`
	Organization struct {
		Fullname         string `json:"fullname,omitempty"`
		ID               int    `json:"id,omitempty"`
		OrganizationCode string `json:"organization_code,omitempty"`
	} `json:"organization,omitempty"`
	Package string `json:"package,omitempty"`
	Pub     struct {
		DeviceSn   string `json:"device_sn,omitempty"`
		Gateway    string `json:"gateway,omitempty"`
		IP         string `json:"ip,omitempty"`
		Mac        string `json:"mac,omitempty"`
		Mask       string `json:"mask,omitempty"`
		SwitchIP   string `json:"switch_ip,omitempty"`
		SwitchPort string `json:"switch_port,omitempty"`
		Type       string `json:"type,omitempty"`
	} `json:"pub,omitempty"`
	UEnd   int `json:"u_end,omitempty"`
	UStart int `json:"u_start,omitempty"`
}

// CPSRecord CPS information
type CPSRecord struct {
	Clazz        string `json:"clazz,omitempty"`
	SN           string `json:"sn,omitempty"`
	IdcId        string `json:"idc_id,omitempty"`
	InstanceID   string `json:"instance_id,omitempty"`
	InstanceName string `json:"instance_name,omitempty"`
	Pin          string `json:"pin,omitempty"`
	Status       string `json:"manage_status,omitempty"`
	IloIP        string `json:"ilo_ip,omitempty"`
	Brand        string `json:"brand,omitempty"`
	Model        string `json:"model,omitempty"`
	Package      string `json:"package,omitempty"`
	Cabinet      string `json:"cabinet,omitempty"`
	UPosition    string `json:"u_position,omitempty"`
	Ipv4         string `json:"ipv4"`
	DeviceType   string `json:"device_type"`
	IloUser      string `json:"ilo_user"`
	IloPass      string `json:"ilo_pass"`
}

// CPSObject store CPS list per IDC ID
type CPSObject struct {
	Idc  string      `json:"idc,omitempty"`
	CPSs []CPSRecord `json:"cps,omitempty"`
}

//JSON  return JSON string
func (c CPSObject) JSON() (string, error) {

	json, err := json.Marshal(c.CPSs)
	if err != nil {
		fmt.Println("CPSObject to JSON error:" + err.Error())
		return "", err
	}
	return string(json), nil
}

// JSON return JSON string
func (c CPSRecord) JSON() (string, error) {

	json, err := json.Marshal(c)
	if err != nil {
		fmt.Println("CPSRecord to JSON error:" + err.Error())
		return "", err
	}
	return string(json), nil
}
