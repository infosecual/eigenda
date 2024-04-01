package inmem

import (
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

func Fuzz_Nosy_HeaderStore_AddHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		h := NewHeaderStore()
		h.AddHeaders(headers)
	})
}

// skipping Fuzz_Nosy_HeaderStore_AttachObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.AccumulatorObject

func Fuzz_Nosy_HeaderStore_GetLatestHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, finalized bool) {
		h := NewHeaderStore()
		h.GetLatestHeader(finalized)
	})
}

// skipping Fuzz_Nosy_HeaderStore_GetLatestObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

// skipping Fuzz_Nosy_HeaderStore_GetObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

func Fuzz_Nosy_HeaderStore_getHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		h := NewHeaderStore()
		h.getHeader(header)
	})
}

func Fuzz_Nosy_HeaderStore_getHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, number uint64) {
		h := NewHeaderStore()
		h.getHeaderByNumber(number)
	})
}

func Fuzz_Nosy_AddPayloads__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		var payloads Payloads
		fill_err = tp.Fill(&payloads)
		if fill_err != nil {
			return
		}

		AddPayloads(headers, payloads)
	})
}

// skipping Fuzz_Nosy_encode__ because parameters include func, chan, or unsupported interface: any
