package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func errPanic(http.ResponseWriter, *http.Request) error {
	panic(123)
}

func TestErrWrapper(t *testing.T) {
	var tests = []struct{
		h appHandler
		code int
		message string
	} {
		{errPanic, 500, "Internal Server Error"},
	}

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("expect (%d %s); got(%d %s)",
				tt.code, tt.message, response.Code, body)
		}
	}
}
