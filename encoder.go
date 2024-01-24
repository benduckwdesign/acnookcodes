package acnookcodes

import (
	"fmt"
	"strings"
)

func TranspositionCipher(data []byte, negate bool, keyType int) []byte {
	sign := 1
	if negate {
		sign = -1
	}
	cipher := TranspositionCipherCharTable[keyType][data[KeyIdx[keyType]]&0xF]
	cipherPos := 0

	for i := 0; i < PASSWORD_DATA_SIZE; i++ {
		// Do not transpose the value at keyIdx
		if i == KeyIdx[keyType] {
			continue
		}

		//cipherByte := cipher[cipherPos%len(cipher)] // Retrieve the cipher byte
		data[i] = byte((int(data[i]) + int(cipher[cipherPos%len(cipher)])*sign) & 0xFF)
		cipherPos++
	}

	return data
}

func BitReverse(data []byte) []byte {
	for i := 0; i < PASSWORD_DATA_SIZE; i++ {
		if i != PASSWORD_BITMIXKEY_IDX {
			data[i] = ^data[i]
		}
	}
	return data
}

func BitArrangeReverse(data []byte) []byte {
	modifiedData := make([]byte, PASSWORD_DATA_SIZE-1)
	readData := make([]byte, PASSWORD_DATA_SIZE-1)

	//for i := 0; i < PASSWORD_DATA_SIZE; i++ {
	//	if i < PASSWORD_BITMIXKEY_IDX {
	//		readData[i] = data[i]
	//	} else if i > PASSWORD_BITMIXKEY_IDX {
	//		readData[i-1] = data[i]
	//	}
	//}
	copy(readData[:PASSWORD_BITMIXKEY_IDX], data[:PASSWORD_BITMIXKEY_IDX])
	copy(readData[PASSWORD_BITMIXKEY_IDX:], data[PASSWORD_BITMIXKEY_IDX+1:])

	//outIdx := 0
	//for i := PASSWORD_DATA_SIZE - 2; i >= 0; i-- {
	//	modifiedData[outIdx] = 0
	//	for b := 7; b >= 0; b-- {
	//		modifiedData[outIdx] |= (readData[i] >> uint(b) & 1) << (7 - b)
	//	}
	//	outIdx++
	//}
	var shiftValues [8]byte
	for i := 0; i < 8; i++ {
		shiftValues[i] = 1 << uint(7-i)
	}

	outIdx := 0
	for i := PASSWORD_DATA_SIZE - 2; i >= 0; i-- {
		val := readData[i]
		var result byte
		for b := 7; b >= 0; b-- {
			bit := (val >> uint(b)) & 1
			result |= bit * shiftValues[b]
		}
		modifiedData[outIdx] = result
		outIdx++
	}

	for i := 0; i < PASSWORD_DATA_SIZE; i++ {
		if i < PASSWORD_BITMIXKEY_IDX {
			data[i] = modifiedData[i]
		} else if i > PASSWORD_BITMIXKEY_IDX {
			data[i] = modifiedData[i-1]
		}
	}

	return data
}

