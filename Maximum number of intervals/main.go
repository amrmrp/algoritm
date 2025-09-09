package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Interval struct {
	l, r int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	intervals := make([]Interval, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &intervals[i].l, &intervals[i].r)
	}

	
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].r == intervals[j].r {
			return intervals[i].l < intervals[j].l
		}
		return intervals[i].r < intervals[j].r
	})

	count := 0
	lastEnd := -1 << 60 
	for _, it := range intervals {
		if it.l >= lastEnd {
			count++
			lastEnd = it.r
		}
	}

	fmt.Println(count)
}
