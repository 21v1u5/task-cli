# task-cli

CLI para gerenciamento de tarefas com persistência local.
Construído como projeto fundacional do meu roadmap de engenharia de software.

## O que resolve

Ferramenta de linha de comando para criar, listar, concluir e remover
tarefas, com persistência em SQLite local — sem servidor, sem dependência
de rede.

## Arquitetura

cmd/main.go → cli.Handler → storage.Store (interface)
                                   ↓
                            storage.SQLiteStore

A camada de storage é abstraída por uma interface, permitindo trocar
SQLite por qualquer outro banco sem alterar os comandos CLI.

## Como rodar

    git clone https://github.com/seu-usuario/task-cli
    make build

    ./bin/task-cli add "Estudar Go" -desc "Capítulo 3 do livro"
    ./bin/task-cli list
    ./bin/task-cli done 1
    ./bin/task-cli delete 2

## Decisões técnicas

**Por que SQLite e não um arquivo JSON?**
SQLite oferece transações ACID e queries. Um arquivo JSON exigiria
carregar tudo em memória e reescrever o arquivo inteiro a cada operação.

**Por que interface Storage em vez de uso direto do SQLite?**
Permite testar os comandos CLI com um MemoryStore em vez de banco real.
Os testes rodam sem I/O e sem estado entre execuções.

**O que eu melhoraria**
- Adicionar prioridade e data de vencimento nas tarefas
- Filtro por status no list (atualmente flag --all)
- Subcomando edit para atualizar título e descrição

## Testes

    make test   # roda com -race detector ativado

CI: GitHub Actions roda testes e build a cada push.