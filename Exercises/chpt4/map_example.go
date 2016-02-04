package main

import "fmt"

func main() {

	something := []string{"coins", "notes", "gold?"}

	thisMap := make(map[string]map[string]int)

	for _, v := range something {
		if _, ok := thisMap[v]; !ok {
			thisMap[v] = make(map[string]int)
		}
		thisMap[v]["random"] = 12
	}

	fmt.Println(thisMap)
}
