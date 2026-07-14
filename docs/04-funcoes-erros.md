# 04 — Funções e erros

## Assinaturas e múltiplos retornos

Go permite retornar vários valores — muito comum retornar **resultado + erro**:

```go
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

func main() {
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Println(resultado)
}
```

## Nomes de retorno (opcional)

```go
func dividir(a, b float64) (quociente float64, err error) {
	if b == 0 {
		err = fmt.Errorf("divisão por zero")
		return
	}
	quociente = a / b
	return
}
```

Use com moderação — pode ajudar em funções curtas ou documentação.

## O tipo `error`

`error` é uma **interface** com um método:

```go
type error interface {
	Error() string
}
```

Valores de erro são comparados com `nil`:

```go
if err != nil {
	// algo deu errado
}
```

> **Idioma Go:** Tratar erro **na hora**, explicitamente. Não há exceções try/catch para fluxo normal.

## Criar erros

```go
return fmt.Errorf("usuário %s não encontrado", nome)

// Encadear erro (Go 1.13+)
return fmt.Errorf("salvar arquivo: %w", err)
```

`%w` permite `errors.Is` e `errors.Unwrap` para verificar causa raiz.

## Erros sentinela

```go
var ErrNaoEncontrado = errors.New("não encontrado")

if errors.Is(err, ErrNaoEncontrado) {
	// ...
}
```

## Funções como valores

```go
var op func(int, int) int = func(a, b int) int { return a + b }
fmt.Println(op(2, 3))
```

Útil para callbacks e estratégias simples.

## Defer

Executa ao sair da função (ordem LIFO):

```go
func ler() {
	f, err := os.Open("dados.txt")
	if err != nil {
		return
	}
	defer f.Close()
	// usar f...
}
```

## Panic e recover

- `panic` — falha grave; use raramente (bugs, estado impossível)
- `recover` — captura panic dentro de `defer` (padrão avançado)

Para fluxo normal, prefira retornar `error`.

> **Erro comum:** Ignorar `err` com `_` sem motivo — só faça isso quando tiver certeza de que o erro não importa.

## Prática

1. Escreva `func SomaSlice(nums []int) (int, error)` que retorna erro se o slice for `nil`
2. Use `fmt.Errorf` com `%w` ao encapsular um erro de `os.Open`
3. Use `defer fmt.Println("fim")` e observe a ordem ao usar vários `defer`

## Próximo passo

[05 — Controle de fluxo](05-controle-de-fluxo.md)

[← Índice](README.md)
