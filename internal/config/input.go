// Função GetURL
// 1º - Solicita a URL do Alvo -> OK
// 2º - Analisa se a URL esta vazia e caso esteja, retorna um erro -> OK
// 3º - Analisa se a URL começa com HTTP:// ou HTTPS:// e caso não comece, retorna um erro
// 4º - Analisa se o ultimo digito da URL é "/", caso não seja, adicione
// 5º - Retorna a URL validada

package config

import "fmt"

func GetURL() (*string, error) {
	// Declaração das variáveis + Zero value (Atribuido pelo GO)
	var URL string
	var loop bool

	// Exibe uma mensagem para o usuario solicitando a URL do ALVO e mostra um exemplo de como o dado deve ser fornecido
	fmt.Println("Insira a URL do alvo, ex: https://www.meualvo.com.br")

	// Solicita a entrada dos dados pelo usuario
	for !loop {
		fmt.Print("URL do Alvo: ")
		_, err := fmt.Scanln(&URL)
		// Se der algum erro na hora de receber os dados, retorna nil para a string e o erro para tratarmos ele
		if err != nil{
			return nil, err
		}

		// Enquanto o usuário não digitar nada ele fica preso no loop

		//FALTA CRIAR E ADICIONAR A FUNÇÃO DE VALIDAR URL (HTTP, HTTPS E /)
		if URL == "" /*|| validateURL == False*/{
			fmt.Println("Erro: a URL não pode estar vazia, tente novamente!")
		} else {
			loop = true
		} 
	}

	return &URL, nil
}
