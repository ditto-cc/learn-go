package main

import (
	"bufio"
	"fmt"
	"learn-go/data-structure/tree/avl"
	"learn-go/data-structure/tree/trie"
	"learn-go/data-structure/utils"
	"os"
	"reflect"
	"strings"
)

type Word string

func (w *Word) String() string {
	return string(*w)
}

func (w1 *Word) Compare(w utils.Comparable) int {
	w2 := w.(*Word)
	s1, s2 := string(*w1), string(*w2)
	if s1 == s2 {
		return 0
	} else if s1 < s2 {
		return -1
	}
	return 1
}

func readFile(filename string, tree *avl.Tree) {
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

func main() {
	avlTree := avl.NewTree()
	readFile("The-catcher-in-the-Rye.txt", avlTree)

	trieTree := trie.NewTrie()
	avlTree.InOrder(func(w utils.Comparable, fre interface{}) {
		trieTree.Add(w.(*Word).String())
	})
	fmt.Println(avlTree.Height())
	fmt.Println(avlTree.Size())

	w := Word("fuck")
	avlTree.Remove(&w)
	count := 0
	avlTree.InOrder(func(comparable utils.Comparable, i interface{}) {
		count++
	})
	fmt.Println(count)

	fmt.Println(avlTree.Height())
	fmt.Println(avlTree.Size())
	reflect.DeepEqual()
	//fmt.Println(trieTree.Size())
	//fmt.Println(trieTree.Search("funny"))
	//fmt.Println(trieTree.Match("f..ny"))
	//fmt.Println(trieTree.Match("f..k"))

	//q := deque.NewDeque(10)
	//fmt.Println("Starting Pushing...")
	//for i := 0; i < 100; i++ {
	//	if i%3 == 0 {
	//		q.PushBack(i)
	//	} else {
	//		q.PushFront(i)
	//	}
	//	fmt.Println(q)
	//}
	//fmt.Println("\nStarting Popping...")
	//for i := 0; i < 100; i++ {
	//	if i%4 == 0 {
	//		q.PopLeft()
	//	} else {
	//		q.PopRight()
	//	}
	//	fmt.Println(q)
	//}

}
