# Taste Match API

Este é o README para o projeto Taste Match API, uma API construída em Go. O guia a seguir descreve como configurar e executar o projeto localmente usando Docker e Docker Compose.

## Tecnologias Utilizadas

  * **Go (Golang):** Linguagem de programação principal da API.
  * **PostgreSQL:** Banco de dados relacional para armazenamento de dados.
  * **Docker & Docker Compose:** Ferramentas para criação e orquestração dos contêineres da aplicação e do banco de dados.
  * **Gin Gonic:** Framework web para Go utilizado na construção das rotas da API.
  * **Gorm:** ORM para Golang

## Pré-requisitos

Antes de começar, certifique-se de que você tem as seguintes ferramentas instaladas na sua máquina:

  * [Docker](https://docs.docker.com/get-docker/)
  * [Docker Compose](https://docs.docker.com/compose/install/)

## Como Executar o Projeto

Siga os passos abaixo para ter o ambiente completo rodando localmente.

### 1\. Clone o Repositório

Primeiro, clone o repositório do projeto para a sua máquina local, caso ainda não o tenha feito.

```bash
git clone git@github.com:Ze-Victor/taste-match-api.git
cd taste-match-api
```

### 2\. Configure as variavéis de ambiente e `docker-compose.yml`

O arquivo `docker-compose.yml` é o coração do nosso ambiente. Ele contém todas as variáveis de ambiente necessárias para a API e para o banco de dados. Verifique se o seu arquivo está configurado corretamente. O modelo final que desenvolvemos é este:

```yaml
# docker-compose.yml

version: '3.8'

services:
  # Serviço para a sua API em Go
  api:
    build: .
    container_name: taste-match-api
    # Mapeando a porta correta que sua API usa
    ports:
      - "8000:8000"
    # Variáveis de ambiente que sua aplicação Go vai usar
    environment:
      - DB_HOST=db # <--- IMPORTANTE: Deve ser o nome do serviço do banco
      - DB_USER=user_taste_match
      - DB_PASS=db_pass_default
      - DB_NAME=taste_match_db
      - DB_PORT=5432
      - SERVER_PORT=8000
    depends_on:
      db:
        condition: service_healthy

  # Serviço para o banco de dados PostgreSQL
  db:
    image: postgres:16-alpine
    container_name: postgres-db
    healthcheck:
      # Verificando com o usuário e banco corretos
      test: ["CMD-SHELL", "pg_isready -U user_taste_match -d taste_match_db"] # <--- MUDANÇA
      interval: 5s
      timeout: 5s
      retries: 5
    # O PostgreSQL usa essas variáveis para criar o banco na primeira vez.
    # Elas PRECISAM ser iguais às da sua API.
    environment:
      - POSTGRES_USER=user_taste_match # <--- MUDANÇA
      - POSTGRES_PASSWORD=db_pass_default # <--- MUDANÇA
      - POSTGRES_DB=taste_match_db # <--- MUDANÇA
    volumes:
      - postgres-data:/var/lib/postgresql/data
    # Você não precisa expor a porta 5432 do banco para o host,
    # a menos que queira conectar com um cliente de DB (ex: DBeaver).
    # Se precisar, descomente a linha abaixo.
    ports:
      - "5432:5432"

volumes:
  postgres-data:
```

```yaml
# docker-compose.yml

APP_NAME=taste-match-api
GO_ENV=development

DB_ENABLE_LOG=true
DB_HOST=db
DB_NAME=taste_match_db
DB_PORT=5432
DB_TYPE=postgresql
DB_PASS=db_pass_default
DB_PASS_MIGRATION=db_pass_migration
DB_USER=user_taste_match
DB_USER_MIGRATION=user_taste_match_migration

SERVER_BASE_PATH=v1
SERVER_PORT=8000
```

**Importante:** Ajuste as senhas e outros valores conforme necessário, mas garanta que as credenciais entre o serviço `api` e `db` sejam sempre idênticas.

### 3\. Suba os Contêineres

Com o arquivo `docker-compose.yml` pronto, execute o seguinte comando na raiz do projeto:

```bash
docker-compose up --build
```

  * `up`: Inicia os contêineres definidos no arquivo.
  * `--build`: Força a reconstrução da imagem da API, garantindo que quaisquer alterações no código Go sejam aplicadas.

Este comando irá baixar a imagem do PostgreSQL, construir a imagem da sua API e iniciar os dois contêineres. Você verá os logs de ambos os serviços no seu terminal.

## Verificação

Para garantir que tudo está funcionando corretamente:

### 1\. Cheque os Contêineres

Abra um **novo terminal** e execute o comando:

```bash
docker ps
```

Você deve ver dois contêineres em execução (`taste-match-api` e `postgres-db`), ambos com o status `Up`.

### 2\. Teste um Endpoint da API

Use uma ferramenta como `curl` ou Postman para fazer uma requisição a um dos seus endpoints. Por exemplo, para a rota que testamos:

```bash
curl http://localhost:8000/v1/health
```

Se tudo estiver correto, você receberá a seguinte resposta JSON:

```json
{"data":{"id":1,"name":"Exemplo"},"message":"Busca realizada com sucesso!"}
```

## Parando a Aplicação

Para parar todos os contêineres, volte ao terminal onde o `docker-compose up` está rodando e pressione `Ctrl + C`. Em seguida, para garantir que a rede e os contêineres sejam removidos, execute:

```bash
docker-compose down
```

Se você também quiser remover o volume de dados do banco de dados (**CUIDADO: isso apagará todos os dados salvos**), use:

```bash
docker-compose down -v
```


