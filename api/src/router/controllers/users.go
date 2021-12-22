package controllers

import (
	"api/db"
	"api/repositories"
	"api/src/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	err = json.Unmarshal(requestBody, &user)
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewUserRepository(db)
	userId, err := repo.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserted: %d", userId)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching users..."))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search user..."))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user..."))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user..."))
}
