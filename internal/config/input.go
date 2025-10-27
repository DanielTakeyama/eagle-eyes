// Função GetURL
// 1º - Solicita a URL do Alvo -> OK
// 2º - Analisa se a URL esta vazia e caso esteja, retorna um erro -> OK
// 3º - Analisa se a URL começa com HTTP:// ou HTTPS:// e caso não comece, retorna um erro
// 4º - Analisa se o ultimo digito da URL é "/", caso não seja, adicione
// 5º - Retorna a URL validada

package config

import (
	"eagle-eyes/internal/utils"
	"fmt"
)

func GetURL() (string, error) {
	// Declaração das variáveis + Zero value (Atribuido pelo GO)
	var URL string
	var loop bool

	// Exibe uma mensagem para o usuario solicitando a URL do ALVO e mostra um exemplo de como o dado deve ser fornecido
	fmt.Println("Insira a URL do alvo, ex: https://www.meualvo.com.br")

	// Solicita a URL do alvo
	for !loop {
		fmt.Print("URL do Alvo: ")
		_, err := fmt.Scanln(&URL)
		// Se der algum erro na hora de receber os dados, retorna uma string vazia e o erro para tratarmos ele no escopo da onde a função foi chamada
		if err != nil{
			return "", err
		}

		// Enquanto o usuário digitar uma URL em branco (Vazia) ele fica preso no loop
		if URL == ""{
			fmt.Println("[erro]: URL Incorreta: A URL não pode estar vazia, tente novamente!")
		} else {
			// Se a URL não estiver vazia, faz uma validação para ver se ela esta no formato correto, caso esteja, sai do loop e retorna a URL validada, caso não esteja, retorna o erro
			URL, err = utils.ValidateURL(URL)
			if err != nil{
				return "", err
			}
			loop = true
		}
	}
	return URL, nil
}
