package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var c int32
	for {
		var maxi, mini, i1, i2 int32
		for i := int32(0); i < int32(len(arr)); i++ {
			if arr[i] < (i+1) && arr[i]-(i+1) < mini {
				mini = arr[i] - (i + 1)
				i1 = i
			}
			if arr[i] > (i+1) && arr[i]-(i+1) >= maxi {
				maxi = arr[i] - (i + 1)
				i2 = i
			}
		}
		fmt.Println(maxi, mini, i1, i2)
		if maxi == 0 && mini == 0 {
			return c
		}
		arr[i1], arr[i2] = arr[i2], arr[i1]
		c++
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Println(res)
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
