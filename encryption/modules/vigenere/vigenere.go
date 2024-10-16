package vigenere

func Encryption(word string, key string) string {
	var ans string
	keyLen := len(key)
	for pos, char := range word {
		shift := rune(key[pos%keyLen] - 'a')
		if char >= 'a' && char <= 'z' {
			ans += string((char-'a'+shift)%26 + 'a')
		} else {
			ans += string(char)
		}
	}

	return ans
}

func Decryption(word string, key string) string {
	var ans string
	keyLen := len(key)
	for pos, char := range word {
		shift := rune(key[pos%keyLen] - 'a')
		if char >= 'a' && char <= 'z' {
			ans += string((char-'a'-shift+26)%26 + 'a')
		} else {
			ans += string(char)
		}
	}

	return ans
}
