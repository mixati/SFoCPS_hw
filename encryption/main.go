package main

import (
	"encryption/modules/XOR"
	"encryption/modules/caesar"
	"encryption/modules/vigenere"
	"fmt"
)

func main() {

	// fmt.Println("Введите алгоритм шифрования:")
	word := "hello"
	key := 23

	encrypt := caesar.CaesarEncryption(word, key)
	dectypt := caesar.CaesarDecryption(encrypt, key)
	fmt.Println(encrypt, dectypt)

	keyVig := "qwre"

	enc := vigenere.VigenereEncryption(word, keyVig)
	dec := vigenere.VigenereDecrypter(enc, keyVig)
	fmt.Println(enc, dec)

	keyXOR := 7
	wordXOR := 1234

	st := xor.XOREncryption(wordXOR, keyXOR)
	cr := xor.XOREncryption(st, keyXOR)
	fmt.Println(st, cr)

}
