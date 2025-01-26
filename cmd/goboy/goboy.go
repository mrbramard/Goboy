package main

import (
	"fmt"
	"goboy/internal/cpu"
	"goboy/internal/memory"
	"os"
)

func Loop() {
	var opcode = memory.ReadROMAtPos(cpu.GetPos())
	fmt.Printf("Opcode: %02X\n", opcode)
	cpu.Instructions[opcode]()
}

func main() {
	var dat, _ = os.ReadFile("./docs/test.gb")
	memory.LoadGameROM(dat)

	for {
		Loop()
	}
}
