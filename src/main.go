package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User - Our struct for all users
type User struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

type Error struct {
	Code    int    `json:"Code"`
	Message string `json:"ErrorMessage"`
}

// Contains all users, acts as dummy DB
var Users []User

func allUsers(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Endpoint Hit: allUsers")
	json.NewEncoder(w).Encode(Users)
}

func retrieveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, user := range Users {
		if user.Id == id {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new User struct
	// append this to our Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	if isUserPresent(user.Id) {
		var error = error(500, "User with id "+strconv.Itoa(user.Id)+" is already present")
		json.NewEncoder(w).Encode(error)
	} else {
		// update our global Users array to include our new User
		Users = append(Users, user)

		json.NewEncoder(w).Encode(user)
	}
}

func error(code int, message string) Error {
	var error Error
	error.Code = code
	error.Message = message
	return error
}

func deleteUser(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, user := range Users {
		if user.Id == id {
			Users = append(Users[:index], Users[index+1:]...)
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if isUserPresent(id) {
		// Remove from the list
		for index, user := range Users {
			if user.Id == id {
				Users = append(Users[:index], Users[index+1:]...)
			}
		}

		reqBody, _ := ioutil.ReadAll(r.Body)
		var user User
		json.Unmarshal(reqBody, &user)
		user.Id = id
		Users = append(Users, user)
	} else {
		var error = error(500, "User with id "+strconv.Itoa(id)+" is not present")
		json.NewEncoder(w).Encode(error)
	}
}

func isUserPresent(userId int) bool {
	for _, user := range Users {
		if user.Id == userId {
			return true
		}
	}
	return false
}

func handleRequests() {
	userRouter := mux.NewRouter().StrictSlash(true)
	userRouter.HandleFunc("/users", allUsers)
	userRouter.HandleFunc("/user", createUser).Methods("POST")
	userRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	userRouter.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	userRouter.HandleFunc("/user/{id}", retrieveUser)
	log.Fatal(http.ListenAndServe(":8001", userRouter))
}

func main() {
	Users = []User{
		{Id: 1, Name: "First User", Address: "First Country"},
		{Id: 2, Name: "Second User", Address: "Second Country"},
	}
	handleRequests()
}
