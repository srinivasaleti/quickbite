package httputils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJSONResponse(t *testing.T) {
	t.Run("success response", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		data := struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			ID:   "10",
			Name: "Chicken Waffle",
		}

		WriteJSONResponse(recorder, data, http.StatusOK)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
		assert.JSONEq(t, `{"id":"10","name":"Chicken Waffle"}`, recorder.Body.String())
	})

	t.Run("error response when encoding fails", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		WriteJSONResponse(recorder, make(chan int), http.StatusInternalServerError)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.JSONEq(t, `{"code":"INTERNAL_SERVER_ERROR","message":"Unable to encode the data"}`, recorder.Body.String())
	})
}

func TestWriteError(t *testing.T) {
	t.Run("error response with custom message", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		WriteError(recorder, "Custom error occurred", InternalServerError, http.StatusBadRequest)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
		assert.JSONEq(t, `{"code":"INTERNAL_SERVER_ERROR","message":"Custom error occurred"}`, recorder.Body.String())
	})
}
