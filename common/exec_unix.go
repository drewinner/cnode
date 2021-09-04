// +build !windows

package common

import (
	"context"
	"fmt"
	"os/exec"
	"syscall"
)

type ExeRs struct {
	rs string
	err    error
}

// Exec 执行shell指令
func Exec(ctx context.Context, command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	resultChan := make(chan ExeRs)
	go func() {
		output, err := cmd.CombinedOutput()
		resultChan <- ExeRs{string(output), err}
	}()
	select {
	case <-ctx.Done():
		if cmd.Process.Pid > 0 {
			_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		return "", fmt.Errorf("cmd:%s,err:%s",command,"timeout")
	case result := <-resultChan:
		return result.rs, result.err
	}
}
