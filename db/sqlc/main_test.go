package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/trantho123/warehouse-management/utils"
)

var TestStore Store

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Connect to database using pgx
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer connPool.Close()

	TestStore = NewStore(connPool)

	os.Exit(m.Run())
}
