package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"gohttpsrv/mockdb"
	"io/ioutil"
	"log"
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	// read req body
	var user mockdb.User
	json.Unmarshal(reqBody, &user)
	user.Id = uuid.New().String()

	// read existing users & append new user
	users := mockdb.ReadUsersFromDb()
	users = append(users, user)
	mockdb.WriteUsersToDb(users)

	enc := json.NewEncoder(w)
	enc.Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := mockdb.ReadUsersFromDb()
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	users := mockdb.ReadUsersFromDb()
	ids, _ := r.URL.Query()["id"]

	var user mockdb.User
	for _, usr := range users {
		if usr.Id == ids[0] {
			user = usr
			break
		}
	}
	json.NewEncoder(w).Encode(user)
}


func updateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	// read req body
	var userToUpdate mockdb.User
	json.Unmarshal(reqBody, &userToUpdate)

	// read existing usersFromDb & append new userToUpdate
	usersFromDb := mockdb.ReadUsersFromDb()

	// update userToUpdate
	var usersPostUpdate []mockdb.User
	var updatedUsr mockdb.User
	for _, usr := range usersFromDb {
		if usr.Id == userToUpdate.Id {
			usr.Name = userToUpdate.Name
			usr.Phone = userToUpdate.Phone
			updatedUsr = usr
		}
		usersPostUpdate = append(usersPostUpdate, usr)
	}

	// write userToUpdate to db
	mockdb.WriteUsersToDb(usersPostUpdate)
	json.NewEncoder(w).Encode(updatedUsr)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	usersFromDb := mockdb.ReadUsersFromDb()

	var usersPostDelete []mockdb.User
	// update userToUpdate

	for _, usr := range usersFromDb {
		if usr.Id != userId {
			usersPostDelete = append(usersPostDelete, usr)
		}
	}

	mockdb.WriteUsersToDb(usersPostDelete)
	w.Write(nil)
}

func handleRequests() {
	http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/getUsers", getAllUsers)
	http.HandleFunc("/getUser", getUser)
	http.HandleFunc("/updateUser", updateUser)
	http.HandleFunc("/deleteUser", deleteUser)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
