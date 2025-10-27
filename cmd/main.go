package main

import (
	"eagle-eyes/internal/utils"
	"fmt"
)

func main() {

	//url, err := config.GetURL()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("URL: ", *url)
	http := "http://www.codeflakes.com.br"
	https := "https://www.codeflakes.com.br/"
	url := "www.codeflakes.com.br"

	urlUm, err := utils.ValidateURL(http)
	if err != nil {
		panic(err)
	}
	fmt.Println(urlUm)
	//////////////////////////////////////

	urlDois, err := utils.ValidateURL(https)
	if err != nil {
		panic(err)
	}
	fmt.Println(urlDois)
	//////////////////////////////////////

	urlTres, err := utils.ValidateURL(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(urlTres)
	//////////////////////////////////////
}
