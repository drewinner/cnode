package client

import (
	pb "github.com/drewinner/gutils/rpc/proto/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"sync"
	"time"
)

var (
	kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}
	Pool = &RpcPool{
		conns: make(map[string]*Client),
	}
)

type Client struct {
	conn       *grpc.ClientConn
	taskClient pb.TaskServiceClient
}

type RpcPool struct {
	conns map[string]*Client
	mu    sync.RWMutex
}

// GetClient /**
/**
*	获取client
*	@param:address 地址 ip||hostname:port 127.0.0.1:8081
 */
func (p *RpcPool) GetClient(address string) (pb.TaskServiceClient, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if client, ok := p.conns[address]; ok {
		return client.taskClient, nil
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	c := pb.NewTaskServiceClient(conn)
	//设置链接map
	Pool.conns[address] = &Client{
		conn:       conn,
		taskClient: c,
	}
	return c, nil
}
