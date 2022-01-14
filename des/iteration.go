package des

import (
	"strconv"
)

// Feistel 加密 16轮置换
func Feistel(data, secretKey string) string {
	L, R := data[:32], data[32:]
	// 子密钥列表
	subkeys := subkeyGenerating(secretKey)

	// 16轮置换
	for i := 0; i < 16; i++ {
		L, R = R, Xor(L, f(R, subkeys[i]))

	}
	// 第16轮将得到的L16和R16两部分整体进行互换，得到最终的L16R16。
	L, R = R, L
	data = L + R

	return data
}


/*
DES的解密过程和DES的加密过程完全类似，
只不过将16圈的子密钥序列K1，K2……K16的顺序倒过来。即第一圈用第16个子密钥K16，第二圈用K15，
*/

//FeistelDecode Feistel 解密 16轮置换
func FeistelDecode(data, secretKey string) string {
	L, R := data[:32], data[32:]
	// 子密钥列表
	subkeys := subkeyGenerating(secretKey)

	// 16轮置换
	for i := 15; i >= 0; i-- {
		L, R = R, Xor(L, f(R, subkeys[i]))

	}
	// 第16轮将得到的L16和R16两部分整体进行互换，得到最终的L16R16。
	L, R = R, L
	data = L + R

	return data
}



// pc-1置换
func pc1Replacement(secretKey string) string {
	newSecretKey := ""
	pc1Matrix := GetPc1Matrix()
	for i := 0; i < len(pc1Matrix); i++ {
		newSecretKey += string(secretKey[pc1Matrix[i]-1])
	}

	return newSecretKey
}

// 左移循环
func ls1(key string, i int) string {
	leftCircle := GetLeftCircle()
	key = key[leftCircle[i]:] + key[:leftCircle[i]]

	return key
}

// pc-2置换
func pc2Replacement(key string) string {
	subkey := ""
	pc2Matrix := GetPc2Matrix()
	for i := 0; i < len(pc2Matrix); i++ {
		subkey += string(key[pc2Matrix[i]-1])
	}

	return subkey
}

// 子密钥生成
func subkeyGenerating(secretKey string) []string {
	var subkeys = make([]string, 0)
	// pc-1置换
	secretKey = pc1Replacement(secretKey)
	C, D := secretKey[:28], secretKey[28:]

	for i := 0; i < 16; i++ {
		C, D = ls1(C, i), ls1(D, i)
		newSubkey := pc2Replacement(C+D)
		subkeys = append(subkeys, newSubkey)
	}

	return subkeys
}

// 扩展置换
func extendReplacement(text string) string {
	newText := ""
	extendMatrix := GetExtendedReplacementMatrix()
	for i := 0; i < len(extendMatrix); i++ {
		newText += string(text[extendMatrix[i]-1])
	}
	return newText
}


// s盒置换
func sBoxReplacement(RExtendedXOR string) string {
	sBoxReplace := ""
	sBox := GetSBox()
	for i := 0; i < 8; i++ {
		a := RExtendedXOR[i*6:(i+1)*6]
		sBoxLineNumber, _ := strconv.ParseInt(a[:1]+a[5:], 2, 0)
		sBoxColumnNumber, _ := strconv.ParseInt(a[1:5], 2, 0)
		partText := strconv.FormatInt(int64(sBox[i][sBoxLineNumber][sBoxColumnNumber]), 2)
		for len(partText) < 4 {
			partText = "0" + partText
		}
		sBoxReplace += partText
	}

	return sBoxReplace
}

// p盒置换
func pBoxReplacement(sBoxReplace string) string {
	pBox := GetPBox()
	pBoxReplace := ""
	for i := 0; i < len(pBox); i++ {
		pBoxReplace += string(sBoxReplace[pBox[i]-1])
	}

	return pBoxReplace
}

/*
非线性函数f
r(i-1) -> 拓展置换 ->48bit -> ^ k -> s置换 -> p置换 -> 32bit
*/
func f(r, k string) string {
	// r拓展置换
	extendR := extendReplacement(r)
	// extendR, k异或
	RExtendedXOR := Xor(extendR, k)
	// s盒置换
	sBoxReplace := sBoxReplacement(RExtendedXOR)
	// p盒置换
	pBoxReplace := pBoxReplacement(sBoxReplace)

	return pBoxReplace
}
