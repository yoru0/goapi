package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yoru0/goapi.git/internal/pkg/common/constant/httpstatus"
)

// Response represents API response payload.
type Response struct {
	Status int           `json:"status"`
	Err    ResponseError `json:"error"`
	Data   interface{}   `json:"data"`
}

// ResponseError is an error in an API request.
type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

// ResponseData is the payload data to be returned for an API request.
type ResponseData struct{}

// NewAPIResponse creates a new API response with default values.
func NewAPIResponse() *Response {
	return &Response{
		Status: httpstatus.OK,
		Err:    ResponseError{},
		Data:   ResponseData{},
	}
}

// NewAPIResponseWithError creates a new API response with an error.
func NewAPIResponseWithError(code, msg string) *Response {
	return &Response{
		Status: httpstatus.OK,
		Err:    ResponseError{code, msg, ""},
		Data:   ResponseData{},
	}
}

// NewAPIResponseWithErrorField creates a new API response with an error and a field.
func NewAPIResponseWithErrorField(code, msg, field string) *Response {
	return &Response{
		Status: httpstatus.OK,
		Err:    ResponseError{code, msg, field},
		Data:   ResponseData{},
	}
}

// SendResponseJSON sends the API response as JSON.
func SendResponseJSON(w http.ResponseWriter, r *Response) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

// SendResponseJSONWithStatusCode sends the API response as JSON with a specific HTTP status code.
func SendResponseJSONWithStatusCode(w http.ResponseWriter, r *Response, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	r.Status = statusCode
	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
