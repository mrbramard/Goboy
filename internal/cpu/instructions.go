package cpu

import (
	"fmt"
	"goboy/internal/memory"
)

var Instructions = map[byte]func(){
	0x00: nop,
	0x01: ld_bc_n16,
	0x02: func() {},
	0x03: func() {},
	0x04: func() {},
	0x05: func() {},
	0x06: func() {},
	0x07: func() {},
	0x08: func() {},
	0x09: func() {},
	0x0A: func() {},
	0x0B: dec_bc,
	0x0C: func() {},
	0x0D: func() {},
	0x0E: func() {},
	0x0F: func() {},
	0x10: func() {},
	0x11: ld_de_n16,
	0x12: func() {},
	0x13: func() {},
	0x14: func() {},
	0x15: func() {},
	0x16: func() {},
	0x17: func() {},
	0x18: jr_n8,
	0x19: func() {},
	0x1A: func() {},
	0x1B: func() {},
	0x1C: func() {},
	0x1D: func() {},
	0x1E: func() {},
	0x1F: func() {},
	0x20: jr_nz_e8,
	0x21: func() {},
	0x22: func() {},
	0x23: func() {},
	0x24: func() {},
	0x25: func() {},
	0x26: func() {},
	0x27: func() {},
	0x28: func() {},
	0x29: func() {},
	0x2A: func() {},
	0x2B: func() {},
	0x2C: func() {},
	0x2D: func() {},
	0x2E: func() {},
	0x2F: func() {},
	0x30: func() {},
	0x31: ld_sp_n16,
	0x32: func() {},
	0x33: func() {},
	0x34: func() {},
	0x35: func() {},
	0x36: func() {},
	0x37: func() {},
	0x38: func() {},
	0x39: func() {},
	0x3A: func() {},
	0x3B: func() {},
	0x3C: func() {},
	0x3D: func() {},
	0x3E: ld_a_n8,
	0x3F: func() {},
	0x40: func() { ldXInY(B, B) },
	0x41: func() { ldXInY(C, B) },
	0x42: func() { ldXInY(D, B) },
	0x43: func() { ldXInY(E, B) },
	0x44: func() { ldXInY(H, B) },
	0x45: func() { ldXInY(L, B) },
	0x46: func() { ldHLinX(B) },
	0x47: func() { ldXInY(A, B) },
	0x48: func() { ldXInY(B, C) },
	0x49: func() { ldXInY(C, C) },
	0x4A: func() { ldXInY(D, C) },
	0x4B: func() { ldXInY(E, C) },
	0x4C: func() { ldXInY(H, C) },
	0x4D: func() { ldXInY(L, C) },
	0x4E: func() { ldHLinX(C) },
	0x4F: func() { ldXInY(A, C) },
	0x50: func() { ldXInY(B, D) },
	0x51: func() { ldXInY(C, D) },
	0x52: func() { ldXInY(D, D) },
	0x53: func() { ldXInY(E, D) },
	0x54: func() { ldXInY(H, D) },
	0x55: func() { ldXInY(L, D) },
	0x56: func() { ldHLinX(D) },
	0x57: func() { ldXInY(A, D) },
	0x58: func() { ldXInY(B, E) },
	0x59: func() { ldXInY(C, E) },
	0x5A: func() { ldXInY(D, E) },
	0x5B: func() { ldXInY(E, E) },
	0x5C: func() { ldXInY(H, E) },
	0x5D: func() { ldXInY(L, E) },
	0x5E: func() { ldHLinX(E) },
	0x5F: func() { ldXInY(A, E) },
	0x60: func() { ldXInY(B, H) },
	0x61: func() { ldXInY(C, H) },
	0x62: func() { ldXInY(D, H) },
	0x63: func() { ldXInY(E, H) },
	0x64: func() { ldXInY(H, H) },
	0x65: func() { ldXInY(L, H) },
	0x66: func() { ldHLinX(H) },
	0x67: func() { ldXInY(A, H) },
	0x68: func() { ldXInY(B, L) },
	0x69: func() { ldXInY(C, L) },
	0x6A: func() { ldXInY(D, L) },
	0x6B: func() { ldXInY(E, L) },
	0x6C: func() { ldXInY(H, L) },
	0x6D: func() { ldXInY(L, L) },
	0x6E: func() { ldHLinX(L) },
	0x6F: func() { ldXInY(A, L) },
	0x70: func() { ldXInHL(B) },
	0x71: func() { ldXInHL(C) },
	0x72: func() { ldXInHL(D) },
	0x73: func() { ldXInHL(E) },
	0x74: func() { ldXInHL(H) },
	0x75: func() { ldXInHL(L) },
	0x76: func() { ldHLinX(B) },
	0x77: func() { ldXInHL(A) },
	0x78: func() { ldXInY(B, A) },
	0x79: func() { ldXInY(C, A) },
	0x7A: func() { ldXInY(D, A) },
	0x7B: func() { ldXInY(E, A) },
	0x7C: func() { ldXInY(H, A) },
	0x7D: func() { ldXInY(L, A) },
	0x7E: func() { ldHLinX(B) },
	0x7F: func() { ldXInY(A, A) },
	0x80: func() { addToA(B) },
	0x81: func() { addToA(C) },
	0x82: func() { addToA(D) },
	0x83: func() { addToA(E) },
	0x84: func() { addToA(H) },
	0x85: func() { addToA(L) },
	0x86: func() { addToA(memory.ReadROMAtPos(uint16(H) + uint16(L)<<8)) },
	0x87: func() { addToA(A) },
	0x88: func() { adcToA(B) },
	0x89: func() { adcToA(C) },
	0x8A: func() { adcToA(D) },
	0x8B: func() { adcToA(E) },
	0x8C: func() { adcToA(H) },
	0x8D: func() { adcToA(L) },
	0x8E: func() { adcToA(memory.ReadROMAtPos(uint16(H) + uint16(L)<<8)) },
	0x8F: func() { adcToA(A) },
	0x90: func() { subToA(B) },
	0x91: func() { subToA(C) },
	0x92: func() { subToA(D) },
	0x93: func() { subToA(E) },
	0x94: func() { subToA(H) },
	0x95: func() { subToA(L) },
	0x96: func() { subToA(memory.ReadROMAtPos(uint16(H) + uint16(L)<<8))},
	0x97: func() { subToA(A) },
	0x98: func() { sbcToA(B) },
	0x99: func() { sbcToA(C) },
	0x9A: func() { sbcToA(D) },
	0x9B: func() { sbcToA(E) },
	0x9C: func() { sbcToA(H) },
	0x9D: func() { sbcToA(L) },
	0x9E: func() { sbcToA(memory.ReadROMAtPos(uint16(H) + uint16(L)<<8)) },
	0x9F: func() { sbcToA(A) },
	0xA0: func() { andToA(B) },
	0xA1: func() { andToA(C) },
	0xA2: func() { andToA(D) },
	0xA3: func() { andToA(E) },
	0xA4: func() { andToA(H) },
	0xA5: func() { andToA(L) },
	0xA6: func() { andToA(memory.ReadROMAtPos(uint16(H) + uint16(L)<<8)) },
	0xA7: func() { andToA(A) },
	0xA8: func() { xorToA(B) },
	0xA9: func() { xorToA(C) },
	0xAA: func() { xorToA(D) },
	0xAB: func() { xorToA(E) },
	0xAC: func() { xorToA(H) },
	0xAD: func() { xorToA(L) },
	0xAE: func() { xorToA(memory.ReadROMAtPos(uint16(H) + uint16(L)<<8)) },
	0xAF: func() { xorToA(A) },
	0xB0: func() { orToA(B) },
	0xB1: func() { orToA(C) },
	0xB2: func() { orToA(D) },
	0xB3: func() { orToA(E) },
	0xB4: func() { orToA(H) },
	0xB5: func() { orToA(L) },
	0xB6: func() { orToA(memory.ReadROMAtPos(uint16(H) + uint16(L)<<8)) },
	0xB7: func() { orToA(A) },
	0xB8: func() { cpToA(B) },
	0xB9: func() { cpToA(C) },
	0xBA: func() { cpToA(D) },
	0xBB: func() { cpToA(E) },
	0xBC: func() { cpToA(H) },
	0xBD: func() { cpToA(L) },
	0xBE: func() { cpToA(memory.ReadROMAtPos(uint16(H) + uint16(L)<<8)) },
	0xBF: func() { cpToA(A) },
	0xC0: func() {},
	0xC1: func() {},
	0xC2: func() {},
	0xC3: jp_n16,
	0xC4: func() {},
	0xC5: func() {},
	0xC6: func() {},
	0xC7: func() {},
	0xC8: func() {},
	0xC9: func() {},
	0xCA: func() {},
	0xCB: func() {},
	0xCC: func() {},
	0xCD: call_a16,
	0xCE: func() {},
	0xCF: func() {},
	0xD0: func() {},
	0xD1: func() {},
	0xD2: func() {},
	0xD4: func() {},
	0xD5: func() {},
	0xD6: func() {},
	0xD7: func() {},
	0xD8: func() {},
	0xD9: func() {},
	0xDA: func() {},
	0xDB: func() {},
	0xDC: func() {},
	0xDE: func() {},
	0xDF: func() {},
	0xE0: ldh_a8_a,
	0xE1: func() {},
	0xE2: func() {},
	0xE3: func() {},
	0xE4: func() {},
	0xE5: func() {},
	0xE6: func() {},
	0xE7: func() {},
	0xE8: func() {},
	0xE9: func() {},
	0xEA: ld_a16_a,
	0xEB: func() {},
	0xEC: func() {},
	0xED: func() {},
	0xEE: func() {},
	0xEF: func() {},
	0xF0: func() {},
	0xF1: func() {},
	0xF2: func() {},
	0xF3: di,
	0xF4: func() {},
	0xF5: func() {},
	0xF6: func() {},
	0xF7: func() {},
	0xF8: func() {},
	0xF9: func() {},
	0xFA: func() {},
	0xFB: func() {},
	0xFC: func() {},
	0xFD: func() {},
	0xFE: cp_n8,
	0xFF: func() {},
}

