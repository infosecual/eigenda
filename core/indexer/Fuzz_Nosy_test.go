package indexer

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/core"
	"github.com/Layr-Labs/eigenda/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func Fuzz_Nosy_IndexedChainState_GetCurrentBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *IndexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.GetCurrentBlockNumber()
	})
}

func Fuzz_Nosy_IndexedChainState_GetIndexedOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *IndexedChainState
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

func Fuzz_Nosy_IndexedChainState_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *IndexedChainState
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

func Fuzz_Nosy_IndexedChainState_getObjects__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ics *IndexedChainState
		fill_err = tp.Fill(&ics)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if ics == nil {
			return
		}

		ics.getObjects(blockNumber)
	})
}

func Fuzz_Nosy_OperatorPubKeysAccumulator_DeserializeObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *OperatorPubKeysAccumulator
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var fork indexer.UpgradeFork
		fill_err = tp.Fill(&fork)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.DeserializeObject(d2, fork)
	})
}

func Fuzz_Nosy_OperatorPubKeysAccumulator_InitializeObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *OperatorPubKeysAccumulator
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var header indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.InitializeObject(header)
	})
}

// skipping Fuzz_Nosy_OperatorPubKeysAccumulator_SerializeObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.AccumulatorObject

// skipping Fuzz_Nosy_OperatorPubKeysAccumulator_UpdateObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.AccumulatorObject

func Fuzz_Nosy_OperatorPubKeysFilterer_FilterFastMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *OperatorPubKeysFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.FilterFastMode(headers)
	})
}

func Fuzz_Nosy_OperatorPubKeysFilterer_FilterHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *OperatorPubKeysFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.FilterHeaders(headers)
	})
}

func Fuzz_Nosy_OperatorPubKeysFilterer_GetSyncPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *OperatorPubKeysFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var latestHeader *indexer.Header
		fill_err = tp.Fill(&latestHeader)
		if fill_err != nil {
			return
		}
		if f1 == nil || latestHeader == nil {
			return
		}

		f1.GetSyncPoint(latestHeader)
	})
}

func Fuzz_Nosy_OperatorPubKeysFilterer_SetSyncPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *OperatorPubKeysFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var latestHeader *indexer.Header
		fill_err = tp.Fill(&latestHeader)
		if fill_err != nil {
			return
		}
		if f1 == nil || latestHeader == nil {
			return
		}

		f1.SetSyncPoint(latestHeader)
	})
}

func Fuzz_Nosy_OperatorSocketsAccumulator_DeserializeObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *OperatorSocketsAccumulator
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var fork indexer.UpgradeFork
		fill_err = tp.Fill(&fork)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.DeserializeObject(d2, fork)
	})
}

func Fuzz_Nosy_OperatorSocketsAccumulator_InitializeObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *OperatorSocketsAccumulator
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var header indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.InitializeObject(header)
	})
}

// skipping Fuzz_Nosy_OperatorSocketsAccumulator_SerializeObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.AccumulatorObject

// skipping Fuzz_Nosy_OperatorSocketsAccumulator_UpdateObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.AccumulatorObject

func Fuzz_Nosy_Upgrader_DetectUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *Upgrader
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		u.DetectUpgrade(headers)
	})
}

func Fuzz_Nosy_Upgrader_GetLatestUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *Upgrader
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var header *indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if u == nil || header == nil {
			return
		}

		u.GetLatestUpgrade(header)
	})
}

func Fuzz_Nosy_operatorSocketsFilterer_FilterFastMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *operatorSocketsFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.FilterFastMode(headers)
	})
}

func Fuzz_Nosy_operatorSocketsFilterer_FilterHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *operatorSocketsFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.FilterHeaders(headers)
	})
}

func Fuzz_Nosy_operatorSocketsFilterer_GetSyncPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *operatorSocketsFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var latestHeader *indexer.Header
		fill_err = tp.Fill(&latestHeader)
		if fill_err != nil {
			return
		}
		if f1 == nil || latestHeader == nil {
			return
		}

		f1.GetSyncPoint(latestHeader)
	})
}

func Fuzz_Nosy_operatorSocketsFilterer_SetSyncPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *operatorSocketsFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var latestHeader *indexer.Header
		fill_err = tp.Fill(&latestHeader)
		if fill_err != nil {
			return
		}
		if f1 == nil || latestHeader == nil {
			return
		}

		f1.SetSyncPoint(latestHeader)
	})
}

func Fuzz_Nosy_operatorSocketsFilterer_WatchOperatorSocketUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *operatorSocketsFilterer
		fill_err = tp.Fill(&f1)
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
		if f1 == nil {
			return
		}

		f1.WatchOperatorSocketUpdate(ctx, operatorId)
	})
}

// skipping Fuzz_Nosy_OperatorSocketsFilterer_FilterFastMode__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core/indexer.OperatorSocketsFilterer

// skipping Fuzz_Nosy_OperatorSocketsFilterer_FilterHeaders__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core/indexer.OperatorSocketsFilterer

// skipping Fuzz_Nosy_OperatorSocketsFilterer_GetSyncPoint__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core/indexer.OperatorSocketsFilterer

// skipping Fuzz_Nosy_OperatorSocketsFilterer_SetSyncPoint__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core/indexer.OperatorSocketsFilterer

// skipping Fuzz_Nosy_OperatorSocketsFilterer_WatchOperatorSocketUpdate__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core/indexer.OperatorSocketsFilterer

func Fuzz_Nosy_operatorPubKeysEventFilterer_FilterEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 operatorPubKeysEventFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if opts == nil {
			return
		}

		f1.FilterEvents(headers, opts)
	})
}

// skipping Fuzz_Nosy_operatorPubKeysEventFilterer_filterEvents__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_operatorPubKeysEventFilterer_filterPubKeyAddedToQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 operatorPubKeysEventFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if opts == nil {
			return
		}

		f1.filterPubKeyAddedToQuorums(headers, opts)
	})
}

func Fuzz_Nosy_operatorPubKeysEventFilterer_filterPubKeyRemovedFromQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 operatorPubKeysEventFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if opts == nil {
			return
		}

		f1.filterPubKeyRemovedFromQuorums(headers, opts)
	})
}

func Fuzz_Nosy_pubkeyRegistrationEventFilterer_addPubkeyRegistration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 pubkeyRegistrationEventFilterer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var events []operatorPubKeysEvent
		fill_err = tp.Fill(&events)
		if fill_err != nil {
			return
		}

		f1.addPubkeyRegistration(events)
	})
}
