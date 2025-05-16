package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    in, out := getBuffers()
    defer out.Flush()

    var testCases int
    fmt.Fscan(in, &testCases)

    // Result complexity is Sum(str)
    for i := 0; i < testCases; i++ {
        var wordsInCase int
        result := 0
        fmt.Fscan(in, &wordsInCase)
        
        evenMap := make(map[string]int)
        oddMap := make(map[string]int)
        equalMap := make(map[string]int)

        // Time complexity = O(s)
        getEvenAndOdd := func(str string) ([]byte, []byte) {
            evenSubStr := make([]byte, 0, len(str) / 2)
            oddSubStr := make([]byte, 0, len(str) / 2)

            for i := 0; i < len(str); i++ {
                if  i % 2 == 0 {
                    evenSubStr = append(evenSubStr, str[i])
                } else {
                    oddSubStr = append(oddSubStr, str[i])
                }
            }
            return evenSubStr, oddSubStr
        }
        
        for j := 0; j < wordsInCase; j++ {
            var word string = ""
            fmt.Fscan(in, &word)

            if len(word) == 1 {
                result += oddMap[word]
                oddMap[word]++
                continue
            }

            even, odd := getEvenAndOdd(word)

            result += evenMap[string(even)]
            result += oddMap[string(odd)]
            result -= equalMap[word]

            equalMap[word]++
            evenMap[string(even)]++
            oddMap[string(odd)]++
        }
        
        fmt.Fprintln(out, result)
    }
}

func getBuffers() (*bufio.Reader, *bufio.Writer) {
    in := bufio.NewReaderSize(os.Stdin, 1024*1024*16) 
    out := bufio.NewWriterSize(os.Stdout, 1024*1024*16)
    return in, out
}