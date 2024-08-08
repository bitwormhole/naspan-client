package main

import (
	"os"

	"github.com/bitwormhole/naspan-client/modules/naspanclient"
	"github.com/starter-go/starter"
)

func main() {
	m := naspanclient.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
