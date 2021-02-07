# Microsserviço CodePix

Esse microsserviço tem o objetivo de ser um hub de transações entre os bancos que simularemos nesta POC.

## Como executar

Utilizo Docker para que todos os serviços utilizados neste projeto sejam executados..

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina apenas execute:
`docker-compose up -d`

### Como executar a aplicação
- Acesse o container da aplicação executando: `docker exec -it codepix_app bash`
- Rode `go run main.go grpc` ou `go run main.go grpc --port 50052`

### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
- PgAdmin
- Apache Kafka
- Criador dos tópicos a serem utilizados pelo Kafka
- Confluent control center
- ZooKeeper

---
## Comandos utilizados 

#### Comando para gerar arquivos pb(/codepix/application/grpc/pb/)
```buildoutcfg
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
```
.
#### Comando interativo do servidor gRPC - container: codepix_app_1
```buildoutcfg
evans -r repl

##examples##
$ call RegisterPixKey
> cpf
> 123456789
> 1

$ call Find
> qrcode
> testeteste
```

#### Comando Cobra (command line interface) - codepix/cmd/grpc.go 
```buildoutcfg
cobra init --pkg-name codepix
```

---

### Autor
@moroleandro

 
