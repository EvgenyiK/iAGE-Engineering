package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
)

//User ...
 type User struct{
	Key string	`json:"key"`
	Value int	`json:"value"`
 }

 var ctx = context.Background()

 //ExampleNewClient ...
 func ExampleNewClient(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
    
    var user User

    err:= json.NewDecoder(r.Body).Decode(&user)
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    err = rdb.Set(ctx, user.Key, user.Value, 0).Err()
    if err != nil {
        log.Fatal(err)
    }

    val, err := rdb.Get(ctx, user.Key).Result()
    if err != nil {
        log.Fatal(err)
    }
    incrVal,err:= strconv.Atoi(val)
    if err != nil {
        log.Fatal(err)
    }
    sum:= incrVal + 1
    fmt.Println("value", sum)

    json.NewEncoder(w).Encode(sum)
}