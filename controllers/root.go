package controllers

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
)

// This is unnecessary, but makes things easier for JSON stuff.
type RootResponse struct {
	Msg string
}

/*********************************************************************************************************************
 * Controller Action: GetRoot
 * -----------------------------
 * HTTP Method: GET
 * API Endpoint: /
 *

 * Description:
 * ------------
 * Returns a simple message stating that the API is reachable.
 *
 * Example:
 * --------
 * curl -X GET http://localhost:8080. This will show the a basic message, "I am TMN-API."

 *********************************************************************************************************************/
func GetRoot(recorder *httptest.ResponseRecorder, request *http.Request, db *gorm.DB) error {
	rootMsg := RootResponse{Msg: "I am TMN-API"}

	json, err := json.Marshal(&rootMsg)
	if err != nil {
		errors.New("Error marshalling struct into JSON: " + err.Error())
	}

	recorder.Write(json)
	return nil
}
