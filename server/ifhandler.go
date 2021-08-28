package server

import "context"

/**
*	业务逻辑处理接口
 */
type Response struct {
	status int32       //状态
	msg    string      //信息
	data   interface{} //数据
	ext    string      //扩展字段
}
type Handler interface {
	HandlerFunc(ctx context.Context, request string) Response
}
