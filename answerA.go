package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextLine() string {
	scanner.Scan()
	return scanner.Text()
}

func nextInt() int {
	i, e := strconv.Atoi(nextLine())
	if e != nil {
		panic(e)
	}
	return i
}

// for large string > 64*1024
var reader = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
    buf := make([]byte, 0, 1000000)
    for {
        l, p, e := reader.ReadLine()
        if e != nil {
            panic(e)
        }
        buf = append(buf, l...)
        if !p {
            break
        }
    }
    return string(buf)
}


func main() {
	scanner.Split(bufio.ScanWords)  // switch to separating by space
	n := nextInt()
	fmt.Println(n)
}
