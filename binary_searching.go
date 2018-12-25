package main

import "fmt"

func binarySearch(n int, ary []int) bool {
	low := 0
	high := len(ary) - 1
	for low <= high {
		mid := low + (high-low)/2
		if ary[mid] == n {
			return true
		} else if ary[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	a := []int{10, 20, 30, 40, 50, 60, 70, 80}
	fmt.Println(binarySearch(10, a))
	fmt.Println(binarySearch(40, a))
	fmt.Println(binarySearch(80, a))
	fmt.Println(binarySearch(0, a))
	fmt.Println(binarySearch(45, a))
	fmt.Println(binarySearch(90, a))
}
