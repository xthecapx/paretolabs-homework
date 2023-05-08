package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"paretolabs-backend/models"
	"paretolabs-backend/utils"
)

func ReadUsersFromFile() (map[int]models.User, []models.User, error) {
	fileData, err := ioutil.ReadFile(utils.GetUserFilePath())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read users.json file: %v", err)
	}

	var userList []models.User
	err = json.Unmarshal(fileData, &userList)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal users.json data: %v", err)
	}
	userHasMap := utils.TransformToHashMap(userList)

	return userHasMap, userList, nil
}