func addToA(reg byte) {
	A += reg
	if A == 0 {
		F |= FlagsMasks["Z"]
	}
	NextPos()
	PrintRegisters()
}

func adcToA(reg byte) {
	A += reg + (F | FlagsMasks["C"])
	if A == 0 {
		F |= FlagsMasks["Z"]
	}
	NextPos()
	PrintRegisters()
}

func subToA(reg byte) {
	A -= reg
	if A == 0 {
		F |= FlagsMasks["Z"]
	}
	F |= FlagsMasks["N"]
	NextPos()
	PrintRegisters()
}

func sbcToA(reg byte) {
	A -= reg + (F | FlagsMasks["C"])
	if A == 0 {
		F |= FlagsMasks["Z"]
	}
	NextPos()
	PrintRegisters()
}

func andToA(reg byte) {
	A &= reg
	if (A) == 0 {
		F |= FlagsMasks["Z"]
	}
}

func xorToA(reg byte) {
	A ^= reg
	if (A) == 0 {
		F |= FlagsMasks["Z"]
	}
	NextPos()
	PrintRegisters()
}

func orToA(reg byte) {
	A |= reg
	if A == 0 {
		F |= FlagsMasks["Z"]
	}
	NextPos()
	PrintRegisters()
}

func cpToA(reg byte) {
	if A-reg == 0 {
		F |= FlagsMasks["Z"]
	}
	F |= FlagsMasks["N"]
	NextPos()
	PrintRegisters()
}

