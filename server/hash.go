package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//Test ...
type Test struct{
	Text string `json:"text"`
	Key string   `json:"key"`
}

//TestHandler ...
func TestHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
	var test Test
	err:= json.NewDecoder(r.Body).Decode(&test)
	if err != nil {
		log.Fatal(err)
	}
	
	token:= jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["text"] = test.Text
	claims["key"] = test.Key

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
	
	json.NewEncoder(w).Encode(t)
}