package tree

import "fmt"
/*
https://www.bilibili.com/video/av24365381/?p=16
结构体的一些用法

为结构生成的方法必须放在同一个包内，但可以是不同的文件
 */

//结构创建在堆上还是栈上？不需要知道都有可能
type Node struct {
	Value       int
	Left, Right *Node
}
//为结构体定义方法，类似java 同一个class里的方法
// 值接受者，是copy一份node给这个函数，就可以使用了，但这里改变node的值是不会在外面生效的
//需要使用指针接受者
func (node Node) Println() {
	fmt.Println(node.Value)
}

//指针接受者，不会copy 值过来，而是把引用传过来，可以改变原变量的值
//在传入的时候，不需要用&取值，因为编译器会自动判断 这点跟方法参数的*不同
func (node *Node) SetValue(value int) {
	node.Value = value
}

//

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

/*func main() {
	var root Node //0 nil nil 等价于root = Node{} 或者 new(Node)
	fmt.Println(root)
	root = Node{Value: 3}
	fmt.Println(root)
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = NodeFactory(11)

	nodes := []Node{
		{Value: 3},
		{4, nil, nil},
		{},
		{6, &root, nil},
	}

	fmt.Println(nodes)
	fmt.Println(CreateNode(1024))

	root.SetValue(33)
	root.Println()

	root.Travers()
}*/

//go没有构造函数一说，如果想要可以通过工厂函数实现

func NodeFactory(value int) *Node {
	return &Node{Value: value} //&返回局部变量地址
}
