package arrayslices

func Sum(arr []int) int {
	sum := 0
	for _, a := range arr {
		sum += a
	}
	return sum
}

func SumAll(numSlices ...[]int) []int {
	// len := len(numSlices)

	// sumSlice := make([]int, len)

	// for i, _ := range numSlices {
	// 	sumSlice[i] = Sum(numSlices[i])
	// }
	// return sumSlice

	var sumSlice []int

	for _, slice := range numSlices {
		sumSlice = append(sumSlice, Sum(slice))
	}
	return sumSlice
}

func tail(slice []int) []int {
	var tail []int
	for i, num := range slice {
		if i != 0 {
			tail = append(tail, num)
		}
	}
	return tail
}
func SumAllTails(slices ...[]int) []int {
	var sum []int

	for _, slice := range slices {
		if len(slice) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(tail(slice)))
		}
	}
	return sum
}
