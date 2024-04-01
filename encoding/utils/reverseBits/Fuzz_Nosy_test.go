package reverseBits

import (
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

func Fuzz_Nosy_ReverseBitsLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, length uint32, value uint32) {
		ReverseBitsLimited(length, value)
	})
}

func Fuzz_Nosy_bitIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		bitIndex(v)
	})
}

func Fuzz_Nosy_reverseBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b uint32) {
		reverseBits(b)
	})
}
