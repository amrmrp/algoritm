package main


import("fmt")


func main(){

	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for  i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n - i - 1; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j+1] = a[j + 1], a[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d",a[i])
		if i != n - 1 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}