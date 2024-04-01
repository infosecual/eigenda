package weth_test

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

func Fuzz_Nosy_Accumulator_DeserializeObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Accumulator
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

func Fuzz_Nosy_Accumulator_InitializeObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Accumulator
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

// skipping Fuzz_Nosy_Accumulator_SerializeObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.AccumulatorObject

// skipping Fuzz_Nosy_Accumulator_UpdateObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.AccumulatorObject

func Fuzz_Nosy_Filterer_FilterFastMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Filterer
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

func Fuzz_Nosy_Filterer_FilterHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Filterer
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

func Fuzz_Nosy_Filterer_GetSyncPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Filterer
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

func Fuzz_Nosy_Filterer_SetSyncPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Filterer
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
