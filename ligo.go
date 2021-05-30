package ligo

import (
	"math/big"
	"math/rand"
	"strconv"
	"strings"
)

var encodedToBinaryReplacer = strings.NewReplacer("l", "1", "I", "0")
var binaryToEncodedReplacer = strings.NewReplacer("1", "l", "0", "I")

func WebsiteToCode(website string, codeLen int) string {
	var x big.Int
	x.SetBytes([]byte(website))
	binaryStr := x.Text(2)

	eightPad := len(binaryStr) % 8
	if eightPad != 0 {
		binaryStr = strings.Repeat("0", 8-eightPad) + binaryStr
	}

	padding := codeLen - len(binaryStr)
	if padding > 0 {
		binaryStr += "00000000"
		padding -= 8
		if padding > 0 {
			for i := 0; i < padding; i++ {
				binaryStr += strconv.Itoa(rand.Intn(2))
			}
		}
	}

	return binaryToEncodedReplacer.Replace(binaryStr)
}

func CodeToWebsite(code string) (string, bool) {
	binaryStr := encodedToBinaryReplacer.Replace(code)

	var x big.Int
	_, ok := x.SetString(binaryStr, 2)
	if !ok {
		return "", false
	}
	str := string(x.Bytes())

	zeroIndex := strings.IndexByte(str, 0)
	if zeroIndex < 0 {
		zeroIndex = len(str)
	}
	return str[:zeroIndex], true
}
