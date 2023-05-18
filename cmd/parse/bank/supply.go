package bank

import (
	"fmt"

	modulestypes "github.com/emrahm/bdjuno/v5/modules/types"

	parsecmdtypes "github.com/emrahm/juno/v5/cmd/parse/types"
	"github.com/emrahm/juno/v5/types/config"
	"github.com/spf13/cobra"

	"github.com/emrahm/bdjuno/v5/database"
	"github.com/emrahm/bdjuno/v5/modules/bank"
)

// supplyCmd returns the Cobra command allowing to refresh x/bank total supply
func supplyCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "supply",
		Short: "Refresh total supply",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			sources, err := modulestypes.BuildSources(config.Cfg.Node, parseCtx.EncodingConfig)
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Build bank module
			bankModule := bank.NewModule(nil, sources.BankSource, parseCtx.EncodingConfig.Marshaler, db)

			err = bankModule.UpdateSupply()
			if err != nil {
				return fmt.Errorf("error while getting latest bank supply: %s", err)
			}

			return nil
		},
	}
}
