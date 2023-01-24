package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Timestruct struct {
	N      int
	Timesl [][2]int
}

func main() {

	for tn := 1; tn < 35; tn++ {

		var nameout, namein string
		if tn < 10 {
			namein = fmt.Sprint(`D:\Education\Route256\sandbox\f\tests\0`, tn)
			nameout = fmt.Sprint("0", tn, ".a")
		} else {
			namein = fmt.Sprint(`D:\Education\Route256\sandbox\f\tests\`, tn)
			nameout = fmt.Sprint(tn, ".a")
		}

		timer := time.Now()

		testin, _ := os.Open(namein)
		defer testin.Close()
		testout, _ := os.Create(nameout)
		defer testout.Close()

		in := bufio.NewReader(testin)   //os.Stdin
		out := bufio.NewWriter(testout) //os.Stdout
		defer out.Flush()
		var t int
		fmt.Fscan(in, &t)
		fmt.Fscanln(in)
		timech := make(chan Timestruct, t)
		for c := 0; c < t; c++ {
			var n uint16
			fmt.Fscan(in, &n)
			fmt.Fscanln(in)

			bbuf := make([]byte, 19*int(n))

			io.ReadFull(in, bbuf)

			go func(b []byte, c int, tch chan Timestruct) {
				slstr := strings.Split(string(b), "\r\n")
				var timesl [][2]int
				var timestr Timestruct
				for _, el := range slstr {
					if len(el) > 1 {
						tspl := strings.Split(el, "-")

						t1, err1 := time.Parse("15:04:05", tspl[0])
						t2, err2 := time.Parse("15:04:05", tspl[1])

						if err1 != nil || err2 != nil || !(t1.Before(t2) || t1.Equal(t2)) {
							timestr = Timestruct{N: c}
							tch <- timestr
							return
						} else {
							n1, _ := strconv.Atoi(strings.ReplaceAll(tspl[0], ":", ""))
							n2, _ := strconv.Atoi(strings.ReplaceAll(tspl[1], ":", ""))
							timesl = append(timesl, [2]int{n1, n2})
						}

					}

				}
				timestr = Timestruct{N: c, Timesl: timesl}
				tch <- timestr
				return
			}(bbuf, c, timech)
		}
		outsl := make([]string, t)
		for count := 1; count <= t; count++ {
			tstr := <-timech
			if len(tstr.Timesl) == 0 {
				outsl[tstr.N] = "NO"
			} else {
				errchan := make(chan bool, len(tstr.Timesl))
				for id, el := range tstr.Timesl {

					go func(el [2]int, id int) {
						for n := id + 1; n < len(tstr.Timesl); n++ {
							elch := tstr.Timesl[n]
							if el[0] == elch[0] || el[0] == elch[1] || el[1] == elch[0] || el[1] == elch[1] {
								errchan <- true
								return
							} else if (el[0] > elch[0] && el[0] < elch[1]) || (el[1] > elch[0] && el[1] < elch[1]) {
								errchan <- true
								return
							}
							/*else if (elch[0].Before(el[0]) && elch[1].After(el[0])) ||
								(elch[0].Before(el[1]) && elch[1].After(el[1])) {
								errchan <- true
								return
							}*/

						}
						errchan <- false
						return
					}(el, id)

				} // end for range timesl

				for a := 0; a < len(tstr.Timesl); a++ {
					err := <-errchan
					if err {
						outsl[tstr.N] = "NO"
					}
				}
				close(errchan)
				if outsl[tstr.N] != "NO" {
					outsl[tstr.N] = "YES"
				}
			}
		}
		close(timech)
		fmt.Println(outsl)

		/*
			if check && n > 1 {
				errchan := make(chan bool, len(timesl))
				var wg sync.WaitGroup

				for id, el := range timesl {
					wg.Add(1)
					go func(el [2]time.Time, id int) {
						for n, elch := range timesl {
							if n != id {
								if el[0].Equal(elch[0]) || el[0].Equal(elch[1]) || el[1].Equal(elch[0]) || el[1].Equal(elch[1]) ||
									(elch[0].Before(el[0]) && elch[1].After(el[0])) || (elch[0].Before(el[1]) && elch[1].After(el[1])) {
									wg.Done()
									errchan <- true
									return
								}
							}
						}
						wg.Done()
						errchan <- false
						return
					}(el, id)

				} // end for range timesl

				wg.Wait()
				close(errchan)
				for a := range errchan {
					if a {
						errf = true
						break
					}
				}
			}

			if !errf && check {
				fmt.Fprint(out, "YES\r\n")
				//fmt.Println("YES")
			} else {
				fmt.Fprint(out, "NO\r\n")
				//fmt.Println("NO")
			}*/

		fmt.Println(time.Until(timer))
		out.Flush()
		testin.Close()
		testout.Close()
	}
	/*
		for tn := 1; tn < 11; tn++ {
			var nameout, nametest string
			if tn < 10 {
				nametest = fmt.Sprint(`D:\Education\Route256\sandbox\f\tests\0`, tn, ".a")
				nameout = fmt.Sprint("0", tn, ".a")
			} else {
				nametest = fmt.Sprint(`D:\Education\Route256\sandbox\f\tests\`, tn, ".a")
				nameout = fmt.Sprint(tn, ".a")
			}
			f1, _ := os.Open(nametest)
			defer f1.Close()
			f2, _ := os.Open(nameout)
			defer f2.Close()
			fmt.Println(tn)
			compareCheckSum(getMD5SumString(f1), getMD5SumString(f2))
		} */

}

func compareCheckSum(sum1, sum2 string) {
	match := "совпадают"
	if sum1 != sum2 {
		match = " не совпадают"
	}
	fmt.Printf("MD5: %s и MD5: %s %s\n", sum1, sum2, match)
}
func getMD5SumString(f *os.File) string {
	file1Sum := md5.New()
	io.Copy(file1Sum, f)
	return fmt.Sprintf("%X", file1Sum.Sum(nil))
}
