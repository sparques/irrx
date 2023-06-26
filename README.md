# IRRX

A package for deciphering and (eventually) sending IR signals using [TinyGo](https://tinygo.org).


## A More Generic Approach

I've taken a more generic approach to IR than some other existing packages. When using the package, you pass a "RxStateMachine" which is what actually
decodes the IR signals, making a wider variety of encoding schemes easier to support.

## Out of the Box Support

Here are what IR schemes are supported out of the box, PRs adding support for other schemes are welcome.

  - HEXBUG 6-button remote controls
    - [In action here](https://www.youtube.com/embed/DqpgLUY_Q5o)

