package acnookcodes

import (
	"bytes"
	"fmt"
)

// TODO: PERFECT
func Change8BitsCode(sixbits []byte) []byte {
	var eight_byte byte = 0
	eight_idx := 0
	six_bit := 0
	eight_bit := 0
	//eight_idx2 := 0
	six_idx2 := 0
	eight_idx2 := 0
	eightbits := &bytes.Buffer{}

	for {
		bit := ((sixbits[six_idx2%len(sixbits)] >> six_bit) & 1) << eight_bit
		eight_bit++
		six_bit++
		eight_byte |= bit

		if eight_bit >= 8 {
			eight_idx++
			eightbits.WriteByte(eight_byte)
			//(eightbits)[eight_idx2%len(eightbits)] = eight_byte
			eight_bit = 0
			eight_idx2++

			if eight_idx >= PASSWORD_STR_SIZE {
				f_eightbits := eightbits.Bytes()
				result := &bytes.Buffer{}
				for i := 0; i < PASSWORD_DATA_SIZE; i++ {
					result.WriteByte(f_eightbits[(eight_idx2+i)%(len(f_eightbits))])
				}
				return result.Bytes()
			}

			eight_byte = 0
		}

		if six_bit >= 6 {
			six_bit = 0
			six_idx2++
		}
	}
}

func DecodeRSACipher(pswd *[]byte) {
	keycode := getRSAKeyCode(*pswd)
	pq := keycode.p * keycode.q
	pq_1 := (keycode.p - 1) * (keycode.q - 1)
	pow := 1
	r := keycode.e
	selectTable_idx := 0

	for {
		n := pow*pq_1 + 1
		if n%r == 0 {
			r = n / r
			break
		}
		pow++
	}

	rsa_keysave := int((*pswd)[PASSWORD_RSA_BITSAVE_IDX])
	if DEBUG {
		fmt.Println("RSA keysave: ", rsa_keysave)
	}
	for i := 0; i < 8; i++ {
		b := int(uint8((*pswd)[keycode.selectTbl[selectTable_idx]])) | (((rsa_keysave >> i) & 1) << 8)
		decrypted := int(b)

		for d := r - 1; d > 0; d-- {
			decrypted = (decrypted * b) % pq
		}

		(*pswd)[keycode.selectTbl[selectTable_idx]] = byte(uint8(decrypted))
		selectTable_idx++
		if selectTable_idx >= len(keycode.selectTbl) {
			selectTable_idx = 0
		}
	}
}

func DecodedToPasswordStruct(data *[]byte) *Password {

	password_obj := Password{}

	password_obj.Checksum = uint32((*data)[0])

	password_obj.String0 = (*data)[2 : 2+PARAM_STRING_SIZE]
	password_obj.String1 = (*data)[10 : 10+PARAM_STRING_SIZE]

	password_obj.ItemId = uint16((*data)[18])<<8 + uint16((*data)[19])

	password_obj.Type = int((*data)[0]>>5) & 7

	if password_obj.Type == Popular || password_obj.Type == CardE {
		password_obj.HitRateIndex = int32(((*data)[0] >> 1) & 3)
		password_obj.NPCType = int32((*data)[0] & 1)
		password_obj.NPCCode = int32((*data)[1])
	} else if password_obj.Type == Magazine {
		password_obj.HitRateIndex = int32(((*data)[0]>>1)&3) | int32(((*data)[0]&1)<<2)
		password_obj.NPCType = int32((*data)[0] & 1)
		password_obj.NPCCode = int32((*data)[1])
	} else {
		password_obj.HitRateIndex = int32(((*data)[0] >> 1) & 3)
		password_obj.NPCType = int32((*data)[0] & 1)
		password_obj.NPCCode = int32((*data)[1])
	}

	encoded := makePasscode(password_obj.Type, int(password_obj.HitRateIndex), int(password_obj.NPCType), int(password_obj.NPCCode), password_obj.String0, password_obj.String1, int(password_obj.ItemId))
	if uint8(encoded[0]) == uint8(password_obj.Checksum) {
		password_obj.ChksmOK = true
	} else {
		if DEBUG {
			fmt.Println("Checksum mismatch:", uint8(encoded[0]))
		}
	}

	return &password_obj
}

