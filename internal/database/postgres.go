package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/om13rajpal/gobackend/config"
)

var Pool *pgxpool.Pool

func ConnectPostgres() {
	var err error
	Pool, err = pgxpool.New(context.Background(), config.POSTGRES_URI)

	if err != nil {
		log.Fatal("Could not create PostgresSQL connection pool", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	err = Pool.Ping(ctx)
	if err != nil {
		log.Fatal("Could not connect to PostgresSQL", err)
	}

	fmt.Println("Connected to PostgresSQL")
}
