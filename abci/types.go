package abci

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AggregationContractPayload represents the data the aggregation contracts will expect as arguments
type AggregationContractPayload struct {
	Data []DataWithVotingPower `json:"data"`
}

type DataWithVotingPower struct {
	Vote        []byte `json:"vote"`
	VotingPower uint64  `json:"ve_power"`
}

// ExtCommitInfoIdx is the expected index for the oracle's extended commit info
const ExtCommitInfoIdx = 0
const InjectedTxs = 1

// VoteExtensionsEnabled checks state for whether or not ves are enabled
func VoteExtensionsEnabled(ctx sdk.Context) bool {
	cp := ctx.ConsensusParams()
	if cp.Abci == nil || cp.Abci.VoteExtensionsEnableHeight == 0 {
		return false
	}

	// Per the cosmos sdk, the first block should not utilize the latest finalize block state. This means
	// vote extensions should NOT be making state changes.
	//
	// Ref: https://github.com/cosmos/cosmos-sdk/blob/2100a73dcea634ce914977dbddb4991a020ee345/baseapp/baseapp.go#L488-L495
	if ctx.BlockHeight() <= 1 {
		return false
	}

	return cp.Abci.VoteExtensionsEnableHeight < ctx.BlockHeight()
}