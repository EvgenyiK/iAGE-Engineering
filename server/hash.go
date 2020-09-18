package server

import (
	//"bytes"
	"encoding/json"
	"net/http"
)


type Test struct{
	Text string `json:"text"`
	Key string   `json:"key"`
}

func TestHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}