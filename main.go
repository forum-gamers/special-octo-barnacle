package main

import (
	"log"
	"net"
	"os"

	"github.com/forum-gamers/special-octo-barnacle/database"
	h "github.com/forum-gamers/special-octo-barnacle/helpers"
	"github.com/forum-gamers/special-octo-barnacle/interceptor"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	h.PanicIfError(godotenv.Load())
	database.Conn()

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "50057"
	}

	lis, err := net.Listen("tcp", ":"+addr)
	if err != nil {
		log.Fatalf("Failed to listen : %s", err.Error())
	}

	interceptor := interceptor.NewInterCeptor()
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptor.Logging, interceptor.UnaryAuthentication),
	)

	log.Printf("Starting to serve in port : %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %s", err.Error())
	}
}
