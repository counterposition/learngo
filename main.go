package main

import (
	"fmt"
)

func main() {
	var cups uint

	fmt.Println("Write how many cups of coffee you will need:")
	for {
		fmt.Printf("> ")
		_, err := fmt.Scanf("%d", &cups)
		if err == nil {
			break
		}
	}

	var water = 200 * cups
	var milk = 50 * cups
	var beans = 15 * cups

	fmt.Printf("For %d cups of coffee you will need:\n", cups)
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", beans)
}
