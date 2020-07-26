package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the surfaceArea function below.
func surfaceArea(A [][]int32) int32 {
	var prev int32 = 0
	var total int32 = 0
	var sideTotal int32 = 0
	var btmTotal int32 = 0
	side := make([]int32, len(A[0]))
	for i := range A {
		for j := range A[i] {
			if A[i][j] != 0 {
				btmTotal += 1
			}
			if j == len(A[i])-1 {
				total = total + int32(math.Abs(float64(A[i][j])-float64(prev))) + 1 + A[i][j]
			} else if j == 0 {
				total = total + A[i][j] + 1
				prev = A[i][j]
			} else {
				total = total + int32(math.Abs(float64(A[i][j])-float64(prev))) + 1
				prev = A[i][j]
			}

			if i == len(A)-1 {
				sideTotal = sideTotal + int32(math.Abs(float64(A[i][j])-float64(side[j]))) + A[i][j]
			} else {
				sideTotal = sideTotal + int32(math.Abs(float64(A[i][j])-float64(side[j])))
			}
		}
		side = A[i]
	}
	return total + sideTotal + btmTotal
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	HW := strings.Split(readLine(reader), " ")

	HTemp, err := strconv.ParseInt(HW[0], 10, 64)
	checkError(err)
	H := int32(HTemp)

	WTemp, err := strconv.ParseInt(HW[1], 10, 64)
	checkError(err)
	W := int32(WTemp)

	var A [][]int32
	for i := 0; i < int(H); i++ {
		ARowTemp := strings.Split(readLine(reader), " ")

		var ARow []int32
		for _, ARowItem := range ARowTemp {
			AItemTemp, err := strconv.ParseInt(ARowItem, 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			ARow = append(ARow, AItem)
		}

		if len(ARow) != int(int(W)) {
			panic("Bad input")
		}

		A = append(A, ARow)
	}

	result := surfaceArea(A)

	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
