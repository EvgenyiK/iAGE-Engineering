package server

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//User ...
 type User struct{
	Key string	`json:"key"`
	Value int	`json:"value"`
 }

 //SetStruct ...
 func SetStruct(c redis.Conn) error {
	 
	usr:= User{
		Key: "age",
		Value: 19,
	}

	// serialize User object to JSON
	json,err:= json.Marshal(usr)
	if err != nil {
		return err
	}

	//SET object
	_,err = c.Do("SET",  usr.Value, json)
	if err != nil {
		return err
	}

	return err

 }

//GetStruct ...
 func GetStruct(c redis.Conn) error {

	key := "age"
	s, err := redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		fmt.Println("User does not exist")
	} else if err != nil {
		return err
	}

	usr := User{}
	err = json.Unmarshal([]byte(s), &usr)

	fmt.Printf("%+v\n", usr)

	return nil

}

//NewPool пул соединений
func NewPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}