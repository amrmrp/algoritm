package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])


	line, _ = reader.ReadString('\n')
	parts = strings.Fields(line)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(parts[i])
	}

	sort.Ints(a)

	for i := 0; i < q; i++ {
		line, _ = reader.ReadString('\n')
		x, _ := strconv.Atoi(strings.TrimSpace(line))
		
		count := sort.SearchInts(a, x)
		fmt.Println(count)
	}
}