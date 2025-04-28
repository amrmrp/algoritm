package main


import("fmt")


func main(){

	var n int 
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	diff := []int{}
	for i := 0; i < n; i++ {
		if arr[i] != i + 1 {
			diff = append(diff, i)
		}
	}

	if len(diff) == 2 && arr[diff[0]] == diff[1] + 1 && arr[diff[1]] == diff[0] + 1 {
		fmt.Println("YES")
	} else if len(diff) == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}