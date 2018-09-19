package queue

//An FIFO Queue
type Queue []interface{}

//func (q *Queue) Push(v interface{}) {
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

//pops ele
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	//return head
	return head.(int)
}

// return if empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
