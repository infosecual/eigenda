package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_Assignment_GetIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Assignment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetIndices()
	})
}

func Fuzz_Nosy_BatchHeader_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *BatchHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Deserialize(d2)
	})
}

func Fuzz_Nosy_BatchHeader_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *BatchHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Encode()
	})
}

func Fuzz_Nosy_BatchHeader_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *BatchHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Serialize()
	})
}

func Fuzz_Nosy_BatchHeader_SetBatchRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *BatchHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var blobHeaders []*BlobHeader
		fill_err = tp.Fill(&blobHeaders)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.SetBatchRoot(blobHeaders)
	})
}

func Fuzz_Nosy_BlobHeader_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *BlobHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Deserialize(d2)
	})
}

func Fuzz_Nosy_BlobHeader_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *BlobHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Encode()
	})
}

func Fuzz_Nosy_BlobHeader_EncodedSizeAllQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobHeader
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.EncodedSizeAllQuorums()
	})
}

func Fuzz_Nosy_BlobHeader_GetQuorumBlobParamsHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *BlobHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.GetQuorumBlobParamsHash()
	})
}

func Fuzz_Nosy_BlobHeader_GetQuorumInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobHeader
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var quorum uint8
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetQuorumInfo(quorum)
	})
}

func Fuzz_Nosy_BlobHeader_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *BlobHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Serialize()
	})
}

func Fuzz_Nosy_G1Point_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		var p2 *G1Point
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil || p2 == nil {
			return
		}

		p := NewG1Point(x, y)
		p.Add(p2)
	})
}

func Fuzz_Nosy_G1Point_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		p := NewG1Point(x, y)
		p.Clone()
	})
}

func Fuzz_Nosy_G1Point_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		p := NewG1Point(x, y)
		p.Deserialize(d3)
	})
}

func Fuzz_Nosy_G1Point_GetOperatorID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		p := NewG1Point(x, y)
		p.GetOperatorID()
	})
}

func Fuzz_Nosy_G1Point_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		p := NewG1Point(x, y)
		p.Hash()
	})
}

func Fuzz_Nosy_G1Point_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		p := NewG1Point(x, y)
		p.Serialize()
	})
}

func Fuzz_Nosy_G1Point_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		var p2 *G1Point
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil || p2 == nil {
			return
		}

		p := NewG1Point(x, y)
		p.Sub(p2)
	})
}

func Fuzz_Nosy_G1Point_VerifyEquivalence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		var p2 *G2Point
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil || p2 == nil {
			return
		}

		p := NewG1Point(x, y)
		p.VerifyEquivalence(p2)
	})
}

func Fuzz_Nosy_G2Point_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *G2Point
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var p2 *G2Point
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p == nil || p2 == nil {
			return
		}

		p.Add(p2)
	})
}

func Fuzz_Nosy_G2Point_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *G2Point
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Clone()
	})
}

func Fuzz_Nosy_G2Point_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *G2Point
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Deserialize(d2)
	})
}

func Fuzz_Nosy_G2Point_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *G2Point
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Serialize()
	})
}

func Fuzz_Nosy_G2Point_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *G2Point
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var p2 *G2Point
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p == nil || p2 == nil {
			return
		}

		p.Sub(p2)
	})
}

func Fuzz_Nosy_KeyPair_MakePubkeyRegistrationData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var operatorAddress common.Address
		fill_err = tp.Fill(&operatorAddress)
		if fill_err != nil {
			return
		}

		k, err := GenRandomBlsKeys()
		if err != nil {
			return
		}
		k.MakePubkeyRegistrationData(operatorAddress)
	})
}

func Fuzz_Nosy_KeyPair_SignHashedToCurveMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1HashedMsg *G1Point
		fill_err = tp.Fill(&g1HashedMsg)
		if fill_err != nil {
			return
		}
		if g1HashedMsg == nil {
			return
		}

		k, err := GenRandomBlsKeys()
		if err != nil {
			return
		}
		k.SignHashedToCurveMessage(g1HashedMsg)
	})
}

func Fuzz_Nosy_KeyPair_SignMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var message [32]byte
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}

		k, err := GenRandomBlsKeys()
		if err != nil {
			return
		}
		k.SignMessage(message)
	})
}

func Fuzz_Nosy_SecurityParam_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SecurityParam
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.String()
	})
}

