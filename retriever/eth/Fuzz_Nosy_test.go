package eth

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_chainClient_FetchBatchHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *chainClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var serviceManagerAddress common.Address
		fill_err = tp.Fill(&serviceManagerAddress)
		if fill_err != nil {
			return
		}
		var batchHeaderHash []byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.FetchBatchHeader(ctx, serviceManagerAddress, batchHeaderHash)
	})
}

// skipping Fuzz_Nosy_ChainClient_FetchBatchHeader__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/retriever/eth.ChainClient
