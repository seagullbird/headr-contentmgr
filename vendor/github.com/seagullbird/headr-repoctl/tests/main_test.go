package tests

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq/dispatch"
	"github.com/seagullbird/headr-repoctl/endpoint"
	"github.com/seagullbird/headr-repoctl/pb"
	"github.com/seagullbird/headr-repoctl/service"
	"github.com/seagullbird/headr-repoctl/transport"
	"google.golang.org/grpc"
	"net"
	"os"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Create server
	go startServer(logger)
	time.Sleep(2 * time.Second)

	// Create client
	client := createClient(logger)
	ctx := context.Background()

	s := Site{
		ctx:    ctx,
		siteID: 1,
	}

	p := Post{
		ctx:      ctx,
		siteID:   1,
		filename: "test-file.md",
		content:  "test content\n",
	}

	// Running tests
	NewSite(t, client, s)
	WritePost(t, client, p)
	ReadPost(t, client, p)
	RemovePost(t, client, p)
	DeleteSite(t, client, s)
}

func startServer(logger log.Logger) {
	dispatcher := dispatch.FakeDispatcher{}
	var (
		svc        = service.New(dispatcher, logger)
		endpoints  = endpoint.New(svc, logger)
		grpcServer = transport.NewGRPCServer(endpoints, logger)
	)
	grpcListener, err := net.Listen("tcp", ":1234")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	logger.Log("transport", "gRPC", "addr", ":1234")
	baseServer := grpc.NewServer()
	pb.RegisterRepoctlServer(baseServer, grpcServer)

	baseServer.Serve(grpcListener)
}

func createClient(logger log.Logger) *service.Service {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		logger.Log("error", err)
		os.Exit(1)
	}
	// repoctl service
	repoctlsvc := transport.NewGRPCClient(conn, logger)
	return &repoctlsvc
}
