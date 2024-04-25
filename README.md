# Merge-Solver

O Merge-Solver é um conjunto de ferramentas que inclui uma API em Go e uma CLI em Rust que facilitam o merge automática de arquivos de código sem precisar alterar manualmente. 
Prioriza as alterações dos arquivos mais novos, enquanto integra perfeitamente o conteúdo das versões mais antigas, garantindo que as atualizações sejam incorporadas e a compatibilidade seja mantida.

## Funcionalidades

- **API (Go):** Serviço para realizar o merge dos arquivos de código.
- **CLI (Rust):** Interage com a API para enviar arquivos de código e receber o resultado com merge feito.

## Primeiros Passos

Estas instruções permitirão que você tenha uma cópia do projeto em execução na sua máquina local para fins de desenvolvimento e teste.

### Pré-requisitos

Antes de começar, certifique-se de que você possui o seguinte instalado:
- [Go](https://golang.org/doc/install) (versão 1.22 ou superior)
- [Rust](https://www.rust-lang.org/tools/install)
- [Git](https://git-scm.com/book/pt-br/v2/Começando-Instalando-o-Git)

### Configuração

Crie um arquivo `.env` no diretório `api` com o seguinte conteúdo:

```py
OPENAI_SECRET_KEY="secret_key"
```

Substitua pela sua chave de API OpenAI real.

### Rodando Localmente

1. **Configurando a API em Go**

   Navegue até o diretório da API e construa o serviço:

   ```bash
   cd api/
   go run cmd/app/main.go # ou apenas `air` para rodar com hot reload
   ```

2. **Configurando a CLI em Rust**

   Navegue até o diretório da CLI e construa o executável:

   ```bash
   cd cli/
   cargo build --release
   ./target/release/cli --old caminho/para/arquivo_antigo.rs --new caminho/para/arquivo_novo.rs
   ```

### Usando a CLI

Com a API em execução, você pode usar a CLI para mesclar dois arquivos:

```bash
./target/release/file_diff --old caminho/para/arquivo_antigo.rs --new caminho/para/arquivo_novo.rs
```

## Como Funciona

- A **API** possui um endpoint POST `/merge` que aceita um JSON com dois campos: `old` e `new`, representando o código antigo e o novo.
- A **CLI** lê o conteúdo dos arquivos fornecidos, envia-os para a API e exibe o resultado com o merge realizado.