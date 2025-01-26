package memory

var memory = make([]byte, 0xFFFFFFFF)
var ROMBank00 = memory[:0x4000]

func LoadGameROM(ROMData []byte) {
	ROMBank00 = ROMData
}

func ReadROMAtPos(pos uint16) byte {
	return ROMBank00[pos]
}

func WriteRAMAtPos(pos uint16, value byte) {
	memory[pos] = value
}
