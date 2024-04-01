package thegraph

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/core"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_indexedChainState_GetIndexedOperatorInfoByOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *indexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorId core.OperatorID
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.GetIndexedOperatorInfoByOperatorId(ctx, operatorId, blockNumber)
	})
}

func Fuzz_Nosy_indexedChainState_GetIndexedOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *indexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.GetIndexedOperatorState(ctx, blockNumber, quorums)
	})
}

func Fuzz_Nosy_indexedChainState_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *indexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.Start(ctx)
	})
}

func Fuzz_Nosy_indexedChainState_getAllOperatorsRegisteredAtBlockNumberWithPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *indexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.getAllOperatorsRegisteredAtBlockNumberWithPagination(ctx, blockNumber)
	})
}

func Fuzz_Nosy_indexedChainState_getQuorumAPK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *indexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.getQuorumAPK(ctx, quorumID, blockNumber)
	})
}

func Fuzz_Nosy_indexedChainState_getQuorumAPKs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *indexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var quorumIDs []uint8
		fill_err = tp.Fill(&quorumIDs)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.getQuorumAPKs(ctx, quorumIDs, blockNumber)
	})
}

func Fuzz_Nosy_indexedChainState_getRegisteredIndexedOperatorInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *indexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.getRegisteredIndexedOperatorInfo(ctx, blockNumber)
	})
}

// skipping Fuzz_Nosy_GraphQLQuerier_Query__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core/thegraph.GraphQLQuerier

// skipping Fuzz_Nosy_IndexedChainState_GetIndexedOperatorInfoByOperatorId__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core/thegraph.IndexedChainState

// skipping Fuzz_Nosy_IndexedChainState_GetIndexedOperatorState__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core/thegraph.IndexedChainState
