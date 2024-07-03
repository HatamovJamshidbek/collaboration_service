package main

import (
	"collaboration_service/config"
	pb "collaboration_service/genproto"
	"collaboration_service/service"
	"collaboration_service/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cnf := config.Config{}
	db, err := postgres.ConnectionDb(&cnf)
	if err != nil {
		log.Fatalf("error:->%s", err.Error())
	}
	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("error:->%s", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCollaborationServiceServer(grpcServer, service.NewCollaborationService(postgres.NewInvasionRepository(db), postgres.NewCollaborationRepositoryRepository(db), postgres.NewCommentRepositoryRepository(db)))
	log.Printf("Listening:%d", listen.Addr())
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("error:->%s", err.Error())
	}
}
