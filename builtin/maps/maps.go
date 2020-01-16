package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "cc",
		"tel":    "13400000000",
		"site":   "xxxx",
		"degree": "graduated",
	}

	fmt.Println(m)

	m2 := make(map[string]string) // m2 == empty map
	var m3 map[string]string      // m3 == nil
	fmt.Println(m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("\nGetting Values")

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("Key does not exists")
	}

	if site, ok := m["cite"]; ok {
		fmt.Println(site)
	} else {
		fmt.Println("Key does not exists")
	}

	fmt.Println("\nDeleting name")
	delete(m, "name")
	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("Key does not exists")
	}

}
