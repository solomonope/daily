package main

import "fmt"

func sort(data []int) {
	for j := 1; j < len(data); j++ {
		key := data[j]
		i := j - 1

		for i >= 0 && data[i] > key {
			data[i+1] = data[i]
			i = i - 1
		}
		data[i+1] = key
	}

	fmt.Println(data)
}
func main() {
	sort([]int{5, 2, 4, 6, 1, 3})
}
