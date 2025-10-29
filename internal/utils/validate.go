package utils
// Trocar a validação de prefixo e sufixo que estão em SLICES para strings.HasPrefix e strings.HasSufix -> OK!
// Trabalhar com ponteiro para ficar evitando copias de variaveis e também para evitar confusão -> OK!
// Retornar o ponteiro e não a copia da URL -> OK!
// Melhorar mensagem de erro, ex: 'url inválida: precisa começar com http:// ou https://' -> OK!

import (
	"fmt"
	"strings"
)

// Valida uma URL para saber se esta no formato correto (com http ou https e / no final)
func ValidateURL(url *string) (*string, error) {
	// Transforma toda a string em letras mínusculas para não ter erro de formatação na hora de comparar
	*url = strings.ToLower(*url)

	// Valida se a string começa com 'http://' ou 'https://'
	if strings.HasPrefix(*url, "http://") || strings.HasPrefix(*url, "https://"){

		// Valida se a string não termina com '/', caso seja verdadeiro, acrescenta uma '/' no final e retorna a URL montada.
		if !strings.HasSuffix(*url, "/"){
			*url = *url + "/"
			return url, nil
		}

		// Caso a URL já tenha '/' no final, apenas retorna a URL já validada
		return url, nil
	} 

	// Caso a URL não tenha 'http://' ou 'https://' retorna uma string vazia e um erro
	return nil, fmt.Errorf("[URL Inválida]: A URL precisa começar com 'http://' ou 'https://'")
}
