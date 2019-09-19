// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // Complete the saveThePrisoner function below.
// func saveThePrisoner(n int32, m int32, s int32) int32 {
// 	if n == 1 {
// 		return n
// 	}
// 	x := m%n + (s - 1)
// 	if x == 0 {
// 		return n
// 	}
// 	if x > n {
// 		return x % n
// 	}
// 	return x
// }

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

// 	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
// 	checkError(err)
// 	t := int32(tTemp)

// 	for tItr := 0; tItr < int(t); tItr++ {
// 		nms := strings.Split(readLine(reader), " ")

// 		nTemp, err := strconv.ParseInt(nms[0], 10, 64)
// 		checkError(err)
// 		n := int32(nTemp)

// 		mTemp, err := strconv.ParseInt(nms[1], 10, 64)
// 		checkError(err)
// 		m := int32(mTemp)

// 		sTemp, err := strconv.ParseInt(nms[2], 10, 64)
// 		checkError(err)
// 		s := int32(sTemp)

// 		result := saveThePrisoner(n, m, s)

// 		fmt.Println(result)
// 	}
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
