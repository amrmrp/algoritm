package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gray(n int) []string {
	if n == 1 {
		return []string{"0", "1"}
	}

	prev := gray(n - 1)
	result := []string{}

	for _, s := range prev {
		result = append(result, "0"+s)
	}

	for i := len(prev) - 1; i >= 0; i-- {
		result = append(result, "1"+prev[i])
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(nStr)

	codes := gray(n)
	for _, code := range codes {
		fmt.Println(code)
	}
}
