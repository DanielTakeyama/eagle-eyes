package main

import (
	"eagle-eyes/internal/config"
	"fmt"
)

func main() {

	url, err := config.GetURL()
	if err != nil {
		panic(err)
	}
	fmt.Println("URL: ", *url)
}
