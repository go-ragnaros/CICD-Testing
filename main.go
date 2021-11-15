package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Data struct {
	D1 string `json:"d1"`
	D2 int    `json:"d2"`
}

func main() {
	d := make(map[string]*Data)

	data1 := &Data{
		D1: "a",
		D2: 1,
	}

	data2 := &Data{
		D1: "b",
		D2: 2,
	}

	d[data1.D1] = data1
	d[data2.D1] = data2

	b,err:=json.Marshal(d)
	if err!=nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))



}
