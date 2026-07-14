# 09 — Testes

Go traz testes integrados. Arquivos de teste terminam em `_test.go` e usam o pacote `testing`.

## Primeiro teste

`calc/calc.go`:

```go
package calc

func Soma(a, b int) int {
	return a + b
}
```

`calc/calc_test.go`:

```go
package calc

import "testing"

func TestSoma(t *testing.T) {
	got := Soma(2, 3)
	want := 5
	if got != want {
		t.Errorf("Soma(2,3) = %d; want %d", got, want)
	}
}
```

Executar:

```powershell
go test ./calc
go test -v ./calc    # verboso
```

## Tabela de testes

Vários casos em um loop:

```go
func TestSomaTabela(t *testing.T) {
	casos := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, c := range casos {
		got := Soma(c.a, c.b)
		if got != c.want {
			t.Errorf("Soma(%d,%d) = %d; want %d", c.a, c.b, got, c.want)
		}
	}
}
```

## Testar funções com erro

```go
func TestDivide(t *testing.T) {
	_, err := Divide(1, 0)
	if err == nil {
		t.Error("esperava erro ao dividir por zero")
	}
}
```

## `t.Helper()`, `t.Fatal`, `t.Skip`

```go
t.Fatal("para aqui")  // falha e interrompe o teste
t.Skip("ainda não implementado")
```

## Benchmarks (intro)

```go
func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(1, 2)
	}
}
```

```powershell
go test -bench=. ./calc
```

## Cobertura

```powershell
go test -cover ./calc
go test -coverprofile=cover.out ./calc
go tool cover -html=cover.out
```

> **Idioma Go:** Teste funções **puras** (sem I/O) primeiro; isole lógica de `main` em pacotes testáveis.

## Prática

1. Crie pacote `calc` com `Sub` e testes em tabela
2. Teste função que retorna erro para input inválido
3. Rode `go test ./...` na raiz do módulo

## Próximo passo

[10 — JSON](10-json.md)

[← Índice](README.md)
