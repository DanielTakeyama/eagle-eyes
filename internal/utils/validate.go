package utils

import "fmt"

// 1 - Se da posição 0:7 == "http://" OU da posição 0:8 == "https://" -> OK
// Então a URL esta certa, se não, retorne um erro. -> OK
// 2 - Após a URL ser validada, verifique se o ultimo digito é um "/", caso seja, ok, caso não, adicione um "/"
// 3 - Retorna a URL validada

func ValidateURL(url string) (string, error) {
	lenURL := len(url) - 1
	if url[0:7] == "http://" || url[0:8] == "https://"{
		
		if url[lenURL:] != "/"{
			url := url + "/"
			return url, nil
		}

		return url, nil
	} 
	
	return "", fmt.Errorf("[erro] URL Inválida: A URL precisa começar com 'http://' ou 'https://'")
}
