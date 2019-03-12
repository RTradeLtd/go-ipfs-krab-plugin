package node_test

import (
	"os"
	"testing"

	"github.com/RTradeLtd/go-ipfs-krab-plugin/node"
)

var (
	homeDir        = os.Getenv("HOME")
	ipfsDir        = homeDir + "/.ipfs"
	ipfsConfigFile = ipfsDir + "/config"
	logDir         = "./tmp"
)

func TestNewNode(t *testing.T) {
	node, err := node.NewNode(ipfsConfigFile, ipfsDir)
	if err != nil {
		t.Fatal(err)
	}
	if err := node.Close(); err != nil {
		t.Fatal(err)
	}
}
