package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
	"github.com/seagullbird/headr-contentmgr/config"
	"github.com/seagullbird/headr-contentmgr/db"
	"github.com/seagullbird/headr-contentmgr/endpoint"
	"github.com/seagullbird/headr-contentmgr/pb"
	"github.com/seagullbird/headr-contentmgr/service"
	"github.com/seagullbird/headr-contentmgr/transport"
	repoctltransport "github.com/seagullbird/headr-repoctl/transport"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Repoctl gRPC service
	conn, err := grpc.Dial("repoctl:2018", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	// repoctl service
	repoctlsvc := repoctltransport.NewGRPCClient(conn, logger)
	// database
	dbConn, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DEVDBHOST, config.DEVDBPORT, config.DEVDBUSER, config.DEVDBNAME, config.DEVDBPASSWORD, config.DEVDBSSLMODE))
	if err != nil {
		logger.Log("error_desc", "Failed to connected to PostgreSQL", "error", err)
	}
	store := db.New(dbConn)
	var (
		service    = service.New(repoctlsvc, store, logger)
		endpoints  = endpoint.New(service, logger)
		grpcServer = transport.NewGRPCServer(endpoints, logger)
	)

	grpcListener, err := net.Listen("tcp", config.PORT)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	logger.Log("transport", "gRPC", "addr", config.PORT)
	baseServer := grpc.NewServer()
	pb.RegisterContentmgrServer(baseServer, grpcServer)

	baseServer.Serve(grpcListener)
}
