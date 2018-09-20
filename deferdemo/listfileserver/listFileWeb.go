package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
	"log"
	_ "net/http/pprof" //性能测试导入包，虽然没有使用，但命令行需要用到所以前面加下划线
)

/*
https://www.bilibili.com/video/av24365381/?p=32
defer 用recover fun(){}接收并处理
自定义接口error
自定义error处理 type assertion
*/

//定义一个函数，根据入参和返回值进行确定，函数名反而用处不大
//这个函数就是webHandler的一个实现
func handleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, mapping) != 0 {
		//return errors.New("url must start with " + mapping)
		return listUserError("url must start with " + mapping) //listUserError是个string可以直接初始化
		//panic(123)
	}

	path := request.URL.Path[len(mapping):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println(file)
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}

const mapping = "/list/"

//定义一个函数，入参和返回值
type webHandler func(http.ResponseWriter, *http.Request) error

//定义一个返回值是函数的函数，把函数作为基本类型，函数里有包含入参和返回值
func errorWrapper(handler webHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Print("panic error->:", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		code := http.StatusOK

		if err != nil {
			if err, ok := err.(userError); ok {
				http.Error(writer, err.Message(), http.StatusBadRequest)
				return //不要漏了
			}
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

//string类型也能实现接口，很不好理解
type listUserError string;

func (e listUserError) Error() string {
	return e.Message()
}

func (e listUserError) Message() string {
	return string(e)
}

func main() {
	/*http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		file, err := os.Open(path)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		defer file.Close()
		fmt.Println(file)
		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		writer.Write(all)
	})*/

	//HandleFunc接受一个函数入参 : handler func(ResponseWriter, *Request) 即上面定义的errorWrapper的返回值
	//errorWrapper的
	http.HandleFunc("/", errorWrapper(handleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
