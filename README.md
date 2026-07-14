# LearningGo — Guia Go para iniciantes

Repositório de estudo da linguagem **Go** com documentação autocontida em português. O objetivo é que você consiga, apenas com este material, escrever programas, organizar projetos e construir um software simples (CLI, testes, JSON, HTTP e concorrência básica).

## Para quem é

- Programadores iniciantes ou vindos de outras linguagens (Java, JavaScript, C#, Python, etc.)
- Quem quer aprender Go do zero com exemplos práticos

## Pré-requisitos

- Terminal (PowerShell no Windows)
- Editor de código (VS Code, Cursor, GoLand, etc.)
- Conexão com a internet para instalar o Go (capítulo 01)

## Como estudar

1. Leia os capítulos na ordem em [docs/README.md](docs/README.md).
2. Digite os exemplos você mesmo — copiar e colar ajuda, mas digitar fixa o aprendizado.
3. Faça os exercícios **Prática** ao final de cada capítulo.
4. Compare com o exercício [1-sortNumber/](1-sortNumber/) quando chegar ao capítulo 05.

## Roteiro de estudo

| Capítulo | Tempo estimado | Depois você consegue… |
|----------|----------------|------------------------|
| [01 — Instalação](docs/01-instalacao.md) | 30 min | Instalar Go e rodar o primeiro programa |
| [02 — Sintaxe básica](docs/02-sintaxe-basica.md) | 1–2 h | Declarar variáveis, tipos e `nil` |
| [03 — Tipos, structs, interfaces](docs/03-tipos-structs-interfaces.md) | 2 h | Modelar dados com structs e slices |
| [04 — Funções e erros](docs/04-funcoes-erros.md) | 1–2 h | Escrever funções e tratar erros |
| [05 — Controle de fluxo](docs/05-controle-de-fluxo.md) | 1 h | Usar `if`, `for`, `switch` e `range` |
| [06 — Pacotes e módulos](docs/06-pacotes-modulos.md) | 1–2 h | Organizar código em pacotes |
| [07 — Biblioteca padrão](docs/07-biblioteca-padrao.md) | 2 h | Usar `fmt`, `os`, `bufio`, etc. |
| [08 — Compilar e executar](docs/08-compilar-executar.md) | 45 min | `go run`, `go build`, gerar binários |
| [09 — Testes](docs/09-testes.md) | 1 h | Escrever testes com `go test` |
| [10 — JSON](docs/10-json.md) | 1 h | Serializar e ler JSON |
| [11 — Concorrência](docs/11-concorrencia.md) | 2 h | Goroutines, channels e `WaitGroup` |
| [12 — HTTP básico](docs/12-http-basico.md) | 1–2 h | Criar um servidor e cliente HTTP |
| [13 — Layout de projeto](docs/13-layout-projeto.md) | 1 h | Estruturar um projeto maior |
| [14 — Projeto completo](docs/14-projeto-completo.md) | 4–8 h | Construir um app de tarefas (CLI + API) |

**Total estimado:** 20–30 horas de estudo com prática.

## Comandos rápidos

Na raiz do repositório (PowerShell):

```powershell
go version
go run .
go run ./1-sortNumber
go build -o app.exe .
go test ./...
```

No PowerShell antigo, use `;` em vez de `&&` para encadear comandos:

```powershell
cd c:\Projects\learningGO; go run .
```

## Estrutura do repositório

```
learningGO/
├── README.md           ← você está aqui
├── go.mod              ← módulo Go (LearningGo)
├── main.go             ← Hello World
├── 1-sortNumber/       ← exercício: adivinhar número
└── docs/               ← guia completo (14 capítulos)
```

## Referência oficial

Este guia não substitui a documentação oficial. Para aprofundar:

- [https://go.dev/doc/](https://go.dev/doc/)
- [https://go.dev/tour/](https://go.dev/tour/) (tour interativo)
- [https://pkg.go.dev/std](https://pkg.go.dev/std) (pacotes da biblioteca padrão)

## Leitura futura (fora deste guia)

Frameworks web (Gin, Echo), bancos de dados, Docker, deploy em nuvem e CI/CD são passos naturais depois do capítulo 14.
