package des

import (
	"bytes"
	"fmt"
	"os"
)

// encode des加密
func encode(text, secretKey string) string {
	//初始置换
	text = InitReplacement(text)
	// 16轮迭代运算
	text = Feistel(text, secretKey)
	//逆初始置换
	cipherText := ReverseReplacement(text)

	return cipherText
}

// decode
func decode(text, secretKey string) string {
	//初始置换
	text = InitReplacement(text)
	// 16轮迭代运算
	text = FeistelDecode(text, secretKey)
	//逆初始置换
	cipherText := ReverseReplacement(text)

	return cipherText
}

// des3 3DES加密 解密
func des3(cipherText string, secretKey string, encodeOrDecode string) string {
	if len(secretKey) != 192 {
		fmt.Println("密钥长度必须为192bit")
		os.Exit(0)
	}
	if encodeOrDecode == "encode" {
		for i := 0; i < 3; i++ {
			cipherText = encode(cipherText, secretKey[i*64:(i+1)*64])
		}
	} else if encodeOrDecode == "decode" {
		for i := 0; i < 3; i++ {
			cipherText = decode(cipherText, secretKey[i*64:(i+1)*64])
		}
	}

	return cipherText
}

// CBCDes3 以CBC分组运行模式进行加密或解密
func CBCDes3(data, secretKey string) string {
	init := string(bytes.Repeat([]byte{byte(48)}, 64))
	cipherData := ""

	iter := len(data) / 64
	for i := 0; i < iter; i++ {
		text := Xor(init, data[i*64:(i+1)*64])
		init = des3(text, secretKey, "encode")
		cipherData += init
	}

	return cipherData
}


func CBCDecodeDes3(data, secretKey string) string {
	init := string(bytes.Repeat([]byte{byte(48)}, 64))
	cipherData := ""

	iter := len(data) / 64
	for i := 0; i < iter; i++ {
		cipherText := data[i*64:(i+1)*64]
		init, cipherText = cipherText, Xor(init, des3(cipherText, secretKey, "decode"))
		cipherData += cipherText
	}

	return cipherData
}