func ldXInY(x byte, y byte) {
	y = x
	NextPos()
	PrintRegisters()
}

func ldHLinX(x byte) {
	x = memory.ReadROMAtPos(uint16(H) + uint16(L)<<8)
	NextPos()
	PrintRegisters()
}

func ldXInHL(x byte) {
	memory.WriteRAMAtPos(uint16(H)+uint16(L)<<8, x)
	NextPos()
	PrintRegisters()
}

func call_a16() {
	fmt.Println("CD : CALL a16")
	//PUSH PC TO STACK
	NextPos()
	NextPos()
	PrintRegisters()
}

func dec_bc() {
	fmt.Println("0B : DEC BC")

	if C == 0 {
		B--
		C = 0xFF
	} else {
		C--
	}

	NextPos()
	PrintRegisters()
}

func di() {
	fmt.Println("F3 : DI")
	//SET IME FLAG TO 1
	NextPos()
	PrintRegisters()
}

func nop() {
	fmt.Println("00 : NOP")
	NextPos()
	PrintRegisters()
}

func jp_n16() {
	fmt.Println("C3 : JP n16")
	PC = [2]byte{
		memory.ReadROMAtPos(GetPos() + 1),
		memory.ReadROMAtPos(GetPos() + 2),
	}
	PrintRegisters()
}

