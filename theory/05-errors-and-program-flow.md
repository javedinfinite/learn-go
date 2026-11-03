# 5. Errors and program flow

A function can return an **error** to tell its caller that something went wrong. The caller should check the error before using the result. In Go this is written out directly, often like this:

    result, err := doSomething()
    if err != nil {
        // handle the problem
    }

See [`errors.go`](../errors.go). Its `divide` function returns an error if the second number is zero. The caller checks `err` and only prints the result if the calculation succeeded.

### Defer, panic, and recover

`defer` schedules a function call to run when the current function is about to finish. It is often used for cleanup. If several calls are deferred, the last one added runs first.

A **panic** stops the normal flow of a program. A deferred function can call `recover` to catch a panic and allow the program to continue unwinding in a controlled way. [`panic_defer_recover.go`](../panic_defer_recover.go) demonstrates all three. Most everyday problems should be returned as errors; panic and recover are for unusual situations where normal error returns are not the right fit.
