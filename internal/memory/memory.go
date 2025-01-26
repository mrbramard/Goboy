package memory

var memory = make([]byte, 0xFFFF)
var ROMBank00 = memory[:0x4000]

func LoadGameROM(ROMData []byte) {
  ROMBank00 = ROMData
}

func ReadROMAtPos(pos int) byte {
  return ROMBank00[pos]
}
