package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	err = json.Unmarshal(requestBody, &user)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	userSaved, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerifyPassword(userSaved.Password, user.Password)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userSaved.Id)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(token)

	responses.JSON(w, http.StatusOK, token)
}
