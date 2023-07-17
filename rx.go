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

func (d *RxDevice) invertedInterruptHandler(interruptPin Pin) {
	ptime := time.Now()
	if d.pin.Get() {
		//pin high
		d.lastHigh = time.Since(d.lastPulse)
	} else {
		d.stateMachine.HandleOnOff(time.Since(d.lastPulse), d.lastHigh)
	}
	d.lastPulse = ptime
}

func (d *RxDevice) interruptHandler(interruptPin Pin) {
	ptime := time.Now()
	if d.pin.Get() {
		//pin high
		d.stateMachine.HandleOnOff(d.lastHigh, time.Since(d.lastPulse))
	} else {
		d.lastHigh = time.Since(d.lastPulse)
	}
	d.lastPulse = ptime
}

// Start sets the interrupt handler and thus starts processing signals.
// Use Start if your RxStateMachine uses on-off pairs, e.g. Hexbug or PPM.
func (d *RxDevice) Start() {
	d.pin.SetInterrupt(PinFalling|PinRising, d.interruptHandler)
}

// StartInverted sets the interrupt handler and thus starts processing signals.
// Use Start if your RxStateMachine uses off-on pairs, e.g. NEC.
func (d *RxDevice) StartInverted() {
	d.pin.SetInterrupt(PinFalling|PinRising, d.invertedInterruptHandler)
}

func (d *RxDevice) Stop() {
	d.pin.SetInterrupt(PinFalling, nil)
}
