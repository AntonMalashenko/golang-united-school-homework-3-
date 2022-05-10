package homework

func average(input [15]float32) (result float32) {
	var sum float32
	var length float32

	for _, value := range input {
		if value != 0 {
			length++
			sum = sum + value
		}
	}
	return sum / length
}
