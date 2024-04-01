package fft

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bn254"
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

func Fuzz_Nosy_FFTSettings_FFT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var vals []fr.Element
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var inv bool
		fill_err = tp.Fill(&inv)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.FFT(vals, inv)
	})
}

func Fuzz_Nosy_FFTSettings_FFTG1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var vals []bn254.G1Affine
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var inv bool
		fill_err = tp.Fill(&inv)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.FFTG1(vals, inv)
	})
}

func Fuzz_Nosy_FFTSettings_InplaceFFT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var vals []fr.Element
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var out []fr.Element
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var inv bool
		fill_err = tp.Fill(&inv)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.InplaceFFT(vals, out, inv)
	})
}

// skipping Fuzz_Nosy_FFTSettings_RecoverPolyFromSamples__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/encoding/fft.ZeroPolyFn

func Fuzz_Nosy_FFTSettings_ShiftPoly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var poly []fr.Element
		fill_err = tp.Fill(&poly)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.ShiftPoly(poly)
	})
}

func Fuzz_Nosy_FFTSettings_UnshiftPoly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var poly []fr.Element
		fill_err = tp.Fill(&poly)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.UnshiftPoly(poly)
	})
}

func Fuzz_Nosy_FFTSettings_ZeroPolyViaMultiplication__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var missingIndices []uint64
		fill_err = tp.Fill(&missingIndices)
		if fill_err != nil {
			return
		}
		var length uint64
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.ZeroPolyViaMultiplication(missingIndices, length)
	})
}

func Fuzz_Nosy_FFTSettings__fft__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var vals []fr.Element
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var valsOffset uint64
		fill_err = tp.Fill(&valsOffset)
		if fill_err != nil {
			return
		}
		var valsStride uint64
		fill_err = tp.Fill(&valsStride)
		if fill_err != nil {
			return
		}
		var rootsOfUnity []fr.Element
		fill_err = tp.Fill(&rootsOfUnity)
		if fill_err != nil {
			return
		}
		var rootsOfUnityStride uint64
		fill_err = tp.Fill(&rootsOfUnityStride)
		if fill_err != nil {
			return
		}
		var out []fr.Element
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs._fft(vals, valsOffset, valsStride, rootsOfUnity, rootsOfUnityStride, out)
	})
}

func Fuzz_Nosy_FFTSettings__fftG1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var vals []bn254.G1Affine
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var valsOffset uint64
		fill_err = tp.Fill(&valsOffset)
		if fill_err != nil {
			return
		}
		var valsStride uint64
		fill_err = tp.Fill(&valsStride)
		if fill_err != nil {
			return
		}
		var rootsOfUnity []fr.Element
		fill_err = tp.Fill(&rootsOfUnity)
		if fill_err != nil {
			return
		}
		var rootsOfUnityStride uint64
		fill_err = tp.Fill(&rootsOfUnityStride)
		if fill_err != nil {
			return
		}
		var out []bn254.G1Affine
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs._fftG1(vals, valsOffset, valsStride, rootsOfUnity, rootsOfUnityStride, out)
	})
}

func Fuzz_Nosy_FFTSettings_makeZeroPolyMulLeaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var dst []fr.Element
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var indices []uint64
		fill_err = tp.Fill(&indices)
		if fill_err != nil {
			return
		}
		var domainStride uint64
		fill_err = tp.Fill(&domainStride)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.makeZeroPolyMulLeaf(dst, indices, domainStride)
	})
}

func Fuzz_Nosy_FFTSettings_reduceLeaves__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var scratch []fr.Element
		fill_err = tp.Fill(&scratch)
		if fill_err != nil {
			return
		}
		var dst []fr.Element
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var ps [][]fr.Element
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.reduceLeaves(scratch, dst, ps)
	})
}

func Fuzz_Nosy_FFTSettings_simpleFT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var vals []fr.Element
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var valsOffset uint64
		fill_err = tp.Fill(&valsOffset)
		if fill_err != nil {
			return
		}
		var valsStride uint64
		fill_err = tp.Fill(&valsStride)
		if fill_err != nil {
			return
		}
		var rootsOfUnity []fr.Element
		fill_err = tp.Fill(&rootsOfUnity)
		if fill_err != nil {
			return
		}
		var rootsOfUnityStride uint64
		fill_err = tp.Fill(&rootsOfUnityStride)
		if fill_err != nil {
			return
		}
		var out []fr.Element
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.simpleFT(vals, valsOffset, valsStride, rootsOfUnity, rootsOfUnityStride, out)
	})
}

func Fuzz_Nosy_FFTSettings_simpleFTG1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxScale uint8
		fill_err = tp.Fill(&maxScale)
		if fill_err != nil {
			return
		}
		var vals []bn254.G1Affine
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var valsOffset uint64
		fill_err = tp.Fill(&valsOffset)
		if fill_err != nil {
			return
		}
		var valsStride uint64
		fill_err = tp.Fill(&valsStride)
		if fill_err != nil {
			return
		}
		var rootsOfUnity []fr.Element
		fill_err = tp.Fill(&rootsOfUnity)
		if fill_err != nil {
			return
		}
		var rootsOfUnityStride uint64
		fill_err = tp.Fill(&rootsOfUnityStride)
		if fill_err != nil {
			return
		}
		var out []bn254.G1Affine
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		fs := NewFFTSettings(maxScale)
		fs.simpleFTG1(vals, valsOffset, valsStride, rootsOfUnity, rootsOfUnityStride, out)
	})
}

func Fuzz_Nosy_EvalPolyAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst *fr.Element
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var coeffs []fr.Element
		fill_err = tp.Fill(&coeffs)
		if fill_err != nil {
			return
		}
		var x *fr.Element
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if dst == nil || x == nil {
			return
		}

		EvalPolyAt(dst, coeffs, x)
	})
}

func Fuzz_Nosy_IsPowerOfTwo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		IsPowerOfTwo(v)
	})
}

func Fuzz_Nosy_expandRootOfUnity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rootOfUnity *fr.Element
		fill_err = tp.Fill(&rootOfUnity)
		if fill_err != nil {
			return
		}
		if rootOfUnity == nil {
			return
		}

		expandRootOfUnity(rootOfUnity)
	})
}

func Fuzz_Nosy_nextPowOf2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		nextPowOf2(v)
	})
}

func Fuzz_Nosy_padPoly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out []fr.Element
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var poly []fr.Element
		fill_err = tp.Fill(&poly)
		if fill_err != nil {
			return
		}

		padPoly(out, poly)
	})
}
