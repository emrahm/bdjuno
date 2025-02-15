package v3

import (
	v3 "github.com/emrahm/juno/v5/cmd/migrate/v3"

	"github.com/emrahm/bdjuno/v5/modules/actions"
)

type Config struct {
	v3.Config `yaml:"-,inline"`

	// The following are there to support modules which config are present if they are enabled

	Actions *actions.Config `yaml:"actions"`
}
