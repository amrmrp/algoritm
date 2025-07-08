package main

import (
	"bufio"
	"fmt"
	"os"
)

var writer *bufio.Writer

func nTuple(n int, prefix string, k int) {
	if k == 0 {
		fmt.Fprintln(writer, prefix)
		return
	}
	for i := 1; i <= n; i++ {
		newPrefix := prefix + fmt.Sprintf("%d ", i)
		nTuple(n, newPrefix, k-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	nTuple(n, "", n)
}
