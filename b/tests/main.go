package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	t, _ := os.Open(`C:\Education\Golang\Route256\sandbox\b\tests\10`)
	defer t.Close()
	rdr := bufio.NewReader(t)
	var count uint16
	var countpr uint32
	timer := time.Now()
	fmt.Fscanln(rdr, &count)
	//fmt.Scan(&count)
	for count > 0 {
		//fmt.Scan(&countpr)
		fmt.Fscanln(rdr, &countpr)
		mappr := make(map[uint32]uint8)
		var sum uint32
		for countpr > 0 {
			var price uint32
			//fmt.Scan(&price)
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
	fmt.Println(time.Until(timer))
}
