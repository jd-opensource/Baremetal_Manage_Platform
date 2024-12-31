package object

import (
	"encoding/json"
	"fmt"
)

// FRUObject type of IPMI FRU Information
type FRUObject struct {
	Manufacturer string
	SN           string
	Model        string
}

// JSON return json string
func (l FRUObject) JSON() (string, error) {

	json, err := json.Marshal(l)
	if err != nil {
		fmt.Println("FRUObject to JSON error:" + err.Error())
		return "", err
	}
	return string(json), nil
}
