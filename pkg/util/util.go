package util

import "github.com/google/uuid"

func UUIDStr() string {
	return uuid.New().String()
}
