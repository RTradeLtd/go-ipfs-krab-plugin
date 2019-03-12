package krab_test

import (
	"os"
	"testing"

	"github.com/RTradeLtd/go-ipfs-krab-plugin/krab"
)

const (
	testPK      = "CAASqgkwggSmAgEAAoIBAQCf0Ps1HJwqgerwmG3a44I0jCx7Im1OK5rribEEVem4T66CaN++MwHKkanmwKldjvHfQRYX6h7rl7sASjcZsVQum2lWyr7BjSpuuQM4yj5krJX5n1WYfXYi0RVphjk0TeD36aCV6vZBffAHAqdNLcuSozrjRjBYzqI8ksmW93AP7krUN4PoVVUy7hcvX4vU5TpB6OMnPo2uD1rFQG2mRhAtQm7DsdMLLsHkl1mwZ159oBiq0HZ8KX511AnF7O4XbfuhYBLP3bA2nMOhLGKNRuEwabDFFHtJSOT8KigcJZqRaDoEvCthH36kU+fOCssVtQ5vilpD4OBFRngxble18R8JAgMBAAECggEBAJibAYVx0F12uVUM+LlvTcHHqCHceeuPD9uiG3o8X3U0ATtd0WjZY/h+p3JEvnUsLiI1EFe9o9DrjhF0zLCn9+6ZUDkkWIqRtdcMq9rqpDGV64/1adK22rvcU0n0dWNQimWwnTsKpiNbknvfOMol1KItY+np9/iVN64HrJ+Pn15Vdl4WkSndnLYGhRU8dAsdFGBMZsqqlezKb3exVg0ZlOpBeYDzgTYffXJWaaypDBeQv51MPJ13GkLoSI75HX5pWl0vbPVjYW80ge2klNLWO0xXi06C/AJym/2proK2/qxOAX5yE8TyTqISO+0Ta4tNVw8nbvFDfd/vBYuL9yPyvgECgYEAwmfLVSgp8+3W3QTvx30Qn16aRrzo6PkGKiMqX5ZRBsVuaRZIZwcQqw0ZIvkvnLl3m2LzNie7t+IOWz10Dh3KxKk61Cr85nNJqT0QtrpCRFTeDTh9m+4rBhl9619ZOcTujyAv32Jh7lEY1R5dnbrxCzjfs/v1m75W5x+AUArsCLkCgYEA0nOq57/lczi9YZ47TUEpvn0Ah7x84SvUiFjGwJ9U7q90tVdauBTOE12F/HNdjW9HkPDUbHyZ7KElkksi94o1ufH4I6ymLjSIXDX4akKJPi/R5zwXERPP6DYVtnoc3MByFYdpSYc7i25XPfP2UUbU9HwZ6ORfYim2OXfVLgouANECgYEAs/olq3vYgySgs9O7LONi/Tg2+eAwfGb3RxFxTDc8Yllro0xm0UMgMkuZBuDNLHoj+i48XdmhF1bn1Z5qEBuSuki11vDJW4xGGLEZBSIg8WPkgzbWSwLOwAHfqoWGdE4WUVkKGPPbGCfqJnvLTZhhSfNoXxeXRY2MpMJsJy5FYLkCgYEAiwQDAgEdtSXm24PlUlvYEk+KOR+GMkt8ofBaocTeGqjl5B/m+jAsDFi7+A3Q32uBj7m16E/KcaflJKTjXTb71G+E5TPXti/XX2n7RV3cQKrf8eocppg6vv5zC03QoPWypowDPaHJ8ImbsX3k2L18UF9l0hSA+VvqVj3VIQtyZfECgYEAkd52znQNEHSgIr7DQCX5ssxcbCGOuW8BWWGChaRMSIc/7MnJTAVRgErwtKSSD4ZhOpnpExZ91a1ItHVcSWNW6UuC6SlcDJyDWTVDvj1ACRl8nhzL0v7ARBq7x/+QgIfIghhw8td382FTDsZDTO+e8MPvU3XWw714NtUggR9eCM4="
	testEncPass = "password123"
	testDSPath  = "tmp"
)

func TestKrabPlugin(t *testing.T) {
	var (
		expectedName     = "krab-keystore"
		expectedVersion  = "v0.0.0"
		expectedTypeName = "krab"
	)
	kb := krab.KBPlugin{}
	if kb.Name() != expectedName {
		t.Fatal("bad name")
	}
	if kb.Version() != expectedVersion {
		t.Fatal("bad version")
	}
	if err := kb.Init(); err != nil {
		t.Fatal(err)
	}
	if kb.KeystoreTypeName() != expectedTypeName {
		t.Fatal("bad keystore type nmae")
	}
	if err := os.Mkdir(testDSPath, os.FileMode(0777)); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(testDSPath)
	kbStore, err := kb.Open(testDSPath, map[string]interface{}{
		"passphrase":     testEncPass,
		"selfPrivateKey": testPK,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if exists, err := kbStore.Has("self"); err != nil {
		t.Fatal(err)
	} else if !exists {
		t.Fatal("key should exist")
	}
}
