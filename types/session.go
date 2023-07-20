package types

type Session struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	VisitorId string `json:"visitorId"`
}
