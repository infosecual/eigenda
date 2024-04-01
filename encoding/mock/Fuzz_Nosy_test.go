package encoding

import (
	"testing"

	"github.com/Layr-Labs/eigenda/encoding"
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

func Fuzz_Nosy_MockEncoder_Decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *MockEncoder
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var chunks []*encoding.Frame
		fill_err = tp.Fill(&chunks)
		if fill_err != nil {
			return
		}
		var indices []uint
		fill_err = tp.Fill(&indices)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var maxInputSize uint64
		fill_err = tp.Fill(&maxInputSize)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Decode(chunks, indices, params, maxInputSize)
	})
}

func Fuzz_Nosy_MockEncoder_EncodeAndProve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *MockEncoder
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.EncodeAndProve(d2, params)
	})
}

func Fuzz_Nosy_MockEncoder_UniversalVerifySubBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *MockEncoder
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var samples []encoding.Sample
		fill_err = tp.Fill(&samples)
		if fill_err != nil {
			return
		}
		var numBlobs int
		fill_err = tp.Fill(&numBlobs)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.UniversalVerifySubBatch(params, samples, numBlobs)
	})
}

func Fuzz_Nosy_MockEncoder_VerifyBlobLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *MockEncoder
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var commitments encoding.BlobCommitments
		fill_err = tp.Fill(&commitments)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.VerifyBlobLength(commitments)
	})
}

func Fuzz_Nosy_MockEncoder_VerifyCommitEquivalenceBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *MockEncoder
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var commitments []encoding.BlobCommitments
		fill_err = tp.Fill(&commitments)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.VerifyCommitEquivalenceBatch(commitments)
	})
}

func Fuzz_Nosy_MockEncoder_VerifyFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *MockEncoder
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var chunks []*encoding.Frame
		fill_err = tp.Fill(&chunks)
		if fill_err != nil {
			return
		}
		var indices []uint
		fill_err = tp.Fill(&indices)
		if fill_err != nil {
			return
		}
		var commitments encoding.BlobCommitments
		fill_err = tp.Fill(&commitments)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.VerifyFrames(chunks, indices, commitments, params)
	})
}
