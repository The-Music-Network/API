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


/*********************************************************************************************************************
 * Controller Action: ShowUser
 * -----------------------------
 * HTTP Method: GET
 * API Endpoint: /users/{id}
 *

 * Description:
 * ------------
 * Returns a single JSON object representing the user with id of {id}
 *
 * Example:
 * --------
 * curl -X GET http://localhost:8080/users/1. This will show the JSON representation of the user with ID=1

 *********************************************************************************************************************/
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

/*********************************************************************************************************************
 * Controller Action: GetUsers
 * -----------------------------
 * HTTP Method: GET
 * API Endpoint: /users
 *

 * Description:
 * ------------
 * Returns an array of all the users in the table
 *
 * Example:
 * --------
 * curl -X GET http://localhost:8080/users

 *********************************************************************************************************************/
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

/*********************************************************************************************************************
 * Controller Action: CreateUser
 * -----------------------------
 * HTTP Method: POST
 * API Endpoint: /users
 *

 * Description:
 * ------------
 * Creates a new user with the given JSON data in the request body
 *
 * Example:
 * --------
 * curl -X POST http://localhost:8080/users -d '{"Name": "Test"}'. This will create a new user with name "Test".
 * Note: the ID of the user is automatically indexed and handled by the ORM.

 *********************************************************************************************************************/
func CreateUser(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return errors.New("Error decoding JSON request body: " + err.Error())
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

/*********************************************************************************************************************
 * Controller Action: UpdateUser
 * -----------------------------
 * HTTP Method: PUT
 * API Endpoint: /users/{id}
 *

 * Description:
 * ------------
 * Updates the user with given {id}, any fields included in the json request body will be overwritten on the given user.
 *
 * Example:
 * --------
 * curl -X PUT http://localhost:8080/users/2 -d '{"Name": "Test1234"}' will only update the Name field of the user.

 *********************************************************************************************************************/
func UpdateUser(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error converting string to integer:", err)
	}
	user := models.User {
		Id: uint(id),
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return errors.New("Error decoding JSON request body: " + err.Error())
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


/*********************************************************************************************************************
 * Controller Action: DeleteUser
 * -----------------------------
 * HTTP Method: DELETE
 * API Endpoint: /users/{id}

 * Description:
 * ------------
 * Finds the user with {id}, deletes them and returns a json response indicating any errors

 * Example:
 * --------
 * curl -X DELETE http://localhost:8080/users/2 This will delete the user with an ID of 2.

 *********************************************************************************************************************/
func DeleteUser(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error {

	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	db.First(&user, id)

	result := db.Delete(&user)
	json, err := json.Marshal(&result)
	if err != nil {
		log.Println("Error marshalling models.User into JSON!:", err)
	}

	recorder.Write(json)
	return nil
}
