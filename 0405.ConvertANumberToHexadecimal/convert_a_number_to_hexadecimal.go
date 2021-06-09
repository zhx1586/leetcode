package convertanumbertohexadecimal

import (
	"strings"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hex := map[int]byte{
		0x0: '0',
		0x1: '1',
		0x2: '2',
		0x3: '3',
		0x4: '4',
		0x5: '5',
		0x6: '6',
		0x7: '7',
		0x8: '8',
		0x9: '9',
		0xa: 'a',
		0xb: 'b',
		0xc: 'c',
		0xd: 'd',
		0xe: 'e',
		0xf: 'f',
	}
	length := 8
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		shift := (length - 1 - i) * 4
		mask := 0xf << shift
		intersection := num & mask
		ret[i] = hex[intersection>>shift]
	}

	return strings.TrimLeft(string(ret), "0")
}
