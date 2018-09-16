package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jaylevin/TMN-API/database"
	"github.com/jaylevin/TMN-API/errs"
	"github.com/jaylevin/TMN-API/models"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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
	id := vars["id"]

	var table models.User
	user, err := database.GetRecord(db, &table, id)
	if err != nil {
		return errs.Stack(err)
	}

	resp, err := json.Marshal(&user)
	if err != nil {
		return errs.New(400, "Error marshalling user into JSON!: "+err.Error())
	}

	recorder.Write(resp)
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

	queries, ok := r.URL.Query()["name"]
	if ok && queries != nil{
		names := strings.Split(queries[0], ",")
		log.Println(names)
		db.Where("name in (?)", names).Find(&users)
	} else {
		db.Find(&users)
	}

	json, err := json.Marshal(&users)
	if err != nil {
		return errs.New(400, "Error marshalling users into JSON!: "+err.Error())
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
		return errs.New(400, "Bad request, error decoding JSON request: "+err.Error())
	}
	defer r.Body.Close()

	result := db.Create(&user)
	if result.Error != nil {
		return errs.New(400, "Bad request, check the request's data")
	}

	json, err := json.Marshal(&user)
	if err != nil {
		return errs.New(400, "Error marshalling users into JSON!: "+err.Error())
	}

	recorder.WriteHeader(201)
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

	var table models.User
	user, err := database.GetRecord(db, &table, id)
	if err != nil {
		return err
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return errs.New(400, "Error decoding JSON request body: "+err.Error())
	}
	defer r.Body.Close()

	result := db.Save(user)
	if result.Error != nil {
		return errs.New(400, result.Error.Error())
	}

	json, err := json.Marshal(&user)
	if err != nil {
		return errs.New(400, "Error marshalling users into JSON!: "+err.Error())
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

	var table models.User
	user, err := database.GetRecord(db, &table, id)
	if err != nil {
		return errs.Stack(err)
	}

	result := db.Delete(user)
	if result.Error != nil {
		errs.New(400, "Bad request")
	}

	recorder.WriteHeader(204)
	return nil
}
