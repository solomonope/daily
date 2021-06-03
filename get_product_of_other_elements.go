package main

import "fmt"

func solve_with_division(data []int) []int {
	totalMultiplication := 1
	for _, d := range data {
		totalMultiplication = totalMultiplication * d
	}

	for i, d := range data {
		data[i] = totalMultiplication / d
	}
	return data
}

func products(data []int) []int {
	prefix := make([]int, len(data))

	//  build prefix array
	for i, d := range data {
		if i == 0 {
			prefix[i] = d
		} else {
			prefix[i] = prefix[i-1] * d
		}
	}

	// build suffix array
	// //edb 5ef eff4f87  uyuyuyyu cfjjjjjjjjjgfgfgfg61  trttrtrtrtrtrytytytytytygtrtrtrtrtrtrtrgg9beahhhhhhhhh9fgfgftrtrtrtrtrtrtrtrgffhghghghgff1ggggggggga998   9002642   055c   0787d
	return data
}
func main() {
	fmt.Println("Hello world")
	shit := []int{1, 2, 3, 4, 5}
	roc := solve_with_division(shit)

	fmt.Println(roc)
	fmt.Println(shit)

}
