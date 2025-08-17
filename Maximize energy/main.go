package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Fruit struct {
    b, a int64
}

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    firstLine := strings.Fields(scanner.Text())
    n, _ := strconv.Atoi(firstLine[0])
    energy, _ := strconv.ParseInt(firstLine[1], 10, 64)

    fruits := make([]Fruit, n)
    for i := 0; i < n; i++ {
        scanner.Scan()
        line := strings.Fields(scanner.Text())
        b, _ := strconv.ParseInt(line[0], 10, 64)
        a, _ := strconv.ParseInt(line[1], 10, 64)
        fruits[i] = Fruit{b, a}
    }

    sort.Slice(fruits, func(i, j int) bool {
        return fruits[i].b < fruits[j].b
    })

    for _, f := range fruits {
        if energy >= f.b && f.b < f.a {
            energy = energy - f.b + f.a
        }
    }

    fmt.Println(energy)
}