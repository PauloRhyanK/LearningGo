# 10 — JSON

O pacote `encoding/json` converte entre structs/maps e JSON.

## Struct tags

```go
type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email,omitempty"` // omite se vazio
	Senha string `json:"-"`               // nunca serializa
}
```

## Marshal (struct → JSON)

```go
import "encoding/json"

u := Usuario{ID: 1, Nome: "Ana"}
b, err := json.Marshal(u)
if err != nil {
	return err
}
fmt.Println(string(b))
// {"id":1,"nome":"Ana"}
```

Indentado:

```go
b, err := json.MarshalIndent(u, "", "  ")
```

## Unmarshal (JSON → struct)

```go
data := []byte(`{"id":2,"nome":"Bob"}`)
var u Usuario
if err := json.Unmarshal(data, &u); err != nil {
	return err
}
```

> **Erro comum:** Passar struct por valor em `Unmarshal` — use **ponteiro** `&u`.

## Map e slice genéricos

```go
var m map[string]any
json.Unmarshal(data, &m)
```

## Arquivo JSON

```go
data, err := os.ReadFile("dados.json")
if err != nil {
	return err
}
var lista []Usuario
if err := json.Unmarshal(data, &lista); err != nil {
	return err
}
```

Salvar:

```go
b, _ := json.MarshalIndent(lista, "", "  ")
os.WriteFile("dados.json", b, 0644)
```

## Erros comuns

| Problema | Solução |
|----------|---------|
| Campo sempre vazio | Tag `json` com nome errado |
| Não atualiza struct | Usar ponteiro no `Unmarshal` |
| Número como string | Tipos incompatíveis na struct |
| `null` em slice | Use ponteiro ou `[]*T` se precisar distinguir |

## Prática

1. Defina struct `Tarefa` com `id`, `titulo`, `feita` e tags JSON
2. Serialize lista de 3 tarefas para string e imprima
3. Leia JSON de string e imprima só títulos não feitas

## Próximo passo

[11 — Concorrência](11-concorrencia.md)

[← Índice](README.md)
