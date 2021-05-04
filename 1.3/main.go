package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Animal struct {
	Name string
	Age  int
}

func main() {
	data := []byte(`{"name":"cat", "age":5}`)
	v := Animal{}
	if err := json.Unmarshal(data, &v); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v.Name, v.Age)
}