func jr_nc_e8() {
	fmt.Println("30 : JR NC, e8")
	if F&FlagsMasks["C"] == 0 {
		steps := int(memory.ReadROMAtPos(GetPos() + 1))
		for i := 0; i < steps; i++ {
			NextPos()
		}
	}
	NextPos()
	NextPos()
	PrintRegisters()
}

func jr_nz_e8() {
	fmt.Println("20 : JR NZ, e8")
	if F&FlagsMasks["Z"] == 0 {
		steps := int(memory.ReadROMAtPos(GetPos() + 1))
		for i := 0; i < steps; i++ {
			NextPos()
		}
	}
	NextPos()
	NextPos()
	PrintRegisters()
}

func jr_n8() {
	fmt.Println("18 : JR n8")
	steps := int(memory.ReadROMAtPos(GetPos() + 1))
	for i := 0; i < steps; i++ {
		NextPos()
	}
	NextPos()
	NextPos()
	PrintRegisters()
}

func cp_n8() {
	fmt.Println("FE : CP n8")
	if A-memory.ReadROMAtPos(GetPos()+1) == 0 {
		F |= FlagsMasks["Z"]
	}
	F |= FlagsMasks["N"]
	NextPos()
	PrintRegisters()
}

func ld_a_n8() {
	fmt.Println("3E : LD A, n8")
	A = memory.ReadROMAtPos(GetPos() + 1)
	NextPos()
	NextPos()
	PrintRegisters()
}

func ld_a16_a() {
	fmt.Println("EA : LD (a16), A")
	memoryPos := uint16(memory.ReadROMAtPos(GetPos()+1)) | uint16(memory.ReadROMAtPos(GetPos()+2))<<8
	memory.WriteRAMAtPos(memoryPos, A)
	NextPos()
	NextPos()
	NextPos()
	PrintRegisters()
}

func ld_bc_n16() {
	fmt.Println("01 : LD BC, n16")
	ld_n16(&B, &C)
	NextPos()
	NextPos()
	NextPos()
}

func ld_de_n16() {
	fmt.Println("11 : LD DE, n16")
	ld_n16(&D, &E)
	NextPos()
	NextPos()
	NextPos()
}

func ld_n16(high *byte, low *byte) {
	*low = memory.ReadROMAtPos(GetPos() + 1)
	*high = memory.ReadROMAtPos(GetPos() + 2)
	NextPos()
	NextPos()
	NextPos()
	PrintRegisters()
}

func ld_sp_n16() {
	fmt.Println("31 : LD SP, n16")
	ld_n16(&SP[1], &SP[0])
	NextPos()
	NextPos()
	NextPos()
	PrintRegisters()
}

func ldh_a8_a() {
	fmt.Println("E0 : LDH (a8), A")
	memoryPos := uint16(memory.ReadROMAtPos(GetPos() + 1))
	memory.WriteRAMAtPos(uint16(0xFF00)+memoryPos, A)
	NextPos()
	NextPos()
	PrintRegisters()
}
