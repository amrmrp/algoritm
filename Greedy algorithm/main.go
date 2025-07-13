package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}

	readInt64 := func() int64 {
		scanner.Scan()
		val, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		return val
	}

	n := readInt()
	r := readInt64()

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = readInt()
	}

	sort.Ints(prices)

	var count int
	var sum int64 = 0
	for _, price := range prices {
		if sum+int64(price) > r {
			break
		}
		sum += int64(price)
		count++
	}

	fmt.Println(count)
}
