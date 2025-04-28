# Aprendendo Go com o Livro "A Linguagem de Programação Go"
Este repositório é dedicado ao meu estudo da linguagem Go, seguindo o livro A Linguagem de Programação Go.

## Sobre o Projeto
A ideia é registrar os códigos, exercícios e projetos realizados durante a leitura do livro, servindo como material de consulta e evolução pessoal.

## Comandos Go

### go run
- O comando ```go run``` compila e executa um programa Go em um único passo.
- Ele pega o(s) arquivo(s) .go, gera o binário temporariamente na memória, executa o programa, e depois descarta o binário.
- Ideal para testes rápidos, pequenos scripts ou enquanto você está desenvolvendo e testando.
- Exemplo: ```go run ch1/helloworld.go```

### go build
- Compila o programa e gera um executável (binário).
- Ideal para gerar o programa para distribuir, rodar depois sem o código-fonte, ou para ambientes de produção.
- Exemplo: ```go build ch1/helloworld.go```
- Esse comando cria um arquivo binário executável chamado helloworld, que pode ser executado a qualquer momento, sem processamento adicionais. ```./helloworld```

### go install
- Instala a ferramenta (binário) para usar no terminal
- Ideal para quando se quer instalar comandos Go externos
- Exemplo: ```go install golang.org/x/tools/cmd/goimports@latest```

## Ferramentas

### goimports

O goimports é uma ferramenta que:

- Organiza automaticamente os imports no seu código Go.
- Adiciona imports que estão faltando.
- Remove imports que não estão sendo usados.
- Reordena os imports no padrão do Go.
- Como instalar: ```go install golang.org/x/tools/cmd/goimports@latest```
- ```@latest``` garante a instalação mais recente da ferramenta.
- Observação: dento de um 'go.mod' pode rodar o ```go install``` sem a necessidade do ```@latest```
- Como usar: ```goimports -w helloworld.go```
- Documentação: [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports).

# Observação
Este projeto é de estudo e aprendizado contínuo. Atualizações serão feitas conforme a evolução na leitura do livro e a realização dos exercícios.
