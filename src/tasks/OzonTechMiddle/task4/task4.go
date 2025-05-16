package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
    X = 2
    O = 1
    No = 1
)

func main() {
    // For normal execution, use stdin and stdout
    runSolution(os.Stdin, os.Stdout)
}

// runSolution is the main logic of the program, made testable by accepting
// any io.Reader as input and io.Writer as output
func runSolution(input io.Reader, output io.Writer) {
    in := bufio.NewReader(input)
    out := bufio.NewWriter(output)
    defer out.Flush()

    var testCases int
    fmt.Fscan(in, &testCases)
    for i := 0; i < testCases; i++ {
        var kLentoWin int
        var nRows, mCols int
        fmt.Fscan(in, &kLentoWin)
        fmt.Fscan(in, &nRows, &mCols)

        table := make([][]int, nRows)

        for n := 0; n < nRows; n++ {
            table[n] = make([]int, mCols)
            var row string
            fmt.Fscan(in, &row)

            for m := 0; m < mCols && m < len(row); m++ {
                switch row[m] {
                case 'X':
                    table[n][m] = X
                case 'O':
                    table[n][m] = O
                default:
                    table[n][m] = No
                }
            }
        }

		
        
        fmt.Fprintln(out, "NO")
    }
}

// getBuffers is no longer needed as we now pass io.Reader and io.Writer directly
func getBuffers()(*bufio.Reader,*bufio.Writer){
    var in *bufio.Reader
    var out *bufio.Writer

    in = bufio.NewReader(os.Stdin)
    out = bufio.NewWriter(os.Stdout)

    return in,out
}