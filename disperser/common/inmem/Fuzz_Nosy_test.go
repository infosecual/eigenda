package inmem

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/core"
	"github.com/Layr-Labs/eigenda/disperser"
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

func Fuzz_Nosy_BlobStore_GetAllBlobMetadataByBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
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
		if q == nil {
			return
		}

		q.GetAllBlobMetadataByBatch(ctx, batchHeaderHash)
	})
}

func Fuzz_Nosy_BlobStore_GetBlobContent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobHash string
		fill_err = tp.Fill(&blobHash)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.GetBlobContent(ctx, blobHash)
	})
}

func Fuzz_Nosy_BlobStore_GetBlobMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.GetBlobMetadata(ctx, blobKey)
	})
}

func Fuzz_Nosy_BlobStore_GetBlobMetadataByStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var status disperser.BlobStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.GetBlobMetadataByStatus(ctx, status)
	})
}

func Fuzz_Nosy_BlobStore_GetBlobMetadataByStatusWithPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var status disperser.BlobStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var limit int32
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var exclusiveStartKey *disperser.BlobStoreExclusiveStartKey
		fill_err = tp.Fill(&exclusiveStartKey)
		if fill_err != nil {
			return
		}
		if q == nil || exclusiveStartKey == nil {
			return
		}

		q.GetBlobMetadataByStatusWithPagination(ctx, status, limit, exclusiveStartKey)
	})
}

func Fuzz_Nosy_BlobStore_GetBlobsByMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadata []*disperser.BlobMetadata
		fill_err = tp.Fill(&metadata)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.GetBlobsByMetadata(ctx, metadata)
	})
}

func Fuzz_Nosy_BlobStore_GetMetadataInBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
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
		var blobIndex uint32
		fill_err = tp.Fill(&blobIndex)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.GetMetadataInBatch(ctx, batchHeaderHash, blobIndex)
	})
}

func Fuzz_Nosy_BlobStore_HandleBlobFailure__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadata *disperser.BlobMetadata
		fill_err = tp.Fill(&metadata)
		if fill_err != nil {
			return
		}
		var maxRetry uint
		fill_err = tp.Fill(&maxRetry)
		if fill_err != nil {
			return
		}
		if q == nil || metadata == nil {
			return
		}

		q.HandleBlobFailure(ctx, metadata, maxRetry)
	})
}

func Fuzz_Nosy_BlobStore_IncrementBlobRetryCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var existingMetadata *disperser.BlobMetadata
		fill_err = tp.Fill(&existingMetadata)
		if fill_err != nil {
			return
		}
		if q == nil || existingMetadata == nil {
			return
		}

		q.IncrementBlobRetryCount(ctx, existingMetadata)
	})
}

func Fuzz_Nosy_BlobStore_MarkBlobConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var existingMetadata *disperser.BlobMetadata
		fill_err = tp.Fill(&existingMetadata)
		if fill_err != nil {
			return
		}
		var confirmationInfo *disperser.ConfirmationInfo
		fill_err = tp.Fill(&confirmationInfo)
		if fill_err != nil {
			return
		}
		if q == nil || existingMetadata == nil || confirmationInfo == nil {
			return
		}

		q.MarkBlobConfirmed(ctx, existingMetadata, confirmationInfo)
	})
}

func Fuzz_Nosy_BlobStore_MarkBlobFailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.MarkBlobFailed(ctx, blobKey)
	})
}

func Fuzz_Nosy_BlobStore_MarkBlobFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.MarkBlobFinalized(ctx, blobKey)
	})
}

func Fuzz_Nosy_BlobStore_MarkBlobInsufficientSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var existingMetadata *disperser.BlobMetadata
		fill_err = tp.Fill(&existingMetadata)
		if fill_err != nil {
			return
		}
		var confirmationInfo *disperser.ConfirmationInfo
		fill_err = tp.Fill(&confirmationInfo)
		if fill_err != nil {
			return
		}
		if q == nil || existingMetadata == nil || confirmationInfo == nil {
			return
		}

		q.MarkBlobInsufficientSignatures(ctx, existingMetadata, confirmationInfo)
	})
}

func Fuzz_Nosy_BlobStore_MarkBlobProcessing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobKey disperser.BlobKey
		fill_err = tp.Fill(&blobKey)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.MarkBlobProcessing(ctx, blobKey)
	})
}

func Fuzz_Nosy_BlobStore_StoreBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blob *core.Blob
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		var requestedAt uint64
		fill_err = tp.Fill(&requestedAt)
		if fill_err != nil {
			return
		}
		if q == nil || blob == nil {
			return
		}

		q.StoreBlob(ctx, blob, requestedAt)
	})
}

func Fuzz_Nosy_BlobStore_getNewBlobHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *BlobStore
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.getNewBlobHash()
	})
}

func Fuzz_Nosy_getMetadataHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, requestedAt uint64) {
		getMetadataHash(requestedAt)
	})
}
