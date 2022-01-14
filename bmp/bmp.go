package bmp

import (
	"io/ioutil"
	"os"
)

type BMP struct {
	fileHeader [14]byte
	bitMapInfo [40]byte
	colorPalette []byte
	bitMapData []byte
}


// ReadImage 读取图片数据
func (bmp *BMP) ReadImage(imagePath string) {
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		panic(err)
	}
	// 给bmp文件头
	copy(bmp.fileHeader[:], data[:14])
	// bmp位图信息
	copy(bmp.bitMapInfo[:], data[14:54])
	// bmp 调色板
	offset := int32(bmp.fileHeader[10]) + int32(bmp.fileHeader[11]) << 8 + int32(bmp.fileHeader[12]) << 16 + int32(bmp.fileHeader[13]) << 24
	bmp.colorPalette = make([]byte, len(data[54:offset]))
	copy(bmp.colorPalette, data[54:offset])

	// bmp位图数据
	bmp.bitMapData = make([]byte, len(data[offset:]))
	copy(bmp.bitMapData, data[offset:])

}

// WriteToImage 将数据写入bmp图片中
func (bmp *BMP) WriteToImage(filename string) {
	data := append(bmp.fileHeader[:], bmp.bitMapInfo[:]...)
	data = append(data, bmp.colorPalette...)
	data = append(data, bmp.bitMapData...)

	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	out.Write(data)
}

// GetBitMapData 获取位图数据
func (bmp *BMP) GetBitMapData() []byte {
	return bmp.bitMapData
}

func (bmp *BMP) UpdateBitMapData(data []byte)  {
	copy(bmp.bitMapData, data)
}