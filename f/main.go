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
	for tn := 1; tn < 36; tn++ {

		var nameout, namein string
		if tn < 10 {
			namein = fmt.Sprint(`D:\Education\Route256\sandbox\f\tests\0`, tn)
			nameout = fmt.Sprint("0", tn, ".a")
		} else {
			namein = fmt.Sprint(`D:\Education\Route256\sandbox\f\tests\`, tn)
			nameout = fmt.Sprint(tn, ".a")
		}

		testin, _ := os.Open(namein)
		defer testin.Close()
		testout, _ := os.Create(nameout)
		defer testout.Close()
		timer := time.Now()
		in := bufio.NewReader(testin)   //os.Stdin
		out := bufio.NewWriter(testout) //os.Stdout
		//defer out.Flush()
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
						var n1, n2, h1, m1, s1, h2, m2, s2 int
						tspl := strings.Split(el, "-")

						comparech := make(chan bool, 3)
						go func(c chan bool) {
							tspl1 := strings.Split(tspl[0], ":")
							h1, _ = strconv.Atoi(tspl1[0])
							m1, _ = strconv.Atoi(tspl1[1])
							s1, _ = strconv.Atoi(tspl1[2])
							if h1 >= 24 || m1 >= 60 || s1 >= 60 {
								comparech <- true
								return
							}
							comparech <- false
							return
						}(comparech)

						go func(c chan bool) {

							tspl2 := strings.Split(tspl[1], ":")
							h2, _ = strconv.Atoi(tspl2[0])
							m2, _ = strconv.Atoi(tspl2[1])
							s2, _ = strconv.Atoi(tspl2[2])
							if h2 >= 24 || m2 >= 60 || s2 >= 60 {
								comparech <- true
								return
							}
							comparech <- false
							return
						}(comparech)

						go func(c chan bool) {
							n1, _ = strconv.Atoi(strings.ReplaceAll(tspl[0], ":", ""))
							n2, _ = strconv.Atoi(strings.ReplaceAll(tspl[1], ":", ""))
							if n1 > n2 {
								comparech <- true
								return
							}
							comparech <- false
							return
						}(comparech)
						var ct bool
						for i := 0; i < 3; i++ {
							ct = <-comparech
							if ct {
								timestr = Timestruct{N: c}
								tch <- timestr
								return
							}
						}

						timesl = append(timesl, [2]int{n1, n2})
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

					go func(el [2]int, id int, errc chan bool) {
						for n := id + 1; n < len(tstr.Timesl); n++ {
							elch := tstr.Timesl[n]
							if el[0] == elch[0] || el[0] == elch[1] || el[1] == elch[0] || el[1] == elch[1] ||
								((el[0] > elch[0] && el[0] < elch[1]) || (el[1] > elch[0] && el[1] < elch[1])) ||
								((el[0] < elch[0] && el[1] > elch[0]) || (el[0] < elch[1] && el[1] > elch[1])) {
								errc <- true
								return

							}
						}
						/*
							errchan1 := make(chan bool, len(tstr.Timesl))
							var wgg sync.WaitGroup

							for n := id + 1; n < len(tstr.Timesl); n++ {
								wgg.Add(1)
								go func(elch, el [2]int, errch chan bool) {
									defer wgg.Done()
									//fmt.Println(id, elch, el)
									if el[0] == elch[0] || el[0] == elch[1] || el[1] == elch[0] || el[1] == elch[1] ||
										((el[0] > elch[0] && el[0] < elch[1]) || (el[1] > elch[0] && el[1] < elch[1])) ||
										((el[0] < elch[0] && el[1] > elch[0]) || (el[0] < elch[1] && el[1] > elch[1])) {
										errch <- true
										return
									}
									errch <- false
									return
								}(tstr.Timesl[n], el1, errchan1)




							}
							wgg.Wait()
							close(errchan1)
							for e := range errchan1 {
								if e {
									errc <- true
									return
								}
							}*/
						/*
							for x := 0; x < len(tstr.Timesl)-id; x++ {
								err := <-errchan1
								fmt.Println(x, err)
								if err {
									errc <- true
									return
								}
							}*/

						errc <- false
						return
					}(el, id, errchan)

				} // end for range timesl

				for a := 0; a < len(tstr.Timesl); a++ {
					err := <-errchan
					if err {
						outsl[tstr.N] = "NO"
						break
					}
				}
				//close(errchan)
				if outsl[tstr.N] != "NO" {
					outsl[tstr.N] = "YES"
				}
			}
		}
		close(timech)
		for _, o := range outsl {
			fmt.Fprint(out, o, "\r\n")
		}
		out.Flush()
		fmt.Println(time.Until(timer))
	}

	for tn := 1; tn < 36; tn++ {
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
	}

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
