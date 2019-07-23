package main

import (
	"fmt"
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

func main() {
	t := util.GetExampleTree()
	t.PrintInOrder()
	fmt.Println()
	t.PrintPreOrder()
	fmt.Println()
	t.PrintPostOrder()
}
