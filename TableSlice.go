package main

import "fmt"

func main() {
	tab := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		tab = append(tab, i*2)
	}
	for index, item := range tab {
		fmt.Println(item, " at index ", index)
	}

	mapa := make(map[string]string)
	mapa["hello"] = "world"
	fmt.Println(mapa["hello"])
	delete(mapa, "hello")
}
