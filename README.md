# Source Code for Backend Services

This service uses Go version 1.20.1 and the fiber web framework with version v2.42.0.
The 3 endpoints return hard-coded JSON responses after a fixed delay of 250 milliseconds.

## Build & Run

```
go build fiber
.\fiber.exe
```

## Build for Linux

Run in Powershell

```
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build
```
