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
	a.Children(b, c)
	a.PrintDepthFirst()
	d := node.New(4)
	b.Children(d)
	a.PrintDepthFirst()
}
