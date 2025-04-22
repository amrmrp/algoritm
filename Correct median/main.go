package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


func main()  {
	
reader := bufio.NewReader(os.Stdin)

line1, _ := reader.ReadString('\n')
n, _ := strconv.Atoi(strings.TrimSpace(line1))

line2, _ := reader.ReadString('\n')
strNums := strings.Fields(line2)

a := make([]int,n)

	for i := 0; i < n; i++ {
		a[i],_ = strconv.Atoi(strNums[i])
	}

	sort.Ints(a)

	var m int 
	if n % 2 == 1 {
		m = a[n / 2]
	}else{
		m = a[(n / 2) - 1]
	}

	totalOps := 0
	for i := 0; i < n; i++ {
		diff := a[i] - m
		if diff < 0 {
			diff = -diff
		}

		totalOps += diff
	}

	fmt.Println(m,totalOps)
}