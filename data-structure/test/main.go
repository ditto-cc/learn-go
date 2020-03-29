package main

import (
	"bufio"
	"fmt"
	"learn-go/data-structure/compare"
	"learn-go/data-structure/heap"
	"learn-go/data-structure/tree/avl"
	"os"
	"strconv"
	"strings"
)

type Word string

func (w *Word) String() string {
	return string(*w)
}

func (w1 *Word) Compare(w compare.Comparable) int {
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

func (w1 *WordFre) Compare(t compare.Comparable) int {
	w2, ok := t.(*WordFre)
	if !ok {
		panic("Type Assertion Failed.")
	}

	res := w1.fre - w2.fre
	if res == 0 {
		if w1.word.String() == w2.word.String() {
			return 0
		} else if w1.word.String() > w2.word.String() {
			return -1
		} else {
			return 1
		}
	}
	return res
}

func (w *WordFre) String() string {
	return w.word.String() + " " + strconv.Itoa(w.fre)
}

func main() {
	tree := avl.CreateAVL()

	readFile("The-catcher-in-the-Rye.txt", tree)

	h := heap.CreateHeap()
	fmt.Println(tree.Size())
	tree.InOrder(func(w compare.Comparable, fre interface{}) {
		h.Push(&WordFre{word: w.(*Word), fre: fre.(int)})
		if h.Size() > 10 {
			h.Pop()
		}
	})

	for h.Size() > 0 {
		fmt.Println(h.Pop())
	}

}
