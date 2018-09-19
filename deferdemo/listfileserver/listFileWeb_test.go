package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
)

// test panic 
func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

//test user error
func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func TestErrWrapper(t *testing.T) {
	testDatas := []struct {
		handler webHandler
		code    int
		message string
	}{
		{errPanic, 500, "Internal Server Error"},
		{errUserError, 400, "user error"},
	}

	for _, tt := range testDatas {
		h := errorWrapper(tt.handler)
		response := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "http://www.immoc.com", nil)
		h(response, req)

		resultByte, _ := ioutil.ReadAll(response.Body)
		result := strings.Trim(string(resultByte), "\n")
		if response.Code != tt.code || result != tt.message {
			t.Errorf("expect (%d,%s),but got (%d,%s)", tt.code, tt.message, response.Code, result)
		}

	}
}

//吧user error复制过来了，最好是提出来，此处为了简单，copy过来了
type testingUserError string;

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}