func BitShift(data []byte, shift int) []byte {
	modifiedData := make([]byte, PASSWORD_DATA_SIZE-1)
	readData := make([]byte, PASSWORD_DATA_SIZE-1)

	for i := 0; i < PASSWORD_DATA_SIZE; i++ {
		if i < PASSWORD_BITMIXKEY_IDX {
			readData[i] = data[i]
		} else if i > PASSWORD_BITMIXKEY_IDX {
			readData[i-1] = data[i]
		}
	}

	if shift > 0 {
		//dstPos := shift / 8
		//dstOfs := shift % 8
		//
		//for i := 0; i < len(modifiedData); i++ {
		//	dst := (i + int(dstPos)) % len(modifiedData)
		//	src := (i + (len(readData) - 1)) % len(readData)
		//	rs := 8 - dstOfs
		//
		//	modifiedData[dst] = (readData[i] << uint(dstOfs)) | (readData[src]>>uint(rs))&0xFF
		//}
		//
		//for i := 0; i < len(data); i++ {
		//	if i < PASSWORD_BITMIXKEY_IDX {
		//		data[i] = modifiedData[i]
		//	} else if i > PASSWORD_BITMIXKEY_IDX {
		//		data[i] = modifiedData[i-1]
		//	}
		//}

		// Precompute constants outside the loop
		dstPos := shift / 8
		dstOfs := shift % 8
		readDataLen := len(readData)
		modifiedDataLen := len(modifiedData)
		mask := uint(0xFF)

		for i := 0; i < len(modifiedData); i++ {
			dst := (i + int(dstPos)) % modifiedDataLen
			src := (i + (readDataLen - 1)) % readDataLen
			rs := 8 - dstOfs

			// Simplify bitwise operations
			shiftedReadData := readData[i] << uint(dstOfs)
			rsData := byte(int(readData[src]) >> uint(rs) & int(mask))

			modifiedData[dst] = shiftedReadData | rsData
		}

		// Simplify the loop that copies elements from modifiedData to data
		for i := 0; i < len(data); i++ {
			switch {
			case i < PASSWORD_BITMIXKEY_IDX:
				data[i] = modifiedData[i]
			case i > PASSWORD_BITMIXKEY_IDX:
				data[i] = modifiedData[i-1]
			}
		}
	} else if shift < 0 {
		for i := 0; i < len(modifiedData); i++ {
			modifiedData[i] = readData[(len(modifiedData)-1)-i]
		}

		shift = -shift
		dstPos := shift / 8
		dstOfs := shift % 8

		for i := 0; i < len(modifiedData); i++ {
			readData[(i+int(dstPos))%len(modifiedData)] = modifiedData[i]
		}

		for i := 0; i < len(modifiedData); i++ {
			src := (i + (len(readData) - 1)) % len(readData)
			modifiedData[i] = readData[i]>>uint(dstOfs) | (readData[src] << uint(8-dstOfs))
			modifiedData[i] = modifiedData[i] & 0xFF
		}

		w := 0
		for i := 0; i < len(data); i++ {
			if i == PASSWORD_BITMIXKEY_IDX {
				w++
			}
			modifiedIdx := len(readData) - 1 - i
			if modifiedIdx < 0 {
				modifiedIdx += len(readData)
			}
			if modifiedIdx >= len(modifiedData) {
				modifiedIdx = modifiedIdx % len(modifiedData)
			}
			data[w%len(data)] = modifiedData[modifiedIdx]
			w++
		}
	}

	return data
}

func getRSAKeyCode(data []byte) RSAKeyInfo {
	var rsa_info uint32 = uint32(uint8(data[PASSWORD_RSA_KEY01_IDX]))
	var p_idx int32 = int32(rsa_info & 3)
	var q_idx int32 = int32((rsa_info >> 2) & 3)

	// Ensure that key0 and key1 differ.
	// NOTE: key0 & key1 will always be two of the following: 17, 19, or 23.
	if p_idx == 3 {
		p_idx = (p_idx ^ q_idx) & 3
		if p_idx == 3 {
			p_idx = 0
		}
	}

	if q_idx == 3 {
		q_idx = (p_idx + 1) & 3
		if q_idx == 3 {
			q_idx = 1
		}
	}

	if p_idx == q_idx {
		q_idx = (p_idx + 1) & 3
		if q_idx == 3 {
			q_idx = 1
		}
	}

	return RSAKeyInfo{
		p:         primeNumbers[p_idx],
		q:         primeNumbers[q_idx],
		selectTbl: selectIdxTable[((rsa_info >> 4) & 0xF)],
		e:         primeNumbers[data[PASSWORD_RSA_EXPONENT_IDX]],
	}
}

