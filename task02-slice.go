package homework

func reverse(input []int64) (result []int64) {
	inputLen := len(input)
	result = make([]int64, inputLen)
	for idx, value := range input {
		index := inputLen - idx - 1
		result[index] = value
	}
	return result
}
