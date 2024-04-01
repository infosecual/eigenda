package batcher

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/core"
	"github.com/Layr-Labs/eigenda/disperser"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/wealdtech/go-merkletree"
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

func Fuzz_Nosy_Batcher_HandleSingleBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Batcher
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.HandleSingleBatch(ctx)
	})
}

func Fuzz_Nosy_Batcher_ProcessConfirmedBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Batcher
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var receiptOrErr *ReceiptOrErr
		fill_err = tp.Fill(&receiptOrErr)
		if fill_err != nil {
			return
		}
		if b == nil || receiptOrErr == nil {
			return
		}

		b.ProcessConfirmedBatch(ctx, receiptOrErr)
	})
}

func Fuzz_Nosy_Batcher_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Batcher
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Start(ctx)
	})
}

func Fuzz_Nosy_Batcher_getBatchID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Batcher
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txReceipt *types.Receipt
		fill_err = tp.Fill(&txReceipt)
		if fill_err != nil {
			return
		}
		if b == nil || txReceipt == nil {
			return
		}

		b.getBatchID(ctx, txReceipt)
	})
}

func Fuzz_Nosy_Batcher_handleFailure__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Batcher
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobMetadatas []*disperser.BlobMetadata
		fill_err = tp.Fill(&blobMetadatas)
		if fill_err != nil {
			return
		}
		var reason FailReason
		fill_err = tp.Fill(&reason)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.handleFailure(ctx, blobMetadatas, reason)
	})
}

func Fuzz_Nosy_Batcher_parseBatchIDFromReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Batcher
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var txReceipt *types.Receipt
		fill_err = tp.Fill(&txReceipt)
		if fill_err != nil {
			return
		}
		if b == nil || txReceipt == nil {
			return
		}

		b.parseBatchIDFromReceipt(txReceipt)
	})
}

func Fuzz_Nosy_Batcher_signalLiveness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Batcher
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.signalLiveness()
	})
}

func Fuzz_Nosy_Batcher_updateConfirmationInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Batcher
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batchData confirmationMetadata
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var txnReceipt *types.Receipt
		fill_err = tp.Fill(&txnReceipt)
		if fill_err != nil {
			return
		}
		if b == nil || txnReceipt == nil {
			return
		}

		b.updateConfirmationInfo(ctx, batchData, txnReceipt)
	})
}

func Fuzz_Nosy_EncodingStreamer_CreateBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamer
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.CreateBatch()
	})
}

func Fuzz_Nosy_EncodingStreamer_MarkBlobPendingConfirmation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamer
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var metadata *disperser.BlobMetadata
		fill_err = tp.Fill(&metadata)
		if fill_err != nil {
			return
		}
		if e == nil || metadata == nil {
			return
		}

		e.MarkBlobPendingConfirmation(metadata)
	})
}

func Fuzz_Nosy_EncodingStreamer_ProcessEncodedBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamer
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var result EncodingResultOrStatus
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ProcessEncodedBlobs(ctx, result)
	})
}

func Fuzz_Nosy_EncodingStreamer_RemoveEncodedBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamer
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var metadata *disperser.BlobMetadata
		fill_err = tp.Fill(&metadata)
		if fill_err != nil {
			return
		}
		if e == nil || metadata == nil {
			return
		}

		e.RemoveEncodedBlob(metadata)
	})
}

// skipping Fuzz_Nosy_EncodingStreamer_RequestEncoding__ because parameters include func, chan, or unsupported interface: chan github.com/Layr-Labs/eigenda/disperser/batcher.EncodingResultOrStatus

// skipping Fuzz_Nosy_EncodingStreamer_RequestEncodingForBlob__ because parameters include func, chan, or unsupported interface: chan github.com/Layr-Labs/eigenda/disperser/batcher.EncodingResultOrStatus

func Fuzz_Nosy_EncodingStreamer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamer
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Start(ctx)
	})
}

func Fuzz_Nosy_EncodingStreamer_dedupRequests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamer
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var metadatas []*disperser.BlobMetadata
		fill_err = tp.Fill(&metadatas)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.dedupRequests(metadatas, referenceBlockNumber)
	})
}

func Fuzz_Nosy_EncodingStreamer_getOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamer
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadatas []*disperser.BlobMetadata
		fill_err = tp.Fill(&metadatas)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.getOperatorState(ctx, metadatas, blockNumber)
	})
}

