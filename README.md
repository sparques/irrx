# IRRX

A package for deciphering and (eventually) sending IR signals using [TinyGo](https://tinygo.org).


## A More Generic Approach

I've taken a more generic approach to IR than some other existing packages. When using the package, you pass a "RxStateMachine" which is what actually
decodes the IR signals, making a wider variety of encoding schemes easier to support.

Simpler (and hopefully faster) than the irremote package (https://github.com/tinygo-org/drivers/blob/release/irremote/receiver.go) that is, IMHO,
amusingly over-written.

## Out of the Box Support

Here are what IR schemes are supported out of the box, PRs adding support for other schemes are welcome.

  - NEC (most common kind of IR remote control, like for TVs)
  - HEXBUG 6-button remote controls
    - [In action here](https://www.youtube.com/embed/DqpgLUY_Q5o)
  - PPM-over-IR (this isn't a thing, but it could be)


## Not just IR

Because of how the demodulating IR receivers work, you could theoretically hook up, say, a
PPM radio control receiver directly to an MCU and use this package plus the PPM StateMachine
to decode the PPM signals.

Basically anything that can be decoded as timed on-off pairs will work here.