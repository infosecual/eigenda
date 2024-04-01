package ratelimit

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/common"
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

func Fuzz_Nosy_rateLimiter_AllowRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *rateLimiter
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params []common.RequestParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.AllowRequest(ctx, params)
	})
}

func Fuzz_Nosy_rateLimiter_checkAllowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *rateLimiter
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params common.RequestParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.checkAllowed(ctx, params)
	})
}

func Fuzz_Nosy_rateLimiter_updateBucketParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *rateLimiter
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params []common.RequestParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var updatedBucketParams []*common.RateBucketParams
		fill_err = tp.Fill(&updatedBucketParams)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.updateBucketParams(ctx, params, updatedBucketParams)
	})
}

func Fuzz_Nosy_RatelimiterCLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string, flagPrefix string) {
		RatelimiterCLIFlags(envPrefix, flagPrefix)
	})
}
