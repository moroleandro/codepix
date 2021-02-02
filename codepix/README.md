## Microsserviço CodePix

Esse microsserviço tem o objetivo de ser um hub de transações entre os bancos que simularemos nesta POC.

## Como executar

Utilizo Docker para que todos os serviços utilizados neste projeto sejam executados..

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina apenas execute:
`docker-compose up -d`

### Como executar a aplicação
- Acesse o container da aplicação executando: `docker exec -it codepix_app bash`
- Rode `go run cmd/codepix/main.go`

### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
- PgAdmin
- Apache Kafka
- Criador dos tópicos a serem utilizados pelo Kafka
- Confluent control center
- ZooKeeper

## Autor
@moroleandro

 
