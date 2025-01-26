package cpu

import "fmt"

var Instructions = map[byte]func(){
  0x00: nop, 
  0xC3: jp_n16,
}

func nop() {
  fmt.Println("NOP")
}

func jp_n16() {
  fmt.Println("JP n16")
}


