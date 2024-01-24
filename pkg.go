package acnookcodes

import (
	"errors"
	"fmt"
)

type CoderOptFunc func(*CoderOpts)

type NookCodeType uint8

const (
	NookCodeType_Famicom NookCodeType = iota
	NookCodeType_Popular
	NookCodeType_CardE
	NookCodeType_Magazine
	NookCodeType_User
	NookCodeType_CardEMini
)

type SpecialVillagerId uint8

type NookCodeNPCType uint8

const (
	NPCType_Normal uint8 = iota
	NPCType_Special
)

const (
	SpecialVillager_Tom_Nook SpecialVillagerId = iota
	SpecialVillager_Wendell
	SpecialVillager_Saharah
	SpecialVillager_Gracie
	SpecialVillager_Joan
	SpecialVillager_Katrina
	SpecialVillager_Copper
	SpecialVillager_Jack
	SpecialVillager_Jingle
	SpecialVillager_Pete
	SpecialVillager_Pelly
	SpecialVillager_Phyllis
	SpecialVillager_Rover
	SpecialVillager_K_K
	SpecialVillager_Chip
	SpecialVillager_Booker
	SpecialVillager_Timmy
	SpecialVillager_Tommy
	SpecialVillager_Redd
	SpecialVillager_Resetti
	SpecialVillager_Gulliver
	SpecialVillager_Porter
	SpecialVillager_Blathers
	SpecialVillager_Kapp_n
	SpecialVillager_Mable
	SpecialVillager_Sable
	SpecialVillager_Tortimer
	SpecialVillager_Wisp
	SpecialVillager_Don
	SpecialVillager_Blanca
	SpecialVillager_Franklin
	SpecialVillager_Farley
)

type VillagerId uint8