func DecodeBitCode(pswd *[]byte) {
	bit_mix_code := int32((*pswd)[PASSWORD_BITMIXKEY_IDX] & 0xF)

	if DEBUG {
		fmt.Println("Bit mix code: ", bit_mix_code)
	}

	if bit_mix_code > 12 {
		*pswd = BitShift(*pswd, int(-bit_mix_code*3))
		if DEBUG {
			fmt.Println("substep: Bit Shift:", *pswd)
		}
		*pswd = BitReverse(*pswd)
		if DEBUG {
			fmt.Println("substep: Bit Reverse:", *pswd)
		}
		*pswd = BitArrangeReverse(*pswd)
		if DEBUG {
			fmt.Println("substep: Bit Arrange Reverse:", *pswd)
		}
	} else if bit_mix_code > 8 {
		*pswd = BitShift(*pswd, int(bit_mix_code*5))
		if DEBUG {
			fmt.Println("substep: Bit Shift:", *pswd)
		}
		*pswd = BitArrangeReverse(*pswd)
		if DEBUG {
			fmt.Println("substep: Bit Arrange Reverse:", *pswd)
		}
	} else if bit_mix_code > 4 {
		*pswd = BitReverse(*pswd)
		if DEBUG {
			fmt.Println("substep: Bit Reverse:", *pswd)
		}
		*pswd = BitShift(*pswd, int(bit_mix_code*5))
		if DEBUG {
			fmt.Println("substep: Bit Shift:", *pswd)
		}
	} else {
		*pswd = BitArrangeReverse(*pswd)
		if DEBUG {
			fmt.Println("substep: Bit Arrange Reverse:", *pswd)
		}
		*pswd = BitShift(*pswd, int(-bit_mix_code*3))
		if DEBUG {
			fmt.Println("substep: Bit Shift:", *pswd)
		}
	}
}

func DecodeSubstitutionCipher(pswd *[]byte) {
	idx := 0
	for i := 0; i < PASSWORD_DATA_SIZE; i++ {
		for j := 0; j < 256; j++ {
			if changeCodeTbl[j] == (*pswd)[idx] {
				(*pswd)[idx] = byte(j)
				idx++
				break
			}
		}
	}
}

func bZeroSlice(slice *[]byte, len int) {
	for i := 0; i < len; i++ {
		(*slice)[i] = 0
	}
}

func CShiftSliceStart(array *[]byte, amount int) []byte {
	shifted := &bytes.Buffer{}
	if DEBUG_CPtrShift {
		fmt.Println("CPtrShift:", amount)
		fmt.Println("CPtrShift:", *array)
	}
	if amount > 0 {
		for i := 0; i < (len(*array) - amount); i++ {
			pos := (amount + i) % (len(*array))
			shifted.WriteByte((*array)[pos])
		}
		for i := 0; i < amount; i++ {
			pos := i % (len(*array))
			shifted.WriteByte((*array)[pos])
		}
		if DEBUG_CPtrShift {
			fmt.Println("CPtrShift:", shifted.Bytes())
		}
		return shifted.Bytes()
	}
	if amount < 0 {
		for i := 0; i < -amount; i++ {
			pos := len(*array) + amount + i
			shifted.WriteByte((*array)[pos])
		}
		for i := 0; i < (len(*array) + amount); i++ {
			pos := i
			shifted.WriteByte((*array)[pos])
		}
		if DEBUG_CPtrShift {
			fmt.Println("CPtrShift:", shifted.Bytes())
		}
		return shifted.Bytes()
	}

	return *array
}

