package playfair

import (
	"strings"
)

func createKeyMatrix(key string) [5][5]rune {
	used := make(map[rune]bool)
	matrix := [5][5]rune{}
	alphabet := "abcdefghiklmnopqrstuvwxyz"
	x := 0
	y := 0
	for _, char := range key {
		if char == 'j' {
			char = 'i'
		}
		if !used[char] {
			matrix[x][y] = char
			used[char] = true
			y++
			if y == 5 {
				y = 0
				x++
			}
		}
	}
	for _, char := range alphabet {
		if !used[char] {
			matrix[x][y] = char
			y++
			if y == 5 {
				y = 0
				x++
			}
		}
	}
	return matrix
}

func prepareText(word string) string {
	var text strings.Builder

	for i, char := range word {
		if i > 1 && text.Len()%2 == 1 && word[i] == word[i-1] {
			text.WriteRune('x')
		}
		if char == 'j' {
			text.WriteRune('i')
		} else {
			text.WriteRune(char)
		}
	}
	if text.Len()%2 != 0 {
		text.WriteRune('x')
	}

	return text.String()
}

func findPosition(matrix [5][5]rune, ch rune) (int, int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if matrix[x][y] == ch {
				return x, y
			}
		}
	}
	return -1, -1
}

func Encryption(word string, key string) string {
	keyMatrix := createKeyMatrix(key)
	text := prepareText(word)

	var encrypted strings.Builder
	for i := 0; i < len(text); i += 2 {
		ch1, ch2 := rune(text[i]), rune(text[i+1])
		x1, y1 := findPosition(keyMatrix, ch1)
		x2, y2 := findPosition(keyMatrix, ch2)
		if x1 == x2 {
			encrypted.WriteRune(keyMatrix[x1][(y1+1)%5])
			encrypted.WriteRune(keyMatrix[x2][(y2+1)%5])
		} else if y1 == y2 {
			encrypted.WriteRune(keyMatrix[(x1+1)%5][y1])
			encrypted.WriteRune(keyMatrix[(x2+1)%5][y2])
		} else {
			encrypted.WriteRune(keyMatrix[x1][y2])
			encrypted.WriteRune(keyMatrix[x2][y1])
		}
	}

	return encrypted.String()
}

func Decryption(word string, key string) string {
	keyMatrix := createKeyMatrix(key)
	text := prepareText(word)

	var decrypted strings.Builder
	for i := 0; i < len(text); i += 2 {
		ch1, ch2 := rune(text[i]), rune(text[i+1])
		x1, y1 := findPosition(keyMatrix, ch1)
		x2, y2 := findPosition(keyMatrix, ch2)
		if x1 == x2 {
			decrypted.WriteRune(keyMatrix[x1][(5+y1-1)%5])
			decrypted.WriteRune(keyMatrix[x2][(5+y2-1)%5])
		} else if y1 == y2 {
			decrypted.WriteRune(keyMatrix[(5+x1-1)%5][y1])
			decrypted.WriteRune(keyMatrix[(5+x2-1)%5][y2])
		} else {
			decrypted.WriteRune(keyMatrix[x1][y2])
			decrypted.WriteRune(keyMatrix[x2][y1])
		}
	}

	return decrypted.String()
}
