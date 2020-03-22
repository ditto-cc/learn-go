package main

import (
	"learn-go/learning/errhandling/filelistserver/filelisting"
	"net/http"
	"os"

	"go.uber.org/zap"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

var logger, _ = zap.NewDevelopment()

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		sugar := logger.Sugar()
		defer sugar.Sync()
		defer func() {
			r := recover()
			if err, ok := r.(error); ok {
				sugar.Warn("Panic: %v", err)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			sugar.Warn(err, request.URL.Path)

			code := http.StatusOK
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

func main() {
	http.HandleFunc(filelisting.Prefix, errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
