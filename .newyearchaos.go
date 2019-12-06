// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int32{1, 2, 5, 3, 7, 8, 6, 4}
// 	x := callResult(arr)
// 	if x != 0 {
// 		fmt.Println(x)
// 	} else {
// 		fmt.Println("Too Chaotic")
// 	}
// }

// func callResult(arr []int32) int32 {
// 	var c int32
// 	org := arr
// 	for {
// 		flag := 0
// 		for j := 0; j < len(arr)-1; j++ {
// 			if org[j] > int32(j+3) {
// 				return 0
// 			}
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 				c++
// 				flag = 1
// 			}
// 		}
// 		if flag == 0 {
// 			return c
// 		}
// 	}
// }
