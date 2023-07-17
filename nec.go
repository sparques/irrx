package irrx

import (
	"time"
)

// NEC implements an irrx.RxStateMachine that can decode NEC signals. If using the commonly available demodulating IR receivers
// that idle high and pull low when getting a signal, you must start the RxDevice using (*irrx.RxDevice).StartInverted() as
// this decoder needs off-on pairs rather than irrx's default on-off. The handler is still called with on as the first argument, though.
//
// The code to verify the received address and command is written, but commented out because my cheap IR remotes don't seem
// to transmit the inverted portion correctly. Unless there's some part of the NEC standard about bit-stuffing I'm missing.
type NEC struct {
	CmdHandler func(addr, cmd byte, repeat bool)

	buf      uint32
	bitcount int
}

func NewNEC(cmdHandler func(addr, cmd byte, repeat bool)) *NEC {
	return &NEC{CmdHandler: cmdHandler}
}

func (nec *NEC) HandleOnOff(on, off time.Duration) {
	switch {
	case on > 200*time.Millisecond:
		return
	case off > 3*time.Millisecond:
		//start of frame
		if on > 3*time.Millisecond {
			nec.buf = 0
			nec.bitcount = 0
		} else {
			// must be a repeat
			nec.CmdHandler(byte(nec.buf&0xFF), byte((nec.buf>>16)&0xFF), true)
		}
		return
	}

	if on > time.Millisecond {
		nec.buf |= 1 << nec.bitcount
	}
	nec.bitcount++

	if nec.bitcount != 32 {
		return
	}

	addr := nec.buf & 0xFF
	// does addr match its inverse?
	// if addr != ^((nec.buf >> 8) & 0xFF) {
	// return
	// }

	cmd := (nec.buf >> 16) & 0xFF
	// does cmd match its inverse?
	// if cmd != ^((nec.buf >> 24) & 0xFF) {
	// return
	// }
	nec.CmdHandler(byte(addr), byte(cmd), false)
}
