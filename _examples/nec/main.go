package main

import (
	"device/arm"
	"fmt"
	. "machine"

	"github.com/sparques/irrx"
)

func dump(addr, cmd byte, repeat bool) {
	fmt.Printf("%X, %X  %v\r\n", addr, cmd, repeat)
}

func main() {
	nec := irrx.NewNEC(dump)
	irrx.NewRxDevice(GPIO14, nec).StartInverted()

	for {
		arm.Asm("wfi")
	}
}
