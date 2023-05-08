package services

import (
	"log"
	"os"
	"paretolabs-backend/models"
	"paretolabs-backend/utils"
)

func GetUsersList() (map[int]models.User, []models.User) {
	if _, err := os.Stat(utils.GetUserFilePath()); os.IsNotExist(err) {
		// users.json does not exist, fetch data and write to file
		userListMap, userList, err := GetDataAndWriteToFile(1, 10)
		if err != nil {
			log.Fatalf("Failed to fetch data and write to file: %v", err)
		}

		return userListMap, userList
	}
	// users.json exists, read data from file
	userListMap, userList, err := ReadUsersFromFile()
	if err != nil {
		log.Fatalf("Failed to read users from file: %v", err)
	}

	return userListMap, userList
}
