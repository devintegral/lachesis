package gossip

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Fantom-foundation/go-lachesis/app"
	"github.com/Fantom-foundation/go-lachesis/evmcore"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
)

const (
	minGasPowerRefund = 800
)

// incGasPowerRefund calculates the origination gas power refund
func (s *Service) incGasPowerRefund(epoch idx.Epoch, evmBlock *evmcore.EvmBlock, receipts types.Receipts, txPositions map[common.Hash]app.TxPosition, sealEpoch bool) {
	// Calc origination scores
	for i, tx := range evmBlock.Transactions {
		txEventPos := txPositions[receipts[i].TxHash]

		if tx.Gas() < receipts[i].GasUsed {
			s.Log.Crit("Transaction gas used is higher than tx gas limit", "tx", receipts[i].TxHash)
		}
		notUsedGas := tx.Gas() - receipts[i].GasUsed
		if notUsedGas >= minGasPowerRefund { // do not refund if refunding is more costly than refunded value
			s.store.IncGasPowerRefund(epoch, txEventPos.Creator, notUsedGas)
		}
	}

	if sealEpoch {
		// prune not needed gas power records
		s.store.DelGasPowerRefunds(epoch - 1)
	}
}

// updateValidationScores calculates the validation scores
func (s *Service) updateValidationScores(block *inter.Block, sealEpoch bool) {
	blockTimeDiff := block.Time - s.store.GetBlock(block.Index-1).Time

	// Calc validation scores
	for _, it := range s.abciApp.GetActiveSfcStakers() {
		// validators only
		if !s.engine.GetValidators().Exists(it.StakerID) {
			continue
		}

		// Check if validator has confirmed events by this Atropos
		missedBlock := !s.blockParticipated[it.StakerID]

		// If have no confirmed events by this Atropos - just add missed blocks for validator
		if missedBlock {
			s.abciApp.IncBlocksMissed(it.StakerID, blockTimeDiff)
			continue
		}

		missedNum := s.abciApp.GetBlocksMissed(it.StakerID).Num
		if missedNum > s.config.Net.Economy.BlockMissedLatency {
			missedNum = s.config.Net.Economy.BlockMissedLatency
		}

		// Add score for previous blocks, but no more than FrameLatency prev blocks
		s.abciApp.AddDirtyValidationScore(it.StakerID, new(big.Int).SetUint64(uint64(blockTimeDiff)))
		for i := idx.Block(1); i <= missedNum && i < block.Index; i++ {
			blockTime := s.store.GetBlock(block.Index - i).Time
			prevBlockTime := s.store.GetBlock(block.Index - i - 1).Time
			timeDiff := blockTime - prevBlockTime
			s.abciApp.AddDirtyValidationScore(it.StakerID, new(big.Int).SetUint64(uint64(timeDiff)))
		}
		s.abciApp.ResetBlocksMissed(it.StakerID)
	}

	if sealEpoch {
		s.abciApp.DelAllActiveValidationScores()
		s.abciApp.MoveDirtyValidationScoresToActive()
	}
}
