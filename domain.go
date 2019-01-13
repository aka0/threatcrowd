package threatcrowd

type Domain struct {
	Emails      []string      `json:"emails"`
	Hashes      []interface{} `json:"hashes"`
	Permalink   string        `json:"permalink"`
	References  []string      `json:"references"`
	Resolutions []struct {
		IPAddress    string `json:"ip_address"`
		LastResolved string `json:"last_resolved"`
	} `json:"resolutions"`
	ResponseCode string   `json:"response_code"`
	Subdomains   []string `json:"subdomains"`
	Votes        int64    `json:"votes"`
}
