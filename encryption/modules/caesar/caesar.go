package caesar

func Encryption(word string, key int) string {
	var ans string
	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			ans += string((char-'a'+rune(key))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			ans += string((char-'A'+rune(key))%26 + 'A')
		} else if char >= 'а' && char <= 'я' {
			ans += string((char-'а'+rune(key))%33 + 'а')
		} else if char >= 'А' && char <= 'Я' {
			ans += string((char-'А'+rune(key))%33 + 'А')
		} else {
			ans += string(char)
		}
	}

	return ans
}

func Decryption(word string, key int) string {
	var ans string
	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			ans += string((char-'a'-rune(key)+26)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			ans += string((char-'A'-rune(key)+26)%26 + 'A')
		} else if char >= 'а' && char <= 'я' {
			ans += string((char-'а'-rune(key)+33)%33 + 'а')
		} else if char >= 'А' && char <= 'Я' {
			ans += string((char-'А'-rune(key)+33)%33 + 'А')
		} else {
			ans += string(char)
		}
	}

	return ans
}
