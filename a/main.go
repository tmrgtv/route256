package main

import "fmt"

func main() {
	var (
		count, i uint16
		a, b     int16
	)
	fmt.Scan(&count)
	for i = 0; i < count; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
