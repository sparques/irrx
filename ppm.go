/*
ppm.go - PPM for IR

Might work exactly the same as though connected to a regular PPM radio receiver.

Note, with a carrier of 38kHz, the most graduations you get per channel is 38, slightly
more than 5 bits worth. This is because ppm uses a 1ms to 2ms pulse per channel and
1ms * 38kHz = 38.

So, fairly limited, but still pretty good.

If using to drive a 180 degree servo, that means each graduation is 4.7 degrees (180/38).


## Example

    const irPin = machine.GPIOX
    ppm := irrx.PPM()
    rx := irrx.NewRxDevice(irPin, ppm)
    rx.Start()

    for {
        X := float32(ppm.Channel(0)-1500)/500
        // X ranges from -1 to 1
        servo.Set(X)
        time.Sleep(100 * time.Millisecond)
    }

*/

package irrx

import "time"

const ppmMinimumTimeBetweenFrames = 6 * time.Millisecond

type ppm struct {
	// where we store the values we've decoded
	channels [16]time.Duration
	// safeChannels are what we set channels to if we exceed Timout
	safeChannels [16]time.Duration

	// if we haven't received a frame in Timeout amount of time, we return
	// values from safeChannels
	Timeout time.Duration

	currentCh int
	last      time.Time
}

func PPM() *ppm {
	def := &ppm{
		Timeout: 10 * ppmMinimumTimeBetweenFrames,
	}
	return def
}

func (p *ppm) HandleOnOff(on, off time.Duration) {
	if on > ppmMinimumTimeBetweenFrames {
		p.last = time.Now()
		p.currentCh = 0
	}

	// prevent out-of-spec signals from panicking us.
	if p.currentCh >= len(p.channels) {
		return
	}

	p.channels[p.currentCh] = off
	p.currentCh++
}

func (p *ppm) SetSafeChannels(sc [16]time.Duration) {
	p.safeChannels = sc
}

// Channel returns the duration of the the pulse for the given channel.
// Converting the time.Duration value into something more useful is left
// to the caller.
// If ppm.Timeout has been exceeded, the safe value for the channel is
// returned.
func (p *ppm) Channel(ch int) time.Duration {
	if time.Since(p.last) > p.Timeout {
		return p.safeChannels[ch]
	}
	return p.channels[ch]
}

// Channels returns all the channels.
// If ppm.Timeout has been exceeded, the safe values are returned
func (p *ppm) Channels() [16]time.Duration {
	if time.Since(p.last) > p.Timeout {
		return p.safeChannels
	}
	return p.channels
}
