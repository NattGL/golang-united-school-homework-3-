package homework

func average(input [15]float32) (result float32) {
	var sum float32
	for _, number := range input {
		sum += number
	}
	sampleCount := float32(len(input))
	return sum / sampleCount
}
