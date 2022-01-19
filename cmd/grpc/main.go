package main

import (
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/keenfury/api/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// --- replace grpc import once - do not remove ---
	// --- replace grpc import - do not remove ---
)

func main() {
	tcpListener, err := net.Listen("tcp", ":"+config.TCPPort)
	if err != nil {
		log.Panic("Unable to start GRPC port:", err)
	}
	defer tcpListener.Close()
	s := grpc.NewServer()

	// --- replace grpc text - do not remove ---

	reflection.Register(s)
	fmt.Printf("Starting GRPC server on port: %s...\n", config.TCPPort)
	s.Serve(tcpListener)
}
