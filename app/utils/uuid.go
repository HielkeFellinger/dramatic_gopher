package utils

import "github.com/google/uuid"

func ParseStringToUuid(uuidAsString string) (uuid.UUID, error) {
	parsedUuid, err := uuid.Parse(uuidAsString)
	if err == nil {
		return parsedUuid, err
	}
	return uuid.UUID{}, err
}
