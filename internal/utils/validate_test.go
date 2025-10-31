package utils

import (
	"testing"
)

type testCase struct{
	input string
	expected string
}

func TestValidateURLInputsBatch(t *testing.T) {
	urlList := []testCase{
		{"htt://www.codeflakes.com.br/", "url inválida: precisa começar com http:// ou https://"},
		{"htts://www.codeflakes.com.br/", "url inválida: precisa começar com http:// ou https://"},
		{"htps://www.codeflakes.com.br/", "url inválida: precisa começar com http:// ou https://"},
		{"https:/www.codeflakes.com.br/", "url inválida: precisa começar com http:// ou https://"},
	}

	for _, obj := range urlList{
		_, err := ValidateURL(&obj.input)
		if err == nil {
			t.Errorf("Esperava um erro, mas não retornou nenhum. Entrada: %v", obj.input)
			continue
		}
		if err.Error() != obj.expected {
			t.Errorf("O erro retornado é diferente do esperado: %v", err)
		}
	}

}
