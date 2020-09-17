package main

import (
	"fmt"
	s "iage/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	r:= mux.NewRouter()
	pool:= s.NewPool()
	conn:= pool.Get()
	defer conn.Close()

	err := s.GetStruct(conn)
	if err != nil {
		fmt.Println(err)
	}
	r.HandleFunc("/postgres/users", s.CreateUser).Methods("POST", "OPTIONS")
	fmt.Println("Starting server on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}