package help

import (
	"encoding/binary"
	"encoding/hex"
)

func B2H(a []byte) string {
	return hex.EncodeToString(a)
}

func B2I(a []byte) uint16 {
	return binary.BigEndian.Uint16(a)
}

func I2B(a uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, a)
	return b
}

func Concat(buffers ...[]byte) []byte {
	var buffer []byte
	for _, b := range buffers {
		buffer = append(buffer, b...)
	}
	return buffer

}
