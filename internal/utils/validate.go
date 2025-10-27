package utils

import (
	"fmt"
	"strings"
)

// Valida uma URL para saber se esta no formato correto (com http ou https e / no final)
func ValidateURL(url string) (string, error) {
	// Pega o ultimo caractere da string
	lenURL := len(url) - 1

	// Transforma toda a string em letras mínusculas para não ter erro de formatação na hora de comparar
	url = strings.ToLower(url)

	// Valida se a string começa com 'http://' ou 'https://'
	if url[0:7] == "http://" || url[0:8] == "https://"{

		// Valida se a string não termina com '/', caso seja verdadeiro, acrescenta uma '/' no final e retorna a URL montada.
		if url[lenURL:] != "/"{
			url := url + "/"
			return url, nil
		}

		// Caso a URL já tenha '/' no final, apenas retorna a URL já validada
		return url, nil
	} 
	
	return "", fmt.Errorf("[erro] URL Inválida: A URL precisa começar com 'http://' ou 'https://'")
}
