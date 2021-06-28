package main

import (
	"fmt"
	"log"
)

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	var x, y int

	fmt.Print("Enter integer #1 -> ")
	if _, err := fmt.Scanln(&x); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter integer #2 -> ")
	if _, err := fmt.Scanln(&y); err != nil {
		log.Fatal(err)
	}

	fmt.Println() // For terminal spacing

	swap(&x, &y)

	fmt.Println("Your swapped values are:")
	fmt.Printf("x -> %d\ny -> %d\n", x, y)
}
