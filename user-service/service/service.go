package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sarracesi/user-service/user-service/env"
	"github.com/sarracesi/user-service/user-service/terminal"
)

const portNum string = ":3333"

var users = env.Users{
	env.User{ID: "1", LastName: "khechine", UserName: "sarrak"},
	env.User{ID: "2", LastName: "lalala", UserName: "blabla"},
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsers(w, r)
	case http.MethodPost:
		AddUser(w, r)
	case http.MethodPut:
		UpdateUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET request received for all users")

	id := r.URL.Query().Get("id")
	if id != "" {
		var foundUser *env.User
		for _, user := range users {
			if user.ID == id {
				foundUser = &user
				break
			}
		}

		if foundUser != nil {
			json.NewEncoder(w).Encode(foundUser)
		} else {
			http.Error(w, "User not found", http.StatusNotFound)
		}
	} else {
		
		fmt.Println("Users:", users)
		json.NewEncoder(w).Encode(users)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser env.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newUser.ID = fmt.Sprintf("%d", len(users)+1)
	users = append(users, newUser)

	fmt.Println("User added:", newUser)
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

	var index int
	for i, user := range users {
		if user.ID == userID {
			index = i
			break
		}
	}

	if index < len(users) {
		var updatedUser env.User
		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		users[index] = updatedUser

		json.NewEncoder(w).Encode(updatedUser)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
	userID := r.URL.Query().Get("id")

	var index int
	for i, user := range users {
		if user.ID == userID {
			index = i
			break
		}
	}

	if index < len(users) {
		users = append(users[:index], users[index+1:]...)
		fmt.Fprintf(w, "User deleted successfully")
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}



func Start() {
	log.Println("Starting our simple http server.")
	http.HandleFunc("/", Home)
	http.HandleFunc("/users", handleUsers)

	go func() {
		err := http.ListenAndServe(portNum, nil)
		if err != nil {
			log.Fatal("Error http server on port ", portNum, err)
		}
	}()

	terminal.Print()
}
