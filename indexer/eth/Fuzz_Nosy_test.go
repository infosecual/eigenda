package eth

import (
	"context"
	"math/big"
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

func Fuzz_Nosy_HeaderService_PullLatestHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HeaderService
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var finalized bool
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.PullLatestHeader(finalized)
	})
}

func Fuzz_Nosy_HeaderService_PullNewHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HeaderService
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var lastHeader *indexer.Header
		fill_err = tp.Fill(&lastHeader)
		if fill_err != nil {
			return
		}
		if h == nil || lastHeader == nil {
			return
		}

		h.PullNewHeaders(lastHeader)
	})
}

func Fuzz_Nosy_HeaderService_getHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HeaderService
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if h == nil || number == nil {
			return
		}

		h.getHeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_HeaderService_headersByRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HeaderService
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.headersByRange(ctx, startHeight, count)
	})
}

func Fuzz_Nosy_toBlockNumArg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if number == nil {
			return
		}

		toBlockNumArg(number)
	})
}
