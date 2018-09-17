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



//上面的遍历有个缺点，遍历的时候只能打印调用Println函数
///现在想让遍历的时候做任何事
//养惯将函数看做一个普通变量的习惯
func (node *Node) TraversWithFunc(doSth func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraversWithFunc(doSth) //如果是java，这里还要判断 if(node != null)
	//在go里nil只是个普通的函数，只要自己判断了处理了就行
	doSth(node)
	node.Right.TraversWithFunc(doSth)
}
