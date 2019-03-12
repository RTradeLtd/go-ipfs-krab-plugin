package node

import (
	"context"
	"encoding/json"
	"os"

	"github.com/RTradeLtd/go-ipfs-krab-plugin/krab"
	"github.com/ipfs/go-ipfs/repo/fsrepo"

	config "github.com/ipfs/go-ipfs-config"
	"github.com/ipfs/go-ipfs/core"
)

// NewNode is used to instantaite our test ipfs node
func NewNode(ipfsConfigPath, repoPath string) (*core.IpfsNode, error) {
	f, err := os.Open(ipfsConfigPath)
	if err != nil {
		return nil, err
	}
	cfg := config.Config{}
	// decode the opened config file into our config struct
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	if err := fsrepo.AddKeystore("krab", krab.Open); err != nil {
		return nil, err
	}
	if err := fsrepo.Init(repoPath, &cfg); err != nil {
		return nil, err
	}
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}
	node, err := core.NewNode(context.Background(), &core.BuildCfg{Repo: repo})
	if err != nil {
		return nil, err
	}
	return node, nil
}
