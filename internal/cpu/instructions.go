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
	0x40: func() {},
	0x41: func() {},
	0x42: func() {},
	0x43: func() {},
	0x44: func() {},
	0x45: func() {},
	0x46: func() {},
	0x47: func() {},
	0x48: func() {},
	0x49: func() {},
	0x4A: func() {},
	0x4B: func() {},
	0x4C: func() {},
	0x4D: func() {},
	0x4E: func() {},
	0x4F: func() {},
	0x50: func() {},
	0x51: func() {},
	0x52: func() {},
	0x53: func() {},
	0x54: func() {},
	0x55: func() {},
	0x56: func() {},
	0x57: func() {},
	0x58: func() {},
	0x59: func() {},
	0x5A: func() {},
	0x5B: func() {},
	0x5C: func() {},
	0x5D: func() {},
	0x5E: func() {},
	0x5F: func() {},
	0x60: func() {},
	0x61: func() {},
	0x62: func() {},
	0x63: func() {},
	0x64: func() {},
	0x65: func() {},
	0x66: func() {},
	0x67: func() {},
	0x68: func() {},
	0x69: func() {},
	0x6A: func() {},
	0x6B: func() {},
	0x6C: func() {},
	0x6D: func() {},
	0x6E: func() {},
	0x6F: func() {},
	0x70: func() {},
	0x71: func() {},
	0x72: func() {},
	0x73: func() {},
	0x74: func() {},
	0x75: func() {},
	0x76: func() {},
	0x77: func() {},
	0x78: ld_a_b,
	0x79: func() {},
	0x7A: func() {},
	0x7B: func() {},
	0x7C: func() {},
	0x7D: func() {},
	0x7E: func() {},
	0x7F: func() {},
	0x80: add_a_b,
	0x81: add_a_c,
	0x82: add_a_d,
	0x83: add_a_e,
	0x84: add_a_h,
	0x85: add_a_l,
	0x87: add_a_a,
	0x88: func() {},
	0x89: func() {},
	0x8A: func() {},
	0x8B: func() {},
	0x8C: func() {},
	0x8D: func() {},
	0x8E: func() {},
	0x8F: func() {},
	0x90: func() {},
	0x91: func() {},
	0x92: func() {},
	0x93: func() {},
	0x94: func() {},
	0x95: func() {},
	0x96: func() {},
	0x97: func() {},
	0x98: func() {},
	0x99: func() {},
	0x9A: func() {},
	0x9B: func() {},
	0x9C: func() {},
	0x9D: func() {},
	0x9E: func() {},
	0x9F: func() {},
	0xA0: func() {},
	0xA1: func() {},
	0xA2: func() {},
	0xA3: func() {},
	0xA4: func() {},
	0xA5: func() {},
	0xA6: func() {},
	0xA7: func() {},
	0xA8: func() {},
	0xA9: func() {},
	0xAA: func() {},
	0xAB: func() {},
	0xAC: func() {},
	0xAD: func() {},
	0xAE: func() {},
	0xB6: func() {},
	0xB8: func() {},
	0xB9: func() {},
	0xBA: func() {},
	0xBB: func() {},
	0xBC: func() {},
	0xBD: func() {},
	0xBE: func() {},
	0xBF: func() {},
	0xC0: func() {},
	0xC1: func() {},
	0xC2: func() {},
	0xC4: func() {},
	0xC5: func() {},
	0xC6: func() {},
	0xC7: func() {},
	0xC8: func() {},
	0xC9: func() {},
	0xCA: func() {},
	0xCB: func() {},
	0xCC: func() {},
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
	0xE1: func() {},
	0xE2: func() {},
	0xE3: func() {},
	0xE4: func() {},
	0xE5: func() {},
	0xE6: func() {},
	0xE7: func() {},
	0xE8: func() {},
	0xE9: func() {},
	0xEB: func() {},
	0xEC: func() {},
	0xED: func() {},
	0xEE: func() {},
	0xEF: func() {},
	0xF0: func() {},
	0xF1: func() {},
	0xF2: func() {},
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
	0xFF: func() {},
	0xAF: xor_a,
	0xB0: or_b,
	0xB1: or_c,
	0xB2: or_d,
	0xB3: or_e,
	0xB4: or_h,
	0xB5: or_l,
	0xB7: or_a,
	0xC3: jp_n16,
	0xCD: call_a16,
	0xE0: ldh_a8_a,
	0xEA: ld_a16_a,
	0xF3: di,
	0xFE: cp_n8,
}

func add_a_a() { add(A) }
func add_a_b() { add(B) }
func add_a_c() { add(C) }
func add_a_d() { add(D) }
func add_a_e() { add(E) }
func add_a_h() { add(H) }
func add_a_l() { add(L) }

func add(reg byte) {
	A += reg
	if A == 0 {
		F |= FlagsMasks["Z"]
	}
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

func or_a() { or(A) }
func or_b() { or(B) }
func or_c() { or(C) }
func or_d() { or(D) }
func or_e() { or(E) }
func or_h() { or(H) }
func or_l() { or(L) }

func or(reg byte) {
	A |= reg
	if A == 0 {
		F |= FlagsMasks["Z"]
	}
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

func ld_a_b() {
	fmt.Println("78 : LD A, B")
	A = B
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

func xor_a() {
	fmt.Println("AF : XOR A")
	A ^= A
	if (A) == 0 {
		F |= FlagsMasks["Z"]
	}
	NextPos()
	PrintRegisters()
}
