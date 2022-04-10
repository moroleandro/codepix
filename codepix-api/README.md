POC - CodePix

## Technologies used
- gRPC
- Apache Kafka
- Docker
- NestJS
- NextJS


## Architecture used

[Hexagonal](https://www.google.com/url?sa=i&url=https%3A%2F%2Ffernandofranzini.wordpress.com%2F2019%2F04%2F09%2Farquitetura-hexagonal%2F&psig=AOvVaw2_JLB5RYyDPmNehob0pNhw&ust=1649603393078000&source=images&cd=vfe&ved=0CAoQjRxqFwoTCKiX8tOhh_cCFQAAAAAdAAAAABAY)

## Structure


## Commands 

- Generate protofiles
$ protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles/ application/grpc/protofiles/*.proto