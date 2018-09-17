package mock

type Retriever struct {
	Content string
}
//实现了get方法就认为实现了retrieve接口 ducking type 具有接口的一些行为就算
func (r Retriever) Get(url string) string {
	return r.Content
}
