package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {

	for tn := 1; tn < 11; tn++ {
		var nameout, namein string
		if tn < 10 {
			namein = fmt.Sprint(`C:\Education\Golang\Route256\sandbox\d\tests\0`, tn)
			nameout = fmt.Sprint("0", tn, ".a")
		} else {
			namein = fmt.Sprint(`C:\Education\Golang\Route256\sandbox\d\tests\`, tn)
			nameout = fmt.Sprint(tn, ".a")
		}

		testin, _ := os.Open(namein)
		testout, _ := os.Create(nameout)

		in := bufio.NewReader(testin)
		out := bufio.NewWriter(testout)

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
				//fmt.Fprintln(out, "Сортировка", sortcol)
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
				/*
					for _, rowp := range sortt {
						for n, el := range rowp {
							if n == len(rowp)-1 {
								fmt.Fprintf(out, "%v\r\n", el)
							} else {
								fmt.Fprint(out, el, " ")
							}
						}
					}
				*/

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
		out.Flush()
		testin.Close()
		testout.Close()
	}

	for tn := 1; tn < 11; tn++ {
		var nameout, nametest string
		if tn < 10 {
			nametest = fmt.Sprint(`C:\Education\Golang\Route256\sandbox\d\tests\0`, tn, ".a")
			nameout = fmt.Sprint("0", tn, ".a")
		} else {
			nametest = fmt.Sprint(`C:\Education\Golang\Route256\sandbox\d\tests\`, tn, ".a")
			nameout = fmt.Sprint(tn, ".a")
		}
		f1, _ := os.Open(nametest)
		defer f1.Close()
		f2, _ := os.Open(nameout)
		defer f2.Close()
		fmt.Println(tn)
		compareCheckSum(getMD5SumString(f1), getMD5SumString(f2))
	}

}

func getMD5SumString(f *os.File) string {
	file1Sum := md5.New()
	io.Copy(file1Sum, f)
	return fmt.Sprintf("%X", file1Sum.Sum(nil))
}

func compareCheckSum(sum1, sum2 string) {
	match := "совпадают"
	if sum1 != sum2 {
		match = " не совпадают"
	}
	fmt.Printf("MD5: %s и MD5: %s %s\n", sum1, sum2, match)
}

/*
	for _, rowp := range t {
		for n, el := range rowp {
			if n == len(rowp)-1 {
				fmt.Printf("%v\r\n", el)
			} else {
				fmt.Print(el, " ")
			}
		}
	}
*/

/*
	for _, rowp := range t {
		for n, el := range rowp {
			if n == len(rowp)-1 {
				fmt.Printf("%v\r\n", el)
			} else {
				fmt.Print(el, " ")
			}
		}
	}
*/
/*
for i := 0; i < len(t); i++ {
						for l := len(t) - 1; l > i; l-- { //for l := len(t) - 1; l > i; l-- { //l := i; l < len(t); l++
							if t[i][sortcol] > t[l][sortcol] {
								t[i], t[l] = t[l], t[i]
							}
						}
					}
*/
