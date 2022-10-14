package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// parse request body
var ParseReqBody = func(r *http.Request, x interface{}) {
	if reqBody, err := io.ReadAll(r.Body); err == nil {
		// Create Json from Object
		if err := json.Unmarshal([]byte(reqBody), x); err != nil {
			return
		}
	}
}
