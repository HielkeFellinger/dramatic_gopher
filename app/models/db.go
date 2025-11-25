package models

import "database/sql"

var DB *sql.DB

var UserService *userService = &userService{}
var CampaignService *campaignService = &campaignService{}

type userService struct {
}
type campaignService struct{}

func (us *userService) GetUserByUsername(username string) (User, error) {
	user := User{}
	sqlGetUserByUsername := `SELECT * FROM users WHERE name = ?`

	result := DB.QueryRow(sqlGetUserByUsername, username)
	return user, result.Scan(&user.Id, &user.Name, &user.DisplayName, &user.Password)
}

func (us *userService) GetUserById(userId int64) (User, error) {
	user := User{}
	sqlGetUserById := `SELECT * FROM users WHERE id = ?`
	result := DB.QueryRow(sqlGetUserById, userId)
	return user, result.Scan(&user.Id, &user.Name, &user.DisplayName, &user.Password)
}

func (cs *campaignService) GetCampaignById(campaignId int64) (Campaign, error) {
	campaign := Campaign{}
	sqlGetCampaignById := `SELECT * FROM campaigns WHERE id = ?`

	result := DB.QueryRow(sqlGetCampaignById, campaignId)
	return campaign, result.Scan(&campaign.Id, &campaign.Name)
}

func (cs *campaignService) GetCampaignsByUserId(userId int64) ([]Campaign, error) {
	campaigns := []Campaign{}
	sqlGetCampaignsByUserId := `SELECT * FROM campaigns WHERE id IN ( SELECT DISTINCT campaign_id FROM campaign_access WHERE user_id = ? )`

	result := DB.QueryRow(sqlGetCampaignsByUserId, userId)
	return campaigns, result.Scan(&campaigns)
}

func (cs *campaignService) CreateCampaign(campaign Campaign) (Campaign, error) {
	returnCampaign := Campaign{}
	sqlCreateCampaign := `INSERT INTO campaigns (name) VALUES (?) RETURNING *`

	result := DB.QueryRow(sqlCreateCampaign, campaign.Name)
	return returnCampaign, result.Scan(&returnCampaign.Id, &returnCampaign.Name)
}
