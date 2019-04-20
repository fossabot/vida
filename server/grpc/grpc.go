package grpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gangachris/vida/config"
	"github.com/gangachris/vida/db"
	"github.com/gangachris/vida/pb"
)

// GRPC represents the GRPC server struct
type GRPC struct {
	config config.Config
	store  *db.Storage
}

// NewServer will return a new instance in GRPC Server
func NewServer(cfg config.Config, store *db.Storage) GRPC {
	return GRPC{
		config: cfg,
		store:  store,
	}
}

// Start starts the GRPC Server
func (g GRPC) Start() error {
	lis, err := net.Listen("tcp", ":"+g.config.GRPCPort())
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMoviesRequestsServer(grpcServer, &moviesRequestServer{store: g.store})
	// tls ?
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}
