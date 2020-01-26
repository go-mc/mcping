# mcping

Ping your mincraft server with a simple command! This program can also download favicon.

## install

```shell
$ go version # installed golang
go version go1.13.4 darwin/amd64
$ go get github.com/go-mc/mcping/mcping # install
```

Please ensure `$GOPATH/bin` or `$GOBIN` is in your `$PATH`.Because mcping will be installed in `$(go env GOPATH)/bin`.

Or you can `go build` and move the executable to your `$PATH`. Or run it directly.

## useage

```shell
mcping host:port			# ping specific server
mcping host					# ping default port (25565)
mcping host -p 404	        # specific protocol version
mcping host -f icon.png		# save server's favicon
mcping host -f="/path with spaces/icon.png"
```

> mcping will lookup SRV record if port is 25565 (include default set) as same as Minecraft itself.

## programming

[![](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-mc/mcping?tab=doc)

```go
import "github.com/go-mc/mcping"
```

