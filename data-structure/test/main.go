package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"learn-go/data-structure/avl"
	_interface "learn-go/data-structure/compare"
	"os"
	"strconv"
	"strings"
)

type Word string

func (w *Word) String() string {
	return string(*w)
}

func (w1 *Word) Compare(w _interface.Comparable) int {
	w2 := w.(*Word)
	s1, s2 := string(*w1), string(*w2)
	if s1 == s2 {
		return 0
	} else if s1 < s2 {
		return -1
	}
	return 1
}

func readFile(filename string, tree *avl.AVLTree) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			s := strings.ToLower(strings.Trim(word, "\",.'!?%$#/-1234567890[](){}+=_~` "))
			if s == "" {
				continue
			}
			w := Word(s)
			if fre := tree.Get(&w); fre != nil {
				tree.Add(&w, fre.(int)+1)
			} else {
				tree.Add(&w, 1)
			}
		}
	}
}

type WordFre struct {
	word *Word
	fre  int
}

func (w *WordFre) String() string {
	return w.word.String() + " " + strconv.Itoa(w.fre)
}

type WordHeap []*WordFre

func (h *WordHeap) Len() int {
	return len(*h)
}

func (h *WordHeap) Less(i, j int) bool {
	return (*h)[i].fre < (*h)[j].fre
}

func (h *WordHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *WordHeap) Push(x interface{}) {
	*h = append(*h, x.(*WordFre))
}

func (h *WordHeap) Pop() interface{} {
	ret := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return ret
}

func main() {
	tree := avl.CreateAVL()

	readFile("The-catcher-in-the-Rye.txt", tree)

	h := WordHeap{}
	fmt.Println(tree.Size())
	tree.InOrder(func(w _interface.Comparable, fre interface{}) {
		heap.Push(&h, &WordFre{word: w.(*Word), fre: fre.(int)})
		if len(h) > 20 {
			heap.Pop(&h)
		}
	})

	for i := 0; i < 20; i++ {
		fmt.Println(heap.Remove(&h, 0))
	}
}
