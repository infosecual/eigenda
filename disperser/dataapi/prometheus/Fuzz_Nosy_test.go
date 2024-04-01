package prometheus

import (
	"context"
	"testing"
	"time"

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

func Fuzz_Nosy_prometheusApi_QueryRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var query string
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var start time.Time
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end time.Time
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		var step time.Duration
		fill_err = tp.Fill(&step)
		if fill_err != nil {
			return
		}

		p, err := NewApi(config)
		if err != nil {
			return
		}
		p.QueryRange(ctx, query, start, end, step)
	})
}

// skipping Fuzz_Nosy_Api_QueryRange__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/prometheus.Api
