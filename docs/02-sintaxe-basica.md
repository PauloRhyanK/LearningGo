# 02 — Sintaxe básica

## Package e imports

```go
package main

import (
	"fmt"
	"os"
)

import "strings" // import único também é válido
```

- Um **pacote** agrupa código relacionado
- **Imports** não usados são **erro de compilação** (o compilador exige que você remova)

> **Idioma Go:** O formatador `go fmt` organiza imports automaticamente.

## Funções

```go
func soma(a int, b int) int {
	return a + b
}

// Tipos iguais podem ser agrupados
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}
```

## Variáveis

```go
// Declaração com tipo
var nome string = "Ana"
var idade int

// Inferência de tipo
var cidade = "São Paulo"

// Declaração curta (só dentro de funções)
pais := "Brasil"
x, y := 10, 20
```

| Forma | Onde usar |
|-------|-----------|
| `var x int` | Variáveis de pacote ou quando precisa do zero value explícito |
| `x := 10` | Dentro de funções, código mais conciso |

## Constantes

```go
const Pi = 3.14159
const (
	StatusOK = 200
	MaxTentativas = 3
)
```

## Tipos básicos

| Tipo | Exemplo | Zero value |
|------|---------|------------|
| `bool` | `true`, `false` | `false` |
| `string` | `"olá"` | `""` |
| `int`, `int64` | `42` | `0` |
| `float64` | `3.14` | `0` |
| `byte` | `'A'` (alias `uint8`) | `0` |
| `rune` | `'世'` (alias `int32`, Unicode) | `0` |

## Conversão de tipos

Go **não** converte tipos implicitamente:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(i)
```

> **Erro comum:** `var f float64 = i` — não compila; use conversão explícita.

## Zero values

Todo tipo tem valor padrão quando você declara sem inicializar:

```go
var n int       // 0
var s string    // ""
var ok bool     // false
var p *int      // nil
```

## O que é `nil`?

`nil` significa **ausência de valor** em:

- ponteiros, slices, maps, channels, funções, interfaces

```go
var p *int
fmt.Println(p == nil) // true
```

Go **não** tem `null` — só `nil`.

> **Erro comum:** Pensar que `int` ou `string` podem ser `nil`. Não podem; usam zero value (`0`, `""`).

## Comentários

```go
// linha única

/*
 bloco
*/
```

## Impressão formatada (`fmt`)

```go
fmt.Println("a", "b")     // com espaços e newline
fmt.Printf("%d\n", 42)    // formatado
fmt.Sprintf("x=%d", 10)   // retorna string
```

Verbos comuns: `%d` inteiro, `%s` string, `%f` float, `%v` valor genérico, `%T` tipo, `%w` wrap de erro (ver capítulo 04).

## Prática

1. Declare variáveis com `var` e `:=` e imprima com `fmt.Printf`
2. Converta um `int` para `float64` e imprima
3. Declare um ponteiro `*string` nil e compare com `nil`

## Próximo passo

[03 — Tipos, structs e interfaces](03-tipos-structs-interfaces.md)

[← Índice](README.md)
