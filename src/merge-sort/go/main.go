package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{10, 8, 9, 6, 7, 4, 5, 2, 3, 1}))
}

func mergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}
	return merge(mergeSort(s[:len(s)/2]), mergeSort(s[len(s)/2:]))
}

func merge(l []int, r []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			result = append(result, l[i])
			i++
		} else {
			result = append(result, r[j])
			j++
		}
	}

	for i < len(l) {
		result = append(result, l[i])
		i++
	}

	for j < len(r) {
		result = append(result, r[j])
		j++
	}

	return result
}
