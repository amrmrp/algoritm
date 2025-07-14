package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	deadlines := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		deadlines[i], _ = strconv.Atoi(scanner.Text())
	}

	sort.Ints(deadlines)

	time := 0
	done := 0
	for _, d := range deadlines {
		if time+1 <= d {
			time++
			done++
		}
	}

	fmt.Println(done)
}
