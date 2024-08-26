package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Data any `json:"data,omitempty"`
	Message string`json:"message,omitempty"`	
}

func Success(w http.ResponseWriter, st string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := Response{
		Status: st,
		Data: data,
	}
	j, _ := json.Marshal(res)
	w.Write(j)
}

func BadRequestResponse(w http.ResponseWriter, st string, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	res := Response{
		Status: st,
		Message: msg,
	}
	j, _ := json.Marshal(res)
	w.Write(j)
}

func InternalServerErrorResponse(w http.ResponseWriter, st string, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	res := Response{
		Status: st,
		Message: msg,
	}
	j, _ := json.Marshal(res)
	w.Write(j)
}