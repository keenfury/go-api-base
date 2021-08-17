package protobuf

import (
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// --- replace protobuf import - do not remove ---
)

func StartTCPPort(port string) {
	fmt.Println("Starting TCP server...")
	tcpListener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panic("Unable to start TCP port:", err)
	}
	defer tcpListener.Close()
	s := grpc.NewServer()
	// --- replace protobuf text - do not remove ---

	reflection.Register(s)
	s.Serve(tcpListener)
}
