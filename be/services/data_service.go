package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"paretolabs-backend/models"
	"paretolabs-backend/utils"
)

func GetDataAndWriteToFile(startFID, endFID int) ([]models.User, error) {
	var userList []models.User

	// Iterate over the fid range
	for fid := startFID; fid <= endFID; fid++ {
		fmt.Printf("Getting data from user: %d\n", fid)
		// Build the URL with the current fid value
		url := fmt.Sprintf("https://api.farcaster.xyz/v2/user?fid=%d", fid)

		// Send a GET request to the URL to fetch the data
		response, err := http.Get(url)
		if err != nil {
			log.Printf("Failed to fetch data for fid %d: %v", fid, err)
			continue // Move to the next iteration
		}
		defer response.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Failed to read response body for fid %d: %v", fid, err)
			continue // Move to the next iteration
		}

		// Parse the received data into the struct
		var data models.HttpResponseData
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Printf("Failed to parse data into struct for fid %d: %v", fid, err)
			continue // Move to the next iteration
		}

		// Append the user information to the list
		fmt.Printf("User data saved: %s\n", data.Result.User.DisplayName)
		userList = append(userList, data.Result.User)

		// Sleep for 200 milliseconds before the next iteration
		time.Sleep(200 * time.Millisecond)
	}

	// Write the user list to a JSON file
	fileData, err := json.MarshalIndent(userList, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal user list: %v", err)
	}

	err = ioutil.WriteFile(utils.GetUserFilePath(), fileData, 0644) // Replace with the desired file path
	if err != nil {
		log.Fatalf("Failed to write data to file: %v", err)
	}

	fmt.Println("Data retrieved and written to file successfully")

	return ReadUsersFromFile()
}
