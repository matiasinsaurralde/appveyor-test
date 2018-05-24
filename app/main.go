package main

import (
	"github.com/matiasinsaurralde/appveyor-test/dotnet"
)

func main() {
	err := dotnet.Init()
	if err != nil {
		panic(err)
	}
}
