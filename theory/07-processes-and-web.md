# 7. Command-line programs and the web

A Go program can receive information from the person running it, read settings from its environment, and communicate with websites. These examples show how.

### Arguments and environment variables

- [`commandLineArgs.go`](../commandLineArgs.go) reads command-line arguments. It also uses Go's `flag` package to define named options such as `-name` and `-num`.
- [`environmentVariable.go`](../environmentVariable.go) sets and reads environment variables. These are named values supplied by the system or the shell running the program.

### URLs and HTTP

A URL has parts: a scheme such as `https`, a host, a path, and sometimes query values such as `?q=golang`. [`urlParsing.go`](../urlParsing.go) uses Go's `net/url` package to read and change these parts.

[`http_call_api.go`](../http_call_api.go) sends a request to a website and reads the response. A request can fail if the network or remote site is unavailable, so check the error. The response body should be closed after reading it.

[`go-create_server.go`](../go-create_server.go) shows the other side: it starts a small web server. A **handler** is a function that runs when a request comes in. It receives the request data and a `ResponseWriter`, which it uses to send a reply. This example listens on port 8080 and keeps running; stop it with Ctrl+C.