func Fuzz_Nosy_EncodingStreamer_validateMetadataQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamer
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var metadatas []*disperser.BlobMetadata
		fill_err = tp.Fill(&metadatas)
		if fill_err != nil {
			return
		}
		var state *core.IndexedOperatorState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if e == nil || state == nil {
			return
		}

		e.validateMetadataQuorums(metadatas, state)
	})
}

func Fuzz_Nosy_EncodingStreamerMetrics_UpdateEncodedBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EncodingStreamerMetrics
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var size uint64
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.UpdateEncodedBlobs(count, size)
	})
}

func Fuzz_Nosy_FinalizerMetrics_IncrementNumBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalizerMetrics
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var state string
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IncrementNumBlobs(state)
	})
}

func Fuzz_Nosy_FinalizerMetrics_ObserveLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalizerMetrics
		fill_err = tp.Fill(&f1)
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
		if f1 == nil {
			return
		}

		f1.ObserveLatency(stage, latencyMs)
	})
}

func Fuzz_Nosy_FinalizerMetrics_UpdateLastSeenFinalizedBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalizerMetrics
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var blockNumber uint64
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.UpdateLastSeenFinalizedBlock(blockNumber)
	})
}

func Fuzz_Nosy_FinalizerMetrics_UpdateNumBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalizerMetrics
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var state string
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.UpdateNumBlobs(state, count)
	})
}

func Fuzz_Nosy_Metrics_IncrementBatchCount__(f *testing.F) {
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
		var size int64
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.IncrementBatchCount(size)
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

		g.ObserveLatency(stage, latencyMs)
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
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Start(ctx)
	})
}

func Fuzz_Nosy_Metrics_UpdateAttestation__(f *testing.F) {
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
		var operatorCount map[uint8]int
		fill_err = tp.Fill(&operatorCount)
		if fill_err != nil {
			return
		}
		var signerCount map[uint8]int
		fill_err = tp.Fill(&signerCount)
		if fill_err != nil {
			return
		}
		var quorumResults map[uint8]*core.QuorumResult
		fill_err = tp.Fill(&quorumResults)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.UpdateAttestation(operatorCount, signerCount, quorumResults)
	})
}

func Fuzz_Nosy_Metrics_UpdateBatchError__(f *testing.F) {
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
		var errType FailReason
		fill_err = tp.Fill(&errType)
		if fill_err != nil {
			return
		}
		var numBlobs int
		fill_err = tp.Fill(&numBlobs)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.UpdateBatchError(errType, numBlobs)
	})
}

func Fuzz_Nosy_Metrics_UpdateCompletedBlob__(f *testing.F) {
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
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		var status disperser.BlobStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.UpdateCompletedBlob(size, status)
	})
}

func Fuzz_Nosy_TxnManagerMetrics_IncrementTxnCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxnManagerMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var state string
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.IncrementTxnCount(state)
	})
}

func Fuzz_Nosy_TxnManagerMetrics_ObserveLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxnManagerMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var latencyMs float64
		fill_err = tp.Fill(&latencyMs)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.ObserveLatency(latencyMs)
	})
}

func Fuzz_Nosy_TxnManagerMetrics_UpdateGasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxnManagerMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var gasUsed uint64
		fill_err = tp.Fill(&gasUsed)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UpdateGasUsed(gasUsed)
	})
}

func Fuzz_Nosy_TxnManagerMetrics_UpdateSpeedUps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxnManagerMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var speedUps int
		fill_err = tp.Fill(&speedUps)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UpdateSpeedUps(speedUps)
	})
}

func Fuzz_Nosy_TxnManagerMetrics_UpdateTxQueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxnManagerMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var txQueue int
		fill_err = tp.Fill(&txQueue)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UpdateTxQueue(txQueue)
	})
}

func Fuzz_Nosy_encodedBlobStore_DeleteEncodingRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.DeleteEncodingRequest(blobKey, quorumID)
	})
}

func Fuzz_Nosy_encodedBlobStore_DeleteEncodingResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.DeleteEncodingResult(blobKey, quorumID)
	})
}

func Fuzz_Nosy_encodedBlobStore_GetEncodedResultSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.GetEncodedResultSize()
	})
}

