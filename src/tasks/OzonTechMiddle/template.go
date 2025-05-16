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
		var c1, c2 int
		fmt.Fscan(in, &c1, &c2)
		fmt.Fprintln(out, "NO")
		
    }
}

func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}