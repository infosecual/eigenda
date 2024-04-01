package indexer

import (
	"context"
	"testing"

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

func Fuzz_Nosy_Header_After__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var prev *Header
		fill_err = tp.Fill(&prev)
		if fill_err != nil {
			return
		}
		if h == nil || prev == nil {
			return
		}

		h.After(prev)
	})
}

func Fuzz_Nosy_Header_BlockHashIs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.BlockHashIs(hash)
	})
}

func Fuzz_Nosy_Header_Equals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var other *Header
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if h == nil || other == nil {
			return
		}

		h.Equals(other)
	})
}

func Fuzz_Nosy_indexer_GetLatestHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *indexer
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var finalized bool
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.GetLatestHeader(finalized)
	})
}

func Fuzz_Nosy_indexer_GetObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *indexer
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var handlerIndex int
		fill_err = tp.Fill(&handlerIndex)
		if fill_err != nil {
			return
		}
		if i == nil || header == nil {
			return
		}

		i.GetObject(header, handlerIndex)
	})
}

// skipping Fuzz_Nosy_indexer_HandleAccumulator__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

func Fuzz_Nosy_indexer_Index__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *indexer
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.Index(ctx)
	})
}

// skipping Fuzz_Nosy_Accumulator_DeserializeObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

// skipping Fuzz_Nosy_Accumulator_InitializeObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

// skipping Fuzz_Nosy_Accumulator_SerializeObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

// skipping Fuzz_Nosy_Accumulator_UpdateObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

// skipping Fuzz_Nosy_Filterer_FilterFastMode__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Filterer

// skipping Fuzz_Nosy_Filterer_FilterHeaders__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Filterer

// skipping Fuzz_Nosy_Filterer_GetSyncPoint__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Filterer

// skipping Fuzz_Nosy_Filterer_SetSyncPoint__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Filterer

// skipping Fuzz_Nosy_HeaderService_PullLatestHeader__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.HeaderService

// skipping Fuzz_Nosy_HeaderService_PullNewHeaders__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.HeaderService

// skipping Fuzz_Nosy_HeaderStore_AddHeaders__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.HeaderStore

// skipping Fuzz_Nosy_HeaderStore_AttachObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.HeaderStore

// skipping Fuzz_Nosy_HeaderStore_FastForward__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.HeaderStore

// skipping Fuzz_Nosy_HeaderStore_GetLatestHeader__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.HeaderStore

// skipping Fuzz_Nosy_HeaderStore_GetLatestObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.HeaderStore

// skipping Fuzz_Nosy_HeaderStore_GetObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.HeaderStore

func Fuzz_Nosy_Headers_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hh Headers
		fill_err = tp.Fill(&hh)
		if fill_err != nil {
			return
		}

		hh.Empty()
	})
}

func Fuzz_Nosy_Headers_First__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hh Headers
		fill_err = tp.Fill(&hh)
		if fill_err != nil {
			return
		}

		hh.First()
	})
}

func Fuzz_Nosy_Headers_GetHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hh Headers
		fill_err = tp.Fill(&hh)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}

		hh.GetHeaderByNumber(number)
	})
}

func Fuzz_Nosy_Headers_IsOrdered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hh Headers
		fill_err = tp.Fill(&hh)
		if fill_err != nil {
			return
		}

		hh.IsOrdered()
	})
}

func Fuzz_Nosy_Headers_Last__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hh Headers
		fill_err = tp.Fill(&hh)
		if fill_err != nil {
			return
		}

		hh.Last()
	})
}

func Fuzz_Nosy_Headers_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hh Headers
		fill_err = tp.Fill(&hh)
		if fill_err != nil {
			return
		}

		hh.Len()
	})
}

func Fuzz_Nosy_Headers_OK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hh Headers
		fill_err = tp.Fill(&hh)
		if fill_err != nil {
			return
		}

		hh.OK()
	})
}

// skipping Fuzz_Nosy_Indexer_GetLatestHeader__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Indexer

// skipping Fuzz_Nosy_Indexer_GetObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Indexer

// skipping Fuzz_Nosy_Indexer_HandleAccumulator__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Indexer

// skipping Fuzz_Nosy_Indexer_Index__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Indexer

// skipping Fuzz_Nosy_UpgradeForkWatcher_DetectUpgrade__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.UpgradeForkWatcher

// skipping Fuzz_Nosy_UpgradeForkWatcher_GetLatestUpgrade__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.UpgradeForkWatcher

func Fuzz_Nosy_CLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
		CLIFlags(envPrefix)
	})
}
