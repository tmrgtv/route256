package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var c, cp uint8
	var x, f int
	fmt.Fscanln(in, &c)
	for c > 0 {
		fmt.Fscanln(in, &cp)
		var str string
		str, _ = in.ReadString(byte('\n'))
		sl := strings.Fields(str)
		mapf := make(map[int]bool)

		for i := 0; i < len(sl); i++ {
			if mapf[i] {
				continue
			}
			p, _ := strconv.Atoi(sl[i])
			for f = 0; f < 100; f++ {
				for i1 := i + 1; i1 < len(sl); i1++ {
					if mapf[i1] {
						continue
					}
					p1, _ := strconv.Atoi(sl[i1])
					switch {
					case p-p1 >= 0:
						x = p - p1
					case p1-p >= 0:
						x = p1 - p
					}
					//fmt.Println("Разница", x, f)
					if x == f { //|| i1 == len(sl)-1
						mapf[i], mapf[i1] = true, true
						fmt.Fprintln(out, i+1, i1+1)
						break
					}
				}
				if x == f {
					break
				}
			}
		}
		fmt.Fprintln(out)
		c--
	}
}
