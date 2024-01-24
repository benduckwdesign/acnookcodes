package acnookcodes

type RSAKeyInfo struct {
	p         int
	q         int
	e         int
	selectTbl []int
}

type Password struct {
	Checksum     uint32
	Type         int
	ItemId       uint16
	NPCType      int32
	NPCCode      int32
	HitRateIndex int32
	String0      []byte
	String1      []byte
	ChksmOK      bool
}
