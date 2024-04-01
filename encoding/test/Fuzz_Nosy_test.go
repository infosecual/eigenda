package main

import (
	"testing"

	"github.com/Layr-Labs/eigenda/encoding"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
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

func Fuzz_Nosy_SampleFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frames []encoding.Frame
		fill_err = tp.Fill(&frames)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}

		SampleFrames(frames, num)
	})
}

func Fuzz_Nosy_printFr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d []fr.Element
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		printFr(d)
	})
}
