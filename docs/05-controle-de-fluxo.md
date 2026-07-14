# 05 — Controle de fluxo

## `if`

```go
if x > 0 {
	fmt.Println("positivo")
} else if x < 0 {
	fmt.Println("negativo")
} else {
	fmt.Println("zero")
}

// if com statement inicial
if err := salvar(); err != nil {
	fmt.Println(err)
}
```

> **Idioma Go:** Chaves `{ }` são **obrigatórias** em `if`, `for`, `switch`.

## `for` — o único loop

Go só tem `for`, em várias formas:

```go
// Clássico
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

// Enquanto
n := 1
for n < 100 {
	n *= 2
}

// Infinito
for {
	break
}
```

## `range`

Itera sobre slice, array, map, string ou channel:

```go
nums := []int{10, 20, 30}
for i, v := range nums {
	fmt.Println(i, v)
}

for k, v := range map[string]int{"a": 1} {
	fmt.Println(k, v)
}

for i, r := range "Go" {
	fmt.Println(i, string(r))
}
```

Se não precisar do índice:

```go
for _, v := range nums {
	fmt.Println(v)
}
```

## `switch`

```go
switch dia {
case "seg":
	fmt.Println("início da semana")
case "sex", "sab":
	fmt.Println("quase fim de semana")
default:
	fmt.Println("outro dia")
}
```

**Switch sem expressão** — cada `case` é um `if` booleano:

```go
switch {
case x < 0:
	fmt.Println("negativo")
case x == 0:
	fmt.Println("zero")
default:
	fmt.Println("positivo")
}
```

> **Idioma Go:** `switch` **não** faz fall-through entre cases (não precisa de `break`).

## `break` e `continue`

```go
for i := 0; i < 10; i++ {
	if i%2 == 0 {
		continue // próxima iteração
	}
	if i > 7 {
		break // sai do loop
	}
}
```

## `return` em `main`

Encerra o programa imediatamente:

```go
if err != nil {
	fmt.Println("erro fatal")
	return
}
```

## Exemplo real: exercício `1-sortNumber`

No repositório, [1-sortNumber/main.go](../1-sortNumber/main.go) combina:

- `for i := range numbers` — percorre índices do array
- `if err != nil` — valida conversão de string para número
- `switch` sem expressão — compara chute com número sorteado
- `continue` — pede novo chute sem sair do loop
- `return` — acerta o número e termina

Trecho ilustrativo:

```go
for i := range numbers {
	// ler entrada...
	chuteInt, err := strconv.ParseInt(chute, 10, 64)
	if err != nil {
		fmt.Println("Digite um número válido!")
		continue
	}
	switch {
	case chuteInt == sortedNumber:
		fmt.Println("Numero Correto!")
		return
	case chuteInt > sortedNumber:
		fmt.Println("Numero Maior que o sorteado!")
		continue
	case chuteInt < sortedNumber:
		fmt.Println("Numero Menor que o sorteado!")
		continue
	}
}
```

Estude o arquivo completo e execute:

```powershell
go run ./1-sortNumber
```

## Prática

1. Imprima números de 1 a 100, mas "Fizz" se múltiplo de 3, "Buzz" se 5, "FizzBuzz" se ambos
2. Some um slice com `for range`
3. Reescreva o jogo de adivinhação com limite de 7 tentativas usando `break`

## Próximo passo

[06 — Pacotes e módulos](06-pacotes-modulos.md)

[← Índice](README.md)