const (
	Villager_Bob VillagerId = iota
	Villager_Olivia
	Villager_Mitzi
	Villager_Kiki
	Villager_Tangy
	Villager_Kabuki
	Villager_Tabby
	Villager_Monique
	Villager_Stinky
	Villager_Purrl
	Villager_Kitty
	Villager_Tom
	Villager_Rosie
	Villager_Nosegay
	Villager_Zoe
	Villager_Pango
	Villager_Cyrano
	Villager_Snooty
	Villager_Teddy
	Villager_Chow
	Villager_Dozer
	Villager_Nate
	Villager_Groucho
	Villager_Tutu
	Villager_Ursula
	Villager_Grizzly
	Villager_Pinky
	Villager_Jay
	Villager_Twiggy
	Villager_Anchovy
	Villager_Piper
	Villager_Admiral
	Villager_Otis
	Villager_Robin
	Villager_Midge
	Villager_Ace
	Villager_Twirp
	Villager_Chuck
	Villager_Stu
	Villager_Goose
	Villager_Betty
	Villager_Hector
	Villager_Egbert
	Villager_Ava
	Villager_Hank
	Villager_Leigh
	Villager_Rhoda
	Villager_Vladimir
	Villager_Murphy
	Villager_Cupcake
	Villager_Kody
	Villager_Maple
	Villager_Pudge
	Villager_Olive
	Villager_Poncho
	Villager_Bluebear
	Villager_Patty
	Villager_Petunia
	Villager_Bessie
	Villager_Belle
	Villager_Alfonso
	Villager_Boots
	Villager_Liz
	Villager_Biskit
	Villager_Goldie
	Villager_Daisy
	Villager_Lucky
	Villager_Portia
	Villager_Maddie
	Villager_Butch
	Villager_Bill
	Villager_Pompom
	Villager_Joey
	Villager_Scoot
	Villager_Derwin
	Villager_Freckles
	Villager_Paolo
	Villager_Dizzy
	Villager_Axel
	Villager_Emerald
	Villager_Tad
	Villager_Wart_Jr
	Villager_Cousteau
	Villager_Puddles
	Villager_Lily
	Villager_Jeremiah
	Villager_Huck
	Villager_Camofrog
	Villager_Ribbot
	Villager_Prince
	Villager_Jambette
	Villager_Billy
	Villager_Chevre
	Villager_Iggy
	Villager_Gruff
	Villager_Sven
	Villager_Velma
	Villager_Jane
	Villager_Cesar
	Villager_Louie
	Villager_Peewee
	Villager_Rollo
	Villager_Bubbles
	Villager_Bertha
	Villager_Elmer
	Villager_Winnie
	Villager_Savannah
	Villager_Ed
	Villager_Cleo
	Villager_Peaches
	Villager_Buck
	Villager_Carrie
	Villager_Mathilda
	Villager_Marcy
	Villager_Kitt
	Villager_Valise
	Villager_Astrid
	Villager_Sydney
	Villager_Gonzo
	Villager_Ozzie
	Villager_Yuka
	Villager_Huggy
	Villager_Rex
	Villager_Aziz
	Villager_Leopold
	Villager_Samson
	Villager_Penny
	Villager_Dora
	Villager_Chico
	Villager_Candi
	Villager_Rizzo
	Villager_Anicotti
	Villager_Limberg
	Villager_Carmen
	Villager_Octavian
	Villager_Sandy
	Villager_Sprocket
	Villager_Rio
	Villager_Queenie
	Villager_Apollo
	Villager_Buzz
	Villager_Quetzal
	Villager_Amelia
	Villager_Pierce
	Villager_Roald
	Villager_Aorora
	Villager_Hopper
	Villager_Cube
	Villager_Puck
	Villager_Gwen
	Villager_Friga
	Villager_Curly
	Villager_Truffles
	Villager_Spork
	Villager_Hugh
	Villager_Rasher
	Villager_Sue_E
	Villager_Hambo
	Villager_Lucy
	Villager_Cobb
	Villager_Boris
	Villager_Bunnie
	Villager_Doc
	Villager_Gaston
	Villager_Coco
	Villager_Gabi
	Villager_Dotty
	Villager_Genji
	Villager_Snake
	Villager_Claude
	Villager_Tank
	Villager_Spike
	Villager_Tiara
	Villager_Vesta
	Villager_Filbert
	Villager_Hazel
	Villager_Peanut
	Villager_Pecan
	Villager_Ricky
	Villager_Static
	Villager_Mint
	Villager_Nibbles
	Villager_Tybalt
	Villager_Rolf
	Villager_Bangle
	Villager_Lobo
	Villager_Freya
	Villager_Chief
	Villager_Weber
	Villager_Mallary
	Villager_Wolfgang
	Villager_Hornsby
	Villager_Oxford
	Villager_T_Bone
	Villager_Biff
	Villager_Opal
	Villager_Bones
	Villager_Bea
	Villager_Bitty
	Villager_Rocco
	Villager_Lulu
	Villager_Blaire
	Villager_Sally
	Villager_Ellie
	Villager_Eloise
	Villager_Alli
	Villager_Pippy
	Villager_Eunice
	Villager_Baabara
	Villager_Fang
	Villager_Deena
	Villager_Pate
	Villager_Stella
	Villager_Cashmere
	Villager_Woolio
	Villager_Cookie
	Villager_Maelle
	Villager_O_Hare
	Villager_Bliss
	Villager_Drift
	Villager_Bud
	Villager_Boomer
	Villager_Elina
	Villager_Flash
	Villager_Dobie
	Villager_Flossie
	Villager_Annalise
	Villager_Plucky
	Villager_Faith
	Villager_Yodel
	Villager_Rowan
	Villager_June
	Villager_Cheri
	Villager_Pigleg
	Villager_Ankha
	Villager_Punchy
)

type CoderOpts struct {
	nookCodeType       *NookCodeType
	stringData         [PARAM_STRING_SIZE * 2]byte
	stringDataChecksum int
	itemId             *uint16
	npcType            *uint8
	specialNpcCode     *SpecialVillagerId
	npcCode            *VillagerId
	hitRate            *uint8
	cardeHitRate       *uint8
}

func defaultCoderOpts() CoderOpts {
	return CoderOpts{
		nookCodeType:       nil,
		stringData:         [16]byte{CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE, CHAR_SPACE},
		stringDataChecksum: 512,
		itemId:             nil,
		npcType:            nil,
		npcCode:            nil,
		hitRate:            nil,
	}
}

const ()

