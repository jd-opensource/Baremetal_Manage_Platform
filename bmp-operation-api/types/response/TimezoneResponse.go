package response

import "bytes"

type GetTimezoneListResponse struct {
	TimeZones OrderedJsonDict `json:"timezones"`
}

type OrderedJsonDict struct {
	M    map[string]string
	Keys []string
}

func (o OrderedJsonDict) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.WriteByte('{')
	for i, k := range o.Keys {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte('"')
		buf.WriteString(k)
		buf.WriteByte('"')
		buf.WriteByte(':')
		buf.WriteByte('"')
		buf.Write([]byte(o.M[k]))
		buf.WriteByte('"')
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}
