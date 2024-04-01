package encoding

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

func Fuzz_Nosy_Frame_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, d2 []byte) {
		c, err := Decode(b)
		if err != nil {
			return
		}
		c.Deserialize(d2)
	})
}

func Fuzz_Nosy_Frame_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		f1, err := Decode(b)
		if err != nil {
			return
		}
		f1.Encode()
	})
}

func Fuzz_Nosy_Frame_Length__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		f1, err := Decode(b)
		if err != nil {
			return
		}
		f1.Length()
	})
}

func Fuzz_Nosy_Frame_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		c, err := Decode(b)
		if err != nil {
			return
		}
		c.Serialize()
	})
}

func Fuzz_Nosy_Frame_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		f1, err := Decode(b)
		if err != nil {
			return
		}
		f1.Size()
	})
}

func Fuzz_Nosy_G1Commitment_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *G1Commitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Deserialize(d2)
	})
}

func Fuzz_Nosy_G1Commitment_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *G1Commitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Serialize()
	})
}

func Fuzz_Nosy_G1Commitment_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *G1Commitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_G2Commitment_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *G2Commitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Deserialize(d2)
	})
}

func Fuzz_Nosy_G2Commitment_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *G2Commitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Serialize()
	})
}

func Fuzz_Nosy_G2Commitment_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *G2Commitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.UnmarshalJSON(d2)
	})
}

// skipping Fuzz_Nosy_Decoder_Decode__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/encoding.Decoder

func Fuzz_Nosy_EncodingParams_ChunkDegree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numSys uint64, numPar uint64, dataSize uint64) {
		p := ParamsFromSysPar(numSys, numPar, dataSize)
		p.ChunkDegree()
	})
}

func Fuzz_Nosy_EncodingParams_NumEvaluations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numSys uint64, numPar uint64, dataSize uint64) {
		p := ParamsFromSysPar(numSys, numPar, dataSize)
		p.NumEvaluations()
	})
}

func Fuzz_Nosy_EncodingParams_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numSys uint64, numPar uint64, dataSize uint64) {
		p := ParamsFromSysPar(numSys, numPar, dataSize)
		p.Validate()
	})
}

// skipping Fuzz_Nosy_Prover_EncodeAndProve__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/encoding.Prover

// skipping Fuzz_Nosy_Verifier_UniversalVerifySubBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/encoding.Verifier

// skipping Fuzz_Nosy_Verifier_VerifyBlobLength__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/encoding.Verifier

// skipping Fuzz_Nosy_Verifier_VerifyCommitEquivalenceBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/encoding.Verifier

// skipping Fuzz_Nosy_Verifier_VerifyFrames__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/encoding.Verifier

func Fuzz_Nosy_GetBlobLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blobSize uint) {
		GetBlobLength(blobSize)
	})
}

func Fuzz_Nosy_GetBlobSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blobLength uint) {
		GetBlobSize(blobLength)
	})
}

func Fuzz_Nosy_GetEncodedBlobLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blobLength uint, quorumThreshold uint8, advThreshold uint8) {
		GetEncodedBlobLength(blobLength, quorumThreshold, advThreshold)
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

// skipping Fuzz_Nosy_encode__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_roundUpDivide__ because parameters include func, chan, or unsupported interface: T
