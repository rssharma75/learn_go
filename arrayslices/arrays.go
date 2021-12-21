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
