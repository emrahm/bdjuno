package source

import (
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

type Source interface {
	Proposal(height int64, id uint64) (v1beta1.Proposal, error)
	ProposalDeposit(height int64, id uint64, depositor string) (v1beta1.Deposit, error)
	TallyResult(height int64, proposalID uint64) (v1beta1.TallyResult, error)
	DepositParams(height int64) (v1beta1.DepositParams, error)
	VotingParams(height int64) (v1beta1.VotingParams, error)
	TallyParams(height int64) (v1beta1.TallyParams, error)
}
