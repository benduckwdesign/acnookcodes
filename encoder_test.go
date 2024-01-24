package acnookcodes

import (
	"testing"
)

func assert(str0 string, str1 string) bool {
	return str0 == str1
}

// Fastcode testing

// TODO: Possibly incorrect code
//func TestBugZapperMagazineFastcode(t *testing.T) {
//	expectedCode := "322444444444444444444444444c"
//
//	inputCodetype := 3
//	inputStr0 := [8]byte{32, 32, 81, 56, 86, 169, 32, 32}
//	inputStr1 := [8]byte{139, 119, 15, 5, 32, 32, 116, 118}
//	inputHitrate := 4
//	inputItemId := 12853
//	inputNpcType := 255
//	inputNpcCode := 255
//
//	var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))
//
//	if !assert(generatedCode, expectedCode) {
//		t.Fatalf(`Code mismatch, got %s expected %s`, generatedCode, expectedCode)
//	}
//}

// TODO: Possibly incorrect code
//func TestGrayTartanMagazineFastcode(t *testing.T) {
//	expectedCode := "QKSDELLLLLLLLLLLLLLLLLLLLLLL"
//
//	c := ASCII2ACBytes("QKSDELLLLLLLLLLLLLLLLLLLLLLL")
//	d, _ := DecodeCode(&c)
//	s := DecodedToPasswordStruct(d)
//	fmt.Printf("%+v\n", *s)
//
//	//inputCodetype := 3
//	//inputStr0 := [8]byte{43, 208, 83, 5, 32, 32, 32, 32}
//	//inputStr1 := [8]byte{144, 25, 8, 169, 163, 54, 91, 32}
//	//inputHitrate := 4
//	//inputItemId := 9424
//	//inputNpcType := 255
//	//inputNpcCode := 255
//	//
//	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))
//
//	coder := NewNookCoder(withMagazine(43, 208, 83, 5, 32, 32, 32, 32, 144, 25, 8, 169, 163, 54, 91, 32), withItemCode(9424), withHitRate(4), withOverrideSpecialVillager(75))
//
//	code, err := coder.Encode()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	generatedCode := code.String()
//
//	if !assert(generatedCode, expectedCode) {
//		t.Fatalf(`Code mismatch, got %s expected %s`, generatedCode, expectedCode)
//	}
//}

// Magazine codes (3)

func TestFroggyChairMagazineCode(t *testing.T) {
	expectedCode := "IAAAAA#Ys%hBgSadjOcczq&orqw8"

	//c := ASCII2ACBytes("IAAAAA#Ys%hBgSadjOcczq&orqw8")
	//d, ok := DecodeCode(&c)
	//if !ok {
	//	s := DecodedToPasswordStruct(d)
	//	fmt.Printf("%+v", *s)
	//}
	//inputCodetype := 3
	//inputStr0 := [8]byte{'p', 'p', 'p', 'n', 'X', 'B', 'c', 'V'}
	//inputStr1 := [8]byte{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'}
	//inputHitrate := 4
	//inputItemId := 0x10A4
	//inputNpcType := 255
	//inputNpcCode := 255

	coder := NewNookCoder(withMagazine(CHAR_p, CHAR_p, CHAR_p, CHAR_p, CHAR_p, CHAR_p, CHAR_p, CHAR_p, CHAR_p, CHAR_p, CHAR_p, CHAR_n, CHAR_X, CHAR_B, CHAR_c, CHAR_V), withItemCode(0x10A4), withHitRate(4), withOverrideSpecialVillager(255))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	generatedCode := code.String()

	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	if !assert(generatedCode, expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, generatedCode, expectedCode)
	}
}

// User codes (4)

func TestFroggyChairUserCode(t *testing.T) {
	expectedCode := "GNx52WeRmpVyXLF3AnVjXOQKaKO&"

	coder := NewNookCoder(withPlayerString("!"), withTownString("!"), withItemCode(0x10A4), withHitRate(4), withOverrideNPCType(true))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	if !assert(code.String(), expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, code, expectedCode)
	}
	
	//inputCodetype := 4
	//inputStr0 := [8]byte{'!', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	//inputStr1 := [8]byte{'!', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	//inputHitrate := 4
	//inputItemId := 0x10A4
	//inputNpcType := 255
	//inputNpcCode := 255
	//
	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	//fmt.Println(string([]byte{52, 103, 108, 54, 67, 77, 68, 79, 75, 105, 83, 119, 76, 117, 77, 104, 102, 122, 56, 110, 87, 57, 89, 115, 88, 86, 78, 67}))

}

// Famicom codes (0)

func TestWariosWoodsFamicomCode(t *testing.T) {
	expectedCode := "cj@Eh@8a3QZ2#AQS8rzMC7gHAeUD"

	//inputCodetype := 0
	//inputStr0 := [8]byte{'!', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	//inputStr1 := [8]byte{'!', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	//inputHitrate := 4
	//inputItemId := 0x1DE0
	//inputNpcType := 255
	//inputNpcCode := 255
	//
	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	coder := NewNookCoder(withCodeType(NookCodeType_Famicom), withItemCode(0x1DE0), withPlayerString("!"), withTownString("!"), withOverrideNPCType(true))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	generatedCode := code.String()

	if !assert(generatedCode, expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, generatedCode, expectedCode)
	}
}