type NookCoder struct {
	dataBuffer []byte
	*CoderOpts
}

type NookCode []byte

func (c *NookCode) String() string {
	if c == nil {
		return ""
	}
	return ConvertBytesToUnicodeString(*c)
}

// code, ok := coder.Encode()
var ErrorNoCodeTypeSet = errors.New("no nook code type set")
var ErrorNoItemIdSet = errors.New("no item id set")
var ErrorNoNpcTypeSet = errors.New("no npc type set")
var ErrorNoNpcCodeSet = errors.New("no npc code set")
var ErrorNoSpecialNpcCodeSet = errors.New("no special npc code set")
var ErrorNoHitRateSet = errors.New("proper hit rate for code type not set")

func (n *NookCoder) Encode() (NookCode, error) {

	if n.itemId == nil {
		return nil, ErrorNoItemIdSet
	}
	var checksum int
	checksum += n.stringDataChecksum
	checksum += int(*n.itemId) & 0xFF

	if n.nookCodeType == nil {
		return nil, ErrorNoCodeTypeSet
	}
	n.dataBuffer[0] = byte((int(*n.nookCodeType) & 7) << 5)

	if *n.nookCodeType == NookCodeType_Famicom {
		n.dataBuffer[0] |= byte(uint32(1) << 1) // hitRateIdx = 1
		if n.npcType == nil {
			n.dataBuffer[0] |= 0
			if n.npcCode == nil {
				n.dataBuffer[1] = 0xFF
			} else {
				n.dataBuffer[1] = byte(*n.npcCode)
			}
		} else {
			n.dataBuffer[0] |= *n.npcType & 1
			if *n.npcType == NPCType_Special {
				if n.specialNpcCode == nil {
					n.dataBuffer[1] = 0xFF
				} else {
					n.dataBuffer[1] = byte(*n.specialNpcCode)
				}
			} else {
				if n.npcCode == nil {
					n.dataBuffer[1] = 0xFF
				} else {
					n.dataBuffer[1] = byte(*n.npcCode)
				}
			}
		}
		checksum += 255
	}

	if *n.nookCodeType == NookCodeType_CardEMini {
		n.dataBuffer[0] |= byte(uint32(1) << 1) // hitRateIdx = 1
		if n.npcType == nil {
			n.dataBuffer[0] |= 0
			if n.npcCode == nil {
				n.dataBuffer[1] = 0xFF
			} else {
				n.dataBuffer[1] = byte(*n.npcCode)
			}
		} else {
			n.dataBuffer[0] |= *n.npcType & 1
			if *n.npcType == NPCType_Special {
				if n.specialNpcCode == nil {
					n.dataBuffer[1] = 0xFF
				} else {
					n.dataBuffer[1] = byte(*n.specialNpcCode)
				}
			} else {
				if n.npcCode == nil {
					n.dataBuffer[1] = 0xFF
				} else {
					n.dataBuffer[1] = byte(*n.npcCode)
				}
			}
		}
		checksum += 255
	}

	if *n.nookCodeType == NookCodeType_User {
		n.dataBuffer[0] |= byte(uint32(1) << 1) // hitRateIdx = 1
		if n.npcType == nil {
			n.dataBuffer[0] |= 0
			if n.npcCode == nil {
				n.dataBuffer[1] = 0xFF
			} else {
				n.dataBuffer[1] = byte(*n.npcCode)
			}
		} else {
			n.dataBuffer[0] |= *n.npcType & 1
			if *n.npcType == NPCType_Special {
				if n.specialNpcCode == nil {
					n.dataBuffer[1] = 0xFF
				} else {
					n.dataBuffer[1] = byte(*n.specialNpcCode)
				}
			} else {
				if n.npcCode == nil {
					n.dataBuffer[1] = 0xFF
				} else {
					n.dataBuffer[1] = byte(*n.npcCode)
				}
			}
		}
		checksum += 255
	}

	if *n.nookCodeType == NookCodeType_Popular {
		n.dataBuffer[0] |= byte(uint32(4) << 1) // hitRateIdx = 4
		if n.npcType == nil {
			n.dataBuffer[0] |= 0
			if n.npcCode == nil {
				return nil, ErrorNoNpcCodeSet
			}
			n.dataBuffer[1] = byte(*n.npcCode)
			checksum += int(*n.npcCode)
		} else {
			n.dataBuffer[0] |= *n.npcType & 1
			if *n.npcType == NPCType_Special {
				if n.specialNpcCode == nil {
					return nil, ErrorNoSpecialNpcCodeSet
				}
				n.dataBuffer[1] = byte(*n.specialNpcCode)
				checksum += int(*n.specialNpcCode)
			} else {
				if n.npcCode == nil {
					return nil, ErrorNoNpcCodeSet
				}
				n.dataBuffer[1] = byte(*n.npcCode)
				checksum += int(*n.npcCode)
			}
		}

	}

	if *n.nookCodeType == NookCodeType_Magazine {
		if n.hitRate == nil {
			return nil, ErrorNoHitRateSet
		}
		if *n.hitRate > 4 {
			return nil, ErrorNoHitRateSet
		}
		n.dataBuffer[0] |= (*n.hitRate >> 2) & 1 // npcType
		n.dataBuffer[0] |= (*n.hitRate & 3) << 1
		if n.npcType == nil {
			if n.npcCode == nil {
				n.dataBuffer[1] = 0xFF
			} else {
				n.dataBuffer[1] = byte(*n.npcCode)
			}
		} else {
			if *n.npcType == NPCType_Special {
				if n.specialNpcCode == nil {
					n.dataBuffer[1] = 0xFF
				} else {
					n.dataBuffer[1] = byte(*n.specialNpcCode)
				}
			} else {
				if n.npcCode == nil {
					n.dataBuffer[1] = 0xFF
				} else {
					n.dataBuffer[1] = byte(*n.npcCode)
				}
			}
		}
		checksum += 255
	}

	if *n.nookCodeType == NookCodeType_CardE {
		if n.cardeHitRate == nil {
			return nil, ErrorNoHitRateSet
		}
		if *n.cardeHitRate > 3 {
			return nil, ErrorNoHitRateSet
		}
		n.dataBuffer[0] |= *n.cardeHitRate << 1
		if n.npcType == nil {
			n.dataBuffer[0] |= 0
			if n.npcCode == nil {
				return nil, ErrorNoNpcCodeSet
			}
			n.dataBuffer[1] = byte(*n.npcCode)
			checksum += int(*n.npcCode)
		} else {
			n.dataBuffer[0] |= *n.npcType & 1
			if *n.npcType == NPCType_Special {
				if n.specialNpcCode == nil {
					return nil, ErrorNoSpecialNpcCodeSet
				}
				n.dataBuffer[1] = byte(*n.specialNpcCode)
				checksum += int(*n.specialNpcCode)
			} else {
				if n.npcCode == nil {
					return nil, ErrorNoNpcCodeSet
				}
				n.dataBuffer[1] = byte(*n.npcCode)
				checksum += int(*n.npcCode)
			}
		}
	}

	copy(n.dataBuffer[2:], n.stringData[:])

	n.dataBuffer[18] = uint8(int32(*n.itemId) >> 8)
	n.dataBuffer[19] = uint8(int32(*n.itemId))

	n.dataBuffer[0] |= byte((checksum & 3) << 3)

	EncodeSubstitutionCipher(n.dataBuffer)
	if DEBUG {
		fmt.Println("EncodeSubstitutionCipher:", n.dataBuffer)
	}
	TranspositionCipher(n.dataBuffer, true, 0)
	if DEBUG {
		fmt.Println("TranspositionCipher:", n.dataBuffer)
	}
	EncodeBitShuffle(n.dataBuffer, 0)
	if DEBUG {
		fmt.Println("EncodeBitShuffle:", n.dataBuffer)
	}
	EncodeChangeRSACipher(n.dataBuffer)
	if DEBUG {
		fmt.Println("EncodeChangeRSACipher:", n.dataBuffer)
	}
	EncodeBitMixCode(n.dataBuffer)
	if DEBUG {
		fmt.Println("EncodeBitMixCode:", n.dataBuffer)
	}
	EncodeBitShuffle(n.dataBuffer, 1)
	if DEBUG {
		fmt.Println("EncodeBitShuffle:", n.dataBuffer)
	}
	TranspositionCipher(n.dataBuffer, false, 1)
	if DEBUG {
		fmt.Println("TranspositionCipher:", n.dataBuffer)
	}
	copiedBuffer := Change6BitsCode(n.dataBuffer)
	if DEBUG {
		fmt.Println("Change6BitsCode:", copiedBuffer)
	}
	copiedBuffer = ChangeCommonFontCode_Uint8Array(copiedBuffer)
	if DEBUG {
		fmt.Println("ChangeCommonFontCode_Uint8Array:", copiedBuffer)
	}
	return copiedBuffer, nil
}

