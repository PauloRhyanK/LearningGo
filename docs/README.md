# Guia Go — Índice

Documentação em português para aprender Go do zero até construir um software simples.

## Ordem de leitura

| # | Capítulo | Descrição |
|---|----------|-----------|
| 01 | [Instalação](01-instalacao.md) | Instalar Go no Windows e primeiro `go run` |
| 02 | [Sintaxe básica](02-sintaxe-basica.md) | `package`, variáveis, tipos, `nil` |
| 03 | [Tipos, structs, interfaces](03-tipos-structs-interfaces.md) | Slices, maps, structs, métodos |
| 04 | [Funções e erros](04-funcoes-erros.md) | Funções, múltiplos retornos, `error` |
| 05 | [Controle de fluxo](05-controle-de-fluxo.md) | `if`, `for`, `switch`, `range` |
| 06 | [Pacotes e módulos](06-pacotes-modulos.md) | `go mod`, imports, exports |
| 07 | [Biblioteca padrão](07-biblioteca-padrao.md) | Pacotes mais usados da stdlib |
| 08 | [Compilar e executar](08-compilar-executar.md) | `go run`, `go build`, binários |
| 09 | [Testes](09-testes.md) | `go test`, tabelas de casos |
| 10 | [JSON](10-json.md) | `encoding/json`, struct tags |
| 11 | [Concorrência](11-concorrencia.md) | Goroutines, channels, `WaitGroup` |
| 12 | [HTTP básico](12-http-basico.md) | Servidor e cliente com `net/http` |
| 13 | [Layout de projeto](13-layout-projeto.md) | `cmd/`, `internal/`, organização |
| 14 | [Projeto completo](14-projeto-completo.md) | App de tarefas: CLI + API |

[← Voltar ao README da raiz](../README.md)

---

## Convenções deste guia

### Blocos de código

- Exemplos completos podem ser salvos como `main.go` e executados com `go run .`
- `// ...` indica código omitido no meio do exemplo
- Comandos de terminal usam PowerShell no Windows

### Caixas especiais

> **Idioma Go** — padrão recomendado pela comunidade Go (idiomático).

> **Erro comum** — armadilha frequente para iniciantes.

> **Dica** — atalho ou boa prática.

### Estrutura de cada capítulo

1. Conceitos explicados com exemplos curtos
2. **Prática** — exercícios para fazer sozinho
3. **Próximo passo** — link para o capítulo seguinte

---

## Glossário

| Termo | Significado |
|-------|-------------|
| **Package (pacote)** | Unidade de organização do código; cada pasta com `.go` do mesmo `package` |
| **Module (módulo)** | Projeto Go versionado; definido em `go.mod` |
| **Main** | Pacote executável; precisa de `func main()` |
| **Export** | Nome que começa com maiúscula; visível fora do pacote |
| **Receiver** | Parâmetro especial de métodos: `func (p Pessoa) Nome()` |
| **Interface** | Conjunto de métodos; tipo implícito em Go |
| **Slice** | “Fatia” dinâmica de array; `[]int` |
| **Zero value** | Valor padrão de um tipo (`0`, `""`, `false`, `nil`) |
| **nil** | Ausência de valor em ponteiros, slices, maps, channels, interfaces, funções |
| **Goroutine** | Função executada em paralelo (leve, gerenciada pelo runtime) |
| **Channel** | Canal de comunicação entre goroutines |
| **Buffer** | Área temporária na memória para I/O mais eficiente (`bufio`) |
| **Handler** | Função que trata uma requisição HTTP |
| **Struct tag** | Metadado em campo de struct, ex.: `` `json:"nome"` `` |
