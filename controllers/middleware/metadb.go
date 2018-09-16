package middleware

import (
	"encoding/json"
	"github.com/jaylevin/TMN-API/errs"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"net/http/httptest"
)

// Acts as a middle man to the http.HandlerFunc function
type MetaDb struct {
	Db       *gorm.DB
	Recorder *httptest.ResponseRecorder
}

type Response struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

type MetaDbFunc func(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error

func (this *MetaDb) Handler(f MetaDbFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var resp interface{} = nil

		err := f(this.Recorder, request, this.Db)
		if err != nil {
			err := err.(*errs.Error)
			this.Recorder.WriteHeader(err.Status())
			resp = err.Root().Error()
		} else if this.Recorder.Body.Bytes() != nil {
			err := json.Unmarshal(this.Recorder.Body.Bytes(), &resp)
			if err != nil {
				log.Println("Error unmarshalling JSON: ", err)
			}
		}

		// Write the response to buffer
		w.Write(this.manufactureResponse(resp, request))

		// Clear the buffer for next API request
		this.Recorder = httptest.NewRecorder()
	}
}

/* Manufactures the JSON response to be returned to the client */
func (this *MetaDb) manufactureResponse(response interface{}, request *http.Request) []byte {
	responseStruct := Response{
		Status:   this.Recorder.Code,
		Response: response,
	}

	// Raw JSON to be returned
	jsonResp, err := json.Marshal(&responseStruct)
	if err != nil {
		log.Println("Error marshalling GenericResponse into json!:", err)
	}

	// If the user wants the output pretty, indent the raw json.
	prettify, ok := request.URL.Query()["prettify"]
	log.Println("Prettify? ", ok)
	if ok {
		if prettify[0] == "true" {
			jsonResp = Prettify(request, jsonResp)
		}
	}

	return jsonResp
}
