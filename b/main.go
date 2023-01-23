package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)
	var count uint16
	var countpr uint32
	fmt.Fscan(rdr, &count)
	for count > 0 {
		fmt.Fscan(rdr, &countpr)
		mappr := make(map[uint32]uint8)
		var sum uint32
		for countpr > 0 {
			var price uint32
			fmt.Fscan(rdr, &price)
			mappr[price]++
			if mappr[price] < 3 {
				sum += price
			} else {
				mappr[price] = 0
			}
			countpr--
		}
		fmt.Println(sum)
		count--
	}
}
