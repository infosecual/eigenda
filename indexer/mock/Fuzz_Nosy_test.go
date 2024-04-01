package mock

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/indexer"
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

func Fuzz_Nosy_MockIndexer_GetLatestHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockIndexer
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var finalized bool
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetLatestHeader(finalized)
	})
}

func Fuzz_Nosy_MockIndexer_GetObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockIndexer
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var header *indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var handlerIndex int
		fill_err = tp.Fill(&handlerIndex)
		if fill_err != nil {
			return
		}
		if m == nil || header == nil {
			return
		}

		m.GetObject(header, handlerIndex)
	})
}

// skipping Fuzz_Nosy_MockIndexer_HandleAccumulator__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

func Fuzz_Nosy_MockIndexer_Index__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockIndexer
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Index(ctx)
	})
}
