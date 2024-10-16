# SFoCPS_hw
## Реализация шифрования Цезаря, Виженера, XOR, Плейфера и RSA

В данном проекте реализованы модули для шифрования и дешифрования c помощью алгоритмов Цезаря, Виженера, XOR, Плейфера и RSA.

### Цезарь, Виженер и Плейфер

Модули [caesar](encryption/modules/caesar/), [vigenere](encryption/modules/vigenere) и [playfair](encryption/modules/playfair/) содержат публичные функции Encryption и Decryption, предназначенные для шифрования и дешифрования текста по алгоритмам Цезаря, Виженера и Плейфера соответственно.

### XOR
Модуль [xor](encryption/modules/xor/) содержит публичную функцию Encryption, предназначенную и для шифрования, и для дешифрования числа по алгоритму XOR.

### RSA

Модуль [rsa](encryption/modules/rsa/) содержит публичные функции Encryption и Decryption, предназначенные для шифрования и дешифрования по алгоритму RSA.

Также в модуле присутсвует фунция GenerateKeys для создания публичного и приватного ключа на основе 2 простых чисел, однако данная функция работает некорректно

## Пример

В данном проекте можно увидеть [пример](encryption/main.go) использования написанных модулей

## Как запустить

Для запуска примера необходимо установить пакеты для работы с языком Golang, перейти в раздел [encryption](encryption) и выполнить команду:
```
go run main.go
```