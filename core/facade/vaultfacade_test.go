package facade

import (
	"testing"

	"github.com/mashmb/1pass/adapter/out/repo/file"
	"github.com/mashmb/1pass/adapter/out/util/crypto"
	"github.com/mashmb/1pass/core/service"
	"github.com/mashmb/1pass/port/out"
)

func setup() *dfltVaultFacade {
	var cryptoUtils out.CrytpoUtils
	var profileRepo out.ProfileRepo

	var keyService service.KeyService

	cryptoUtils = crypto.NewPbkdf2CryptoUtils()
	profileRepo = file.NewFileProfileRepo("../../assets/onepassword_data")

	keyService = service.NewDfltKeyService(cryptoUtils, profileRepo)

	return NewDfltVaultFacade(keyService)
}

func TestLock(t *testing.T) {
	facade := setup()
	pass := "freddy"
	facade.Unlock(pass)

	if facade.keys == nil {
		t.Errorf("Unlock() should provide keys")
	}

	facade.Lock()

	if facade.keys != nil {
		t.Errorf("Lock() should clear keys")
	}
}

func TestUnlock(t *testing.T) {
	facade := setup()
	goodPass := "freddy"
	badPass := ""
	err := facade.Unlock(badPass)

	if err == nil {
		t.Error("Unlock() should fail because of invalid password")
	}

	err = facade.Unlock(goodPass)

	if err != nil {
		t.Error("Unlock() should pass because of valid password")
	}
}
