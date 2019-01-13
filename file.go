package threatcrowd

type File struct {
	Domains      []string      `json:"domains"`
	Ips          []string      `json:"ips"`
	Md5          string        `json:"md5"`
	Permalink    string        `json:"permalink"`
	References   []interface{} `json:"references"`
	ResponseCode string        `json:"response_code"`
	Scans        []string      `json:"scans"`
	Sha1         string        `json:"sha1"`
}
