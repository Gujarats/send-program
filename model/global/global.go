package global

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string
	Message string
	Latency float64
	Data    interface{}
}

func SetBadResponse(w http.ResponseWriter, Status string, Message string) {
	w.WriteHeader(http.StatusBadRequest)
	resp := Response{}
	resp.Status = Status
	resp.Message = Message
	json.NewEncoder(w).Encode(resp)
}

func SetOkResponse(w http.ResponseWriter, Status string, Message string, Data interface{}) {
	w.WriteHeader(http.StatusOK)
	resp := Response{}
	resp.Status = Status
	resp.Message = Message
	resp.Data = Data
	json.NewEncoder(w).Encode(resp)
}
