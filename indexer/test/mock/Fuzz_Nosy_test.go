package mock

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
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

// skipping Fuzz_Nosy_ContractSimulator_Start__ because parameters include func, chan, or unsupported interface: context.CancelFunc

func Fuzz_Nosy_simulatedBackend_BatchCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sb *simulatedBackend
		fill_err = tp.Fill(&sb)
		if fill_err != nil {
			return
		}
		var b []rpc.BatchElem
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if sb == nil {
			return
		}

		sb.BatchCall(b)
	})
}

func Fuzz_Nosy_simulatedBackend_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sb *simulatedBackend
		fill_err = tp.Fill(&sb)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var b []rpc.BatchElem
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if sb == nil {
			return
		}

		sb.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_simulatedBackend_Call__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_simulatedBackend_CallContext__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_simulatedBackend_getBlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sb *simulatedBackend
		fill_err = tp.Fill(&sb)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var result *types.Header
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		var blockNum string
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		if sb == nil || result == nil {
			return
		}

		sb.getBlockByNumber(ctx, result, blockNum)
	})
}

// skipping Fuzz_Nosy_SimulatedBackend_AdjustTime__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_BalanceAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_BlockByHash__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_CallContract__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_Close__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_CodeAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_Commit__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_EstimateGas__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_FilterLogs__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_Fork__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_HeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_NonceAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_PendingCallContract__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_PendingCodeAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_PendingNonceAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_Rollback__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_SendTransaction__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_StorageAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_SuggestGasPrice__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_SuggestGasTipCap__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_TransactionByHash__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_TransactionCount__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_TransactionInBlock__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend

// skipping Fuzz_Nosy_SimulatedBackend_TransactionReceipt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer/test/mock.SimulatedBackend