func Fuzz_Nosy_encodedBlobStore_GetEncodingResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.GetEncodingResult(blobKey, quorumID)
	})
}

func Fuzz_Nosy_encodedBlobStore_GetNewAndDeleteStaleEncodingResults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.GetNewAndDeleteStaleEncodingResults(blockNumber)
	})
}

func Fuzz_Nosy_encodedBlobStore_HasEncodingRequested__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.HasEncodingRequested(blobKey, quorumID, referenceBlockNumber)
	})
}

func Fuzz_Nosy_encodedBlobStore_MarkEncodedResultPendingConfirmation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.MarkEncodedResultPendingConfirmation(blobKey, quorumID)
	})
}

func Fuzz_Nosy_encodedBlobStore_PutEncodingRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.PutEncodingRequest(blobKey, quorumID)
	})
}

func Fuzz_Nosy_encodedBlobStore_PutEncodingResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *encodedBlobStore
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var result *EncodingResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if e == nil || result == nil {
			return
		}

		e.PutEncodingResult(result)
	})
}

func Fuzz_Nosy_finalizer_FinalizeBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *finalizer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.FinalizeBlobs(ctx)
	})
}

func Fuzz_Nosy_finalizer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *finalizer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Start(ctx)
	})
}

func Fuzz_Nosy_finalizer_getLatestFinalizedBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *finalizer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.getLatestFinalizedBlock(ctx)
	})
}

func Fuzz_Nosy_finalizer_getTransactionBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *finalizer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.getTransactionBlockNumber(ctx, hash)
	})
}

func Fuzz_Nosy_finalizer_updateBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *finalizer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadatas []*disperser.BlobMetadata
		fill_err = tp.Fill(&metadatas)
		if fill_err != nil {
			return
		}
		var lastFinalBlock uint64
		fill_err = tp.Fill(&lastFinalBlock)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.updateBlobs(ctx, metadatas, lastFinalBlock)
	})
}

func Fuzz_Nosy_txnManager_ProcessTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *txnManager
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *TxnRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if t1 == nil || req == nil {
			return
		}

		t1.ProcessTransaction(ctx, req)
	})
}

func Fuzz_Nosy_txnManager_ReceiptChan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *txnManager
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.ReceiptChan()
	})
}

func Fuzz_Nosy_txnManager_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *txnManager
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Start(ctx)
	})
}

func Fuzz_Nosy_txnManager_ensureAnyTransactionEvaled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *txnManager
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txs []*transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.ensureAnyTransactionEvaled(ctx, txs)
	})
}

func Fuzz_Nosy_txnManager_monitorTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *txnManager
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *TxnRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if t1 == nil || req == nil {
			return
		}

		t1.monitorTransaction(ctx, req)
	})
}

func Fuzz_Nosy_txnManager_speedUpTxn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *txnManager
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var tag string
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		if t1 == nil || tx == nil {
			return
		}

		t1.speedUpTxn(ctx, tx, tag)
	})
}

// skipping Fuzz_Nosy_Finalizer_FinalizeBlobs__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/batcher.Finalizer

// skipping Fuzz_Nosy_Finalizer_Start__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/batcher.Finalizer

// skipping Fuzz_Nosy_TxnManager_ProcessTransaction__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/batcher.TxnManager

// skipping Fuzz_Nosy_TxnManager_ReceiptChan__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/batcher.TxnManager

// skipping Fuzz_Nosy_TxnManager_Start__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/batcher.TxnManager

func Fuzz_Nosy_getChunksSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var result *EncodingResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if result == nil {
			return
		}

		getChunksSize(result)
	})
}

func Fuzz_Nosy_isBlobAttested__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signedQuorums map[uint8]*core.QuorumResult
		fill_err = tp.Fill(&signedQuorums)
		if fill_err != nil {
			return
		}
		var header *core.BlobHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		isBlobAttested(signedQuorums, header)
	})
}

func Fuzz_Nosy_numBlobsAttested__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signedQuorums map[uint8]*core.QuorumResult
		fill_err = tp.Fill(&signedQuorums)
		if fill_err != nil {
			return
		}
		var headers []*core.BlobHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		numBlobsAttested(signedQuorums, headers)
	})
}

func Fuzz_Nosy_serializeProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var proof *merkletree.Proof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if proof == nil {
			return
		}

		serializeProof(proof)
	})
}
