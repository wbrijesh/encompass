package types

type Visitor struct {
	Id              string `json:"id"`
	Browser         string `json:"browser"`
	OperatingSystem string `json:"operating_system"`
	Resolution      string `json:"resolution"`
	Device          string `json:"device"`
	Language        string `json:"language"`
	Country         string `json:"country"`
	WebsiteId       string `json:"website_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
