package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// https://blog.logrocket.com/exploring-structs-interfaces-go/#convert-interface-struct-golang
func ParseBody(r *http.Request, x interface{}) {
	// basically read all from request body and if err is nil, split the body into a slice and store into 'x' and return

	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
