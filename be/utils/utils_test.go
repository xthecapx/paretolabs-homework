package utils

import (
	"paretolabs-backend/models"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserFilePath(t *testing.T) {
	expectedPath := filepath.Join(".", "db", "users.json")
	actualPath := GetUserFilePath()

	assert.Equal(t, expectedPath, actualPath)
}

func TestTransformToHashMap(t *testing.T) {
	users := []models.User{
		{FID: 1, Username: "user1"},
		{FID: 2, Username: "user2"},
		{FID: 3, Username: "user3"},
	}

	expectedMap := map[int]models.User{
		1: {FID: 1, Username: "user1"},
		2: {FID: 2, Username: "user2"},
		3: {FID: 3, Username: "user3"},
	}

	actualMap := TransformToHashMap(users)

	assert.Equal(t, expectedMap, actualMap)
}
