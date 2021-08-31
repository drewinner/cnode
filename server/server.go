package server

import (
	"context"
	"fmt"
	"github.com/drewinner/gnode/common"
	"time"
)
import pb "github.com/drewinner/gnode/proto/rpc"

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
	case common.RUN_SCHEMA_GLUE_GO:
		handler, err := Get(req.JobHandler)
		if err != nil {
			return nil, err
		}
		r := handler.HandlerFunc(ctx, req.Params)
		rs.Status = r.status
		rs.LogMsg = r.msg
	case common.RUN_SCHEMA_GLUE_SHELL:
		output, err := common.Exec(ctx, req.GetParams())
		if err == nil {
			rs.Status = 1
			rs.LogMsg = output
		} else {
			rs.LogMsg = err.Error()
		}
		rs.ExecEndTime = time.Now().String()
	case common.RUN_SCHEMA_GLUE_HTTP:
		fmt.Print("http")
	default:
		fmt.Println("default..")
	}
	return rs, err
}
