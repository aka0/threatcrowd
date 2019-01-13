package threatcrowd

import (
	"encoding/json"
	"log"
	"net/http"
)

type ThreatCrowdClient struct {
	Client  *http.Client
	BaseURL string
	Delay   int64
}

func NewClient() *ThreatCrowdClient {
	BaseURL := "https://www.threatcrowd.org/searchApi/v2"
	Delay := int64(10)

	client := &ThreatCrowdClient{
		Client:  &http.Client{},
		BaseURL: BaseURL,
		Delay:   Delay,
	}

	return client
}

func (c *ThreatCrowdClient) NewAPIRequest(endpoint string, param string, target string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+endpoint, nil)

	if err != nil {
		return req, err
	}

	q := req.URL.Query()
	q.Add(param, target)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (c *ThreatCrowdClient) Do(r *http.Request, v interface{}) error {

	resp, err := c.Client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&v)

	if err != nil {
		return err
	}

	return nil
}

func (c *ThreatCrowdClient) EmailReport(email string) Email {
	req, err := c.NewAPIRequest("/email/report/", "email", email)

	if err != nil {
		log.Fatalln("unable to build email request")
	}

	emailReport := Email{}
	err = c.Do(req, &emailReport)

	if err != nil {
		log.Fatalln("unable to request email report")
	}

	return emailReport
}

func (c *ThreatCrowdClient) DomainReport(domain string) Domain {
	req, err := c.NewAPIRequest("/domain/report/", "domain", domain)

	if err != nil {
		log.Fatalln("unable to build domain request")
	}

	domainReport := Domain{}
	err = c.Do(req, &domainReport)

	if err != nil {
		log.Fatalln("unable to request domain report")
	}

	return domainReport

}

func (c *ThreatCrowdClient) FileReport(resource string) File {
	req, err := c.NewAPIRequest("/file/report/", "resource", resource)

	if err != nil {
		log.Fatalln("unable to build file request")
	}

	report := File{}
	err = c.Do(req, &report)

	if err != nil {
		log.Fatalln("unable to request file report")
	}

	return report
}

func (c *ThreatCrowdClient) AntiVirusReport(antivirus string) AntiVirus {
	req, err := c.NewAPIRequest("/antivirus/report/", "antivirus", antivirus)

	if err != nil {
		log.Fatalln("unable to build antivirus request")
	}

	report := AntiVirus{}
	err = c.Do(req, &report)

	if err != nil {
		log.Fatalln("unable to request antivirus report")
	}

	return report
}

func (c *ThreatCrowdClient) IPReport(ip string) IP {

	req, err := http.NewRequest(http.MethodGet, c.BaseURL+"/ip/report/", nil)

	if err != nil {
		log.Fatalln("error building query")
	}

	q := req.URL.Query()
	q.Add("ip", ip)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("error querying %s", req.URL.String())
	}
	defer resp.Body.Close()

	ipReport := IP{}
	err = json.NewDecoder(resp.Body).Decode(&ipReport)

	if err != nil {
		log.Fatalln("error reading response")
	}

	return ipReport
}
