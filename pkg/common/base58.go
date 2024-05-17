package common

import (
	"crypto/sha256"
	"fmt"
	"runtime/debug"

	"github.com/shengdoushi/base58"
)

const addressLength = 20
const prefixMainnet = 0x41

func Encode(input []byte) string {
	return base58.Encode(input, base58.BitcoinAlphabet)
}

func EncodeCheck(input []byte) string {
	h256h0 := sha256.New()
	h256h0.Write(input)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	inputCheck := input
	inputCheck = append(inputCheck, h1[:4]...)

	return Encode(inputCheck)
}

func Decode(input string) ([]byte, error) {
	return base58.Decode(input, base58.BitcoinAlphabet)
}

func DecodeCheck(input string) ([]byte, error) {
	decodeCheck, err := Decode(input)
	if err != nil {
		return nil, err
	}

	fmt.Println("DecodeCheck ============ decodeCheck: ", decodeCheck)
	fmt.Println()
	if len(decodeCheck) < 4 {
		fmt.Println("DecodeCheck ============ decodeCheck 长度低于 4")
		debug.PrintStack()
		return nil, fmt.Errorf("b58 check error")
	}

	// tron address should should have 20 bytes + 4 checksum + 1 Prefix
	if len(decodeCheck) != addressLength+4+1 {
		return nil, fmt.Errorf("invalid address length: %d", len(decodeCheck))
	}

	// check prefix
	if decodeCheck[0] != prefixMainnet {
		return nil, fmt.Errorf("invalid prefix")
	}

	decodeData := decodeCheck[:len(decodeCheck)-4]

	h256h0 := sha256.New()
	h256h0.Write(decodeData)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	fmt.Println("DecodeCheck ============ h1[0] check: ", h1[0] == decodeCheck[len(decodeData)])
	fmt.Println("DecodeCheck ============ h1[1] check: ", h1[1] == decodeCheck[len(decodeData)+1])
	fmt.Println("DecodeCheck ============ h1[2] check: ", h1[2] == decodeCheck[len(decodeData)+2])
	fmt.Println("DecodeCheck ============ h1[3] check: ", h1[2] == decodeCheck[len(decodeData)+3])
	if h1[0] == decodeCheck[len(decodeData)] &&
		h1[1] == decodeCheck[len(decodeData)+1] &&
		h1[2] == decodeCheck[len(decodeData)+2] &&
		h1[3] == decodeCheck[len(decodeData)+3] {

		return decodeData, nil
	}

	debug.PrintStack()
	return nil, fmt.Errorf("b58 check error")
}
