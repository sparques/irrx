package irrx

import (
	"time"
)

// Cheapo implements an RX statemachine for an cheap, uknown brand IR remote
type Cheapo struct {
	CmdHandler func(uint32)

	buf      uint32
	bitcount int
}

func NewCheapo(cmdHandler func(uint32)) *Cheapo {
	return &Cheapo{CmdHandler: cmdHandler}
}

func (c *Cheapo) HandleOnOff(on, off time.Duration) {
	switch {
	case on > 7*time.Millisecond: // 9ms start of frame
		c.buf = 0
		c.bitcount = 0
		return
	case off > time.Millisecond:
		// one
		c.buf |= 1 << c.bitcount
		fallthrough
	case off < time.Millisecond:
		// zero
		c.bitcount++
	}

	if c.bitcount < 10 {
		return
	}
	c.CmdHandler(c.buf)
}
