package main

import (
	"context"
	"hw7/internal/http"
	"hw7/internal/store/inmemory"
	"log"
	"sync"
)

func main() {
	store := inmemory.NewDB()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go RunGRPC(store, ":8080", wg)

	srv := http.NewServer(context.Background(), ":8081", store)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()
	wg.Wait()
}
