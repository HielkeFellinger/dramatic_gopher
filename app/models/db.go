package models

import (
	"database/sql"

	"github.com/HielkeFellinger/dramatic_gopher/app/utils"
)

var DB *sql.DB

var UserService *userService = &userService{}
var CampaignService *campaignService = &campaignService{}

type userService struct {
}
type campaignService struct{}

func (us *userService) GetUserByUsername(username string) (User, error) {
	user := User{}
	row := DB.QueryRow(`SELECT * FROM users WHERE name = ?;`, username)
	return user, row.Scan(&user.Id, &user.Name, &user.DisplayName, &user.Password, &user.Role)
}

func (us *userService) GetUserById(userId int64) (User, error) {
	user := User{}
	row := DB.QueryRow(`SELECT * FROM users WHERE id = ?;`, userId)
	return user, row.Scan(&user.Id, &user.Name, &user.DisplayName, &user.Password, &user.Role)
}

func (us *userService) InsertUser(user User) (User, error) {
	encryptPass, encryptErr := utils.HashPassword(user.Password)
	if encryptErr != nil {
		return User{}, encryptErr
	}

	returnUser := User{}
	sqlInsertUser := `INSERT INTO users(name, display_name, role, password) VALUES (?, ?, 'user', ?) RETURNING *;`
	row := DB.QueryRow(sqlInsertUser, user.Name, user.DisplayName, encryptPass)
	return returnUser, row.Scan(&user.Id, &user.Name, &user.DisplayName, &user.Password, &user.Role)
}

func (cs *campaignService) GetCampaignById(campaignId int64) (Campaign, error) {
	campaign := Campaign{}
	result := DB.QueryRow(`SELECT * FROM campaigns WHERE id = ?;`, campaignId)
	return campaign, result.Scan(&campaign.Id, &campaign.Name, &campaign.State, &campaign.Password)
}

func (cs *campaignService) GetCampaignsByUserId(userId int64) ([]Campaign, error) {
	var campaigns []Campaign
	sqlGetCampaignsByUserId := `SELECT * FROM campaigns WHERE id IN ( SELECT DISTINCT campaign_id FROM campaign_access WHERE user_id = ? );`

	rows, queryErr := DB.Query(sqlGetCampaignsByUserId, userId)
	defer rows.Close()
	if queryErr != nil {
		return campaigns, queryErr
	}

	for rows.Next() {
		campaign := Campaign{}
		if rowErr := rows.Scan(&campaign.Id, &campaign.Name, &campaign.State, &campaign.Password); rowErr != nil {
			return campaigns, rowErr
		}
		campaigns = append(campaigns, campaign)
	}
	return campaigns, rows.Err()
}

func (cs *campaignService) InsertCampaign(campaign Campaign) (Campaign, error) {
	encryptPass, encryptErr := utils.HashPassword(campaign.Password)
	if encryptErr != nil {
		return Campaign{}, encryptErr
	}

	returnCampaign := Campaign{}
	sqlInsertCampaign := `INSERT INTO campaigns (name, password) VALUES (?, ?) RETURNING *;`
	row := DB.QueryRow(sqlInsertCampaign, campaign.Name, string(encryptPass))
	return returnCampaign, row.Scan(&returnCampaign.Id, &returnCampaign.Name, &returnCampaign.State, &returnCampaign.Password)
}
