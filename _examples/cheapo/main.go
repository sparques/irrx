package main

import (
	"device/arm"
	"fmt"
	. "machine"

	"github.com/sparques/irrx"
)

func dump(cmd uint32) {
	fmt.Printf("%011b\r\n", cmd)
}

func main() {
	nec := irrx.NewCheapo(dump)
	irrx.NewRxDevice(GPIO14, nec).Start()

	for {
		arm.Asm("wfi")
	}
}
