package main

import (
	"fmt"
	"learn-go/retriever/mock"
	"learn-go/retriever/real"
	"time"
)

const url string = "www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type Session interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func post(p Poster) {
	p.Post("www.baidu.com", map[string]string{
		"contents": "12345",
	})
}

func session(s Session) string {
	fmt.Println("Trying Session.")
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}


func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n> Type switch: ", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	var r Retriever
	retriever := &mock.Retriever{"this is a mock retriever."}
	r = retriever
	inspect(r)
	fmt.Println()
	fmt.Println(download(retriever))
	fmt.Println(session(retriever))
	fmt.Println()
	r = &real.Retriever{UserAgent: "Mozilla/5.0", Timeout: time.Minute}
	inspect(r)
	fmt.Println()
	// Type Assertion
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.Timeout)
	} else {
		fmt.Println("not *real retriever.")
	}
}
