package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
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
