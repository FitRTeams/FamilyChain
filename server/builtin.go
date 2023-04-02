package server

import (
	"https://github.com/FitRTeams/familychain/chain"
	"https://github.com/FitRTeams/familychain/consensus"
	consensusDev "https://github.com/FitRTeams/familychain/consensus/dev"
	consensusDummy "https://github.com/FitRTeams/familychain/consensus/dummy"
	consensusIBFT "https://github.com/FitRTeams/familychain/consensus/ibft"
	consensusPolyBFT "https://github.com/FitRTeams/familychain/consensus/polybft"
	"https://github.com/FitRTeams/familychain/secrets"
	"https://github.com/FitRTeams/familychain/secrets/awsssm"
	"https://github.com/FitRTeams/familychain/secrets/gcpssm"
	"https://github.com/FitRTeams/familychain/secrets/hashicorpvault"
	"https://github.com/FitRTeams/familychain/secrets/local"
	"https://github.com/FitRTeams/familychain/state"
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
