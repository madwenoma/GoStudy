package main

import (
	"goStudy/tree"
	"fmt"
)

/**
https://www.bilibili.com/video/av24365381/?p=18
通过组合的方式对struct进行扩展
 */


type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postTranves() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postTranves()
	right := myTreeNode{myNode.node.Right}
	right.postTranves()
}

func main() {
	myNode := myTreeNode{&tree.Node{Value: 3}}
	fmt.Println(myNode.node.Value)
	myNode.postTranves()

}
