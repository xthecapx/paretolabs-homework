package utils

import (
	"paretolabs-backend/models"
	"path/filepath"
)

func GetUserFilePath() string {
	return filepath.Join(".", "db", "users.json")
}

func TransformToHashMap(users []models.User) map[int]models.User {
	userMap := make(map[int]models.User)
	for _, user := range users {
		userMap[user.FID] = user
	}
	return userMap
}