// TODO: seems good
func DecodeBitShuffle(pswd *[]byte, stage int) {
	workBuf := make([]byte, PASSWORD_DATA_SIZE-1)
	buf := make([]byte, PASSWORD_DATA_SIZE-1)
	keyIdx := 0
	test2 := 0
	var test []byte
	count := 0
	selectTable := 0

	count = PASSWORD_DATA_SIZE - 1
	keyIdx = 2

	if stage == 0 {
		keyIdx = 13
		count = PASSWORD_DATA_SIZE - 2
	}

	copy(workBuf[0:keyIdx], (*pswd)[0:keyIdx])
	test = CShiftSliceStart(pswd, keyIdx+1) // ?
	test2 = (PASSWORD_DATA_SIZE - 1) - keyIdx
	workBuf = CShiftSliceStart(&workBuf, keyIdx)
	copy(workBuf[:test2], test[:test2])
	bZeroSlice(&buf, count)
	workBuf = CShiftSliceStart(&workBuf, -keyIdx)

	selectTable = int((*pswd)[keyIdx] & 3)

	for i := 0; i < count; i++ {
		for bit := 0; bit < 8; bit++ {
			dstIdx := i + selectIdxTable[selectTable][bit]

			if dstIdx >= count {
				dstIdx -= count
			}

			buf[i] |= uint8(((workBuf[dstIdx] >> bit) & 1) << bit)
		}
	}

	copy((*pswd)[0:keyIdx], buf[0:keyIdx])
	buf = CShiftSliceStart(&buf, keyIdx)
	*pswd = CShiftSliceStart(pswd, keyIdx+1)
	if test2 > len(*pswd) || test2 > len(buf) {
		fmt.Println("undefined behavior")
	}
	copy((*pswd)[0:test2], buf[0:test2]) // TODO: Undefined behavior
	*pswd = CShiftSliceStart(pswd, -(keyIdx + 1))
}

func RuneToAscii(char int32) uint8 {
	switch char {
	case 'a':
		return CHAR_a
	case 'b':
		return CHAR_b
	case 'c':
		return CHAR_c
	case 'd':
		return CHAR_d
	case 'e':
		return CHAR_e
	case 'f':
		return CHAR_f
	case 'g':
		return CHAR_g
	case 'h':
		return CHAR_h
	case 'i':
		return CHAR_i
	case 'j':
		return CHAR_j
	case 'k':
		return CHAR_k
	case 'l':
		return CHAR_l
	case 'm':
		return CHAR_m
	case 'n':
		return CHAR_n
	case 'o':
		return CHAR_o
	case 'p':
		return CHAR_p
	case 'q':
		return CHAR_q
	case 'r':
		return CHAR_r
	case 's':
		return CHAR_s
	case 't':
		return CHAR_t
	case 'u':
		return CHAR_u
	case 'v':
		return CHAR_v
	case 'w':
		return CHAR_w
	case 'x':
		return CHAR_x
	case 'y':
		return CHAR_y
	case 'z':
		return CHAR_z
	case 'A':
		return CHAR_A
	case 'B':
		return CHAR_B
	case 'C':
		return CHAR_C
	case 'D':
		return CHAR_D
	case 'E':
		return CHAR_E
	case 'F':
		return CHAR_F
	case 'G':
		return CHAR_G
	case 'H':
		return CHAR_H
	case 'I':
		return CHAR_I
	case 'J':
		return CHAR_J
	case 'K':
		return CHAR_K
	case 'L':
		return CHAR_L
	case 'M':
		return CHAR_M
	case 'N':
		return CHAR_N
	case 'O':
		return CHAR_O
	case 'P':
		return CHAR_P
	case 'Q':
		return CHAR_Q
	case 'R':
		return CHAR_R
	case 'S':
		return CHAR_S
	case 'T':
		return CHAR_T
	case 'U':
		return CHAR_U
	case 'V':
		return CHAR_V
	case 'W':
		return CHAR_W
	case 'X':
		return CHAR_X
	case 'Y':
		return CHAR_Y
	case 'Z':
		return CHAR_Z
	case '!':
		return CHAR_EXCLAMATION
	case '#':
		return CHAR_HASHTAG
	case '@':
		return CHAR_AT_SIGN
	case '&':
		return CHAR_AMPERSAND
	case '(':
		return CHAR_OPEN_PARENTHESIS
	case ')':
		return CHAR_CLOSE_PARENTHESIS
	case '%':
		return CHAR_PERCENT
	case '0':
		return CHAR_ZERO
	case '1':
		return CHAR_ONE
	case '2':
		return CHAR_TWO
	case '3':
		return CHAR_THREE
	case '4':
		return CHAR_FOUR
	case '5':
		return CHAR_FIVE
	case '6':
		return CHAR_SIX
	case '7':
		return CHAR_SEVEN
	case '8':
		return CHAR_EIGHT
	case '9':
		return CHAR_NINE
	default:
		return 0xFF
	}
}

