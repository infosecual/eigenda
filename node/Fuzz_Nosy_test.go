package node

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/api/grpc/node"
	"github.com/Layr-Labs/eigenda/core"
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

func Fuzz_Nosy_Metrics_AcceptBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var status string
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var batchSize int64
		fill_err = tp.Fill(&batchSize)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.AcceptBatches(status, batchSize)
	})
}

func Fuzz_Nosy_Metrics_AcceptBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var quorumId uint8
		fill_err = tp.Fill(&quorumId)
		if fill_err != nil {
			return
		}
		var blobSize int64
		fill_err = tp.Fill(&blobSize)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.AcceptBlobs(quorumId, blobSize)
	})
}

func Fuzz_Nosy_Metrics_ObserveLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		var stage string
		fill_err = tp.Fill(&stage)
		if fill_err != nil {
			return
		}
		var latencyMs float64
		fill_err = tp.Fill(&latencyMs)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.ObserveLatency(method, stage, latencyMs)
	})
}

func Fuzz_Nosy_Metrics_RecordRPCRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		var status string
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.RecordRPCRequest(method, status)
	})
}

func Fuzz_Nosy_Metrics_RecordSocketAddressChange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.RecordSocketAddressChange()
	})
}

func Fuzz_Nosy_Metrics_RemoveNCurrentBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var numBatches int
		fill_err = tp.Fill(&numBatches)
		if fill_err != nil {
			return
		}
		var totalBatchSize int64
		fill_err = tp.Fill(&totalBatchSize)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.RemoveNCurrentBatch(numBatches, totalBatchSize)
	})
}

func Fuzz_Nosy_Metrics_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Start()
	})
}

func Fuzz_Nosy_Node_ProcessBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var header *core.BatchHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var blobs []*core.BlobMessage
		fill_err = tp.Fill(&blobs)
		if fill_err != nil {
			return
		}
		var rawBlobs []*node.Blob
		fill_err = tp.Fill(&rawBlobs)
		if fill_err != nil {
			return
		}
		if n == nil || header == nil {
			return
		}

		n.ProcessBatch(ctx, header, blobs, rawBlobs)
	})
}

func Fuzz_Nosy_Node_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Start(ctx)
	})
}

func Fuzz_Nosy_Node_ValidateBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var header *core.BatchHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var blobs []*core.BlobMessage
		fill_err = tp.Fill(&blobs)
		if fill_err != nil {
			return
		}
		if n == nil || header == nil {
			return
		}

		n.ValidateBatch(ctx, header, blobs)
	})
}

func Fuzz_Nosy_Node_checkCurrentNodeIp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.checkCurrentNodeIp(ctx)
	})
}

func Fuzz_Nosy_Node_checkRegisteredNodeIpOnChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.checkRegisteredNodeIpOnChain(ctx)
	})
}

func Fuzz_Nosy_Node_expireLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.expireLoop()
	})
}

func Fuzz_Nosy_Node_updateSocketAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var newSocketAddr string
		fill_err = tp.Fill(&newSocketAddr)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.updateSocketAddress(ctx, newSocketAddr)
	})
}

// skipping Fuzz_Nosy_Operator_getQuorumIdsToRegister__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/core.Transactor

func Fuzz_Nosy_Store_DeleteExpiredEntries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Store
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var currentTimeUnixSec int64
		fill_err = tp.Fill(&currentTimeUnixSec)
		if fill_err != nil {
			return
		}
		var timeLimitSec uint64
		fill_err = tp.Fill(&timeLimitSec)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.DeleteExpiredEntries(currentTimeUnixSec, timeLimitSec)
	})
}

func Fuzz_Nosy_Store_DeleteKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Store
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var keys *[][]byte
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}
		if s == nil || keys == nil {
			return
		}

		s.DeleteKeys(ctx, keys)
	})
}

func Fuzz_Nosy_Store_GetBatchHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Store
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batchHeaderHash [32]byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBatchHeader(ctx, batchHeaderHash)
	})
}

func Fuzz_Nosy_Store_GetBlobHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Store
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batchHeaderHash [32]byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}
		var blobIndex int
		fill_err = tp.Fill(&blobIndex)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlobHeader(ctx, batchHeaderHash, blobIndex)
	})
}

