package server

import (
	"github.com/familychain/family/chain"
	"github.com/familychain/family/consensus"
	consensusDev "github.com/familychain/family/consensus/dev"
	consensusDummy "github.com/familychain/family/consensus/dummy"
	consensusIBFT "github.com/familychain/family/consensus/ibft"
	consensusPolyBFT "github.com/familychain/family/consensus/polybft"
	"github.com/familychain/family/secrets"
	"github.com/familychain/family/secrets/awsssm"
	"github.com/familychain/family/secrets/gcpssm"
	"github.com/familychain/family/secrets/hashicorpvault"
	"github.com/familychain/family/secrets/local"
	"github.com/familychain/family/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = "polybft"
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
