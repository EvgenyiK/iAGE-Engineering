package main

import(
	"fmt"
	s"iage/server"
)


func main() {
	pool:= s.NewPool()
	conn:= pool.Get()
	defer conn.Close()

	err := s.GetStruct(conn)
	if err != nil {
		fmt.Println(err)
	}
}