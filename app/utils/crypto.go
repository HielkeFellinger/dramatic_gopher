package utils

import (
	"strconv"

	"github.com/HielkeFellinger/dramatic_gopher/app/config"
	"golang.org/x/crypto/bcrypt"
)

var minCryptoCost = 16

// HashPassword Hash the models.User password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), getCryptoCost())
}

func getCryptoCost() int {
	// Get env settings
	envCost := config.CurrentConfig.CryptCost

	envIntCost, err := strconv.ParseInt(envCost, 10, 8)
	if err != nil || int(envIntCost) < minCryptoCost {
		return minCryptoCost
	}
	return int(envIntCost)
}
