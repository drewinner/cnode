package server

import (
	"context"
	"fmt"
	"testing"
)

//type HandlerFunc func(ctx context.Context,request string) Response
func Exec(ctx context.Context, request string) Response {
	fmt.Println("exec:", request, ctx)
	return Response{
		status: 1,
		msg:    "exec...",
		data:   nil,
		ext:    "",
	}
}

func TestServer_Run(t *testing.T) {
	//初始化执行器、这部分是调用方写死
	Set("TestHandler2", HandlerFunc(Exec))
	address := "127.0.0.1:8090"
	Start(address)
	fmt.Println("start finish..")
	select {}
}