func ASCII2ACBytes(str string) []byte {
	//data := make([]byte, PARAM_STRING_SIZE)
	//for i := 0; i < PARAM_STRING_SIZE; i++ {
	//	if i < len(str) {
	//		character := str[i : i+1]
	//		idx := byte(characterMapIndex(character))
	//		if idx == 0xFF {
	//			data[i] = 0x20
	//		} else {
	//			data[i] = idx
	//		}
	//	} else {
	//		data[i] = 0x20
	//	}
	//}
	//
	//return data
	data := make([]byte, PARAM_STRING_SIZE)
	strLen := len(str)

	for i := 0; i < PARAM_STRING_SIZE; i++ {
		var idx byte
		if i < strLen {
			character := str[i]
			idx = byte(characterMapIndex(string(character)))
		} else {
			idx = 0xFF
		}

		if idx == 0xFF {
			data[i] = 0x20
		} else {
			data[i] = idx
		}
	}

	return data
}

func Uint8Array2ACBytes(str []byte) []byte {
	data := make([]byte, PARAM_STRING_SIZE)
	for i := 0; i < PARAM_STRING_SIZE; i++ {
		if i < len(str) {
			data[i] = str[i] & 0xFF
		} else {
			data[i] = 0x20
		}
	}

	return data
}

func makePasscode(codeType, hitRate, npcType, npcCode int, str0, str1 []byte, itemID int) []byte {

	npcCode &= 0xFF
	hitRateIdx := hitRate

	if codeType == Famicom {
		hitRateIdx = 1
		npcCode = 0xFF
	}

	if codeType == CardEMini {
		hitRateIdx = 1
		npcCode = 0xFF
	}

	if codeType == User {
		hitRateIdx = 1
		npcCode = 0xFF
	}

	if codeType == Popular {
		hitRateIdx = 4
	}

	if codeType == Magazine {
		npcType = (hitRateIdx >> 2) & 1
		hitRateIdx &= 3
		npcCode = 0xFF
	}

	data := make([]byte, PASSWORD_DATA_SIZE)

	data[0] = byte((codeType & 7) << 5)
	data[0] |= byte(uint32(hitRateIdx) << 1)
	data[0] |= byte(npcType & 1)

	data[1] = byte(npcCode)

	copy(data[2:], []byte(str0))
	copy(data[10:], []byte(str1))

	item_id := itemID
	data[18] = uint8(int32(item_id) >> 8)
	data[19] = uint8(int32(item_id))

	var checksum int

	for i := 0; i < PARAM_STRING_SIZE; i++ {
		checksum += int(data[2+i])
	}

	for i := 0; i < PARAM_STRING_SIZE; i++ {
		checksum += int(data[10+i])
	}

	checksum += int(itemID) & 0xFF
	checksum += int(npcCode)

	data[0] |= byte((checksum & 3) << 3)

	if DEBUG {
		fmt.Println("codeType:", codeType, "hitRate:", hitRateIdx, "itemId:", itemID, "str0:", str0, "str1:", str1, "checksum:", data[0])
	}

	return data
}

func EncodeSubstitutionCipher(data []byte) []byte {
	for i := 0; i < PASSWORD_DATA_SIZE; i++ {
		data[i] = changeCodeTbl[data[i]]
	}
	return data
}

var encodeBitShuffle_ModOffsets_20 map[int][]int = func() map[int][]int {

	self := map[int][]int{}

	for i, _ := range selectIdxTable {
		modOffsets := make([]int, len(selectIdxTable[0]))
		for j, selectVal := range selectIdxTable[i] {
			modOffsets[j] = selectVal % 20
		}
		self[i] = modOffsets
	}

	return self

}()

var encodeBitShuffle_ModOffsets_19 map[int][]int = func() map[int][]int {

	self := map[int][]int{}

	for i, _ := range selectIdxTable {
		modOffsets := make([]int, len(selectIdxTable[0]))
		for j, selectVal := range selectIdxTable[i] {
			modOffsets[j] = selectVal % 19
		}
		self[i] = modOffsets
	}

	return self

}()

