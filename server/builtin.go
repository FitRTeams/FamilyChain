package server

import (
	"github.com/FitRTeams/familychain/chain"
	"github.com/FitRTeams/familychain/consensus"
	consensusDev "github.com/FitRTeams/familychain/consensus/dev"
	consensusDummy "github.com/FitRTeams/familychain/consensus/dummy"
	consensusIBFT "github.com/FitRTeams/familychain/consensus/ibft"
	consensusPolyBFT "github.com/FitRTeams/familychain/consensus/polybft"
	"github.com/FitRTeams/familychain/secrets"
	"github.com/FitRTeams/familychain/secrets/awsssm"
	"github.com/FitRTeams/familychain/secrets/gcpssm"
	"github.com/FitRTeams/familychain/secrets/hashicorpvault"
	"github.com/FitRTeams/familychain/secrets/local"
	"github.com/FitRTeams/familychain/state"
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
