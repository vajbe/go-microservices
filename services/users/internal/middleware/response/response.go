package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(w http.ResponseWriter, message string, data interface{}) {
	sendResponse(w, http.StatusOK, Response{
		Status:  "Success",
		Message: message,
		Data:    data,
	})
}

func Fail(w http.ResponseWriter, message string, data interface{}) {
	sendResponse(w, http.StatusOK, Response{
		Status:  "Fail",
		Message: message,
		Errors:  data,
	})
}

func Error(w http.ResponseWriter, message string, code int) {
	sendResponse(w, code, Response{
		Status:  "Error",
		Message: message,
	})
}

func sendResponse(w http.ResponseWriter, code int, payload Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
