package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// خواندن n و k
	line1, _ := reader.ReadString('\n')
	parts := strings.Fields(line1)
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	// خواندن آرایه a
	line2, _ := reader.ReadString('\n')
	nums := strings.Fields(line2)

	x := make([]int64, n)

	for i := 0; i < n; i++ {
		a, _ := strconv.ParseInt(nums[i], 10, 64)
		x[i] = a - int64(i)*int64(k)
	}

	// پیدا کردن میانه
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	median := x[n/2]

	// محاسبه تعداد عملیات‌ها
	var totalOps int64 = 0
	for i := 0; i < n; i++ {
		totalOps += abs(x[i] - median)
	}

	fmt.Println(totalOps)
}
