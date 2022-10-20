package hex

import (
	"encoding/hex"
	"fmt"
	"hex/endian"
	"strconv"
)

type Hex string

func NewHex(hexString string) (*Hex, error) {
	if hexString[:2] == "0x" {
		hexString = hexString[2:]
	}

	dst := make([]byte, hex.DecodedLen(len(hexString)))

	if _, err := hex.Decode(dst, []byte(hexString)); err != nil {
		return nil, err
	}

	result := new(Hex)

	*result = Hex(hexString)

	return result, nil
}

func (h *Hex) CountBytes() int64 {
	return int64(len(*h) / 2)
}

func (h *Hex) ToBigEndian() *endian.BigEndian {
	data := []byte{}

	for i := 0; i < len(*h); i += 2 {
		currentByte, _ := strconv.ParseInt(string(*h)[i:i+2], 16, 64)
		data = append(data, byte(currentByte))
	}

	result := new(endian.BigEndian)
	*result = data

	return result
}

func (h *Hex) ToLittleEndian() *endian.LittleEndian {
	data := []byte{}

	for i := len(*h) - 2; i >= 0; i -= 2 {
		currentByte, _ := strconv.ParseInt(string(*h)[i:i+2], 16, 64)
		data = append(data, byte(currentByte))
	}

	result := new(endian.LittleEndian)
	*result = data

	return result
}

func FromBigEndianToHex(bE *endian.BigEndian) *Hex {
	hexString := ""

	for _, currentByte := range *bE {
		hexString += fmt.Sprintf("%02x", currentByte)
	}

	result := new(Hex)
	*result = Hex(hexString)
	return result
}

func reverse(lst []byte) chan byte {
	ret := make(chan byte)
	go func() {
		for i, _ := range lst {
			ret <- lst[len(lst)-1-i]
		}
		close(ret)
	}()
	return ret
}

func FromLittleEndianToHex(lE *endian.LittleEndian) *Hex {
	hexString := ""

	for currentByte := range reverse(*lE) {
		hexString += fmt.Sprintf("%02x", currentByte)
	}

	result := new(Hex)
	*result = Hex(hexString)
	return result
}
