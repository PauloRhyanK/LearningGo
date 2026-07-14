# 01 — Instalação e primeiro programa

## O que é Go?

**Go** (ou Golang) é uma linguagem criada no Google, compilada, com tipagem estática, foco em simplicidade, concorrência nativa e ferramentas integradas (`go build`, `go test`, `go fmt`).

## Instalar no Windows

1. Acesse [https://go.dev/dl/](https://go.dev/dl/)
2. Baixe o instalador **Windows** (`.msi`)
3. Execute e siga o assistente (instala em `C:\Program Files\Go` por padrão)
4. Feche e abra o terminal

Verifique:

```powershell
go version
```

Saída esperada (versão pode variar):

```
go version go1.26.3 windows/amd64
```

## Variáveis de ambiente (conceito)

| Variável | Papel |
|----------|--------|
| Instalação do Go | Onde estão o compilador e a stdlib |
| `GOPATH` | Antigo workspace; hoje menos central com **módulos** |
| Módulos (`go.mod`) | Forma moderna de gerenciar dependências e projetos |

> **Idioma Go:** Projetos novos usam **Go modules**. Você não precisa configurar `GOPATH` manualmente para começar.

## Primeiro programa

Crie uma pasta e um arquivo `main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Olá, Go!")
}
```

Inicialize o módulo (nome do módulo pode ser qualquer identificador válido):

```powershell
go mod init meuapp
go run .
```

O ponto `.` significa “pacote atual”.

## Ferramentas úteis desde o início

```powershell
go env          # variáveis de ambiente do Go
go fmt .        # formata código automaticamente
go vet .        # análise estática simples
```

## Estrutura mínima de um programa executável

```go
package main   // pacote executável

import "fmt"   // importar pacotes

func main() {  // ponto de entrada
	fmt.Println("...")
}
```

- Todo arquivo `.go` começa com `package Nome`
- Programas que rodam no terminal usam `package main` e `func main()`

> **Erro comum:** Esquecer `package main` ou `func main()` — o compilador não gera executável.

## Prática

1. Instale o Go e confirme com `go version`
2. Crie uma pasta fora deste repo, rode `go mod init teste` e imprima seu nome com `fmt.Println`
3. Execute `go fmt` e `go vet` no arquivo

## Próximo passo

[02 — Sintaxe básica](02-sintaxe-basica.md)

[← Índice](README.md)
