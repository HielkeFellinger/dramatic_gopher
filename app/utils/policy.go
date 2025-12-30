package utils

import (
	"strings"

	"github.com/HielkeFellinger/dramatic_gopher/app/config"
)

func CheckIfPasswordMeetsPolicy(password string, mustNotEqual string) bool {
	passTrimmed := strings.TrimSpace(password)

	if len(passTrimmed) < config.CurrentConfig.PasswordMinLength {
		return false
	}
	if strings.ToLower(passTrimmed) == strings.ToLower(strings.TrimSpace(mustNotEqual)) {
		return false
	}

	return true
}

func CheckIfNameMeetsPolicy(username string) bool {
	nameTrimmed := strings.TrimSpace(username)

	if len(nameTrimmed) < 4 {
		return false
	}

	return true
}
