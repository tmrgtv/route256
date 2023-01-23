package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//timer := time.Now()
	//testin, _ := os.Open(`C:\Education\Golang\Route256\sandbox\e\tests\34`)
	//defer testin.Close()
	//testout, _ := os.Create("34.a")
	//defer testout.Close()
	in := bufio.NewReader(os.Stdin)   //os.Stdin
	out := bufio.NewWriter(os.Stdout) //os.Stdout
	defer out.Flush()
	var t uint8
	fmt.Fscan(in, &t)
	for t > 0 {
		var n uint16
		fmt.Fscan(in, &n)
		fmt.Fscanln(in)
		var str string
		str, _ = in.ReadString(byte('\r'))
		sl := strings.Fields(str)
		//fmt.Println(sl)
		mapf := make(map[string]int)
		for n, el := range sl {
			if num, ok := mapf[el]; !ok {
				mapf[el] = n
			} else {
				if n-num > 1 {
					fmt.Fprint(out, "NO\r\n")
					break
				} else {
					mapf[el] = n
				}
			}
			if n+1 == len(sl) {
				fmt.Fprint(out, "YES\r\n")
			}

		}
		t--
	}
	//fmt.Println(time.Until(timer))
}
