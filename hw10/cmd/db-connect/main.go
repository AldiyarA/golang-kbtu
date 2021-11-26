package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func main() {
	ctx := context.Background()
	urlExample := "postgres://shikimori:adminadmin@localhost:5432/shikimori"
	conn, err := pgx.Connect(ctx, urlExample)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
	log.Println("Pinged DB")
}