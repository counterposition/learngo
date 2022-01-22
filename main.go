package main

import "fmt"

func main() {
	lines := []string{
		"Starting to make a coffee",
		"Grinding coffee beans",
		"Boiling water",
		"Mixing boiled water with crushed coffee beans",
		"Pouring coffee into the cup",
		"Pouring some milk into the cup",
		"Coffee is ready!",
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
