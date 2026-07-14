# 03 — Tipos, structs e interfaces

## Arrays e slices

**Array** — tamanho fixo:

```go
var a [3]int = [3]int{1, 2, 3}
b := [...]int{10, 20} // tamanho inferido
```

**Slice** — tamanho dinâmico (mais usado):

```go
s := []int{1, 2, 3}
s = append(s, 4)

sub := s[1:3] // fatia: índices [1, 3)
len(s)        // comprimento
cap(s)        // capacidade interna
```

> **Idioma Go:** Prefira **slices** em vez de arrays na maioria dos casos.

Slice `nil` vs vazio:

```go
var a []int           // nil, len 0
b := []int{}          // não nil, len 0
```

## Maps

```go
idades := map[string]int{
	"Ana": 30,
	"Bob": 25,
}

idades["Carlos"] = 28
v, ok := idades["Ana"] // ok == true se existir
delete(idades, "Bob")
```

> **Erro comum:** Escrever em map `nil` — causa **panic**. Use `make` ou literal.

```go
m := make(map[string]int)
m["chave"] = 1
```

## Structs

Agrupam campos relacionados:

```go
type Pessoa struct {
	Nome  string
	Idade int
}

p := Pessoa{Nome: "Ana", Idade: 30}
p.Idade = 31
```

## Ponteiros

```go
p := &Pessoa{Nome: "Bob", Idade: 20} // p é *Pessoa
p.Idade = 21                          // equivalente a (*p).Idade
```

| Situação | Use ponteiro? |
|----------|----------------|
| Alterar struct dentro de função | Sim |
| Struct grande (evitar cópia) | Às vezes |
| Tipos pequenos (`int`, `bool`) | Geralmente não |

## Métodos

Funções ligadas a um tipo:

```go
func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Sou %s, %d anos", p.Nome, p.Idade)
}

// Com ponteiro: pode modificar o struct
func (p *Pessoa) Aniversario() {
	p.Idade++
}
```

O parâmetro entre parênteses chama-se **receiver** (`p`).

## Interfaces

Conjunto de métodos. Um tipo implementa uma interface **implicitamente** (sem `implements`):

```go
type Falador interface {
	Falar() string
}

type Cachorro struct{ Nome string }

func (c Cachorro) Falar() string {
	return "Au! Sou " + c.Nome
}

func ouvir(f Falador) {
	fmt.Println(f.Falar())
}
```

Interface vazia `any` (alias de `interface{}`) aceita qualquer valor — use com moderação.

## Type alias e definitions

```go
type ID int64
type Celsius float64
```

## Prática

1. Crie um slice de strings, adicione nomes com `append` e itere com `range`
2. Crie um `map[string]float64` de notas e busque com `ok`
3. Defina struct `Livro` com método `Resumo() string`

## Próximo passo

[04 — Funções e erros](04-funcoes-erros.md)

[← Índice](README.md)
