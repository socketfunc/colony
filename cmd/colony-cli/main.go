package main

import (
	"context"
	"fmt"
	"log"

	v1beta1 "github.com/socketfunc/colony/proto/worker/v1beta1"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := v1beta1.NewWorkerClient(conn)
	v := []interface{}{int32(60), int32(40)}
	args, err := msgpack.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	in := &v1beta1.InvokeRequest{
		Name:      "add",
		Namespace: "default",
		Args:      args,
	}
	out, err := client.Invoke(context.Background(), in)
	if err != nil {
		log.Fatal(err)
	}
	var ret interface{}
	if err := msgpack.Unmarshal(out.Result, &ret); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}
