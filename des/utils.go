package des

import (
	"bytes"
	"fmt"
)

const blockSize int = 8

// PKCS5Padding PKCS5填充
func PKCS5Padding(cipherText []byte) []byte {
	padding := blockSize - len(cipherText) % blockSize
	if padding < 8 {
		padText := bytes.Repeat([]byte{byte(padding)}, padding)
		cipherText = append(cipherText, padText...)
	}

	return cipherText
}

// ToBinary 字节切片转换为二进制
func ToBinary(data []byte) string {
	binaryText := ""
	for _, v := range data {
		binaryText += fmt.Sprintf("%.8b", v)
	}
	return binaryText
}

// BinaryToByte 将二进制数据转换为字节数组
func BinaryToByte(bin string) []byte {
	byteData := make([]byte, 0)
	for i := 0; i < len(bin); i+=8 {
		byteData = append(byteData, sToI(bin[i:i+8]))
	}

	return byteData
}

// 8为二进制字符串转化为字节
func sToI(str string) uint8 {
	return (str[0]-48)<<7 + (str[1]-48)<<6 + (str[2]-48)<<5 + (str[3]-48)<<4 + (str[4]-48)<<3 +
		(str[5]-48)<<2 + (str[6]-48)<<1 + (str[7] - 48)
}

// Xor 异或
func Xor(l, r string) string {
	RExtendedXOR := ""
	for i := 0; i < len(r); i++ {
		RExtendedXOR += string(((l[i]-'0') ^ (r[i]-'0')) + '0')
	}
	return RExtendedXOR
}