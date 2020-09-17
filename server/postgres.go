package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//Postgres ...
type Postgres struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type response struct {
	ID      int64  `json:"id.omitempty"`
	Message string `json:"message,omitempty"`
}

func connection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	//проверка соединения
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	return db
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	var user Postgres

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}
	insertID := insertUser(user)

	res := response{
		ID:      insertID,
		Message: "User to create ",
	}

	json.NewEncoder(w).Encode(res)

}

//реализация вставки sql
func insertUser(user Postgres) int64 {
	db := connection()
	defer db.Close()
	sqlstatement := `INSERT INTO users(name,age)values($1,$2)returning id`

	var id int64

	err := db.QueryRow(sqlstatement, user.ID).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)

	return id
}
