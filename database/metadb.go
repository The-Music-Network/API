package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"github.com/jaylevin/TMN-API/controllers/middleware"
)

// MetaDb is an attempt to simplify the use of the data source in a
// semi-functional (dependency injected) way.
type MetaDb struct {
	Db       *gorm.DB
	Recorder *httptest.ResponseRecorder
}

type MetaDbFunc func(recorder *httptest.ResponseRecorder, r *http.Request, db *gorm.DB) (error)

func (this *MetaDb) Handler(f MetaDbFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

		err := f(this.Recorder, request, this.Db)
		if err != nil {
			log.Println("Error encountered handling API endpoint:", err)
		}

		var out []byte = this.Recorder.Body.Bytes()
		prettify, ok := request.URL.Query()["prettify"]

		if ok {
			if prettify[0] == "true" {
				out = middleware.Prettify(request, this.Recorder.Body.Bytes())
			}
		}

		w.Write(out)
		this.Recorder = httptest.NewRecorder()
	}
}
