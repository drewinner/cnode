package server

import (
	"context"
	"errors"
	"sync"
)

type HandlerFunc func(ctx context.Context, request string) Response

// Exec /** 实现接口

func (f HandlerFunc) HandlerFunc(ctx context.Context, req string) Response {
	return f(ctx, req)
}

//字符串-接口对应关系、将来执行
var (
	funNameMap map[string]Handler
	l          sync.Mutex
)

func init() {
	funNameMap = make(map[string]Handler)
}

func Get(name string) (Handler, error) {
	l.Lock()
	defer l.Unlock()
	if f, ok := funNameMap[name]; ok {
		return f, nil
	}
	return nil, errors.New("handler not found")
}

// Set /** 设置名字和接口对应关系、此函数存在对应关系、不进行覆盖
func Set(name string, f Handler) bool {
	l.Lock()
	defer l.Unlock()
	if _, ok := funNameMap[name]; !ok {
		funNameMap[name] = f
	}
	return true
}
