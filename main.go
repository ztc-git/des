package main

import (
	"des/bmp"
	"des/des"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	// 申明bmp结构体，获取bmp图片数据
	var bmp = new(bmp.BMP)
	bmp.ReadImage("./gopher.bmp")

	//需要加密的位图数据部分
	data := bmp.GetBitMapData()
	length := len(data)
	// 加密密钥
	secretKey := []byte("qwertyuiiuytrewqqwertyui")

	// PKcs5 使加密明文长度是8的倍数
	data = des.PKCS5Padding(data)

	//将明文数据和密钥的转化为二进制
	cipherText := des.ToBinary(data)
	key := des.ToBinary(secretKey)

	// des3 加密
	cipherText = des.CBCDes3(cipherText, key)
	// 二进制数据转化为字节数据，并写入图片
	data = des.BinaryToByte(cipherText)
	bmp.UpdateBitMapData(data[:length])
	bmp.WriteToImage("./encode.bmp")
	fmt.Println("加密完成")

	// des3 解密
	cipherText = des.CBCDecodeDes3(cipherText, key)
	// 二进制数据转化为字节数据，并写入图片
	data = des.BinaryToByte(cipherText)
	bmp.UpdateBitMapData(data[:length])

	bmp.WriteToImage("./decode.bmp")
	fmt.Println("解密完成")
	fmt.Println(time.Now())
}
