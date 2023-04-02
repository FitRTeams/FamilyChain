package memory

import (
	"testing"

	"https://github.com/FitRTeams/familychain/blockchain/storage"
)

func TestStorage(t *testing.T) {
	t.Helper()

	f := func(t *testing.T) (storage.Storage, func()) {
		t.Helper()

		s, _ := NewMemoryStorage(nil)

		return s, func() {}
	}
	storage.TestStorage(t, f)
}
