package controllers

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
)

type RootResponse struct {
	Msg string
}

func GetRoot(recorder *httptest.ResponseRecorder, request *http.Request, db *gorm.DB) error {
	rootMsg := RootResponse{Msg: "I am TMN-API"}

	json, err := json.Marshal(&rootMsg)
	if err != nil {
		errors.New("Error marshalling struct into JSON: " + err.Error())
	}

	recorder.Write(json)

	return nil
}

