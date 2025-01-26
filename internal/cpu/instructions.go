package cpu

import (
	"fmt"
	"goboy/internal/memory"
)

var Instructions = map[byte]func(){
	0x00: nop,
	0x01: ld_bc_n16,
	0x0B: dec_bc,
	0x11: ld_de_n16,
	0x18: jr_n8,
	0x20: jr_nz_e8,
	0x30: jr_nc_e8,
	0x31: ld_sp_n16,
	0x3E: ld_a_n8,
	0x78: ld_a_b,
	0x80: add_a_b,
	0x81: add_a_c,
	0x82: add_a_d,
	0x83: add_a_e,
	0x84: add_a_h,
	0x85: add_a_l,
	0x87: add_a_a,
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
