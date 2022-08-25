package main

import "fmt"

var a = []int{3, 5, 8, 10, 14, 17, 21, 39}
var n int = 8

func main() {
	fmt.Println(binarySearch(10))
	fmt.Println(binarySearch(39))
	fmt.Println(binarySearch(100))

}

// 目的の値keyの添字を返す
func binarySearch(key int) int {
	left := 0
	right := len(a) - 1
	for right >= left {
		mid := left + (right - left)
		if a[mid] == key {
			return mid
		} else if a[mid] > key {
			right = mid - 1
		} else if a[mid] < key {
			left = mid + 1
		}
	}
	return -1
}
