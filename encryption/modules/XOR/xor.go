package xor

func XOREncryption(num, key int) int {
	ans := 0
	// level := 1
	for level := 1; num != 0; level *= 10 {
		ans = ans + ((num%10)^key)*level
		num /= 10
	}
	return ans
}
