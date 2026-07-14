# 14 — Projeto completo: app de tarefas

Este capítulo consolida tudo: **CLI**, **JSON**, **testes**, **HTTP** (opcional) e **layout** de projeto.

Implemente em uma pasta nova, por exemplo `cmd/todo/` e `internal/todo/`, sem alterar os exercícios existentes.

## Objetivo

Aplicativo de **tarefas** (todo list) que:

- Adiciona, lista e marca tarefas como feitas
- Persiste em arquivo `tarefas.json`
- (Opcional) Expõe API REST na porta 8080

## Checklist: você construiu um software se…

- [ ] Tem `go.mod` e compila com `go build`
- [ ] Lógica separada de `main` e coberta por testes
- [ ] Trata erros (`if err != nil`) em I/O e JSON
- [ ] Dados sobrevivem ao reiniciar (arquivo JSON)
- [ ] README explica como rodar
- [ ] (Opcional) API responde JSON com códigos HTTP corretos

## Modelo de dados

```go
package todo

type Tarefa struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Feita  bool   `json:"feita"`
}
```

## Interface Store

```go
type Store interface {
	Listar() ([]Tarefa, error)
	Adicionar(titulo string) (Tarefa, error)
	MarcarFeita(id int) error
}
```

## Implementação JSON (esboço)

```go
type JSONStore struct {
	path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

func (s *JSONStore) load() ([]Tarefa, error) {
	data, err := os.ReadFile(s.path)
	if os.IsNotExist(err) {
		return []Tarefa{}, nil
	}
	if err != nil {
		return nil, err
	}
	var lista []Tarefa
	if err := json.Unmarshal(data, &lista); err != nil {
		return nil, err
	}
	return lista, nil
}

func (s *JSONStore) save(lista []Tarefa) error {
	b, err := json.MarshalIndent(lista, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, b, 0644)
}
```

Complete `Adicionar` (gerar ID incremental) e `MarcarFeita`.

## Camada App (CLI)

```go
type App struct {
	store Store
}

func (a *App) RunCLI(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("uso: todo <add|list|done> ...")
	}
	switch args[1] {
	case "add":
		if len(args) < 3 {
			return fmt.Errorf("uso: todo add <título>")
		}
		t, err := a.store.Adicionar(strings.Join(args[2:], " "))
		if err != nil {
			return err
		}
		fmt.Printf("criada #%d: %s\n", t.ID, t.Titulo)
	case "list":
		lista, err := a.store.Listar()
		if err != nil {
			return err
		}
		for _, t := range lista {
			mark := " "
			if t.Feita {
				mark = "x"
			}
			fmt.Printf("[%s] %d: %s\n", mark, t.ID, t.Titulo)
		}
	case "done":
		// parse id, chamar MarcarFeita
	default:
		return fmt.Errorf("comando desconhecido: %s", args[1])
	}
	return nil
}
```

## Testes sugeridos

```go
func TestAdicionar_Listar(t *testing.T) {
	dir := t.TempDir()
	s := NewJSONStore(filepath.Join(dir, "t.json"))
	t1, err := s.Adicionar("comprar leite")
	if err != nil {
		t.Fatal(err)
	}
	lista, err := s.Listar()
	if err != nil {
		t.Fatal(err)
	}
	if len(lista) != 1 || lista[0].ID != t1.ID {
		t.Errorf("lista inesperada: %+v", lista)
	}
}
```

## Variante API HTTP (opcional)

Rotas mínimas:

| Método | Rota | Ação |
|--------|------|------|
| GET | `/tarefas` | Listar JSON |
| POST | `/tarefas` | Body `{"titulo":"..."}` |
| PATCH | `/tarefas/{id}/feita` | Marcar feita |

```go
func (a *App) RegisterHTTP(mux *http.ServeMux) {
	mux.HandleFunc("GET /tarefas", a.handleListar)
	mux.HandleFunc("POST /tarefas", a.handleCriar)
	// ...
}

func main() {
	store := todo.NewJSONStore("tarefas.json")
	app := todo.NewApp(store)
	mux := http.NewServeMux()
	app.RegisterHTTP(mux)
	http.ListenAndServe(":8080", mux)
}
```

Teste:

```powershell
curl http://localhost:8080/tarefas
curl -X POST http://localhost:8080/tarefas -d "{\"titulo\":\"estudar Go\"}"
```

## Passos de implementação (ordem)

1. Criar `internal/todo` com struct `Tarefa`
2. Implementar `JSONStore` + testes
3. Implementar `App.RunCLI`
4. `cmd/todo/main.go` fino
5. README com comandos
6. (Opcional) Handlers HTTP

## Comandos finais

```powershell
go test ./internal/todo/...
go run ./cmd/todo add "Aprender Go"
go run ./cmd/todo list
go run ./cmd/todo done 1
go build -o todo.exe ./cmd/todo
```

## Relação com este repositório

| O que você já fez | Onde no guia |
|-------------------|--------------|
| Ler teclado, `strconv`, `switch` | [1-sortNumber](../1-sortNumber/), cap. 05–07 |
| Módulo `LearningGo` | [go.mod](../go.mod), cap. 06 |
| Hello World | [main.go](../main.go), cap. 01–02 |

O app de tarefas é o **próximo nível**: persistência, testes e estrutura de pastas.

## Prática final

1. Implemente o todo CLI completo
2. Adicione teste para `MarcarFeita` com ID inexistente (deve retornar erro)
3. (Desafio) API HTTP + CLI compartilhando o mesmo `JSONStore`

## Parabéns

Com os 14 capítulos você tem base para:

- Ferramentas de linha de comando
- APIs REST simples
- Serviços com concorrência
- Projetos organizados e testáveis

Continue em [go.dev/doc](https://go.dev/doc) e na comunidade [Gopher Brasil](https://go.dev/community).

[← Índice](README.md) | [← README da raiz](../README.md)
