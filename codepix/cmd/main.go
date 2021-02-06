package main

import (
    "github.com/moroleandro/poc-codepix/codepix/infrastructure/grpc"
    "github.com/moroleandro/poc-codepix/codepix/infrastructure/db"
    "gorm.io/jinzhu/gorm"
    "os"
)

var database *gorm.DB

func main() {
    database = db.ConnectDB(os.Getenv("ENV"))
    grpc.StartGrpcServer(database, 50051)
}