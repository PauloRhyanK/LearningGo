# 07 — Biblioteca padrão

A **stdlib** de Go é rica e estável. Você não precisa instalar pacotes externos para tarefas comuns.

Documentação completa: [pkg.go.dev/std](https://pkg.go.dev/std)

## Entrada e saída

| Pacote | Para que serve |
|--------|----------------|
| **fmt** | Imprimir, formatar strings (`Println`, `Printf`, `Sprintf`) |
| **os** | Sistema operacional: arquivos, `Stdin`/`Stdout`, variáveis de ambiente |
| **bufio** | I/O com **buffer**; `NewScanner` para ler linhas do teclado |
| **strings** | Manipular strings (`TrimSpace`, `Split`, `Contains`) |
| **strconv** | Converter string ↔ número (`Atoi`, `ParseInt`, `Itoa`) |

### Exemplo: ler do teclado

```go
scanner := bufio.NewScanner(os.Stdin)
fmt.Println("Digite algo:")
scanner.Scan()
texto := strings.TrimSpace(scanner.Text())
```

Usado no exercício [1-sortNumber](../1-sortNumber/main.go).

## Números e aleatoriedade

| Pacote | Para que serve |
|--------|----------------|
| **math** | Funções matemáticas (`Sqrt`, `Abs`) |
| **math/rand/v2** | Números pseudoaleatórios (Go 1.22+) |

```go
import "math/rand/v2"

n := rand.IntN(101) // 0 a 100
```

## Tempo

| Pacote | Para que serve |
|--------|----------------|
| **time** | Datas, durações, timers, `Sleep` |

```go
time.Sleep(2 * time.Second)
agora := time.Now()
```

## Dados estruturados

| Pacote | Para que serve |
|--------|----------------|
| **encoding/json** | JSON ↔ struct/map |
| **encoding/csv** | Arquivos CSV |
| **bytes** | Buffer de bytes |

## Rede

| Pacote | Para que serve |
|--------|----------------|
| **net/http** | Cliente e servidor HTTP |
| **net/url** | Parse de URLs |

## Concorrência

| Pacote | Para que serve |
|--------|----------------|
| **sync** | `Mutex`, `WaitGroup`, `Once` |
| **context** | Cancelamento e deadlines em APIs |

## Erros

| Pacote | Para que serve |
|--------|----------------|
| **errors** | `New`, `Is`, `As`, `Unwrap` |
| **fmt** | `Errorf` com `%w` |

## Utilitários

| Pacote | Para que serve |
|--------|----------------|
| **sort** | Ordenar slices |
| **io** | Interfaces `Reader`/`Writer` |
| **path/filepath** | Caminhos multiplataforma |
| **log** | Log simples |
| **log/slog** | Log estruturado (Go 1.21+) |

## Tabela rápida por tarefa

| Tarefa | Pacotes |
|--------|---------|
| CLI que lê input | `bufio`, `os`, `fmt`, `strings`, `strconv` |
| Ler/escrever arquivo | `os`, `io`, `bufio` |
| API REST | `net/http`, `encoding/json` |
| Testes | `testing` (capítulo 09) |
| Paralelismo | `sync`, goroutines (capítulo 11) |

## Prática

1. Leia nome e idade do teclado e imprima uma frase formatada
2. Gere 5 números aleatórios entre 1 e 6 (simular dado)
3. Liste variáveis de ambiente com `os.Environ()` (cuidado: saída longa)

## Próximo passo

[08 — Compilar e executar](08-compilar-executar.md)

[← Índice](README.md)