func ToAC_CodeOnly_ASCII(pswd string) []byte {
	workBuf := &bytes.Buffer{}
	for _, char := range pswd {
		asciiCode := RuneToAscii(char)
		if asciiCode == 0xFF {
			return nil
		}
		workBuf.WriteByte(asciiCode)
	}
	return workBuf.Bytes()
}

func AdjustLetter(pswd *[]byte) {
	for i := 0; i < PASSWORD_STR_SIZE; i++ {
		switch (*pswd)[i] {
		case CHAR_ZERO:
			(*pswd)[i] = CHAR_O
			break
		case CHAR_ONE:
			(*pswd)[i] = CHAR_l
			break
		default:
			break
		}
	}
}

func ChangePasswordFontCodeSub(c byte) uint8 {
	var res uint8 = 0xFF

	for i := 0; i < 64; i++ {
		if int(uint8(c)) == usable2fontnum[i] {
			res = uint8(i)
			//if DEBUG {
			//	fmt.Println(res)
			//}
			// break
		}
	}

	return uint8(res)
}

const (
	TRUE = iota
	FALSE
)

func ChangePasswordFontCode(pswd *[]byte) uint8 {
	changedPswd := make([]byte, PASSWORD_STR_SIZE)
	var result uint8 = TRUE

	for i := 0; i < PASSWORD_STR_SIZE; i++ {
		var newCode uint8 = ChangePasswordFontCodeSub((*pswd)[i])
		if newCode == 0xFF {
			result = FALSE
			break
		}
		changedPswd[i] = newCode
	}

	if result == TRUE {
		copy(*pswd, changedPswd)
	}

	return result
}

func DecodeCode(crypticPassword *[]byte) (*[]byte, bool) {
	workBuf := make([]byte, PASSWORD_STR_SIZE)
	result := false

	if DEBUG {
		fmt.Println(*crypticPassword)
	}

	copy(workBuf, *crypticPassword)
	if DEBUG {
		fmt.Println("AdjustLetter:", workBuf)
	}
	AdjustLetter(&workBuf)

	if DEBUG {
		fmt.Println("ChangeFontCode:", workBuf)
	}
	if ChangePasswordFontCode(&workBuf) == TRUE {
		if DEBUG {
			fmt.Println("Change8BitsCode:", workBuf)
		}
		workBuf = Change8BitsCode(workBuf) // perfect
		if DEBUG {
			fmt.Println("TranspositionCipher:", workBuf)
		}
		workBuf = TranspositionCipher(workBuf, true, 1)
		if DEBUG {
			fmt.Println("DecodeBitShuffle:", workBuf)
		}
		DecodeBitShuffle(&workBuf, 1)
		if DEBUG {
			fmt.Println("DecodeBitCode:", workBuf)
		}
		DecodeBitCode(&workBuf)
		if DEBUG {
			fmt.Println("DecodeRSACipher:", workBuf)
		}
		DecodeRSACipher(&workBuf)
		if DEBUG {
			fmt.Println("DecodeBitShuffle:", workBuf)
		}
		DecodeBitShuffle(&workBuf, 0)
		if DEBUG {
			fmt.Println("TranspositionCipher:", workBuf)
		}
		workBuf = TranspositionCipher(workBuf, false, 0)
		if DEBUG {
			fmt.Println("DecodeSubstitutionCipher:", workBuf)
		}
		DecodeSubstitutionCipher(&workBuf)

		result = true
	}

	return &workBuf, result
}
