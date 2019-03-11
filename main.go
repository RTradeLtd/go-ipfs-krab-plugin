package main

import (
	"github.com/RTradeLtd/go-ipfs-krab-plugin/krab"
	"github.com/ipfs/go-ipfs/plugin"
)

// Plugins is an exported list of plugins that will be loaded by go-ipfs.
var Plugins = []plugin.Plugin{
	&krab.KBPlugin{},
}

func main() {}
