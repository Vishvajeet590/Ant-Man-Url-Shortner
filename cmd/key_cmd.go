package main

import (
	"Ant-Man-Url/infrastructure/repository"
	"Ant-Man-Url/usecase/key"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"log"
	"os"
)

func main() {
	//dbUrl := "postgres://vishwajeet:vishvapriya123@localhost:5432/keystore"
	url := "postgres://upoaygnbwzrvic:8b648c57bafc6fb2cece95182a10f46af1224b9d09339aa1571906ed9cf54d91@ec2-52-31-217-108.eu-west-1.compute.amazonaws.com:5432/db40dr9bkunn0j"
	connection, _ := pq.ParseURL(url)
	connPgx, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		println("Errrorr...")
	}
	defer connPgx.Close(context.Background())

	repo := repository.NewKeyOutSql(connPgx)
	service := key.NewService(repo)
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("illegal length of arguments\n")
		return
	}

	if args[1] == "addInstance" {
		start, end, err := service.GetRangeConfig(args[2])
		if err != nil {
			fmt.Printf("Error : %s", err)
			return
		}
		fmt.Printf("Cofig \nStart : %d\nEnd : %d", start, end)
		fmt.Printf("Adding instance...\n")

		err = service.AddKeyOut(start, end)
		if err != nil {
			fmt.Printf("Error : %s\n", err)
			return
		}
		log.Printf("Keys for the instance are added...\nYou can run the new instance....")

	}

}
