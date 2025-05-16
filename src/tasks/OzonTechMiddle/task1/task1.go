package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    in, out := getBuffers()
    defer out.Flush()

    var count int
    fmt.Fscan(in, &count)
    for i:= 0; i<count; i++ {
		var s string
        fmt.Fscan(in, &s)
		if len(s) == 1 {
			fmt.Fprintln(out, "YES")
			continue
		}
		if len(s) == 2 && s[0] == s[1] {
			fmt.Fprintln(out, "YES")
			continue
		}
		if s[0] != s[len(s) - 1] {
			fmt.Fprintln(out, "NO")
			continue
		}

		var originalSymbol = s[0]
		var res string
		for j:= 1; j<len(s) - 1; j++ {
			if s[j] != originalSymbol {
				if s[j - 1] != s[j+1] {
					res = "NO"
					break
				}
			}
		}
		if res == "NO" {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}