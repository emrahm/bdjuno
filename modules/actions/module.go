package actions

import (
	"cosmossdk.io/simapp/params"
	"github.com/emrahm/juno/v5/modules"
	"github.com/emrahm/juno/v5/node"
	"github.com/emrahm/juno/v5/node/builder"
	nodeconfig "github.com/emrahm/juno/v5/node/config"
	"github.com/emrahm/juno/v5/types/config"

	modulestypes "github.com/emrahm/bdjuno/v5/modules/types"
)

const (
	ModuleName = "actions"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

type Module struct {
	cfg     *Config
	node    node.Node
	sources *modulestypes.Sources
}

func NewModule(cfg config.Config, encodingConfig *params.EncodingConfig) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}

	actionsCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}

	nodeCfg := cfg.Node
	if actionsCfg.Node != nil {
		nodeCfg = nodeconfig.NewConfig(nodeconfig.TypeRemote, actionsCfg.Node)
	}

	// Build the node
	junoNode, err := builder.BuildNode(nodeCfg, encodingConfig)
	if err != nil {
		panic(err)
	}

	// Build the sources
	sources, err := modulestypes.BuildSources(nodeCfg, encodingConfig)
	if err != nil {
		panic(err)
	}

	return &Module{
		cfg:     actionsCfg,
		node:    junoNode,
		sources: sources,
	}
}

func (m *Module) Name() string {
	return ModuleName
}