func NewNookCoder(opts ...CoderOptFunc) *NookCoder {
	o := defaultCoderOpts()
	for _, of := range opts {
		of(&o)
	}
	return &NookCoder{
		CoderOpts:  &o,
		dataBuffer: make([]byte, PASSWORD_DATA_SIZE),
	}
}

//func example() {
//	coder := NewNookCoder(withTownName("Joe Town"), withPlayerName("Joe Mama"), withVillager(Villager_Anchovy))
//}

// This function will properly format the magazine name,
// so you don't need to split and reverse it manually.
func withMagazineName(name string) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_Magazine
			opts.nookCodeType = &t
		}
		for i, v := range []byte(name) {
			// Reverse the halves of the magazine name.
			if i < PARAM_STRING_SIZE {
				opts.stringDataChecksum -= int(opts.stringData[i+PARAM_STRING_SIZE])
				opts.stringData[i+PARAM_STRING_SIZE] = v
				opts.stringDataChecksum += int(v)
			} else if i > PARAM_STRING_SIZE-1 {
				opts.stringDataChecksum -= int(opts.stringData[i-PARAM_STRING_SIZE])
				opts.stringData[i-PARAM_STRING_SIZE] = v
				opts.stringDataChecksum += int(v)
			}
		}
	}
}

// This function will properly format the magazine name,
// so you don't need to split and reverse it manually.
func withMagazine(name ...int) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_Magazine
			opts.nookCodeType = &t
		}
		for i, v := range name {
			// Reverse the halves of the magazine name.
			if i < PARAM_STRING_SIZE {
				opts.stringDataChecksum -= int(opts.stringData[i+PARAM_STRING_SIZE])
				opts.stringData[i+PARAM_STRING_SIZE] = uint8(v)
				opts.stringDataChecksum += v
			} else if i > PARAM_STRING_SIZE-1 {
				opts.stringDataChecksum -= int(opts.stringData[i-PARAM_STRING_SIZE])
				opts.stringData[i-PARAM_STRING_SIZE] = uint8(v)
				opts.stringDataChecksum += v
			}
		}
	}
}

