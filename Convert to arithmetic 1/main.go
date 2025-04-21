package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	line1 , _  := reader.ReadString('\n');
	tokens := strings.Fields(line1)
	n, _ := strconv.Atoi(tokens[0])
	k, _ := strconv.Atoi(tokens[1])
	
	line2, _ := reader.ReadString('\n')
	tokens = strings.Fields(line2)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] , _ = strconv.Atoi(tokens[i])
	}

	minOps := math.MaxInt


	for i := 0; i < n; i++ {
		strart := a[i] - i * k
		ops    := 0

		for j := 0; j < n; j++ {

			target := strart + j * k
			diff   := a[j] - target

			if diff < 0 {
				 diff = - diff
			}
			
			ops += diff
		}

		if ops < minOps {
			minOps = ops
		}
	}


	fmt.Println(minOps)
}