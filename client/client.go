package client

import (
	"context"
	pb "github.com/drewinner/gutils/rpc/proto/rpc"
	"time"
)

// Invoke /**
/**
*	调用服务端
*	@param:id int32 任务id
*	@param:logId 日志id
*	@param:taskHandler 任务标识
*	@param:params 参数
*	@param:timeout 超时时间、如果设置为0、设置为1秒过期
 */
func Invoke(address string, id, logId int32, taskHandler, params string, timeout int) (resp *pb.TaskResp, err error) {
	if timeout == 0 {
		timeout = 1
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := Pool.GetClient(address)
	if err != nil {
		return nil, err
	}
	r, err := client.Call(ctx, &pb.TaskReq{
		Id:         id,
		LogId:      logId,
		JobHandler: taskHandler,
		Params:     params,
	})
	return r, err
}
