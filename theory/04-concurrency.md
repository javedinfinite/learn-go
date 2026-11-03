# 4. Concurrency

Concurrency lets a Go program work on more than one task at a time. A **goroutine** is a function running concurrently with other work. Put `go` before a function call to start it as a goroutine. The caller continues without waiting for that function to finish.

Read [`go_routine.go`](../go_routine.go). It first uses a short sleep to keep the program alive, then shows a `WaitGroup`. A wait group lets the main work wait until the goroutines say they are done.

### Channels

A **channel** passes values between goroutines. One goroutine sends a value and another receives it. An unbuffered channel makes the sender wait until a receiver is ready. A buffered channel can hold a fixed number of values while the receiver catches up.

[`channels.go`](../channels.go) shows both kinds. `PrintChannels1` is an example of what can go wrong: it sends on an unbuffered channel before starting a receiver, so the program gets stuck. In `PrintChannels2`, another goroutine receives the values. The sender closes a channel when it has no more values to send; then a `range` loop over that channel can finish.

### Choosing, timing, and sharing work

- [`select_timeout.go`](../select_timeout.go) uses `select` to wait for whichever channel is ready first. It also shows how to stop waiting after a timeout, and how timers and tickers send time values.
- [`worker_pools.go`](../worker_pools.go) starts several workers that take jobs from the same channel and send results back.
- [`rate_limitter.go`](../rate_limitter.go) uses a channel and a timer to limit how quickly requests are sent.

These examples combine goroutines and channels: goroutines do the work, and channels help pass data or coordinate when that work can continue.
