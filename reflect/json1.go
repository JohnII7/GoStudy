package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors""`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 23, []string{"周星驰", "张柏芝"}}

	// 结构体转为Json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// Json转为结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
	}
	fmt.Printf("%v\n", myMovie)
}