func TestCluCluLandFamicomCode(t *testing.T) {
	expectedCode := "dVkjlLzGpDmHQEi9CjeMJUPKIB5v"

	coder := NewNookCoder(withCodeType(NookCodeType_Famicom), withPlayer(CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART), withTown(CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE), withItemCode(0x1DA8), withOverrideNPCType(true))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	//inputCodetype := 0
	//inputStr0 := [8]byte{CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART}
	//inputStr1 := [8]byte{CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE}
	//inputHitrate := 4
	//inputItemId := 0x1DA8
	//inputNpcType := 255
	//inputNpcCode := 255
	//
	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	generatedCode := code.String()

	if !assert(generatedCode, expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, generatedCode, expectedCode)
	}
}

func TestCluCluLandOtherFamicomCode(t *testing.T) {
	expectedCode := "#B@Ipdg@zfrA6hLJf9GmNDUenN6Q"

	coder := NewNookCoder(withCodeType(NookCodeType_Famicom), withPlayer(CHAR_B, CHAR_e, CHAR_n), withTown(CHAR_P, CHAR_a, CHAR_r, CHAR_i, CHAR_s), withItemCode(0x1DA8))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	//inputCodetype := 0
	//inputStr0 := [8]byte{CHAR_B, CHAR_e, CHAR_n, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE}
	//inputStr1 := [8]byte{CHAR_P, CHAR_a, CHAR_r, CHAR_i, CHAR_s, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE}
	//inputHitrate := 4
	//inputItemId := 0x1DA8
	//inputNpcType := 0 //must be zero
	//inputNpcCode := 255
	//
	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	if !assert(code.String(), expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, code.String(), expectedCode)
	}
}

// #B@Ipdg@zfrA6h famicom, Nookipedia
// LJf9GmNDUenN6Q

// #B@Ipd7@j2rA6h user code, JS
// LJfMMmN@UPnN6Q

// 4B@IpdnEz%9A6h famicom code, JS
// LQf9GmNDUenN6Q

// Popular Codes (1)

func TestKikiCabanaTablePopularCode(t *testing.T) {
	expectedCode := "b%waEznWZl%Mj#MBUaSwAZtQfjXn"

	//inputCodetype := Popular
	//inputStr0 := [8]byte{CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART}
	//inputStr1 := [8]byte{CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE}
	//inputHitrate := 4
	//inputItemId := 0x151C
	//inputNpcType := 0
	//inputNpcCode := 3

	// Set to Popular type automatically.
	coder := NewNookCoder(withVillager(Villager_Kiki), withPlayer(CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART), withTown(CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE), withItemCode(0x151C))

	// Kiki = 3

	// If NPC is special, or the code type is for CardE,
	// NpcType is set to 1. Otherwise, NpcType is set to zero.
	// It is usually set to 0xFF for other code types internally.

	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	if !assert(code.String(), expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, code.String(), expectedCode)
	}
}

func TestTokekeOnlyMePopularCode(t *testing.T) {
	expectedCode := "3ixu@x7M794xq3&56x7qlRyX5y9C"

	//inputCodetype := Popular
	//inputStr0 := [8]byte{CHAR_A, CHAR_B, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART}
	//inputStr1 := [8]byte{CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE}
	//inputHitrate := 4
	//inputItemId := 0x2A27
	//inputNpcType := 1
	//inputNpcCode := 13

	coder := NewNookCoder(withSpecialVillager(SpecialVillager_K_K), withPlayer(CHAR_A, CHAR_B, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART), withTown(CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE), withItemCode(0x2A27))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	// Tokeke = 13 (Special List)

	// If NPC is special, or the code type is for CardE,
	// NpcType is set to 1. Otherwise, NpcType is set to zero.
	// It is usually set to 0xFF for other code types internally.

	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	generatedCode := code.String()

	if !assert(generatedCode, expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, generatedCode, expectedCode)
	}
}

// CardE Code

func TestFarleyAppleTVCardECode(t *testing.T) {
	expectedCode := "wt6tmGPqBRe4oizqcXyDhRiP6lYo"

	//inputCodetype := CardE
	//inputStr0 := [8]byte{CHAR_A, CHAR_B, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART}
	//inputStr1 := [8]byte{CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE}
	//inputHitrate := CARDE_HITRATE_EIGHTY_80
	//inputItemId := 0x12F8
	//inputNpcType := 1
	//inputNpcCode := 31
	//
	//// Farley = 31 (Special List)
	//
	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	coder := NewNookCoder(withCardEHitRate(0), withSpecialVillager(SpecialVillager_Farley), withPlayer(CHAR_A, CHAR_B, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART), withTown(CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE), withItemCode(0x12F8))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	generatedCode := code.String()

	if !assert(generatedCode, expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, generatedCode, expectedCode)
	}
}

// CardEMini Code

func TestAppleTVCardEMiniCode(t *testing.T) {
	expectedCode := "2t6tmLp5c#Uio#zqfGjEGR4BpnVo"

	//inputCodetype := CardEMini
	//inputStr0 := [8]byte{CHAR_A, CHAR_B, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART}
	//inputStr1 := [8]byte{CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE}
	//inputHitrate := HITRATE_ONE_HUNDRED_100
	//inputItemId := 0x12F8
	//inputNpcType := 0 // Set to 0 for CardEMini to pass. Unknown if this is intended behavior.
	//inputNpcCode := 255
	//
	//var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))

	coder := NewNookCoder(withMinigame(CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_A, CHAR_B, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART), withItemCode(0x12F8))

	code, err := coder.Encode()
	if err != nil {
		t.Fatal(err)
	}

	generatedCode := code.String()

	if !assert(generatedCode, expectedCode) {
		t.Fatalf(`Code mismatch, got %s expected %s`, generatedCode, expectedCode)
	}
}