func Fuzz_Nosy_SecurityParam_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sp *SecurityParam
		fill_err = tp.Fill(&sp)
		if fill_err != nil {
			return
		}
		if sp == nil {
			return
		}

		sp.Validate()
	})
}

func Fuzz_Nosy_Signature_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Signature
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pubkey *G2Point
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var message [32]byte
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		if s == nil || pubkey == nil {
			return
		}

		s.Verify(pubkey, message)
	})
}

func Fuzz_Nosy_StdAssignmentCoordinator_CalculateChunkLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StdAssignmentCoordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var state *OperatorState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var blobLength uint
		fill_err = tp.Fill(&blobLength)
		if fill_err != nil {
			return
		}
		var targetNumChunks uint
		fill_err = tp.Fill(&targetNumChunks)
		if fill_err != nil {
			return
		}
		var param *SecurityParam
		fill_err = tp.Fill(&param)
		if fill_err != nil {
			return
		}
		if c == nil || state == nil || param == nil {
			return
		}

		c.CalculateChunkLength(state, blobLength, targetNumChunks, param)
	})
}

func Fuzz_Nosy_StdAssignmentCoordinator_GetAssignments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StdAssignmentCoordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var state *OperatorState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var blobLength uint
		fill_err = tp.Fill(&blobLength)
		if fill_err != nil {
			return
		}
		var info *BlobQuorumInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		if c == nil || state == nil || info == nil {
			return
		}

		c.GetAssignments(state, blobLength, info)
	})
}

func Fuzz_Nosy_StdAssignmentCoordinator_GetOperatorAssignment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StdAssignmentCoordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var state *OperatorState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var header *BlobHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var quorum uint8
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var id OperatorID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if c == nil || state == nil || header == nil {
			return
		}

		c.GetOperatorAssignment(state, header, quorum, id)
	})
}

func Fuzz_Nosy_StdAssignmentCoordinator_ValidateChunkLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StdAssignmentCoordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var state *OperatorState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var blobLength uint
		fill_err = tp.Fill(&blobLength)
		if fill_err != nil {
			return
		}
		var info *BlobQuorumInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		if c == nil || state == nil || info == nil {
			return
		}

		c.ValidateChunkLength(state, blobLength, info)
	})
}

// skipping Fuzz_Nosy_StdSignatureAggregator_AggregateSignatures__ because parameters include func, chan, or unsupported interface: chan github.com/Layr-Labs/eigenda/core.SignerMessage

func Fuzz_Nosy_shardValidator_UpdateOperatorID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *shardValidator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var operatorID OperatorID
		fill_err = tp.Fill(&operatorID)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.UpdateOperatorID(operatorID)
	})
}

// skipping Fuzz_Nosy_shardValidator_ValidateBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_shardValidator_VerifyBlobLengthWorker__ because parameters include func, chan, or unsupported interface: chan error

// skipping Fuzz_Nosy_shardValidator_universalVerifyWorker__ because parameters include func, chan, or unsupported interface: chan error

func Fuzz_Nosy_shardValidator_validateBlobQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *shardValidator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var quorumHeader *BlobQuorumInfo
		fill_err = tp.Fill(&quorumHeader)
		if fill_err != nil {
			return
		}
		var blob *BlobMessage
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		var operatorState *OperatorState
		fill_err = tp.Fill(&operatorState)
		if fill_err != nil {
			return
		}
		if v == nil || quorumHeader == nil || blob == nil || operatorState == nil {
			return
		}

		v.validateBlobQuorum(quorumHeader, blob, operatorState)
	})
}

// skipping Fuzz_Nosy_AssignmentCoordinator_CalculateChunkLength__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.AssignmentCoordinator

// skipping Fuzz_Nosy_AssignmentCoordinator_GetAssignments__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.AssignmentCoordinator

// skipping Fuzz_Nosy_AssignmentCoordinator_GetOperatorAssignment__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.AssignmentCoordinator

// skipping Fuzz_Nosy_AssignmentCoordinator_ValidateChunkLength__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.AssignmentCoordinator

func Fuzz_Nosy_BatchHeader_GetBatchHeaderHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h BatchHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.GetBatchHeaderHash()
	})
}

func Fuzz_Nosy_BlobHeader_GetBlobHeaderHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h BlobHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.GetBlobHeaderHash()
	})
}

// skipping Fuzz_Nosy_BlobRequestAuthenticator_AuthenticateBlobRequest__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.BlobRequestAuthenticator

// skipping Fuzz_Nosy_BlobRequestSigner_GetAccountID__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.BlobRequestSigner

