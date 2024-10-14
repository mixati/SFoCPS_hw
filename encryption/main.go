package main

import (
	"encryption/modules/caesar"
	"encryption/modules/playfair"
	"encryption/modules/rsa"
	"encryption/modules/vigenere"
	"encryption/modules/xor"
	"fmt"
)

func main() {
	/* Пример работы шифра Цезаря */
	fmt.Println("Пример работы шифра Цезаря:")
	wordCaesar := "hello"
	keyCaesar := 23
	fmt.Printf("Исходное слово: %s;\nКлюч: %d\n", wordCaesar, keyCaesar)
	encryptCaecar := caesar.Encryption(wordCaesar, keyCaesar)
	dectyptCaecar := caesar.Decryption(encryptCaecar, keyCaesar)
	fmt.Printf("Зашифрованный текст: %s\nРасшифровка: %s\n\n", encryptCaecar, dectyptCaecar)

	/* Пример работы шифра Виженера */
	fmt.Println("Пример работы шифра Виженера:")
	wordVigenere := "friend"
	keyVigenere := "qwre"
	fmt.Printf("Исходное слово: %s\nКлюч: %s\n", wordVigenere, keyVigenere)
	encryptVigenere := vigenere.Encryption(wordVigenere, keyVigenere)
	decryptVigenere := vigenere.Decryption(encryptVigenere, keyVigenere)
	fmt.Printf("Зашифрованный текст: %s\nРасшифровка: %s\n\n", encryptVigenere, decryptVigenere)

	/* Пример работы шифра XOR */
	fmt.Println("Пример работы шифра XOR:")
	wordXOR := 1234
	keyXOR := 7
	fmt.Printf("Исходное число: %d\nКлюч: %d\n", wordXOR, keyXOR)
	encryptXOR := xor.Encryption(wordXOR, keyXOR)
	decryptXOR := xor.Encryption(encryptXOR, keyXOR)
	fmt.Printf("Зашифрованное число: %d\nРасшифровка: %d\n\n", encryptXOR, decryptXOR)

	/* Пример работы шифра Плейфера */
	fmt.Println("Пример работы шифра Плейфера:")
	wordPlay := "helloworld"
	keyPlay := "qwer"
	fmt.Printf("Исходное слово: %s\nКлюч: %s\n", wordPlay, keyPlay)
	encryptPlay := playfair.Encryption(wordPlay, keyPlay)
	decryptPlay := playfair.Decryption(encryptPlay, keyPlay)
	fmt.Printf("Зашифрованный текст: %s\nРасшифровка: %s\n\n", encryptPlay, decryptPlay)

	/* Пример работы шифра RSA */
	fmt.Println("Пример работы шифра RSA:")
	wordRSA := "hellomyfriend"
	// pubKey, privKey := rsa.GenerateKeys(19, 41)
	pubKey, privKey := rsa.Key{A: 691, B: 779}, rsa.Key{A: 571, B: 779}
	fmt.Printf("Исходное слово: %s\n", wordRSA)
	fmt.Printf("Публичный ключ: %v\n", pubKey)
	fmt.Printf("Приватный ключ: %v\n", privKey)
	encryptRSA := rsa.Encryption(wordRSA, pubKey)
	decryptRSA := rsa.Decryption(encryptRSA, privKey)
	fmt.Printf("Зашифрованный текст: %v\n", encryptRSA)
	fmt.Printf("Расшифровка: %s\n\n", decryptRSA)
}
