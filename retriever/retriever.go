package main

import (
	"fmt"
	"learn-go/retriever/mock"
	"learn-go/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = &mock.Retriever{"this is a mock retriever."}
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5.0", Timeout: time.Minute}
	inspect(r)
	// Type Assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.Timeout)
	} else {
		fmt.Println("not *real retriever.")
	}
	//fmt.Println(download(r))
}
