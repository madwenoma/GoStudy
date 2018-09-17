package tree

/**
 在别的文件里定义struct的方法
	需要在同一个包下
 */


func (node *Node) Travers() {
	if node == nil {
		return
	}
	node.Left.Travers() //如果是java，这里还要判断 if(node != null)
	//在go里nil只是个普通的函数，只要自己判断了处理了就行
	node.Println()
	node.Right.Travers()
}
