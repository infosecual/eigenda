package verifier

import (
	"testing"

	"github.com/Layr-Labs/eigenda/encoding"
	"github.com/Layr-Labs/eigenda/encoding/kzg"
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

func Fuzz_Nosy_ParametrizedVerifier_VerifyFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ParametrizedVerifier
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var commit *bn254.G1Affine
		fill_err = tp.Fill(&commit)
		if fill_err != nil {
			return
		}
		var f3 *encoding.Frame
		fill_err = tp.Fill(&f3)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if v == nil || commit == nil || f3 == nil {
			return
		}

		v.VerifyFrame(commit, f3, index)
	})
}

func Fuzz_Nosy_Verifier_BatchVerifyCommitEquivalence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var commitmentsPair []CommitmentPair
		fill_err = tp.Fill(&commitmentsPair)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		group, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		group.BatchVerifyCommitEquivalence(commitmentsPair)
	})
}

func Fuzz_Nosy_Verifier_Decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
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
		if config == nil {
			return
		}

		v, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		v.Decode(chunks, indices, params, maxInputSize)
	})
}

func Fuzz_Nosy_Verifier_GetKzgVerifier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		g, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		g.GetKzgVerifier(params)
	})
}

func Fuzz_Nosy_Verifier_NewKzgVerifier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		g, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		g.NewKzgVerifier(params)
	})
}

func Fuzz_Nosy_Verifier_UniversalVerify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var samples []Sample
		fill_err = tp.Fill(&samples)
		if fill_err != nil {
			return
		}
		var m int
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		v, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		v.UniversalVerify(params, samples, m)
	})
}

func Fuzz_Nosy_Verifier_UniversalVerifySubBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var samplesCore []encoding.Sample
		fill_err = tp.Fill(&samplesCore)
		if fill_err != nil {
			return
		}
		var numBlobs int
		fill_err = tp.Fill(&numBlobs)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		v, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		v.UniversalVerifySubBatch(params, samplesCore, numBlobs)
	})
}

func Fuzz_Nosy_Verifier_VerifyBlobLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var commitments encoding.BlobCommitments
		fill_err = tp.Fill(&commitments)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		v, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		v.VerifyBlobLength(commitments)
	})
}

func Fuzz_Nosy_Verifier_VerifyCommit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var lengthCommit *bn254.G2Affine
		fill_err = tp.Fill(&lengthCommit)
		if fill_err != nil {
			return
		}
		var legnthProof *bn254.G2Affine
		fill_err = tp.Fill(&legnthProof)
		if fill_err != nil {
			return
		}
		var length uint64
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if config == nil || lengthCommit == nil || legnthProof == nil {
			return
		}

		v, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		v.VerifyCommit(lengthCommit, legnthProof, length)
	})
}

func Fuzz_Nosy_Verifier_VerifyCommitEquivalenceBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var commitments []encoding.BlobCommitments
		fill_err = tp.Fill(&commitments)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		v, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		v.VerifyCommitEquivalenceBatch(commitments)
	})
}

func Fuzz_Nosy_Verifier_VerifyFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var frames []*encoding.Frame
		fill_err = tp.Fill(&frames)
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
		if config == nil {
			return
		}

		v, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		v.VerifyFrames(frames, indices, commitments, params)
	})
}

func Fuzz_Nosy_Verifier_newKzgVerifier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *kzg.KzgConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var loadG2Points bool
		fill_err = tp.Fill(&loadG2Points)
		if fill_err != nil {
			return
		}
		var params encoding.EncodingParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		g, err := NewVerifier(config, loadG2Points)
		if err != nil {
			return
		}
		g.newKzgVerifier(params)
	})
}

func Fuzz_Nosy_CreateRandomnessVector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		CreateRandomnessVector(n)
	})
}

func Fuzz_Nosy_toUint64Array__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chunkIndices []uint
		fill_err = tp.Fill(&chunkIndices)
		if fill_err != nil {
			return
		}

		toUint64Array(chunkIndices)
	})
}
