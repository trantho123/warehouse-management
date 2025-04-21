package main

import (
	"context"
	"fmt"
	"log"

	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/trantho123/warehouse-management/api"
	db "github.com/trantho123/warehouse-management/db/sqlc"
	"github.com/trantho123/warehouse-management/utils"
)

var interruptSignals = []os.Signal{os.Interrupt, os.Kill}

func main() {
	fmt.Println("Welcome to the Warehouse Management System!")

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ctx := context.Background()

	connPool, err := ConnectDB(ctx, &config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	store := db.NewStore(connPool)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}
	server.Start(config.HTTP_SERVER_ADDRESS)

}

func ConnectDB(ctx context.Context, config *utils.Config) (*pgxpool.Pool, error) {

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// Test the connection
	if err := connPool.Ping(ctx); err != nil {
		return nil, err
	}

	return connPool, nil
}
