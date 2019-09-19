package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the circularArrayRotation function below.
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	arrlen := int32(len(a))
	res := make([]int32, len(queries))
	newArr := make([]int32, arrlen)
	var rot int32 = k
	if rot >= arrlen {
		rot = rot % arrlen
	}
	for i := int32(0); i < arrlen; i++ {
		if i+rot > arrlen-1 {
			newArr[i+rot-arrlen] = a[i]
		} else {
			newArr[i+rot] = a[i]
		}
	}
	for i, val := range queries {
		res[i] = newArr[val]
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nkq := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nkq[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nkq[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	qTemp, err := strconv.ParseInt(nkq[2], 10, 64)
	checkError(err)
	q := int32(qTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	var queries []int32

	for i := 0; i < int(q); i++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := circularArrayRotation(a, k, queries)

	for i, resultItem := range result {
		fmt.Println(resultItem)

		if i != len(result)-1 {
			// fmt.Println()
		}
	}

	// fmt.Println()
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
