package client

import (
	"fmt"
	"testing"
)

func TestInvoke(t *testing.T) {
	for i := 0; i < 10; i++ {
		resp, err := Invoke("127.0.0.1:8090", 222, 2, "test", "test...", 0)
		if err != nil {
			fmt.Println("invoke err", err.Error())
			return
		}
		fmt.Printf("resp:%+v\n", resp)
	}
}
