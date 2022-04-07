package main

import (
	"fmt"
	"net"
	"os"

	"github.com/keenfury/go-api-base/config"
	mig "github.com/keenfury/go-api-base/tools/migration/src"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// --- replace grpc import once - do not remove ---
	// --- replace grpc import - do not remove ---
)

func main() {
	if config.UseMigration {
		err := os.MkdirAll(config.MigrationPath, 0744)
		if err != nil {
			fmt.Printf("Unable to make scripts/migrations directory structure: %s\n", err)
		}

		errVerify := mig.VerifyDBInit(config.DBDB, config.DBHost, config.DBUser, config.DBPass)
		if errVerify != nil {
			panic(errVerify)
		}
		mig.RunMigration(config.MigrationPath, config.DBHost, config.DBUser, config.DBPass, config.DBDB)
	}

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
