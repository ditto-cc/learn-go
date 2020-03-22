package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {contents=%s}", r.Contents)
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, m map[string]string) string {
	r.Contents = m["contents"]
	return r.Contents
}
