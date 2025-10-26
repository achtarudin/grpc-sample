package main

import (
	hellov1 "buf-grpc-sample/protogen/hello/v1"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})

	helloReq := hellov1.SayHelloRequest{
		Name: "Hello, Baz!",
	}

	helloReq2 := &hellov1.SayManyHellosRequest{
		Name: &hellov1.SayHelloRequest{
			Name: "Hello, Qux!",
		},
	}

	log.Println(helloReq.GetName())
	log.Println(helloReq2.GetName().GetName())

}
