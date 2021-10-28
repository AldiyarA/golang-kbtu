package main

import (
	"context"
	"google.golang.org/grpc"
	"hw7/api"
	"log"
	"time"
)

const (
	port = ":8080"
)

func main() {
	ctx := context.Background()

	connStartTime := time.Now()
	conn, err := grpc.Dial("localhost" + port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", port, err)
	}
	log.Printf("connected in %d microsec", time.Now().Sub(connStartTime).Microseconds())

	computerClient := api.NewComputerServiceClient(conn)
	//phoneClient := api.NewPhoneServiceClient(conn)
	id1 := int64(2)
	//id2 := int64(2)

	computer1, err := computerClient.Get(ctx, &api.Id{Id: id1})
	if err != nil{
		log.Fatal(err)
	}else{
		log.Println(computer1)
	}

	//computer2, err := computerClient.Get(ctx, &api.Id{Id: id2})
	//if err != nil{
	//	log.Fatal(err)
	//}else{
	//	log.Println(computer2)
	//}
	//
	//phone1, err := phoneClient.Get(ctx, &api.Id{Id: id1})
	//if err != nil{
	//	log.Fatal(err)
	//}else{
	//	log.Println(phone1)
	//}
	//phone2, err := phoneClient.Get(ctx, &api.Id{Id: id2})
	//if err != nil{
	//	log.Fatal(err)
	//}else{
	//	log.Println(phone2)
	//}
}
