package server

import (
	"context"
	"fmt"
	"time"
)
import pb "github.com/drewinner/gnode/proto/rpc"

const (
	GLUE_GO = iota
	GLUE_SHELL
	GLUE_HTTP
)

type Server struct {
	pb.UnimplementedTaskServiceServer
}

/**
*	runSchema判断命令类型 运行模式 1 glue模式go 2 glue shell模式 3 glue http模式
 */
func (s *Server) Call(ctx context.Context, req *pb.TaskReq) (*pb.TaskResp, error) {
	var err error
	rs := &pb.TaskResp{
		Id:            req.Id,
		LogId:         req.LogId,
		Status:        3,
		ExecStartTime: time.Now().String(),
		ExecEndTime:   "",
		LogMsg:        "执行失败",
	}
	switch req.RunSchema {
	case GLUE_GO:
		handler, err := Get(req.JobHandler)
		if err != nil {
			return nil, err
		}
		r := handler.HandlerFunc(ctx, req.Params)
		rs.Status = r.status
		rs.ExecEndTime = time.Now().String()
		rs.LogMsg = r.msg
	case GLUE_SHELL:
		fmt.Println("shell")
	case GLUE_HTTP:
		fmt.Print("http")
	default:
		fmt.Println("default..")
	}
	return rs, err
}
