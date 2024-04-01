package mock

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

func Fuzz_Nosy_ChurnerClient_Churn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChurnerClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorAddress string
		fill_err = tp.Fill(&operatorAddress)
		if fill_err != nil {
			return
		}
		var keyPair *core.KeyPair
		fill_err = tp.Fill(&keyPair)
		if fill_err != nil {
			return
		}
		var quorumIDs []uint8
		fill_err = tp.Fill(&quorumIDs)
		if fill_err != nil {
			return
		}
		if c == nil || keyPair == nil {
			return
		}

		c.Churn(ctx, operatorAddress, keyPair, quorumIDs)
	})
}
