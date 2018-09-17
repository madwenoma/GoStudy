package main

import (
	"goStudy/tree"
	"fmt"
)

func main() {
	var root tree.Node //0 nil nil 等价于root = Node{} 或者 new(Node)
	fmt.Println(root)
	root = tree.Node{Value: 3}
	fmt.Println(root)
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.NodeFactory(11)

	nodes := []tree.Node{
		{Value: 3},
		{4, nil, nil},
		{},
		{6, &root, nil},
	}

	fmt.Println(nodes)
	fmt.Println(tree.CreateNode(1024))

	root.SetValue(33)
	root.Println()

	root.Travers()
	//
	fmt.Println("test travers with func argument")
	nodeCount := 0
	//能做很多事
	root.TraversWithFunc(func(node *tree.Node) {
		fmt.Println(node.Value)//打印 这里的node相当于在遍历方法里回传过来的
		nodeCount++ //计算个数
	})
	fmt.Println(nodeCount)

}
