package server

import (
	"context"
	"time"
)
import pb "github.com/drewinner/gutils/rpc/proto/rpc"

type Server struct {
	pb.UnimplementedTaskServiceServer
}

func (s *Server) Call(ctx context.Context, req *pb.TaskReq) (*pb.TaskResp, error) {
	handler, err := Get(req.JobHandler)
	if err != nil {
		return nil, err
	}
	start := time.Now().String()
	r := handler.HandlerFunc(ctx, req.Params)
	return &pb.TaskResp{
		Id:            req.Id,
		LogId:         req.LogId,
		Status:        r.status,
		ExecStartTime: start,
		ExecEndTime:   time.Now().String(),
		LogMsg:        r.msg,
	}, nil
}
