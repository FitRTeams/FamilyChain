package server

import (
	"github.com/FamilyChain/family/chain"
	"github.com/FamilyChain/family/consensus"
	consensusDev "github.com/FamilyChain/family/consensus/dev"
	consensusDummy "github.com/FamilyChain/family/consensus/dummy"
	consensusIBFT "github.com/FamilyChain/family/consensus/ibft"
	consensusPolyBFT "github.com/FamilyChain/family/consensus/polybft"
	"github.com/FamilyChain/family/secrets"
	"github.com/FamilyChain/family/secrets/awsssm"
	"github.com/FamilyChain/family/secrets/gcpssm"
	"github.com/FamilyChain/family/secrets/hashicorpvault"
	"github.com/FamilyChain/family/secrets/local"
	"github.com/FamilyChain/family/state"
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
