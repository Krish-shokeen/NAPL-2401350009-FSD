package mathutil

func Factorial(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func Power(base int, exponent int) int {
	result := 1

	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result
}
