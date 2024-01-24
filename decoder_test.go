package acnookcodes

import (
	"bytes"
	"fmt"
	"testing"
)

func assertBytes(byt0 []byte, byt1 []byte) bool {
	return bytes.Equal(byt0, byt1)
}

func TestFishingRodDecode2(t *testing.T) {
	//code_bytes := ToAC_CodeOnly_ASCII("uVVeEe@sjJxJprBFrb2HOxxqzabM")
	code_bytes := ToAC_CodeOnly_ASCII("aO@KJh3aKQaQN&RhKqq3SEFG3AS4")
	decoded_code, ok := DecodeCode(&code_bytes)
	if !ok {
		t.Fatal("failed")
	}
	password := DecodedToPasswordStruct(decoded_code)
	t.Logf("%+v", password)
}

func TestFishingRodDecode(t *testing.T) {
	inputCodetype := 4
	inputStr0 := [8]byte{'!', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	inputStr1 := [8]byte{'!', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	inputHitrate := HITRATE_ONE_HUNDRED_100
	inputItemId := 0x2203
	inputNpcType := 1
	inputNpcCode := 255

	type Pair struct {
		score int
		code  string
	}

	var pair Pair
	best_score := -200
	best_code := ""
	var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, inputHitrate, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))
	pair = Pair{score: scoreSequentialCharacters(generatedCode), code: generatedCode}
	code_bytes := ToAC_CodeOnly_ASCII("aO@KJh3aKQaQN&RhKqq3SEFG3AS4")
	decoded_code, ok := DecodeCode(&code_bytes)
	if !ok {
		//failed
		t.Fatal("Failed")
	}
	password := DecodedToPasswordStruct(decoded_code)
	if password.ChksmOK {
		if pair.score > best_score {
			best_score = pair.score
			best_code = pair.code
			fmt.Println(best_code)
		}
	}

	//inputCodetype := Popular
	//inputStr0 := [8]byte{'!', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	//inputStr1 := [8]byte{'!', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
	////inputHitrate := HITRATE_ONE_HUNDRED_100
	//inputItemId := 0x2203
	//inputNpcType := int(NPCType_Special)
	//inputNpcCode := int(SpecialVillager_Tom_Nook)
	//
	//type Pair struct {
	//	score int
	//	code  string
	//}
	//
	//best_score := 0
	//best_code := "NO CODES"
	//
	//var pair Pair
	//for H := 0; H < 256; H++ {
	//	var generatedCode string = ConvertBytesToUnicodeString(MakePassword(inputCodetype, H, inputStr0[:], inputStr1[:], inputItemId, inputNpcType, inputNpcCode))
	//	pair = Pair{score: scoreSequentialCharacters(generatedCode), code: generatedCode}
	//	if pair.score > best_score {
	//		best_score = pair.score
	//		best_code = pair.code
	//	}
	//}
	//
	//fmt.Println(best_code)
}

func TestAppleTVCardEMiniDecode(t *testing.T) {
	var code = ToAC_CodeOnly_ASCII("2t6tmLp5c#Uio#zqfGjEGR4BpnVo")
	//178  255  65  66  43  43  43  43  43  43  47  47  201  47  47  47  47  47  18  248  173
	//var expected_decrypt = []byte{178, 255, 65, 66, 43, 43, 43, 43, 43, 43, 47, 47, 201, 47, 47, 47, 47, 47, 18, 248, 0}

	inputCodetype := CardEMini
	inputStr0 := [8]byte{CHAR_A, CHAR_B, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART, CHAR_SYMBOL_HEART}
	inputStr1 := [8]byte{CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_RABBIT, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE, CHAR_SYMBOL_MUSIC_NOTE}
	inputHitrate := HITRATE_ONE_HUNDRED_100
	inputItemId := 0x12F8
	inputNpcType := 0 // Set to 0 for CardEMini to pass. Unknown if this is intended behavior.
	inputNpcCode := 255
	var expected_decrypt = makePasscode(inputCodetype, inputHitrate, inputNpcType, inputNpcCode, inputStr0[:], inputStr1[:], inputItemId)

	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Log("Decrypted:")
		t.Logf("%v", *decryptedPassword)
	}

	if !assertBytes((*decryptedPassword)[:19], expected_decrypt[:19]) {
		t.Logf("%+v", DecodedToPasswordStruct(decryptedPassword))
		t.Logf("%+v", DecodedToPasswordStruct(&expected_decrypt))
		t.Fatalf(`Code mismatch, got %v expected %v`, *decryptedPassword, expected_decrypt)
	}
}

