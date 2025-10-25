# 🦅 Eagle Eyes

**Eagle Eyes** é uma ferramenta de **cibersegurança open-source** escrita em **Go**, criada para realizar **escaneamento de diretórios** em aplicações web — semelhante a ferramentas como **DirSearch**, **Dirb** e **Gobuster**.  
O principal objetivo deste projeto é **educacional e de pesquisa**, focado em aprimorar o conhecimento sobre segurança ofensiva e defensiva de aplicações web.

---

## 🚀 Objetivo

O **Eagle Eyes** foi desenvolvido com a missão de ajudar pesquisadores, estudantes e profissionais de segurança a:
- Identificar **diretórios e endpoints ocultos** em aplicações web.
- Compreender o funcionamento interno de scanners HTTP em Go.
- Contribuir para um **ecossistema open-source ético** na área de segurança cibernética.

---

## ⚙️ Estrutura do Projeto

```bash
eagle-eyes/
├─ cmd/
│  └─ main.go        # ponto de entrada (main)
│
├─ internal/
│  ├─ scan/              # lógica principal do scanner (core)
│  ├─ utils/             # funções auxiliares (logs, parsing, etc)
│  └─ config/            # leitura e validação de configurações
│
├─ pkg/
│  └─ reporter/          # geração de relatórios (JSON, CSV, etc)
│
├─ tests/                # testes unitários e integração (ambiente local)
│
├─ docs/                 # documentação, arquitetura e exemplos de uso
│
└─ README.md
```

---

## 🛠️ Tecnologias Utilizadas
- Go (Golang) → linguagem principal do projeto
- HTTP standard library → para comunicação e requisições

---

## 💡 Como Contribuir

Contribuições são muito bem-vindas e você pode ajudar com:
- Correções e melhorias no código.
- Otimizações de performance.
- Novas features e modos de relatório.
- Sugestões de melhorias na documentação.

# 1. Faça um fork do repositório
git clone https://github.com/DanielTakeyama/eagle-eyes.git

# 2. Crie uma branch para sua feature
git checkout -b feature/minha-feature

# 3. Commit suas mudanças
git commit -m "Adiciona nova feature"

# 4. Envie o PR para revisão
git push origin feature/minha-feature

---

## ⚠️ Aviso Legal e Ético

O Eagle Eyes foi criado **exclusivamente para fins educacionais, de pesquisa e testes de segurança autorizados.**
O uso desta ferramenta **sem autorização explícita do proprietário** de um sistema **é ilegal** e pode **violar leis locais e internacionais**
(como o Marco Civil da Internet e o Código Penal Brasileiro, art. 154-A).

**O autor NÃO se responsabiliza por qualquer dano, invasão ou prejuízo causado por terceiros que utilizem o Eagle Eyes de forma indevida.**
Ao utilizar este software, você concorda que o uso é por sua própria conta e risco.

**Use com ética. 🧠**
Teste apenas sistemas que você possui ou tem permissão explícita para testar.

---

## 📜 Licença
Este projeto está licenciado sob os termos da MIT License.
Você é livre para usar, copiar, modificar e distribuir este software, desde que mantenha os créditos originais e o aviso de licença.

Leia o arquivo LICENSE para mais detalhes.

---

## 👨‍💻 Autor

Daniel Takeyama
Desenvolvedor fullstack e entusiasta de cibersegurança e software livre.

**Lembre-se: “A melhor defesa é o conhecimento.” – Eagle Eyes**