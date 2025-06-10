package main

import (
    "fmt"
    "strconv"
)

func main() {
    var X string
    fmt.Scan(&X) 
    sumDigits := 0
    for _, digit := range X {
        num, _ := strconv.Atoi(string(digit)) 
        sumDigits += num 
    }

    fmt.Println(sumDigits)
}
