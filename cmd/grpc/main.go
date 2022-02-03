package main

import (
	"fmt"
	"net"

	"github.com/keenfury/go-api-base/config"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// --- replace grpc import once - do not remove ---
	// --- replace grpc import - do not remove ---
)

func main() {
	tcpListener, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		log.Panic("Unable to start GRPC port:", err)
	}
	defer tcpListener.Close()
	s := grpc.NewServer()

	// --- replace grpc text - do not remove ---

	reflection.Register(s)
	fmt.Printf("Starting GRPC server on port: %s...\n", config.GrpcPort)
	s.Serve(tcpListener)
}
