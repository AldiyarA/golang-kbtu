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

var(
	users []*api.User = []*api.User{
		&api.User{
			Id:        1,
			Username:  "AxDI",
			Email:     "adelov.aldiyar@gmail.com",
			Password:  "qweasd123",
			Firstname: "Aldiyar",
			Lastname:  "Adelov",
		},
		&api.User{
			Id:        2,
			Username:  "Yeroma",
			Email:     "zhanibekov02@gmail.com",
			Password:  "qweasd123",
			Firstname: "Yergeldi",
			Lastname:  "Zhanibek",
		},
		&api.User{
			Id:        3,
			Username:  "Kambar_Z",
			Email:     "zhamauov02@mail.ru",
			Password:  "qweasd123",
			Firstname: "Kambar",
			Lastname:  "Zhamauov",
		},
	}
)

func main() {
	ctx := context.Background()

	connStartTime := time.Now()
	conn, err := grpc.Dial("localhost" + port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", port, err)
	}
	log.Printf("connected in %d microsec", time.Now().Sub(connStartTime).Microseconds())


	UserClient := api.NewUserServiceClient(conn)
	for _, user:= range users{
		_, err := UserClient.Create(ctx, user)
		if err != nil {
			return
		}
	}

	computerClient := api.NewComputerServiceClient(conn)
	id1 := int64(2)

	computer1, err := computerClient.Get(ctx, &api.Id{Id: id1})
	if err != nil{
		log.Fatal(err)
	}else{
		log.Println(computer1)
	}
}
