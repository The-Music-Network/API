package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Prettify(r *http.Request, original []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, original, "", "  ")
	if err != nil {
		log.Println("Error prettifying JSON: ", err)
		return nil
	}
	return out.Bytes()
}
