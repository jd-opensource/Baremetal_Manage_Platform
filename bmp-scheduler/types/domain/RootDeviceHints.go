package domain

type RootDeviceHints struct { //空值传给agent会err，所以设置omitempty
	Size int64 `json:"size,omitempty"`

	Rotational bool `json:"rotational,omitempty"`

	Wwn string `json:"wwn,omitempty"`

	Name string `json:"name,omitempty"`

	WwnVendorExtension string `json:"wwn_vendor_extension,omitempty"`

	WwnWithExtension string `json:"wwn_with_extension,omitempty"`

	Model string `json:"model,omitempty"`

	Serial string `json:"serial,omitempty"`

	Vendor string `json:"vendor,omitempty"`

	Hctl string `json:"hctl,omitempty"`

	ByPath string `json:"by_path,omitempty"`
}
