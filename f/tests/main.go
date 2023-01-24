package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.Now()

	t1, _ := time.Parse("15:04:05", "23:59:59")
	t2, _ := time.Parse("15:04:05", "23:59:59")

	fmt.Println(t1.After(t2))
	fmt.Println(t1.Before(t2))
	fmt.Println(t1.After(t2))
	fmt.Println(t1.Before(t2))
	fmt.Println(t1.Equal(t2))
	fmt.Println(t1.Equal(t2))
	fmt.Println(t1.Equal(t2))

	fmt.Println(t1.Equal(t2))

	fmt.Println(time.Until(timer))
}