func EncodeBitShuffle(data []byte, key int) []byte {
	//var keyIdx, count int
	//if key == 0 {
	//	keyIdx = 13
	//	count = 19
	//} else {
	//	keyIdx = 2
	//	count = 20
	//}
	keyIdx := 2
	count := 20

	if key == 0 {
		keyIdx = 13
		count = 19
	}

	readData := make([]byte, PASSWORD_DATA_SIZE-1)
	modData := make([]byte, PASSWORD_DATA_SIZE-1)

	//for i := 0; i < PASSWORD_DATA_SIZE; i++ {
	//	if i < keyIdx {
	//		readData[i] = data[i]
	//	} else if i > keyIdx {
	//		readData[i-1] = data[i]
	//	}
	//}
	if keyIdx > 0 {
		copy(readData[:keyIdx], data[:keyIdx])
	}
	if keyIdx < PASSWORD_DATA_SIZE-1 {
		copy(readData[keyIdx:], data[keyIdx+1:])
	}

	//selectTbl := selectIdxTable[data[keyIdx]&3]

	//for i := 0; i < count; i++ {
	//	byteVal := readData[i]
	//	for j := 0; j < 8; j++ {
	//		outputOfs := (selectTbl[j] + i) % count
	//		bit := byteVal & (1 << j)
	//		modData[outputOfs] |= bit
	//	}
	//}

	var modOffsets []int

	if count == 20 {
		modOffsets = encodeBitShuffle_ModOffsets_20[int(data[keyIdx]&3)]
	} else if count == 19 {
		modOffsets = encodeBitShuffle_ModOffsets_19[int(data[keyIdx]&3)]
	}

	// Loop through data and apply modifications
	for i := 0; i < count; i++ {
		byteVal := readData[i]

		if count == 20 {
			for j, offset := range modOffsets {
				bit := byteVal & (1 << uint(j))
				outputOfs := (offset + i) % count
				modData[outputOfs] |= bit
			}
		} else if count == 19 {
			for j, offset := range modOffsets {
				bit := byteVal & (1 << uint(j))
				outputOfs := (offset + i) % count
				modData[outputOfs] |= bit
			}
		}
	}

	//for i := 0; i < len(data); i++ {
	//	if i < keyIdx {
	//		data[i] = modData[i]
	//	} else if i > keyIdx {
	//		data[i] = modData[i-1]
	//	}
	//}
	if keyIdx > 0 {
		copy(data[:keyIdx], modData[:keyIdx])
	}
	if keyIdx < len(data)-1 {
		copy(data[keyIdx+1:], modData[keyIdx:])
	}

	return data
}

func modularExponentiation(m, e, n int) int {
	result := 1
	base := m % n

	for e > 0 {
		if e&1 == 1 {
			result = (result * base) % n
		}
		e >>= 1
		base = (base * base) % n
	}

	return result
}

func EncodeChangeRSACipher(data []byte) []byte {
	rsaData := getRSAKeyCode(data)

	rsaBitSave := 0                // Each bit represents the 9th bit in our ciphertext values
	n := rsaData.p * rsaData.q     // The multiple of our primes
	e := rsaData.e                 // Our exponent
	selectTbl := rsaData.selectTbl // Array of 8 byte indexes which we will apply the RSA encryption to

	for i := 0; i < 8; i++ {
		c := int(data[selectTbl[i]])
		//m := c

		c = modularExponentiation(c, e, n)
		//// Modular Exponentiation from [2, e]
		//for j := 0; j < e-1; j++ {
		//	// c will always be below one of the following: 17*19 (323), 17*23 (391), 19*23 (437)
		//	// In other words, c is in the range [0, p*q).
		//	c = (c * m) % n
		//}

		data[selectTbl[i]] = byte(c & 0xFF)
		rsaBitSave |= (int(c>>8) & 1) << i // Save the 9th bit in case the ciphertext went over 255.
	}

	if DEBUG {
		fmt.Println("RSA keysave:", byte(rsaBitSave&0xFF), "||", rsaBitSave)
	}
	data[PASSWORD_RSA_BITSAVE_IDX] = byte(rsaBitSave & 0xFF)
	return data
}

