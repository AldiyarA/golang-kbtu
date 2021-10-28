package main

import (
	"google.golang.org/grpc"
	"hw7/api"
	"hw7/internal/store"
	"log"
	"net"
	"sync"
)

func RunGRPC(store store.Store, port string, wg *sync.WaitGroup)  {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot listen to %s: %v", port, err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	computerService := store.Electronics().Computers()
	phoneService := store.Electronics().Phones()
	userService := store.Users()

	api.RegisterComputerServiceServer(grpcServer, computerService)
	api.RegisterPhoneServiceServer(grpcServer, phoneService)
	api.RegisterUserServiceServer(grpcServer, userService)

	log.Printf("Serving grpc on %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve on %v: %v", listener.Addr(), err)
	}
	wg.Done()
}
