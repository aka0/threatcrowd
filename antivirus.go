package threatcrowd

type AntiVirus struct {
	Hashes       []string      `json:"hashes"`
	Permalink    string        `json:"permalink"`
	References   []interface{} `json:"references"`
	ResponseCode string        `json:"response_code"`
}
