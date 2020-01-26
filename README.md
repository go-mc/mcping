# mcping

Ping your mincraft server with a simple command!  

```shell
$ mcping play.miaoscraft.cn
MCPING (play.miaoscraft.cn:25565):
server: Spigot 1.15.2
protocol: 578
description: 
Miaoscraft!
delay: 42.200607ms
list: 2/18
- [defevt] 30e0098b-20a7-4067-b552-73517ad146dc
- [Wizard_BOSS] 664b582d-8838-4761-b011-c3852004f47d

```

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
mcping host -protocol=404	# specific protocol version
```

> mcping will lookup SRV record if port is 25565 (include default set) as same as Minecraft itself.

