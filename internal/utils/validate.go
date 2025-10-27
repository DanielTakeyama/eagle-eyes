package utils

import "fmt"

// 1 - Se da posição 0:7 == "http://" OU da posição 0:8 == "https://"
// Então a URL esta certa, se não, retorne um erro.
// 2 - Após a URL ser validada, verifique se o ultimo digito é um "/", caso seja, ok, caso não, adicione um "/"
// 3 - Retorna a URL validada

func ValidateURL(url string) (string, error) {
	if url[0:7] == "http://" || url[0:8] == "https://"{
		fmt.Println("URL Válida!")
		return url, nil
	} 
	
	return "", fmt.Errorf("[erro]: URL Inválida")
}
