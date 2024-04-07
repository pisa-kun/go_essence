package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var content = `
{
	"groups": "放課後",
	"concept": "Girls",
	"informations": {
		"color": "orange",
		"member": 5
	}
}
`

type Informations struct {
	Color  string `json:"color"`
	Member int    `json:"member"`
}

type Data struct {
	Groups       string       `json:"groups"`
	Concept      string       `json:"concept"`
	Informations Informations `json:"informations"`
}

func main() {
	// var data Data
	// err := json.Unmarshal([]byte(content), &data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(data)

	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data Data
	err = json.NewDecoder(f).Decode(&data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