// skipping Fuzz_Nosy_BlobRequestSigner_SignBlobRequest__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.BlobRequestSigner

func Fuzz_Nosy_Bundle_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bundle
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.Size()
	})
}

func Fuzz_Nosy_Bundles_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb Bundles
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}

		cb.Serialize()
	})
}

func Fuzz_Nosy_Bundles_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb Bundles
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}

		cb.Size()
	})
}

// skipping Fuzz_Nosy_ChainState_GetCurrentBlockNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.ChainState

// skipping Fuzz_Nosy_ChainState_GetOperatorState__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.ChainState

// skipping Fuzz_Nosy_ChainState_GetOperatorStateByOperator__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.ChainState

// skipping Fuzz_Nosy_IndexedChainState_GetIndexedOperatorState__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.IndexedChainState

// skipping Fuzz_Nosy_IndexedChainState_Start__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.IndexedChainState

func Fuzz_Nosy_OperatorID_Hex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id OperatorID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.Hex()
	})
}

func Fuzz_Nosy_OperatorSocket_GetDispersalSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, nodeIP string, dispersalPort string, retrievalPort string) {
		s := MakeOperatorSocket(nodeIP, dispersalPort, retrievalPort)
		s.GetDispersalSocket()
	})
}

func Fuzz_Nosy_OperatorSocket_GetRetrievalSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, nodeIP string, dispersalPort string, retrievalPort string) {
		s := MakeOperatorSocket(nodeIP, dispersalPort, retrievalPort)
		s.GetRetrievalSocket()
	})
}

func Fuzz_Nosy_OperatorSocket_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, nodeIP string, dispersalPort string, retrievalPort string) {
		s := MakeOperatorSocket(nodeIP, dispersalPort, retrievalPort)
		s.String()
	})
}

// skipping Fuzz_Nosy_ShardValidator_UpdateOperatorID__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.ShardValidator

// skipping Fuzz_Nosy_ShardValidator_ValidateBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.ShardValidator

// skipping Fuzz_Nosy_SignatureAggregator_AggregateSignatures__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.SignatureAggregator

// skipping Fuzz_Nosy_Transactor_BatchOperatorIDToAddress__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_BuildConfirmBatchTxn__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_CalculateOperatorChurnApprovalDigestHash__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_ConfirmBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_DeregisterOperator__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetBlockStaleMeasure__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetCurrentBlockNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetCurrentQuorumBitmapByOperatorId__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetNumberOfRegisteredOperatorForQuorum__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetOperatorSetParams__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetOperatorStakes__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetOperatorStakesForQuorums__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetQuorumBitmapForOperatorsAtBlockNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetQuorumCount__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetQuorumSecurityParams__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetRegisteredQuorumIdsForOperator__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetRequiredQuorumNumbers__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_GetStoreDurationBlocks__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_OperatorIDToAddress__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_RegisterOperator__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_RegisterOperatorWithChurn__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_StakeRegistry__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_UpdateOperatorSocket__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

// skipping Fuzz_Nosy_Transactor_WeightOfOperatorForQuorum__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

func Fuzz_Nosy_ComputeSignatoryRecordHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		var nonSignerKeys []*G1Point
		fill_err = tp.Fill(&nonSignerKeys)
		if fill_err != nil {
			return
		}

		ComputeSignatoryRecordHash(referenceBlockNumber, nonSignerKeys)
	})
}

func Fuzz_Nosy_GetSignedPercentage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var state *OperatorState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var quorum uint8
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var stakeAmount *big.Int
		fill_err = tp.Fill(&stakeAmount)
		if fill_err != nil {
			return
		}
		if state == nil || stakeAmount == nil {
			return
		}

		GetSignedPercentage(state, quorum, stakeAmount)
	})
}

func Fuzz_Nosy_HashBatchHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchHeader contractEigenDAServiceManager.IEigenDAServiceManagerBatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}

		HashBatchHeader(batchHeader)
	})
}

func Fuzz_Nosy_ParseOperatorSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, socket string) {
		ParseOperatorSocket(socket)
	})
}

// skipping Fuzz_Nosy_encode__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_extractIPAndPorts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		extractIPAndPorts(s)
	})
}

func Fuzz_Nosy_nextPowerOf2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d uint64) {
		nextPowerOf2(d)
	})
}

func Fuzz_Nosy_roundUpDivide__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint, b uint) {
		roundUpDivide(a, b)
	})
}
