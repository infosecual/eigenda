package traffic

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

func Fuzz_Nosy_TrafficGenerator_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		g, err := NewTrafficGenerator(config)
		if err != nil {
			return
		}
		g.Run()
	})
}

func Fuzz_Nosy_TrafficGenerator_StartTraffic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		g, err := NewTrafficGenerator(config)
		if err != nil {
			return
		}
		g.StartTraffic(ctx)
	})
}

func Fuzz_Nosy_TrafficGenerator_sendRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		g, err := NewTrafficGenerator(config)
		if err != nil {
			return
		}
		g.sendRequest(ctx, d3)
	})
}
