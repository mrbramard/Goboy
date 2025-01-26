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
var PC = [2]byte{0, 0}

var Pos = 0x100

func PrintRegisters() {
  fmt.Println("AF BC DE HL SP PC")
  fmt.Println(fmt.Sprintf("%X%X %X%X %X%X %X%X %X %X", A, F, B, C, D, E, H, L, SP, PC))
}

func NextPos() {
  Pos++
}
