package math

func Adder(xs ...int) int {
	res := 0
	for _, v := range xs {
		res += v
	}
	return res
}

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}