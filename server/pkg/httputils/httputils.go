package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Errorcode string

const (
	InternalServerError Errorcode = "INTERNAL_SERVER_ERROR"
	NotFound            Errorcode = "NOT_FOUND"
)

// APIError defines the structure of an API error response.
type APIError struct {
	Code    Errorcode `json:"code"`
	Message string    `json:"message"`
}

// WriteError sends an error response to the client.
func WriteError(w http.ResponseWriter, message string, code Errorcode, status int) {
	er := APIError{
		Code:    code,
		Message: message,
	}

	WriteJSONResponse(w, er, status)
}

// WriteJSONResponse sends a successful JSON response.
func WriteJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		WriteError(w, "Unable to encode the data", InternalServerError, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Fprintf(w, "Failed to write JSON response")
	}
}
