package main

import (
	"goboy/internal/cartridge"
	"os"
)

func main() {
	dat, _ := os.ReadFile("./docs/test.gb")
	cartridge.PrintHeader(dat)
}
