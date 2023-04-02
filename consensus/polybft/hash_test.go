package polybft

import (
	"testing"

	"github.com/FitRTeams/familychain/consensus/polybft/bitmap"
	bls "github.com/FitRTeams/familychain/consensus/polybft/signer"
	"github.com/FitRTeams/familychain/consensus/polybft/wallet"
	"github.com/FitRTeams/familychain/types"
	"github.com/stretchr/testify/assert"
)

func Test_setupHeaderHashFunc(t *testing.T) {
	extra := &Extra{
		Validators: &ValidatorSetDelta{Removed: bitmap.Bitmap{1}},
		Parent:     createSignature(t, []*wallet.Account{generateTestAccount(t)}, types.ZeroHash, bls.DomainCheckpointManager),
		Checkpoint: &CheckpointData{},
		Seal:       []byte{},
		Committed:  &Signature{},
	}

	header := &types.Header{
		Number:    2,
		GasLimit:  10000003,
		Timestamp: 18,
	}

	header.ExtraData = append(make([]byte, ExtraVanity), extra.MarshalRLPTo(nil)...)
	notFullExtraHash := types.HeaderHash(header)

	extra.Seal = []byte{1, 2, 3, 255}
	extra.Committed = createSignature(t, []*wallet.Account{generateTestAccount(t)}, types.ZeroHash, bls.DomainCheckpointManager)
	header.ExtraData = append(make([]byte, ExtraVanity), extra.MarshalRLPTo(nil)...)
	fullExtraHash := types.HeaderHash(header)

	assert.Equal(t, notFullExtraHash, fullExtraHash)

	header.ExtraData = []byte{1, 2, 3, 4, 100, 200, 255}
	assert.Equal(t, types.ZeroHash, types.HeaderHash(header)) // to small extra data
}
