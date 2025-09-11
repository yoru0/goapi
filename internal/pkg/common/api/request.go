package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yoru0/goapi.git/internal/pkg/common/constant/httpstatus"
)

// RequestParam is an interface for request parameters that require validation.
type RequestParam interface {
	Validate() (msg, field string)
}

// DecodeBodyJSON decodes the JSON body into the provided struct.
func DecodeBodyJSON(w http.ResponseWriter, r *http.Request, v RequestParam) (ok bool) {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		log.Println(err.Error())
		msg := "Request body format is invalid"
		response := NewAPIResponseWithError("INVALID_PARAM", msg)
		SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return false
	}

	if msg, field := v.Validate(); msg != "" {
		response := NewAPIResponseWithErrorField("INVALID_PARAM", msg, field)
		SendResponseJSONWithStatusCode(w, response, httpstatus.BadRequest)
		return false
	}

	return true
}
