package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"github.com/f-go/grpc-playground"
)

func main(){
	cc, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}

	// c := service.NewHintClient(cc)
	// req := service.ContrivedMessageRequest{
	// 	Body: "How's the weather? 🤓",
	// }
	//
	// res, err := c.ContrivedMethod(context.Background(), &req)
	// if err != nil {
	// 	log.Fatal("error while calling: %v", err)
	// }
	// fmt.Println(res.Body)
}