// This function will properly format the minigame name,
// so you don't need to split and reverse it manually.
func withMinigameString(name string) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_CardEMini
			opts.nookCodeType = &t
		}
		for i, v := range []byte(name) {
			// Reverse the halves of the magazine name.
			if i < PARAM_STRING_SIZE {
				opts.stringDataChecksum -= int(opts.stringData[i+PARAM_STRING_SIZE])
				opts.stringData[i+PARAM_STRING_SIZE] = v
				opts.stringDataChecksum += int(v)
			} else if i > PARAM_STRING_SIZE-1 {
				opts.stringDataChecksum -= int(opts.stringData[i-PARAM_STRING_SIZE])
				opts.stringData[i-PARAM_STRING_SIZE] = v
				opts.stringDataChecksum += int(v)
			}
		}
	}
}

// This function will properly format the minigame name,
// so you don't need to split and reverse it manually.
func withMinigame(name ...int) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_CardEMini
			opts.nookCodeType = &t
		}
		for i, v := range name {
			// Reverse the halves of the magazine name.
			if i < PARAM_STRING_SIZE {
				opts.stringDataChecksum -= int(opts.stringData[i+PARAM_STRING_SIZE])
				opts.stringData[i+PARAM_STRING_SIZE] = uint8(v)
				opts.stringDataChecksum += v
			} else if i > PARAM_STRING_SIZE-1 {
				opts.stringDataChecksum -= int(opts.stringData[i-PARAM_STRING_SIZE])
				opts.stringData[i-PARAM_STRING_SIZE] = uint8(v)
				opts.stringDataChecksum += v
			}
		}
	}
}

