package models

type Campaign struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CampaignAccess struct {
	Id         string `json:"id"`
	Role       string `json:"role"`
	UserId     string `json:"user_id"`
	CampaignId string `json:"campaign_id"`
}
