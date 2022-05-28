package main

import (
	"fmt"
	"time"
)

func main() {
	//grpc_testing.StreamingInputCallResponse{}
	fmt.Printf("Hello, World! Now: ", string(time.Now().UnixNano()))
}
