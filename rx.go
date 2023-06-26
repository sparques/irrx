package irrx

import (
	. "machine"
	"time"
)

type RxDevice struct {
	pin Pin

	pulseCount int
	lastPulse  time.Time
	lastHigh   time.Duration

	stateMachine RxStateMachine
}

type RxStateMachine interface {
	HandleOnOff(on, off time.Duration)
}

func NewRxDevice(pin Pin, rsm RxStateMachine) *RxDevice {
	// pull-up? I think the IR module has a pull up already
	pin.Configure(PinConfig{Mode: PinInput})

	return &RxDevice{
		pin:          pin,
		stateMachine: rsm,
	}
}

func (d *RxDevice) interruptHandler(interruptPin Pin) {
	ptime := time.Now()
	if interruptPin.Get() {
		//pin high
		d.stateMachine.HandleOnOff(d.lastHigh, time.Since(d.lastPulse))
	} else {
		d.lastHigh = time.Since(ptime)
	}
	d.lastPulse = ptime
}

func (d *RxDevice) Start() {
	d.pin.SetInterrupt(PinFalling|PinRising, d.interruptHandler)
}

func (d *RxDevice) Stop() {
	d.pin.SetInterrupt(PinFalling, nil)
}
