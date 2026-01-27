package models

type Campaign struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	Password    string `json:"-"`
}

type CampaignAccess struct {
	Id          string `json:"id"`
	Role        string `json:"role"`
	UserId      string `json:"user_id"`
	DisplayName string `json:"display_name"`
	CampaignId  string `json:"campaign_id"`
}

type CampaignToData struct {
	Id         string `json:"id"`
	CampaignId string `json:"campaign_id"`
	DataDir    string `json:"data_dir"`
}