func withPlayerString(name string) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_User
			opts.nookCodeType = &t
		}
		for i, v := range []byte(name) {
			if i < PARAM_STRING_SIZE {
				opts.stringDataChecksum -= int(opts.stringData[i])
				opts.stringData[i] = v
				opts.stringDataChecksum += int(v)
			}
		}
	}
}

func withPlayer(name ...int) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_User
			opts.nookCodeType = &t
		}
		for i, v := range name {
			if i < PARAM_STRING_SIZE {
				opts.stringDataChecksum -= int(opts.stringData[i])
				opts.stringData[i] = uint8(v)
				opts.stringDataChecksum += v
			}
		}
	}
}

func withTown(name ...int) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_User
			opts.nookCodeType = &t
		}
		for i, v := range name {
			if i < PARAM_STRING_SIZE {
				opts.stringDataChecksum -= int(opts.stringData[i+PARAM_STRING_SIZE])
				opts.stringData[i+PARAM_STRING_SIZE] = uint8(v)
				opts.stringDataChecksum += v
			}
		}
	}
}

func withTownString(name string) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_User
			opts.nookCodeType = &t
		}
		for i, v := range []byte(name) {
			if i < PARAM_STRING_SIZE {
				opts.stringDataChecksum -= int(opts.stringData[i+PARAM_STRING_SIZE])
				opts.stringData[i+PARAM_STRING_SIZE] = v
				opts.stringDataChecksum += int(v)
			}
		}
	}
}

func withSpecialVillager(villager SpecialVillagerId) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_Popular
			opts.nookCodeType = &t
		}
		var t uint8 = NPCType_Special
		opts.npcType = &t
		opts.specialNpcCode = &villager
	}
}

func withVillager(villager VillagerId) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_Popular
			opts.nookCodeType = &t
		}
		var t uint8 = NPCType_Normal
		opts.npcType = &t
		opts.npcCode = &villager
	}
}

func withOverrideVillager(value uint8) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_User
			opts.nookCodeType = &t
		}
		if opts.npcType == nil {
			var t uint8 = NPCType_Normal
			opts.npcType = &t
		} else {
			if *opts.npcType != NPCType_Normal {
				var t uint8 = NPCType_Normal
				opts.npcType = &t
			}
		}
		opts.npcCode = (*VillagerId)(&value)
	}
}

func withOverrideSpecialVillager(value uint8) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_User
			opts.nookCodeType = &t
		}
		if opts.npcType == nil {
			var t uint8 = NPCType_Special
			opts.npcType = &t
		} else {
			if *opts.npcType != NPCType_Special {
				var t uint8 = NPCType_Special
				opts.npcType = &t
			}
		}
		opts.specialNpcCode = (*SpecialVillagerId)(&value)
	}
}

// true = 1, false = 0
func withOverrideNPCType(value bool) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_User
			opts.nookCodeType = &t
		}
		var t uint8 = NPCType_Normal
		if value {
			t = NPCType_Special
		}
		opts.npcType = &t
	}
}

func withItemCode(item uint16) CoderOptFunc {
	return func(opts *CoderOpts) {
		opts.itemId = &item
	}
}

func withCodeType(codetype NookCodeType) CoderOptFunc {
	return func(opts *CoderOpts) {
		opts.nookCodeType = &codetype
	}
}

// 4 - 100%
// 3 - 0%
// 2 - 30%
// 1 - 60%
// 0 - 80%
func withHitRate(hitrate_level uint8) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_Magazine
			opts.nookCodeType = &t
		}
		opts.hitRate = &hitrate_level
	}
}

// 3 - 20%
// 2 - 40%
// 1 - 60%
// 0 - 80%
func withCardEHitRate(e_hitrate_level uint8) CoderOptFunc {
	return func(opts *CoderOpts) {
		if opts.nookCodeType == nil {
			t := NookCodeType_CardE
			opts.nookCodeType = &t
		}
		opts.cardeHitRate = &e_hitrate_level
	}
}
