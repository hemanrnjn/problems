package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(calcPermutation(int64(n)))
	}
}

func calcPermutation(n int64) int64 {
	m := int64(1)
	for i := int64(1); i <= n; i++ {
		m *= i
	}
	sum := m * (((n - 1) * n) / 2)
	return sum
}
