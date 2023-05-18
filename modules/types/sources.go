package types

import (
	"fmt"

	"cosmossdk.io/simapp/params"
	"github.com/emrahm/juno/v5/node/remote"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	nodeconfig "github.com/emrahm/juno/v5/node/config"

	banksource "github.com/emrahm/bdjuno/v5/modules/bank/source"
	remotebanksource "github.com/emrahm/bdjuno/v5/modules/bank/source/remote"
	distrsource "github.com/emrahm/bdjuno/v5/modules/distribution/source"
	remotedistrsource "github.com/emrahm/bdjuno/v5/modules/distribution/source/remote"
	govsource "github.com/emrahm/bdjuno/v5/modules/gov/source"
	remotegovsource "github.com/emrahm/bdjuno/v5/modules/gov/source/remote"
	mintsource "github.com/emrahm/bdjuno/v5/modules/mint/source"
	remotemintsource "github.com/emrahm/bdjuno/v5/modules/mint/source/remote"
	slashingsource "github.com/emrahm/bdjuno/v5/modules/slashing/source"
	remoteslashingsource "github.com/emrahm/bdjuno/v5/modules/slashing/source/remote"
	stakingsource "github.com/emrahm/bdjuno/v5/modules/staking/source"
	remotestakingsource "github.com/emrahm/bdjuno/v5/modules/staking/source/remote"
)

type Sources struct {
	BankSource     banksource.Source
	DistrSource    distrsource.Source
	GovSource      govsource.Source
	MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	StakingSource  stakingsource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource:     remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
		DistrSource:    remotedistrsource.NewSource(source, distrtypes.NewQueryClient(source.GrpcConn)),
		GovSource:      remotegovsource.NewSource(source, govtypes.NewQueryClient(source.GrpcConn)),
		MintSource:     remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource: remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		StakingSource:  remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
