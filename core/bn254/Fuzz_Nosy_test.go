package bn254

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254"
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

func Fuzz_Nosy_CheckG1AndG2DiscreteLogEquality__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pointG1 *bn254.G1Affine
		fill_err = tp.Fill(&pointG1)
		if fill_err != nil {
			return
		}
		var pointG2 *bn254.G2Affine
		fill_err = tp.Fill(&pointG2)
		if fill_err != nil {
			return
		}
		if pointG1 == nil || pointG2 == nil {
			return
		}

		CheckG1AndG2DiscreteLogEquality(pointG1, pointG2)
	})
}

func Fuzz_Nosy_SerializeG1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *bn254.G1Affine
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		SerializeG1(p)
	})
}

func Fuzz_Nosy_SerializeG2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *bn254.G2Affine
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		SerializeG2(p)
	})
}

func Fuzz_Nosy_VerifySig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sig *bn254.G1Affine
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var pubkey *bn254.G2Affine
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var msgBytes [32]byte
		fill_err = tp.Fill(&msgBytes)
		if fill_err != nil {
			return
		}
		if sig == nil || pubkey == nil {
			return
		}

		VerifySig(sig, pubkey, msgBytes)
	})
}
