package config

import (
	"fmt"
)

func GetURL() (*string, error) {
	// Cria uma variável para receber a URL
	var url string

	// Le o STDIN (Entrada Padrão do terminal) e armazena o valor na memória da variável URL
	for url == ""{
		
		// Exibe uma mensagem solicitando a URL do Alvo
		fmt.Println("Digite a URL do alvo, ex: https://www.suaurl.com.br")
		fmt.Print("URL do ALVO: ")
		
		// Pega os dados digitado pelo usuario
		_, err := fmt.Scanln(&url)
		
		// Verifica se houve algum erro, caso tenha retorna o erro
		if err != nil {
			return nil, fmt.Errorf("falha ao ler a URL do usuario")
		}
	}

	// Retorna o valor para quem chamou a função
	return &url, nil
}
