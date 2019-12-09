package leetcode

func reverse(n int) int {
	newNumber := 0
	for n > 0 {
		remainder := n % 10
		newNumber *= 10
		newNumber += remainder
		n /= 10

	}
	return newNumber
}
