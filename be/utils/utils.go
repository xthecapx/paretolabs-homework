package utils

import (
	"path/filepath"
)

func GetUserFilePath() string {
	return filepath.Join(".", "db", "users.json")
}
