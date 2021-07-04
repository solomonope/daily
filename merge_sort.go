package main

import "fmt"

func sort(data []int) {
	mergSort(data, 0, len(data)-1)
}

func mergSort(data []int, left, right int) {
	if left > right {
		mid := int(left + right/2)
		mergSort(data, left, mid)
		mergSort(data, mid+1, right)
		merge(data, left, mid, right)
	}
}

func merge(data []int, left, mid, right int) {
	counter := 0
	temp := make([]int, left+right+1)
	l, r := 0, 0

	for l <= mid || r <= right {
	
	}

	fmt.Println(temp)
}

func main() {
	fmt.Println("Merge Sort")
}
