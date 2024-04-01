package kzg

import (
	"bufio"
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

func Fuzz_Nosy_KZGSettings_CommitToPoly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var srs *SRS
		fill_err = tp.Fill(&srs)
		if fill_err != nil {
			return
		}
		var coeffs []fr.Element
		fill_err = tp.Fill(&coeffs)
		if fill_err != nil {
			return
		}
		if fs == nil || srs == nil {
			return
		}

		ks, err := NewKZGSettings(fs, srs)
		if err != nil {
			return
		}
		ks.CommitToPoly(coeffs)
	})
}

func Fuzz_Nosy_CLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
		CLIFlags(envPrefix)
	})
}

func Fuzz_Nosy_GenerateTestingSetup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, secret string, n uint64) {
		GenerateTestingSetup(secret, n)
	})
}

func Fuzz_Nosy_ReadDesiredBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader *bufio.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}
		var numBytesToRead uint64
		fill_err = tp.Fill(&numBytesToRead)
		if fill_err != nil {
			return
		}
		if reader == nil {
			return
		}

		ReadDesiredBytes(reader, numBytesToRead)
	})
}

func Fuzz_Nosy_ReadG1PointSection__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filepath string, from uint64, to uint64, numWorker uint64) {
		ReadG1PointSection(filepath, from, to, numWorker)
	})
}

func Fuzz_Nosy_ReadG1Points__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filepath string, n uint64, numWorker uint64) {
		ReadG1Points(filepath, n, numWorker)
	})
}

func Fuzz_Nosy_ReadG2PointSection__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filepath string, from uint64, to uint64, numWorker uint64) {
		ReadG2PointSection(filepath, from, to, numWorker)
	})
}

func Fuzz_Nosy_ReadG2Points__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filepath string, n uint64, numWorker uint64) {
		ReadG2Points(filepath, n, numWorker)
	})
}

// skipping Fuzz_Nosy_readG1Worker__ because parameters include func, chan, or unsupported interface: chan<- error

// skipping Fuzz_Nosy_readG2Worker__ because parameters include func, chan, or unsupported interface: chan<- error
