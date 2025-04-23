package main

import "fmt"

// TC - O(n) && SC - O(n)
func destinationCity(paths [][]string) string {
	outgoing := make(map[string]bool)

	for _, path := range paths {
		outgoing[path[0]] = true
	}

	for _, path := range paths {
		if !outgoing[path[1]] {
			return path[1]
		}
	}

	return ""
}

func main() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(destinationCity(paths)) // Output: Sao Paulo
}
