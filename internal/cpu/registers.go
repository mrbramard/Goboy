package cpu

import "fmt"

// Registers
var A = byte(0)
var F = byte(0)
var B = byte(0)
var C = byte(0)
var D = byte(0)
var E = byte(0)
var H = byte(0)
var L = byte(0)
var SP = [2]byte{0, 0}
var PC = [2]byte{0x00, 0x01}

// Flags masks
var FlagsMasks = map[string]byte{
	"Z": byte(1 << 7),
	"N": byte(1 << 6),
	"H": byte(1 << 5),
	"C": byte(1 << 4),
}

func PrintRegisters() {
	fmt.Println("AF   BC   DE   HL   SP   PC")
	fmt.Println(fmt.Sprintf("%02X%02X %02X%02X %02X%02X %02X%02X %02X %02X", A, F, B, C, D, E, H, L, SP, PC))
	fmt.Println()
}

func NextPos() {
	PC[0]++
	if PC[0] == 0 {
		PC[1]++
	}
}

func GetPos() uint16 {
	readablePos := uint16(PC[1])<<8 + uint16(PC[0])
	// fmt.Printf("pos uint16: %v | pos hex: %02X%02X\n", readablePos, PC[1], PC[0])
	return readablePos
}