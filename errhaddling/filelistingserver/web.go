package main

import (
	"github.com/crawler/errhaddling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type userError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//panic
		defer func() {
			if r := recover(); r != nil {
			log.Printf("Panic: %v", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error handling request: %s", err.Error())

			//user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			//system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
