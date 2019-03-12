package krab

import (
	"errors"

	krab "github.com/RTradeLtd/rtfs/krab"
	"github.com/ipfs/go-ipfs/keystore"
	"github.com/ipfs/go-ipfs/plugin"
	prompt "github.com/ipfs/go-prompt"
	ci "github.com/libp2p/go-libp2p-crypto"
)

const (
	name     = "krab-keystore"
	version  = "v0.0.0"
	typeName = "krab"
)

// KBPlugin is the core type of the krab keystore plugin
type KBPlugin struct{}

var _ plugin.PluginKeystore = (*KBPlugin)(nil)

// Name returns the plugin name
func (kb *KBPlugin) Name() string {
	return name
}

// Version returns the plugin version
func (kb *KBPlugin) Version() string {
	return version
}

// Init is used to hold initialization logic
func (kb *KBPlugin) Init() error {
	return nil
}

// KeystoreTypeName returns the type of the keystore
func (kb *KBPlugin) KeystoreTypeName() string {
	return typeName
}

// Open is used to open our connection to the krab keystore
func (kb *KBPlugin) Open(repoPath string, config map[string]interface{}, prompter prompt.Prompter) (keystore.Keystore, error) {
	encPass, ok := config["passphrase"].(string)
	if !ok {
		return nil, errors.New("passphrase not a valid type")
	}
	selfPrivateKey, ok := config["selfPrivateKey"].(string)
	if !ok {
		return nil, errors.New("selfPrivateKey not a valid type")
	}
	kbKeystore, err := krab.NewKrab(krab.Opts{Passphrase: encPass, DSPath: repoPath})
	if err != nil {
		return nil, err
	}
	if exists, err := kbKeystore.Has("self"); err != nil {
		decoded, err := ci.ConfigDecodeKey(selfPrivateKey)
		if err != nil {
			return nil, err
		}
		pk, err := ci.UnmarshalPrivateKey(decoded)
		if err != nil {
			return nil, err
		}
		if err := kbKeystore.Put("self", pk); err != nil {
			return nil, err
		}
	} else if !exists {
		return nil, errors.New("unexpected error")
	}
	return kbKeystore, nil
}
