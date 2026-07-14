# 06 — Pacotes e módulos

## Pacote vs módulo

| Conceito | O que é |
|----------|---------|
| **Pacote** | Pasta + arquivos `.go` com o mesmo `package` no topo |
| **Módulo** | Projeto versionado; definido em `go.mod` na raiz |

Este repositório usa o módulo `LearningGo` ([go.mod](../go.mod)).

## `package main` vs biblioteca

```go
// Executável — gera binário com go build
package main

func main() { }
```

```go
// Biblioteca — importada por outros pacotes
package mathutil

func Dobro(n int) int {
	return n * 2
}
```

## Exportação (visibilidade)

Nome com **primeira letra maiúscula** = exportado (público no pacote):

```go
func Publica() {}   // outros pacotes podem usar
func privada() {}  // só dentro do mesmo pacote
```

> **Idioma Go:** Visibilidade é por **capitalização**, não por `public`/`private`.

## Imports

```go
import "LearningGo/internal/util"  // caminho do módulo + pasta

import (
	"fmt"
	"strings"
)
```

Caminho de import = **módulo** + **caminho da pasta** dentro do módulo.

## `go.mod`

```go
module LearningGo

go 1.26.3
```

Comandos:

```powershell
go mod init NomeDoModulo   # criar módulo novo
go mod tidy                # ajustar dependências
go get github.com/pkg/errors@latest
```

## Várias pastas no mesmo módulo

Neste repo:

```
learningGO/
├── main.go              # package main na raiz
└── 1-sortNumber/
    └── main.go          # outro package main
```

Executar o subprograma:

```powershell
go run ./1-sortNumber
```

Cada pasta com `package main` e `func main()` é um executável separado.

## Organizar código em pacotes

Exemplo futuro:

```
LearningGo/
├── go.mod
├── cmd/
│   └── todo/
│       └── main.go      # package main — fino
└── internal/
    └── todo/
        └── todo.go      # package todo — lógica
```

`internal/` — convenção: código **não** importável de fora do módulo.

## Prática

1. Crie pasta `calc/` com `calc.go` (`package calc`) e função exportada `Soma`
2. Em `main.go` na raiz, importe e use `calc.Soma`
3. Rode `go mod tidy` e `go run .`

## Próximo passo

[07 — Biblioteca padrão](07-biblioteca-padrao.md)

[← Índice](README.md)
