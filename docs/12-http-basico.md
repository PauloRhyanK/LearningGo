# 12 — HTTP básico

O pacote `net/http` implementa cliente e servidor HTTP.

## Servidor mínimo

```go
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("servindo em :8080")
	http.ListenAndServe(":8080", nil)
}
```

Teste no navegador ou:

```powershell
curl http://localhost:8080/Go
```

## Handler com método

```go
func tarefas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "listar")
	case http.MethodPost:
		fmt.Fprintln(w, "criar")
	default:
		http.Error(w, "método não permitido", http.StatusMethodNotAllowed)
	}
}
```

## JSON na API

```go
import "encoding/json"

type Resposta struct {
	OK   bool   `json:"ok"`
	Dado string `json:"dado"`
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Resposta{OK: true, Dado: "teste"})
}
```

`NewEncoder` escreve direto na resposta — eficiente para APIs.

## Ler corpo da requisição POST

```go
var body struct {
	Titulo string `json:"titulo"`
}
if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}
defer r.Body.Close()
```

## Cliente HTTP (GET)

```go
resp, err := http.Get("https://api.github.com")
if err != nil {
	return err
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
if err != nil {
	return err
}
fmt.Println(resp.Status, len(body))
```

## `http.Server` customizado (opcional)

```go
mux := http.NewServeMux()
mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
})

srv := &http.Server{Addr: ":8080", Handler: mux}
srv.ListenAndServe()
```

## Códigos de status comuns

| Código | Significado |
|--------|-------------|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 404 | Not Found |
| 500 | Internal Server Error |

```go
http.Error(w, "não encontrado", http.StatusNotFound)
w.WriteHeader(http.StatusCreated)
```

## Prática

1. Servidor em `:8080` com rota `/ping` retornando `pong`
2. Rota POST `/echo` que devolve o JSON recebido
3. Cliente que faz GET em sua rota `/ping` e imprime o corpo

## Próximo passo

[13 — Layout de projeto](13-layout-projeto.md)

[← Índice](README.md)
