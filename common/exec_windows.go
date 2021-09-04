// +build windows

package common

import (
	"context"
	"os/exec"
	"strconv"
	"syscall"
)


type ExeRs struct {
	rs string
	err    error
}

// 执行shell命令，可设置执行超时时间
func Exec(ctx context.Context, command string) (string, error) {
	cmd := exec.Command("cmd", "/C", command)
	// 隐藏cmd窗口
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	var resultChan = make(chan ExeRs)
	go func() {
		output, err := cmd.CombinedOutput()
		resultChan <- ExeRs{string(output), err}
	}()
	select {
	case <-ctx.Done():
		if cmd.Process.Pid > 0 {
			_ = exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(cmd.Process.Pid)).Run()
			_ = cmd.Process.Kill()
		}
		return "", fmt.Errorf("cmd:%s,err:%s",command,"timeout")
	case result := <-resultChan:
		return convertEncoding(result.output), result.err
	}
}

func convertEncoding(outputGBK string) string {
	// 转换为utf8编码
	outputUTF8, ok := GBK2UTF8(outputGBK)
	if ok {
		return outputUTF8
	}
	return outputGBK
}
