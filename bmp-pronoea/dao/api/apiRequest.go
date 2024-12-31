package api

type IBaseApiRequest interface {
	Decode() (map[string]interface{}, error)
}