func EncodeChangeRSACipherOverride(data []byte, keysave int) []byte {
	rsaData := getRSAKeyCode(data)

	rsaBitSave := 0                // Each bit represents the 9th bit in our ciphertext values
	n := rsaData.p * rsaData.q     // The multiple of our primes
	e := rsaData.e                 // Our exponent
	selectTbl := rsaData.selectTbl // Array of 8 byte indexes which we will apply the RSA encryption to

	for i := 0; i < 8; i++ {
		c := int(data[selectTbl[i]])
		m := c

		// Modular Exponentiation from [2, e]
		for j := 0; j < e-1; j++ {
			// c will always be below one of the following: 17*19 (323), 17*23 (391), 19*23 (437)
			// In other words, c is in the range [0, p*q).
			c = (c * m) % n
		}

		data[selectTbl[i]] = byte(c & 0xFF)
		rsaBitSave |= (int(c>>8) & 1) << i // Save the 9th bit in case the ciphertext went over 255.
	}

	if DEBUG {
		fmt.Println("RSA keysave:", byte(keysave&0xFF), "||", keysave)
	}
	data[PASSWORD_RSA_BITSAVE_IDX] = byte(keysave & 0xFF)
	return data
}

func EncodeBitMixCode(data []byte) []byte {
	code := int(uint8(data[PASSWORD_BITMIXKEY_IDX]) & 0xF)

	if DEBUG {
		fmt.Println("Bit mix code:", code)
	}

	if code > 12 { // [13, 15] (3 or 18.75%)
		BitArrangeReverse(data)
		if DEBUG {
			fmt.Println("substep: BitArrangeReverse:", data)
		}
		BitReverse(data)
		if DEBUG {
			fmt.Println("substep: BitReverse:", data)
		}
		BitShift(data, code*3)
		if DEBUG {
			fmt.Println("substep: BitShift:", data)
		}
	} else if code > 8 { // [9, 12] (4 or 25%)
		BitArrangeReverse(data)
		if DEBUG {
			fmt.Println("substep: BitArrangeReverse:", data)
		}
		BitShift(data, code*-5)
		if DEBUG {
			fmt.Println("substep: BitShift:", data)
		}
	} else if code > 4 { // [5, 8] (4 or 25%)
		BitShift(data, code*-5)
		if DEBUG {
			fmt.Println("substep: BitShift:", data)
		}
		BitReverse(data)
		if DEBUG {
			fmt.Println("substep: BitReverse:", data)
		}
	} else { // [0, 4] (5 or 31.25%)
		BitShift(data, code*3)
		if DEBUG {
			fmt.Println("substep: BitShift:", data)
		}
		BitArrangeReverse(data)
		if DEBUG {
			fmt.Println("substep: BitArrangeReverse:", data)
		}
	}

	return data
}

func Change6BitsCode(eightbits []byte) []byte {
	var six_byte byte = 0
	six_idx := 0
	eight_bit := 0
	six_bit := 0

	var sixbits []byte = make([]byte, PASSWORD_STR_SIZE)

	for {
		bit := (eightbits[0] >> eight_bit & 1) << six_bit
		eight_bit++
		six_bit++
		six_byte |= bit

		if six_bit >= 6 {
			(sixbits)[six_idx] = six_byte
			six_idx++
			six_bit = 0

			if six_idx >= PASSWORD_STR_SIZE {
				return sixbits
			}

			six_byte = 0
		}

		if eight_bit >= 8 {
			eight_bit = 0
			eightbits = eightbits[1:]
		}
	}
}

