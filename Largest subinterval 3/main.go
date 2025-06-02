package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int64, n)
	line, _ := reader.ReadString('\n') 
	line, _ = reader.ReadString('\n')
	nums := strings.Fields(line)

	for i := 0; i < n; i++ {
		a[i], _ = strconv.ParseInt(nums[i], 10, 64)
	}

	maxsum := a[0]
	ans := a[0]

	for i := 1; i < n; i++ {
		maxsum = max(maxsum+a[i], a[i])
		ans = max(ans, maxsum)
	}

	fmt.Println(ans)
}
