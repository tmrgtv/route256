package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count, sorts uint8
	fmt.Fscanln(in, &count)
	fmt.Fscanln(in)
	for count > 0 {
		var r, c int
		fmt.Fscanln(in, &r, &c)
		var t [][]uint8
		for r > 0 {
			row := make([]uint8, c)
			for i := 0; i < c; i++ {
				var n uint8
				fmt.Fscan(in, &n)
				row[i] = n
			}
			t = append(t, row)
			r--
		}
		fmt.Fscan(in, &sorts)
		var rep int = -1
		for sorts > 0 {
			var sortcol int
			var sortt [][]uint8
			fmt.Fscan(in, &sortcol)
			sortcol--
			if rep != sortcol {
				mapsort := make(map[int]bool)
				for i := 0; i < len(t); i++ {
					var min uint8 = 101
					var poloj int
					for l := 0; l < len(t); l++ {
						if t[l][sortcol] < min && !mapsort[l] {
							min = t[l][sortcol]
							poloj = l
						}
					}
					mapsort[poloj] = true
					sortt = append(sortt, t[poloj])
				}
				rep = sortcol
				t = sortt
			}
			sorts--
		}
		for _, rowp := range t {
			for n, el := range rowp {
				if n == len(rowp)-1 {
					fmt.Fprintf(out, "%v\r\n", el)
				} else {
					fmt.Fprint(out, el, " ")
				}
			}
		}
		fmt.Fprint(out, "\r\n")
		fmt.Fscanln(in)
		fmt.Fscanln(in)
		count--
	}

}
