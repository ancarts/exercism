package diffsquares

func SquareOfSum(n int) int {
	sum := (n + 1) * n / 2

	square := sum * sum

	return square
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
