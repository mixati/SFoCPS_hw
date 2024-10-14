package rsa

import (
	"math/rand"
	"strings"
)

type Key struct {
	A int
	B int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcModule(p, q int) int {
	return p * q
}

func calcEuler(p, q int) int {
	return (p - 1) * (q - 1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func generateE(euler int) int {
	var num int
	flag := true
	for flag {
		num = rand.Intn(euler-1) + 1
		if gcd(num, euler) == 1 {
			flag = false
		}
	}
	return num
}

func calcD(euler, e int) int {
	a := euler
	b := e
	var q, r int
	x1 := 0
	x2 := 1
	y1 := 1
	y2 := 0
	for b > 0 {
		q = a / b
		r = a - q*b
		x := x2 - q*x1
		y := y2 - q*y1
		a = b
		b = r
		x2 = x1
		x1 = x
		y2 = y1
		y1 = y
	}
	d := euler - abs(min(x2, y2))

	return d
}

func GenerateKeys(p, q int) (Key, Key) {
	n := calcModule(p, q)
	euler := calcEuler(p, q)
	e := generateE(euler)
	d := calcD(euler, e)
	return Key{e, n}, Key{d, n}
}

func convertText(word string) []int {
	var convertedText []int
	for _, char := range word {
		convertedText = append(convertedText, int(char-'a'))
	}
	return convertedText
}

func encryptSymbol(num int, key Key) int {
	res := 1
	for i := 0; i < key.A; i++ {
		res = (res * num) % key.B
	}
	return res
}

func Encryption(word string, key Key) []int {
	convertedText := convertText(word)
	for i := 0; i < len(word); i++ {
		convertedText[i] = encryptSymbol(convertedText[i], key)
	}
	return convertedText
}

func Decryption(word []int, key Key) string {
	for i := 0; i < len(word); i++ {
		word[i] = encryptSymbol(word[i], key)
	}
	var text strings.Builder
	for i := 0; i < len(word); i++ {
		text.WriteRune(rune('a' + word[i]))
	}
	return text.String()
}
