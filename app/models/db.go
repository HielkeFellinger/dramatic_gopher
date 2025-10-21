package models

import "database/sql"

var DB *sql.DB

var UserService *userService = &userService{}

type userService struct {
}

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
