package terminal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sarracesi/user-service/user-service/env"
)

func ShowMenu() {
	fmt.Println("Select an operation:")
	fmt.Println("1. Get all users")
	fmt.Println("2. Add a new user")
	fmt.Println("3. Update a user")
	fmt.Println("4. Delete a user")
	fmt.Println("5. Get user by ID")
	fmt.Println("0. Exit")
}

func GetAllUsersAndPrint() {
	fmt.Println("Getting all users:")

	resp, err := http.Get("http://localhost:3333/users")
	if err != nil {
		log.Fatal("Error getting users:", err)
	}
	defer resp.Body.Close()

}

func GetUserByIDAndPrint() {
	fmt.Print("Enter the ID of the user you want to retrieve: ")
	var userID string
	fmt.Scan(&userID)

	resp, err := http.Get(fmt.Sprintf("http://localhost:3333/users?id=%s", userID))
	if err != nil {
		log.Fatal("Error getting user:", err)
	}
	defer resp.Body.Close()

}

func AddNewUserAndPrint() {
	fmt.Println("Enter the details for the new user:")
	var newUser env.User

	fmt.Print("Last Name: ")
	fmt.Scan(&newUser.LastName)
	fmt.Print("Username: ")
	fmt.Scan(&newUser.UserName)

	// Convert user data to JSON
	userData, err := json.Marshal(newUser)
	if err != nil {
		log.Fatal("Error marshalling user data:", err)
	}

	resp, err := http.Post("http://localhost:3333/users", "application/json", bytes.NewBuffer(userData))
	if err != nil {
		log.Fatal("Error sending POST request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User added successfully.")
	} else {
		fmt.Println("Error adding user. Status code:", resp.StatusCode)
	}
}

func UpdateExistingUserAndPrint() {
	fmt.Print("Enter the ID of the user you want to update: ")
	var userID string
	fmt.Scan(&userID)

	fmt.Println("Enter the updated details for the user:")
	var updatedUser env.User
	fmt.Print("Last Name: ")
	fmt.Scan(&updatedUser.LastName)
	fmt.Print("Username: ")
	fmt.Scan(&updatedUser.UserName)
	fmt.Print("id: ")
	fmt.Scan(&updatedUser.ID)

	jsonData, err := json.Marshal(updatedUser)
	if err != nil {
		log.Fatal("Error marshaling JSON data:", err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:3333/users?id=%s", userID), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error creating PUT request:", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error updating user:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func DeleteUserByIDAndPrint() {
	fmt.Print("Enter the ID of the user you want to delete: ")
	var userID string
	fmt.Scan(&userID)

	// Perform DELETE request to delete a user by ID
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:3333/users?id=%s", userID), nil)
	if err != nil {
		log.Fatal("Error creating DELETE request:", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error deleting user:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func Print() {
	for {
		ShowMenu()
		var choice int
		fmt.Print("Enter your choice (0-5): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			GetAllUsersAndPrint()
		case 2:
			AddNewUserAndPrint()
		case 3:
			UpdateExistingUserAndPrint()
		case 4:
			DeleteUserByIDAndPrint()
		case 5:
			GetUserByIDAndPrint()
		case 0:
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 0 and 5.")
		}
	}
}
