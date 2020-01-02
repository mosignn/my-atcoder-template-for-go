package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type io struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func newIO() *io {
	return &io{
		scanner: bufio.NewScanner(os.Stdin),
		writer:  bufio.NewWriter(os.Stdout),
	}
}

func (io *io) nextLine() string {
	io.scanner.Scan()
	return io.scanner.Text()
}

func (io *io) nextInt() int {
	i, e := strconv.Atoi(io.nextLine())
	if e != nil {
		panic(e)
	}
	return i
}

func (io *io) scanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		num = io.nextInt()
		nums = append(nums, num)
	}
	return
}

func (io *io) printLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sumSlice(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func binarySearch(target int, list []int) int {
	left := 0
	right := len(list) - 1
	mid := (left + right) / 2

	for {
		if list[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) / 2
		if left >= right {
			break
		}
	}
	return mid
}

// IO will be orverwrited in test
var IO = newIO()

func main() {
	IO.scanner.Split(bufio.ScanWords) // switch to separating by space
	// scanner.Buffer([]byte{}, 1000000009) // switch to read large size input
	// defer io.Flush()
}
