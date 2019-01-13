package threatcrowd

type IP struct {
	ResponseCode string `json:"response_code"`
	Resolutions  []struct {
		LastResolved string `json:"last_resolved"`
		Domain       string `json:"domain"`
	} `json:"resolutions"`
	Hashes     []string      `json:"hashes"`
	References []interface{} `json:"references"`
	Votes      int           `json:"votes"`
	Permalink  string        `json:"permalink"`
}
