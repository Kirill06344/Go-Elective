package calculations

func Calculate(n int64, flag bool) int64 {
	res := int64(1)
	for i := int64(2); i <= n; i++ {
		res *= i
	}
	return res
}
