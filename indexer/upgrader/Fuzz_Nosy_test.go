package eigendaUpgrader

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
		var headers []indexer.Header
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
		var header indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		u.GetLatestUpgrade(header)
	})
}