func Fuzz_Nosy_Store_GetChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Store
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batchHeaderHash [32]byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}
		var blobIndex int
		fill_err = tp.Fill(&blobIndex)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetChunks(ctx, batchHeaderHash, blobIndex, quorumID)
	})
}

func Fuzz_Nosy_Store_HasKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Store
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.HasKey(ctx, key)
	})
}

func Fuzz_Nosy_Store_StoreBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Store
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var header *core.BatchHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var blobs []*core.BlobMessage
		fill_err = tp.Fill(&blobs)
		if fill_err != nil {
			return
		}
		var blobsProto []*node.Blob
		fill_err = tp.Fill(&blobsProto)
		if fill_err != nil {
			return
		}
		if s == nil || header == nil {
			return
		}

		s.StoreBatch(ctx, header, blobs, blobsProto)
	})
}

func Fuzz_Nosy_Store_deleteNBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Store
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var currentTimeUnixSec int64
		fill_err = tp.Fill(&currentTimeUnixSec)
		if fill_err != nil {
			return
		}
		var numBatches int
		fill_err = tp.Fill(&numBatches)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.deleteNBatches(currentTimeUnixSec, numBatches)
	})
}

func Fuzz_Nosy_churnerClient_Churn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *churnerClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorAddress string
		fill_err = tp.Fill(&operatorAddress)
		if fill_err != nil {
			return
		}
		var keyPair *core.KeyPair
		fill_err = tp.Fill(&keyPair)
		if fill_err != nil {
			return
		}
		var quorumIDs []uint8
		fill_err = tp.Fill(&quorumIDs)
		if fill_err != nil {
			return
		}
		if c == nil || keyPair == nil {
			return
		}

		c.Churn(ctx, operatorAddress, keyPair, quorumIDs)
	})
}

// skipping Fuzz_Nosy_ChurnerClient_Churn__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/node.ChurnerClient

// skipping Fuzz_Nosy_DB_Delete__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/node.DB

// skipping Fuzz_Nosy_DB_DeleteBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/node.DB

// skipping Fuzz_Nosy_DB_Get__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/node.DB

// skipping Fuzz_Nosy_DB_NewIterator__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/node.DB

// skipping Fuzz_Nosy_DB_Put__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/node.DB

// skipping Fuzz_Nosy_DB_WriteBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/node.DB

func Fuzz_Nosy_DecodeBatchExpirationKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		DecodeBatchExpirationKey(key)
	})
}

func Fuzz_Nosy_EncodeBatchExpirationKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, expirationTime int64) {
		EncodeBatchExpirationKey(expirationTime)
	})
}

func Fuzz_Nosy_EncodeBatchHeaderKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchHeaderHash [32]byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}

		EncodeBatchHeaderKey(batchHeaderHash)
	})
}

func Fuzz_Nosy_EncodeBlobHeaderKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchHeaderHash [32]byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}
		var blobIndex int
		fill_err = tp.Fill(&blobIndex)
		if fill_err != nil {
			return
		}

		EncodeBlobHeaderKey(batchHeaderHash, blobIndex)
	})
}

func Fuzz_Nosy_EncodeBlobHeaderKeyPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchHeaderHash [32]byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}

		EncodeBlobHeaderKeyPrefix(batchHeaderHash)
	})
}

func Fuzz_Nosy_EncodeBlobKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchHeaderHash [32]byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}
		var blobIndex int
		fill_err = tp.Fill(&blobIndex)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}

		EncodeBlobKey(batchHeaderHash, blobIndex, quorumID)
	})
}

// skipping Fuzz_Nosy_SocketAddress__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common/pubip.Provider

func Fuzz_Nosy_copyBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, src []byte) {
		copyBytes(src)
	})
}

func Fuzz_Nosy_decodeChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		decodeChunks(d1)
	})
}

func Fuzz_Nosy_encodeChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chunks [][]byte
		fill_err = tp.Fill(&chunks)
		if fill_err != nil {
			return
		}

		encodeChunks(chunks)
	})
}
