package prover

import (
	"testing"

	"github.com/Layr-Labs/eigenda/encoding"
	"github.com/Layr-Labs/eigenda/encoding/fft"
	"github.com/Layr-Labs/eigenda/encoding/kzg"
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

func Fuzz_Nosy_ParametrizedProver_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ParametrizedProver
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var polyFr []fr.Element
		fill_err = tp.Fill(&polyFr)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Commit(polyFr)
	})
}

func Fuzz_Nosy_ParametrizedProver_Decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ParametrizedProver
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var frames []encoding.Frame
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
		if g == nil {
			return
		}

		g.Decode(frames, indices, maxInputSize)
	})
}

func Fuzz_Nosy_ParametrizedProver_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ParametrizedProver
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var inputFr []fr.Element
		fill_err = tp.Fill(&inputFr)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Encode(inputFr)
	})
}

func Fuzz_Nosy_ParametrizedProver_EncodeBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *ParametrizedProver
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var inputBytes []byte
		fill_err = tp.Fill(&inputBytes)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.EncodeBytes(inputBytes)
	})
}

func Fuzz_Nosy_ParametrizedProver_GetSlicesCoeff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ParametrizedProver
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var polyFr []fr.Element
		fill_err = tp.Fill(&polyFr)
		if fill_err != nil {
			return
		}
		var dimE uint64
		fill_err = tp.Fill(&dimE)
		if fill_err != nil {
			return
		}
		var j uint64
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		var l uint64
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetSlicesCoeff(polyFr, dimE, j, l)
	})
}

func Fuzz_Nosy_ParametrizedProver_ProveAllCosetThreads__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ParametrizedProver
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var polyFr []fr.Element
		fill_err = tp.Fill(&polyFr)
		if fill_err != nil {
			return
		}
		var numChunks uint64
		fill_err = tp.Fill(&numChunks)
		if fill_err != nil {
			return
		}
		var chunkLen uint64
		fill_err = tp.Fill(&chunkLen)
		if fill_err != nil {
			return
		}
		var numWorker uint64
		fill_err = tp.Fill(&numWorker)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ProveAllCosetThreads(polyFr, numChunks, chunkLen, numWorker)
	})
}

// skipping Fuzz_Nosy_ParametrizedProver_proofWorker__ because parameters include func, chan, or unsupported interface: <-chan uint64

func Fuzz_Nosy_Prover_Decode__(f *testing.F) {
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

		p, err := NewProver(config, loadG2Points)
		if err != nil {
			return
		}
		p.Decode(chunks, indices, params, maxInputSize)
	})
}

func Fuzz_Nosy_Prover_EncodeAndProve__(f *testing.F) {
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
		var d3 []byte
		fill_err = tp.Fill(&d3)
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

		e, err := NewProver(config, loadG2Points)
		if err != nil {
			return
		}
		e.EncodeAndProve(d3, params)
	})
}

func Fuzz_Nosy_Prover_GetKzgEncoder__(f *testing.F) {
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

		g, err := NewProver(config, loadG2Points)
		if err != nil {
			return
		}
		g.GetKzgEncoder(params)
	})
}

func Fuzz_Nosy_Prover_PreloadAllEncoders__(f *testing.F) {
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
		if config == nil {
			return
		}

		g, err := NewProver(config, loadG2Points)
		if err != nil {
			return
		}
		g.PreloadAllEncoders()
	})
}

func Fuzz_Nosy_Prover_newProver__(f *testing.F) {
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

		g, err := NewProver(config, loadG2Points)
		if err != nil {
			return
		}
		g.newProver(params)
	})
}

func Fuzz_Nosy_SRSTable_GetSubTables__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tableDir string
		fill_err = tp.Fill(&tableDir)
		if fill_err != nil {
			return
		}
		var s1 []bn254.G1Affine
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var numWorker uint64
		fill_err = tp.Fill(&numWorker)
		if fill_err != nil {
			return
		}
		var numChunks uint64
		fill_err = tp.Fill(&numChunks)
		if fill_err != nil {
			return
		}
		var chunkLen uint64
		fill_err = tp.Fill(&chunkLen)
		if fill_err != nil {
			return
		}

		p, err := NewSRSTable(tableDir, s1, numWorker)
		if err != nil {
			return
		}
		p.GetSubTables(numChunks, chunkLen)
	})
}

