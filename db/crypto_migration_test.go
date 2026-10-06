package db

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
	"testing"
)

func TestLegacyAndNewPasswords(t *testing.T) {
	for _, salt := range [][]byte{appCryptoSalt, legacyCryptoSalt} {
		key, _ := deriveKeyWithSalt(salt)
		block, _ := aes.NewCipher(key)
		gcm, _ := cipher.NewGCM(block)
		nonce := make([]byte, gcm.NonceSize())
		old := base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte("legacy password"), nil))
		got, err := DecryptPassword(old)
		if err != nil || got != "legacy password" {
			t.Fatalf("legacy decrypt: %q %v", got, err)
		}
	}
	a, _ := EncryptPassword("new password")
	b, _ := EncryptPassword("new password")
	if !strings.HasPrefix(a, "g2:") || a == b {
		t.Fatal("missing version or random salt")
	}
	raw, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(a, "g2:"))
	raw[len(raw)-1] ^= 1
	if _, err := DecryptPassword("g2:" + base64.StdEncoding.EncodeToString(raw)); err == nil {
		t.Fatal("tampering accepted")
	}
}
