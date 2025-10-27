package main

import "eagle-eyes/internal/utils"

func main() {

	//url, err := config.GetURL()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("URL: ", *url)
	http := "https://www.codeflakes.com.br"
	https := "https://www.codeflakes.com.br"
	url := "www.codeflakes.com.br"

	_, err := utils.ValidateURL(http)
	if err != nil {
		panic(err)
	}
	//////////////////////////////////////

	_, err = utils.ValidateURL(https)
	if err != nil {
		panic(err)
	}
	//////////////////////////////////////

	_, err = utils.ValidateURL(url)
	if err != nil {
		panic(err)
	}
	//////////////////////////////////////
}
