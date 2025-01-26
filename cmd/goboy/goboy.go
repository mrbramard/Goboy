package main

import (
	"os"
  "fmt"
  "goboy/internal/cpu"
  "goboy/internal/memory"
)

func PrintState(dat []byte) {
  cpu.PrintRegisters()
  fmt.Println()
  fmt.Println(fmt.Sprintf("%X", dat[cpu.Pos]))
  cpu.NextPos()
  fmt.Println()
}

func PrintNextBytes() {
  fmt.Printf("%X %X %X %X %X\n",
    memory.ReadROMAtPos(cpu.Pos),
    memory.ReadROMAtPos(cpu.Pos+1),
    memory.ReadROMAtPos(cpu.Pos+2),
    memory.ReadROMAtPos(cpu.Pos+3),
    memory.ReadROMAtPos(cpu.Pos+4),
  )
}

func Loop() {
  cpu.PrintRegisters()
  PrintNextBytes()
  cpu.Instructions[memory.ReadROMAtPos(cpu.Pos)]()
  cpu.NextPos()
}

func main() {
  var dat, _ = os.ReadFile("./docs/test.gb")
  memory.LoadGameROM(dat)

  Loop()
  Loop()
}
