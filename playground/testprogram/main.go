package main

import (
	"fmt"

	p "github.com/matiasinsaurralde/go-dotnet/playground/pkg"
)

func main() {
	// dotnet.Init()
	/*
		d := p.Hello(2)
		fmt.Println(2, d)

		d2 := p.Hello(4)
		fmt.Println(2, d2)
	*/

	p.Hello(666)
	output := p.GetString()
	fmt.Println("output is", output)
	// fmt.Println("calling Hello with 3")
	// output := p.Hello(3)
	// fmt.Println("receiving ", output)
}