func TestDecodeFunnyCode1(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("GodIsTheGreatestBeingForever")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// Bit mix: 0
	// Prize: Dingloid (5808)
	// wide space, (, tab, key symbol, s, Y, I(two dots), acute Y,
	// wide space, G, P, circle, O(two dots), circle, wide space
	// Valid to Nook, Invalid to decoder

	correct := uint16(5811)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFunnyCode2(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("ThqPoliceDogIsCopperHeIsCool")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// Bit mix: 6
	// Prize: coral shirt (9435)
	// right whisker, bug symbol, y(two dots), double vertical lines, i, wide space, mail symbol, acute E, V, cross, acute Y, reverse acute i, D with a horizontal dash, tab
	// Valid to Nook

	correct := uint16(9435)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFunnyCode3(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("MupersmaspbnoSSupersmashbroS")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 9
	// Prize: ponderosa bonsai
	// double vertical lines, left whisker, acute E, semicolon, wide space, strange colon, wide space, cubed, percent, lowercase ae, wide space, k, f, long dash, wide space, sun symbol
	// Valid to Nook

	correct := uint16(5107) //5104
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFunnyCode4(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("NintendoGamecubearenumberone")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 15
	// Prize: regal bookcase
	// b, y(two dots), ), S, #, wide space, 3,8,a with n, c with comma underneath, wide space, up-left acute E, cloud symbol, space, upside down ?, Y
	// Valid to Nook

	correct := uint16(4499)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFunnyCode5(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("Toad&MushDoomsToad&Mushrooms")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 1
	// Prize: ranch tea table
	// squared, j, i, cloud symbol, wide space, w, wide space, E, T, tilde O, right whisker, I with n, bunny symbol, acute E, L, happy symbol :D
	// Valid to Nook

	correct := uint16(0x11A0)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFunnyCode6(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("WzatswrongwitHWhatswrongwitH")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 5
	// Prize: exotic shirt
	// E(two dots), downward acute o, wide space, downward acute o, wide space, !, SS, o with n, E, O, <<, logical not, downward A, acute U, l, double vertical bars
	// Valid to Nook

	correct := uint16(0x24E4)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeHarvestLampCode(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("LYE@rAKm@MZVpJRRerwI4c5vXVQs")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 0 PASS

	correct := uint16(0x32d0)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeHarvestSofaCode(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("GsoCkARQLtldP9Pu7HT8bFBe7hdM")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 0

	correct := uint16(0x32f4)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeHarvestClockCode(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("8@lCwCxXkwaxtpiCA&QIMd3@C@KE")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 0

	correct := uint16(0x32f0)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeBlueCornerCode(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("DY%T%sh&GBRQyPxnQSyHB8z8JPme")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 6

	correct := uint16(0x3344)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeOwlClockCode(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("SqLSwI7xBhDDytqZySKeAXb8QDCc")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	// bit mix: 6

	correct := uint16(0x1E70)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

// Decode fast codes

func TestDecodeFastCode1(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("aa22W%%%%%%%%%%%%%%%%%%%%%%%")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(0x24C9)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode2(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("QLFCsssssssssssssssssssssssT")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(0x24CF)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode3(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("3eNDCWWWWWWWWWWWWWWWWWWWWWWW")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(0x260C)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode4(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("QwBsBErrrrrrrrrrrrrrrrrrrrrr")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(0x2A05)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode5(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("44rUt5vvvvvvvvvvvvvvvvvvvvvv")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(0x2231)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode6(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("dmNffdSSSSSSSSSSSSSSSSSSSSSS")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(0x20D0)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode7(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("BeMMcAAaaaaaaaaaaaaaaaaaaaaa")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(5927)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode8(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("RfwBNPPPPPPPPPPPPPPPPPPPPPPP")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(5821)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode9(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("AwweB@@@@@@@@@@@@@@@@@@@@@@@")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(4963)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}

func TestDecodeFastCode10(t *testing.T) {

	var code = ToAC_CodeOnly_ASCII("KKLgEEEEEEEEEEeeeeeeeeeeeeee")
	decryptedPassword, ok := DecodeCode(&code)
	if !ok {
		t.Fatalf("Failed to decode password")
	} else {
		t.Logf("%v", *decryptedPassword)
	}

	correct := uint16(0x13C0)
	password := DecodedToPasswordStruct(decryptedPassword)
	t.Logf("%+v", password)
	if password.ItemId != correct {
		t.Fatal("Inaccurate decode, wrong item id, got", password.ItemId, "expected", correct)
	}

}
