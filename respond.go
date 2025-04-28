package respond

import (
	"encoding/json"
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func renderJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func OK(w http.ResponseWriter) {
	renderJson(w, http.StatusOK, nil)
}

func Okay(w http.ResponseWriter, data interface{}) {
	renderJson(w, http.StatusOK, data)
}

func Respond(w http.ResponseWriter, status int, data interface{}) {
	renderJson(w, status, data)
}

func Error(w http.ResponseWriter, status int, message string) {
	renderJson(w, status, RestError{Message: message, Status: status})
}

func BadRequest(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Bad Request"
	}
	Error(w, http.StatusBadRequest, message)
}

func InternalError(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Internal Server Error"
	}
	Error(w, http.StatusInternalServerError, message)
}

func NotFound(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Not Found"
	}
	Error(w, http.StatusNotFound, message)
}
