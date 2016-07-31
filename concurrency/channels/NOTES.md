# Channels in Go

A popular adage in Go is: 

_Do not communicate by sharing memory; instead, share memory by communicating._

...and this sharing memory by communicating is accomplished with Go’s *channels,* which synchronize and coordinate independently running go-routines. Go still does contain other more well=known tools for dealing with concurrency, such as wait-groups, mutexes, and atomicity for handling shared variables; however, Go's channels are far more powerful. 

Channels are created with the `make()` function, and sending and receiving on a channel is accomplished with the channel operand: `<-`

`mychannel := make(chan string)`: this creates a channel called "mychannel" which takes string values. 

*Sending*: the channel operand separates the channel and the value being sent; here the string "The Dude Abides" is sent to the channel mychannel! 
 
 `mychannel <- "The Dude Abides"`
 
*Receiving:* the value comes before the channel operand; here x receives whatever strings were on the mychannel channel: 

`x := <- mychannel`

A channel can also discard its results; here, we discard the results of mychannel: 

`<- mychannel`
 
To *close* a channel, we use the `close()` function. When we close a channel, no more values can be sent to it. If any values are sent to a channel after it has been closed, a panic will occur. 

`close(mychannel)`

## Unbuffered vs. Buffered Channels

*Unbuffered channels:* are channels with no capacity to hold any value before it is received. Unbuffered channels require both a send and a receive goroutine to be ready at the same time before any send/receive operations can complete. Unbuffered channels are created by omitting any capacity value in the make statement, i.e. `make(chan string).` 

TODO EXAMPLE 

*Buffered channels:* are a channel with a capacity to hold one or more values before they are received. These are created by providing the channel capacity as a parameter to the make function. Thus: 

`make(chan string, 3)` 

...creates a buffered channel capable of holding 3 string values. 

TODO EXAMPLE

