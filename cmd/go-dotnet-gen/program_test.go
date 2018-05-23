package main

import (
	"fmt"

	"github.com/matiasinsaurralde/go-dotnet/cmd/go-dotnet-gen/mypackage"
)

func main() {
	runtime, err := mypackage.NewRuntime(mypackage.RuntimeParams{})
	if err != nil {
		panic(err)
	}
	defer runtime.Shutdown()
	n := mypackage.HelloWorld(16)
	fmt.Println("HelloWorld = ", n)
	// mypackage.HelloWorld(12)
}
