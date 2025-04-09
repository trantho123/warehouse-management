package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/trantho123/warehouse-management/utils"
)

var TestQuerier *Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// Connect to database
	db, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	TestQuerier = New(db)

	os.Exit(m.Run())
}
