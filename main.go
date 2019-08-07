package main

import (
	"DAGTest/dag"
	"fmt"
)

func main() {
	fmt.Println("pog")
	a := node.New(1)
	b := node.New(2)
	c := node.New(3)
	a.AddChildren(b, c)
	a.PrintDepthFirst()
	d := node.New(4)
	b.AddChildren(d)
	a.PrintDepthFirst()
}
