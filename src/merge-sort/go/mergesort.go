package mergesort

func MergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}

	l := MergeSort(s[:len(s)/2])
	r := MergeSort(s[len(s)/2:])

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
