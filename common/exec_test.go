package common

import (
	"context"
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	str,err := Exec(context.TODO(),"ls -lh")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str)
}

