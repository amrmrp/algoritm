package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MOD int64 = 1e9 + 7

func main() {
	reader := bufio.NewReader(os.Stdin)

	line1, _ := reader.ReadString('\n')
	parts := strings.Fields(line1)
	n, _ := strconv.Atoi(parts[0])
	x, _ := strconv.ParseInt(parts[1], 10, 64)

	line2, _ := reader.ReadString('\n')
	coeffStr := strings.Fields(line2)

	coeffs := make([]int64, n+1)
	for i := 0; i <= n; i++ {
		coeffs[i], _ = strconv.ParseInt(coeffStr[i], 10, 64)
	}

	res := coeffs[0]
	for i := 1; i <= n; i++ {
		res = (res*x + coeffs[i]) % MOD
		if res < 0 {
			res += MOD
		}
	}

	fmt.Println(res)
}
