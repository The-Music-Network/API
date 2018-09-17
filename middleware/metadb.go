package middleware

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"net/http/httptest"

	// inner package dependencies
	"github.com/The-Music-Network/TMN-API/errs"
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

func newResponse(status int, response interface{}) *Response {
	return &Response {
		Status: status,
		Response: response,
	}
}

type MetaDbFunc func(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) error

func (this *MetaDb) Handler(f MetaDbFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var legitResp interface{} = nil // the legitimate response that is used by the http client

		err := f(this.Recorder, request, this.Db)
		var prelimResp = this.Recorder.Body.Bytes() // pre-liminary response before any middleware introduction
		log.Println("prelim resp",string(prelimResp))
		if err != nil {
			/*
				If there is an error, then we should be assuming there is a corresponding 4XX http status code
			    (Bad request)
			 */
			err := err.(*errs.Error)
			this.Recorder.WriteHeader(err.Status())
			legitResp = err.Root().Error()

		} else if prelimResp != nil {
			/* Basically, this is executed when the http request was 2XX (Successful),
				We unmarshal the json into
			*/

				log.Println(string(prelimResp))
			/*
			 Unmarshalls JSON array/object (prelimResp) into golang array/struct (legitResp)
			*/
			err := json.Unmarshal(prelimResp, &legitResp)
			if err != nil {
				log.Println("Error unmarshalling JSON: ", err)
			}
		}

		// Write the response to buffer
		w.Write(this.manufactureResponse(legitResp, request))

		// Clear the buffer for next API request
		this.Recorder = httptest.NewRecorder()
	}
}

/* Manufactures the JSON response to be returned to the client */
func (this *MetaDb) manufactureResponse(response interface{}, request *http.Request) []byte {
	structResponse := newResponse(this.Recorder.Code, response)
	log.Println(structResponse)

	// Encodes golang structResponse into Raw JSON to be returned to the http client
	jsonResp, err := json.Marshal(&structResponse)
	if err != nil {
		log.Println("Error marshalling Response into json!:", err)
	}

	log.Println("jsonResp prelim:", string(jsonResp))
	// If the client wants the output pretty, indent the raw json.
	prettify, ok := request.URL.Query()["prettify"]
	if ok {
		if prettify[0] == "true" {
			jsonResp = Prettify(request, jsonResp)
		}
	}

	log.Println("json Resp:", string(jsonResp))
	return jsonResp
}
