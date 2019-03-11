package krab

import (
	"errors"

	krab "github.com/RTradeLtd/rtfs/krab"
	"github.com/ipfs/go-ipfs/keystore"
	"github.com/ipfs/go-ipfs/plugin"
	prompt "github.com/ipfs/go-prompt"
)

// KBPlugin is the core type of the krab keystore plugin
type KBPlugin struct{}

var _ plugin.PluginKeystore = (*KBPlugin)(nil)

// Name returns the plugin name
func (kb *KBPlugin) Name() string {
	return "krab-keystore"
}

// Version returns the plugin version
func (kb *KBPlugin) Version() string {
	return "v0.0.0"
}

// Init is used to hold initialization logic
func (kb *KBPlugin) Init() error {
	return nil
}

// KeystoreTypeName returns the type of the keystore
func (kb *KBPlugin) KeystoreTypeName() string {
	return "krab"
}

// Open is used to open our connection to the krab keystore
func (kb *KBPlugin) Open(repoPath string, config map[string]interface{}, prompter prompt.Prompter) (keystore.Keystore, error) {
	encPass, ok := config["passphrase"].(string)
	if !ok {
		return nil, errors.New("passphrase not a valid type")
	}
	kbKeystore, err := krab.NewKrab(krab.Opts{Passphrase: encPass, DSPath: repoPath})
	if err != nil {
		return nil, err
	}
	return kbKeystore, nil
}
