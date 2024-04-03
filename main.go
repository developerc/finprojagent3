package main

import (
	"grpc/client"
	"grpc/server"
)

func main() {
	client.RegisterAgent()
	server.CreateOrchGRPCserver()
}
