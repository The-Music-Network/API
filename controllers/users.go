package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jaylevin/TMN-API/models"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"net/http/httptest"
)

func ShowUser(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error {

	vars := mux.Vars(r)
	idStr := vars["id"]

	var user models.User
	db.Find(&user, idStr)

	json, err := json.Marshal(&user)
	if err != nil {
		return errors.New("Error marshalling users into JSON!: " + err.Error())
	}

	recorder.Write(json)
	return nil
}

func GetUsers(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error {

	users := make([]models.User, 0)
	db.Find(&users)

	json, err := json.Marshal(&users)
	if err != nil {
		return errors.New("Error marshalling users into JSON!: " + err.Error())
	}

	recorder.Write(json)
	return nil
}

func CreateUser(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return errors.New("Error decoding JSON data: " + err.Error())
	}
	defer r.Body.Close()

	result := db.Create(&models.User{Name: "Test12"})
	json, err := json.Marshal(&result)
	if err != nil {
		log.Println("Error marshalling models.User into JSON!:", err)
	}

	recorder.Write(json)
	return nil
}

func UpdateUser(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error converting string to int:", err)
	}
	user := models.User{
		Id: uint(id),
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return errors.New("Error decoding JSON data: " + err.Error())
	}
	defer r.Body.Close()

	result := db.Save(&user)
	json, err := json.Marshal(&result)
	if err != nil {
		log.Println("Error marshalling models.User into JSON!:", err)
	}

	recorder.Write(json)
	return nil
}

func DeleteUser(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error {

	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	db.First(&user, id) // find the user with the given ID from http URL parameter (http://localhost:8080/users/{id})

	result := db.Delete(&user)
	json, err := json.Marshal(&result)
	if err != nil {
		log.Println("Error marshalling models.User into JSON!:", err)
	}

	recorder.Write(json)
	return nil
}
