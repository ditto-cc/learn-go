package main

import "io/ioutil"

import "fmt"

import "strconv"

import "os"

import "bufio"

func readFile() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d", score))
	}
	return g
}

func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}

	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

func printFile() {
	if file, err := os.Open("abc.txt"); err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		panic(err)
	}
}

func main() {
	readFile()

	printFile()

	fmt.Println(
		convertToBin(3),
		convertToBin(13),
		convertToBin(32),
	)

	fmt.Println(
		grade(0),
		grade(59),
		grade(79),
		grade(99),
	)

}