func Fuzz_Nosy_SRSTable_Precompute__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tableDir string
		fill_err = tp.Fill(&tableDir)
		if fill_err != nil {
			return
		}
		var s1 []bn254.G1Affine
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var n3 uint64
		fill_err = tp.Fill(&n3)
		if fill_err != nil {
			return
		}
		var dim uint64
		fill_err = tp.Fill(&dim)
		if fill_err != nil {
			return
		}
		var dimE uint64
		fill_err = tp.Fill(&dimE)
		if fill_err != nil {
			return
		}
		var l uint64
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var m uint64
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var filePath string
		fill_err = tp.Fill(&filePath)
		if fill_err != nil {
			return
		}
		var n9 uint64
		fill_err = tp.Fill(&n9)
		if fill_err != nil {
			return
		}

		p, err := NewSRSTable(tableDir, s1, n3)
		if err != nil {
			return
		}
		p.Precompute(dim, dimE, l, m, filePath, n9)
	})
}

func Fuzz_Nosy_SRSTable_PrecomputeSubTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tableDir string
		fill_err = tp.Fill(&tableDir)
		if fill_err != nil {
			return
		}
		var s1 []bn254.G1Affine
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var numWorker uint64
		fill_err = tp.Fill(&numWorker)
		if fill_err != nil {
			return
		}
		var fs *fft.FFTSettings
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var m uint64
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dim uint64
		fill_err = tp.Fill(&dim)
		if fill_err != nil {
			return
		}
		var dimE uint64
		fill_err = tp.Fill(&dimE)
		if fill_err != nil {
			return
		}
		var j uint64
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		var l uint64
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if fs == nil {
			return
		}

		p, err := NewSRSTable(tableDir, s1, numWorker)
		if err != nil {
			return
		}
		p.PrecomputeSubTable(fs, m, dim, dimE, j, l)
	})
}

func Fuzz_Nosy_SRSTable_TableReaderThreads__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tableDir string
		fill_err = tp.Fill(&tableDir)
		if fill_err != nil {
			return
		}
		var s1 []bn254.G1Affine
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var n3 uint64
		fill_err = tp.Fill(&n3)
		if fill_err != nil {
			return
		}
		var filePath string
		fill_err = tp.Fill(&filePath)
		if fill_err != nil {
			return
		}
		var dimE uint64
		fill_err = tp.Fill(&dimE)
		if fill_err != nil {
			return
		}
		var l uint64
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var n7 uint64
		fill_err = tp.Fill(&n7)
		if fill_err != nil {
			return
		}

		p, err := NewSRSTable(tableDir, s1, n3)
		if err != nil {
			return
		}
		p.TableReaderThreads(filePath, dimE, l, n7)
	})
}

func Fuzz_Nosy_SRSTable_TableWriter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tableDir string
		fill_err = tp.Fill(&tableDir)
		if fill_err != nil {
			return
		}
		var s1 []bn254.G1Affine
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var numWorker uint64
		fill_err = tp.Fill(&numWorker)
		if fill_err != nil {
			return
		}
		var fftPoints [][]bn254.G1Affine
		fill_err = tp.Fill(&fftPoints)
		if fill_err != nil {
			return
		}
		var dimE uint64
		fill_err = tp.Fill(&dimE)
		if fill_err != nil {
			return
		}
		var filePath string
		fill_err = tp.Fill(&filePath)
		if fill_err != nil {
			return
		}

		p, err := NewSRSTable(tableDir, s1, numWorker)
		if err != nil {
			return
		}
		p.TableWriter(fftPoints, dimE, filePath)
	})
}

// skipping Fuzz_Nosy_SRSTable_precomputeWorker__ because parameters include func, chan, or unsupported interface: <-chan uint64

// skipping Fuzz_Nosy_SRSTable_readWorker__ because parameters include func, chan, or unsupported interface: <-chan github.com/Layr-Labs/eigenda/encoding/kzg/prover.Boundary

func Fuzz_Nosy_CeilIntPowerOf2Num__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d uint64) {
		CeilIntPowerOf2Num(d)
	})
}

func Fuzz_Nosy_GetAllPrecomputedSrsMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, tableDir string) {
		GetAllPrecomputedSrsMap(tableDir)
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
