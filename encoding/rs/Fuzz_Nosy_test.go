package rs

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

func Fuzz_Nosy_Encoder_Decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var verbose bool
		fill_err = tp.Fill(&verbose)
		if fill_err != nil {
			return
		}
		var frames []Frame
		fill_err = tp.Fill(&frames)
		if fill_err != nil {
			return
		}
		var indices []uint64
		fill_err = tp.Fill(&indices)
		if fill_err != nil {
			return
		}
		var maxInputSize uint64
		fill_err = tp.Fill(&maxInputSize)
		if fill_err != nil {
			return
		}

		g, err := NewEncoder(params, verbose)
		if err != nil {
			return
		}
		g.Decode(frames, indices, maxInputSize)
	})
}

func Fuzz_Nosy_Encoder_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var verbose bool
		fill_err = tp.Fill(&verbose)
		if fill_err != nil {
			return
		}
		var inputFr []fr.Element
		fill_err = tp.Fill(&inputFr)
		if fill_err != nil {
			return
		}

		g, err := NewEncoder(params, verbose)
		if err != nil {
			return
		}
		g.Encode(inputFr)
	})
}

func Fuzz_Nosy_Encoder_EncodeBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var verbose bool
		fill_err = tp.Fill(&verbose)
		if fill_err != nil {
			return
		}
		var inputBytes []byte
		fill_err = tp.Fill(&inputBytes)
		if fill_err != nil {
			return
		}

		g, err := NewEncoder(params, verbose)
		if err != nil {
			return
		}
		g.EncodeBytes(inputBytes)
	})
}

func Fuzz_Nosy_Encoder_ExtendPolyEval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var verbose bool
		fill_err = tp.Fill(&verbose)
		if fill_err != nil {
			return
		}
		var coeffs []fr.Element
		fill_err = tp.Fill(&coeffs)
		if fill_err != nil {
			return
		}

		g, err := NewEncoder(params, verbose)
		if err != nil {
			return
		}
		g.ExtendPolyEval(coeffs)
	})
}

func Fuzz_Nosy_Encoder_GetInterpolationPolyCoeff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var verbose bool
		fill_err = tp.Fill(&verbose)
		if fill_err != nil {
			return
		}
		var chunk []fr.Element
		fill_err = tp.Fill(&chunk)
		if fill_err != nil {
			return
		}
		var k uint32
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		g, err := NewEncoder(params, verbose)
		if err != nil {
			return
		}
		g.GetInterpolationPolyCoeff(chunk, k)
	})
}

func Fuzz_Nosy_Encoder_GetInterpolationPolyEval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var verbose bool
		fill_err = tp.Fill(&verbose)
		if fill_err != nil {
			return
		}
		var interpolationPoly []fr.Element
		fill_err = tp.Fill(&interpolationPoly)
		if fill_err != nil {
			return
		}
		var j uint32
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		g, err := NewEncoder(params, verbose)
		if err != nil {
			return
		}
		g.GetInterpolationPolyEval(interpolationPoly, j)
	})
}

func Fuzz_Nosy_Encoder_MakeFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var verbose bool
		fill_err = tp.Fill(&verbose)
		if fill_err != nil {
			return
		}
		var polyEvals []fr.Element
		fill_err = tp.Fill(&polyEvals)
		if fill_err != nil {
			return
		}

		g, err := NewEncoder(params, verbose)
		if err != nil {
			return
		}
		g.MakeFrames(polyEvals)
	})
}

// skipping Fuzz_Nosy_Encoder_interpolyWorker__ because parameters include func, chan, or unsupported interface: <-chan github.com/Layr-Labs/eigenda/encoding/rs.JobRequest

func Fuzz_Nosy_Frame_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		f1, err := Decode(b)
		if err != nil {
			return
		}
		f1.Encode()
	})
}

func Fuzz_Nosy_EncodingParams_ChunkDegree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numChunks uint64, chunkLen uint64) {
		p := ParamsFromMins(numChunks, chunkLen)
		p.ChunkDegree()
	})
}

func Fuzz_Nosy_EncodingParams_NumEvaluations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numChunks uint64, chunkLen uint64) {
		p := ParamsFromMins(numChunks, chunkLen)
		p.NumEvaluations()
	})
}

func Fuzz_Nosy_EncodingParams_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numChunks uint64, chunkLen uint64) {
		p := ParamsFromMins(numChunks, chunkLen)
		p.Validate()
	})
}

func Fuzz_Nosy_GetLeadingCosetIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64, numChunks uint64) {
		GetLeadingCosetIndex(i, numChunks)
	})
}

func Fuzz_Nosy_GetNumElement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dataLen uint64, CS int) {
		GetNumElement(dataLen, CS)
	})
}

func Fuzz_Nosy_GetNumSys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dataSize uint64, chunkLen uint64) {
		GetNumSys(dataSize, chunkLen)
	})
}

func Fuzz_Nosy_NextPowerOf2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d uint64) {
		NextPowerOf2(d)
	})
}

func Fuzz_Nosy_RoundUpDivision__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64) {
		RoundUpDivision(a, b)
	})
}

func Fuzz_Nosy_ToByteArray__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dataFr []fr.Element
		fill_err = tp.Fill(&dataFr)
		if fill_err != nil {
			return
		}
		var maxDataSize uint64
		fill_err = tp.Fill(&maxDataSize)
		if fill_err != nil {
			return
		}

		ToByteArray(dataFr, maxDataSize)
	})
}

func Fuzz_Nosy_ToFrArray__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		ToFrArray(d1)
	})
}
