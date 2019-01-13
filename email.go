package threatcrowd

type Email struct {
	Domains      []string      `json:"domains"`
	Permalink    string        `json:"permalink"`
	References   []interface{} `json:"references"`
	ResponseCode string        `json:"response_code"`
}
