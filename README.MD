1. 代码生成：
    1. 进入：cnode/proto
    2. 执行：protoc -I=. --go_out=./ --go-grpc_out=./ ./rpc/task.proto