func ChangeCommonFontCode_Uint8Array(passwordData []byte) []byte {
	password := make([]byte, PASSWORD_STR_SIZE)
	for i := 0; i < PASSWORD_STR_SIZE; i++ {
		password[i] = byte(usable2fontnum[passwordData[i]])
	}

	return password
}

func MakePassword(codeType int, hitRateIdx int, str0 []byte, str1 []byte, itemID int, npcType int, npcCode int) []byte {
	data := makePasscode(codeType, hitRateIdx, npcType, npcCode, str0, str1, itemID)
	if DEBUG {
		fmt.Println("makePasscode:", data)
	}
	EncodeSubstitutionCipher(data)
	if DEBUG {
		fmt.Println("EncodeSubstitutionCipher:", data)
	}
	TranspositionCipher(data, true, 0)
	if DEBUG {
		fmt.Println("TranspositionCipher:", data)
	}
	EncodeBitShuffle(data, 0)
	if DEBUG {
		fmt.Println("EncodeBitShuffle:", data)
	}
	EncodeChangeRSACipher(data)
	if DEBUG {
		fmt.Println("EncodeChangeRSACipher:", data)
	}
	EncodeBitMixCode(data)
	if DEBUG {
		fmt.Println("EncodeBitMixCode:", data)
	}
	EncodeBitShuffle(data, 1)
	if DEBUG {
		fmt.Println("EncodeBitShuffle:", data)
	}
	TranspositionCipher(data, false, 1)
	if DEBUG {
		fmt.Println("TranspositionCipher:", data)
	}
	data = Change6BitsCode(data)
	if DEBUG {
		fmt.Println("Change6BitsCode:", data)
	}
	data = ChangeCommonFontCode_Uint8Array(data)
	if DEBUG {
		fmt.Println("ChangeCommonFontCode_Uint8Array:", data)
	}
	return data
}

func MakePasswordKeysave(codeType int, hitRateIdx int, str0 []byte, str1 []byte, itemID int, npcType int, npcCode int, keysave int) []byte {
	data := makePasscode(codeType, hitRateIdx, npcType, npcCode, str0, str1, itemID)
	if DEBUG {
		fmt.Println("makePasscode:", data)
	}
	EncodeSubstitutionCipher(data)
	if DEBUG {
		fmt.Println("EncodeSubstitutionCipher:", data)
	}
	TranspositionCipher(data, true, 0)
	if DEBUG {
		fmt.Println("TranspositionCipher:", data)
	}
	EncodeBitShuffle(data, 0)
	if DEBUG {
		fmt.Println("EncodeBitShuffle:", data)
	}
	EncodeChangeRSACipherOverride(data, keysave)
	if DEBUG {
		fmt.Println("EncodeChangeRSACipher:", data)
	}
	EncodeBitMixCode(data)
	if DEBUG {
		fmt.Println("EncodeBitMixCode:", data)
	}
	EncodeBitShuffle(data, 1)
	if DEBUG {
		fmt.Println("EncodeBitShuffle:", data)
	}
	TranspositionCipher(data, false, 1)
	if DEBUG {
		fmt.Println("TranspositionCipher:", data)
	}
	data = Change6BitsCode(data)
	if DEBUG {
		fmt.Println("Change6BitsCode:", data)
	}
	data = ChangeCommonFontCode_Uint8Array(data)
	if DEBUG {
		fmt.Println("ChangeCommonFontCode_Uint8Array:", data)
	}
	return data
}

func characterMapIndex(char string) int {
	for i, c := range characterMap {
		if c == char {
			return i
		}
	}
	return -1
}

func ConvertBytesToUnicodeString(bytes []byte) string {
	var builder strings.Builder

	// Iterate over the bytes and map them to Unicode characters
	for _, b := range bytes {
		builder.WriteString(characterMap[int(b)])
	}

	return builder.String()
}
