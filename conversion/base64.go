package conversion

import "strings"

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Base64Encode(input []byte) string {
	var sb strings.Builder

	var padding string
	for i := len(input) % 3; i > 0 && i < 3; i++ {
		var zeroByte byte
		input = append(input, zeroByte)
		padding += "="
	}

	for i := 0; i < len(input); i += 3 {
		group := [4]byte{
			input[i] >> 2,
			(input[i]<<4)&0x3F + input[i+1]>>4,
			(input[i+1]<<2)&0x3F + input[i+2]>>6,
			input[i+2] & 0x3F,
		}

		for _, b := range group {
			sb.WriteString(string(Alphabet[int(b)]))
		}
	}

	encoded := sb.String()

	encoded = encoded[:len(encoded)-len(padding)] + padding[:]

	return encoded
}

func Base64Decode(input string) []byte {
	padding := strings.Count(input, "=")
	var decoded []byte

	for i := 0; i < len(input); i += 4 {
		byteInput := [4]byte{
			byte(strings.IndexByte(Alphabet, input[i])),
			byte(strings.IndexByte(Alphabet, input[i+1])),
			byte(strings.IndexByte(Alphabet, input[i+2])),
			byte(strings.IndexByte(Alphabet, input[i+3])),
		}

		group := [3]byte{
			byteInput[0]<<2 + byteInput[1]>>4,
			byteInput[1]<<4 + byteInput[2]>>2,
			byteInput[2]<<6 + byteInput[3],
		}

		decoded = append(decoded, group[:]...)
	}

	return decoded[:len(decoded)-padding]
}
