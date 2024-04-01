package toeplitz

import (
	"testing"

	"github.com/Layr-Labs/eigenda/encoding/fft"
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

func Fuzz_Nosy_Circular_GetCoeff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		c := NewCircular(v, fs)
		c.GetCoeff()
	})
}

func Fuzz_Nosy_Circular_GetFFTCoeff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		c := NewCircular(v, fs)
		c.GetFFTCoeff()
	})
}

func Fuzz_Nosy_Circular_Multiply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var x []fr.Element
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		c := NewCircular(v, fs)
		c.Multiply(x)
	})
}

func Fuzz_Nosy_Toeplitz_DirectMultiply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var x []fr.Element
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		t1, err := NewToeplitz(v, fs)
		if err != nil {
			return
		}
		t1.DirectMultiply(x)
	})
}

func Fuzz_Nosy_Toeplitz_ExtendCircularVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		t1, err := NewToeplitz(v, fs)
		if err != nil {
			return
		}
		t1.ExtendCircularVec()
	})
}

func Fuzz_Nosy_Toeplitz_FromColVToRowV__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var cv []fr.Element
		fill_err = tp.Fill(&cv)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		t1, err := NewToeplitz(v, fs)
		if err != nil {
			return
		}
		t1.FromColVToRowV(cv)
	})
}

func Fuzz_Nosy_Toeplitz_GetCoeff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		t1, err := NewToeplitz(v, fs)
		if err != nil {
			return
		}
		t1.GetCoeff()
	})
}

func Fuzz_Nosy_Toeplitz_GetFFTCoeff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		t1, err := NewToeplitz(v, fs)
		if err != nil {
			return
		}
		t1.GetFFTCoeff()
	})
}

func Fuzz_Nosy_Toeplitz_GetMatDim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		t1, err := NewToeplitz(v, fs)
		if err != nil {
			return
		}
		t1.GetMatDim()
	})
}

func Fuzz_Nosy_Toeplitz_Multiply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v []fr.Element
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var x []fr.Element
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		t1, err := NewToeplitz(v, fs)
		if err != nil {
			return
		}
		t1.Multiply(x)
	})
}
