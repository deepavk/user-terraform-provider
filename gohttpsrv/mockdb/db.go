package mockdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var jsonfile = "../gohttpsrv/mockdb/users.json"

func ReadUsersFromDb() []User {
	jsonFile, err := os.Open(jsonfile)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []User
	json.Unmarshal(byteValue, &users)
	return users
}

func WriteUsersToDb(users []User) {
	file, _ := json.MarshalIndent(users, "", " ")
	err := ioutil.WriteFile(jsonfile, file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
