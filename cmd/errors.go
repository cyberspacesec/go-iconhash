package cmd

import (
	"fmt"
	"os"
)

// HandleError 处理错误并退出程序
func HandleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
