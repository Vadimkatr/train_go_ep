package factorial

// Factorial is function to calculate factorial of numbers
func Factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	var fact uint = 1
	for i := uint(1); i <= n; i++ {
		fact *= i
	}
	
	return fact
}
