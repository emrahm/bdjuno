package daily_refetch

import (
	"github.com/emrahm/juno/v5/node"

	bdjunodb "github.com/emrahm/bdjuno/v5/database"

	"github.com/emrahm/juno/v5/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	node     node.Node
	database *bdjunodb.Db
}

// NewModule builds a new Module instance
func NewModule(
	node node.Node,
	database *bdjunodb.Db,
) *Module {
	return &Module{
		node:     node,
		database: database,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "daily refetch"
}
