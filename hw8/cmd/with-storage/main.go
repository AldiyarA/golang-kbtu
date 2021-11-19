package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hw8/internal/http"
	"hw8/internal/store/postgres"
)

func main() {
	urlExample := "postgres://shikimori:adminadmin@localhost:5432/shikimori"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	//cache, err := lru.New2Q(6)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//if err != nil {
	//	panic(err)
	//}
	srv := http.NewServer(
		context.Background(),
		http.WithAddress(":8080"),
		http.WithStore(store),
		//http.WithCache(cache),
		http.WithRedis(rdb),
		)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
