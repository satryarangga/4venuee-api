package helpers

import (
	"net/http"
	"encoding/json"
)

type Response struct {
	Code	int `json:"code"`
	Message		string `json:"message"`
	Data	interface{} `json:"data"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{}, msg)
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}, msg string) {
	data := Response{code, msg, payload}
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}