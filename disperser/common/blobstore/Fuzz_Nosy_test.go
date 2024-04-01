package blobstore

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

func Fuzz_Nosy_BlobMetadataStore_GetAllBlobMetadataByBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
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

		s.GetAllBlobMetadataByBatch(ctx, batchHeaderHash)
	})
}

func Fuzz_Nosy_BlobMetadataStore_GetBlobMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadataKey disperser.BlobKey
		fill_err = tp.Fill(&metadataKey)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlobMetadata(ctx, metadataKey)
	})
}

func Fuzz_Nosy_BlobMetadataStore_GetBlobMetadataByStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetBlobMetadataByStatus(ctx, status)
	})
}

func Fuzz_Nosy_BlobMetadataStore_GetBlobMetadataByStatusCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetBlobMetadataByStatusCount(ctx, status)
	})
}

func Fuzz_Nosy_BlobMetadataStore_GetBlobMetadataByStatusWithPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
		fill_err = tp.Fill(&s)
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
		if s == nil || exclusiveStartKey == nil {
			return
		}

		s.GetBlobMetadataByStatusWithPagination(ctx, status, limit, exclusiveStartKey)
	})
}

func Fuzz_Nosy_BlobMetadataStore_GetBlobMetadataInBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
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
		var blobIndex uint32
		fill_err = tp.Fill(&blobIndex)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlobMetadataInBatch(ctx, batchHeaderHash, blobIndex)
	})
}

func Fuzz_Nosy_BlobMetadataStore_IncrementNumRetries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
		fill_err = tp.Fill(&s)
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
		if s == nil || existingMetadata == nil {
			return
		}

		s.IncrementNumRetries(ctx, existingMetadata)
	})
}

func Fuzz_Nosy_BlobMetadataStore_QueueNewBlobMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobMetadata *disperser.BlobMetadata
		fill_err = tp.Fill(&blobMetadata)
		if fill_err != nil {
			return
		}
		if s == nil || blobMetadata == nil {
			return
		}

		s.QueueNewBlobMetadata(ctx, blobMetadata)
	})
}

func Fuzz_Nosy_BlobMetadataStore_SetBlobStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadataKey disperser.BlobKey
		fill_err = tp.Fill(&metadataKey)
		if fill_err != nil {
			return
		}
		var status disperser.BlobStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetBlobStatus(ctx, metadataKey, status)
	})
}

func Fuzz_Nosy_BlobMetadataStore_UpdateBlobMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlobMetadataStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadataKey disperser.BlobKey
		fill_err = tp.Fill(&metadataKey)
		if fill_err != nil {
			return
		}
		var updated *disperser.BlobMetadata
		fill_err = tp.Fill(&updated)
		if fill_err != nil {
			return
		}
		if s == nil || updated == nil {
			return
		}

		s.UpdateBlobMetadata(ctx, metadataKey, updated)
	})
}

func Fuzz_Nosy_SharedBlobStore_GetAllBlobMetadataByBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
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

		s.GetAllBlobMetadataByBatch(ctx, batchHeaderHash)
	})
}

func Fuzz_Nosy_SharedBlobStore_GetBlobContent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetBlobContent(ctx, blobHash)
	})
}

func Fuzz_Nosy_SharedBlobStore_GetBlobMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadataKey disperser.BlobKey
		fill_err = tp.Fill(&metadataKey)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlobMetadata(ctx, metadataKey)
	})
}

func Fuzz_Nosy_SharedBlobStore_GetBlobMetadataByStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobStatus disperser.BlobStatus
		fill_err = tp.Fill(&blobStatus)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlobMetadataByStatus(ctx, blobStatus)
	})
}

func Fuzz_Nosy_SharedBlobStore_GetBlobMetadataByStatusWithPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobStatus disperser.BlobStatus
		fill_err = tp.Fill(&blobStatus)
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
		if s == nil || exclusiveStartKey == nil {
			return
		}

		s.GetBlobMetadataByStatusWithPagination(ctx, blobStatus, limit, exclusiveStartKey)
	})
}

func Fuzz_Nosy_SharedBlobStore_GetBlobsByMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetBlobsByMetadata(ctx, metadata)
	})
}

func Fuzz_Nosy_SharedBlobStore_GetMetadataInBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
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
		var blobIndex uint32
		fill_err = tp.Fill(&blobIndex)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetMetadataInBatch(ctx, batchHeaderHash, blobIndex)
	})
}

func Fuzz_Nosy_SharedBlobStore_HandleBlobFailure__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
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
		if s == nil || metadata == nil {
			return
		}

		s.HandleBlobFailure(ctx, metadata, maxRetry)
	})
}

func Fuzz_Nosy_SharedBlobStore_IncrementBlobRetryCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
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
		if s == nil || existingMetadata == nil {
			return
		}

		s.IncrementBlobRetryCount(ctx, existingMetadata)
	})
}

func Fuzz_Nosy_SharedBlobStore_MarkBlobConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
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
		if s == nil || existingMetadata == nil || confirmationInfo == nil {
			return
		}

		s.MarkBlobConfirmed(ctx, existingMetadata, confirmationInfo)
	})
}

func Fuzz_Nosy_SharedBlobStore_MarkBlobFailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadataKey disperser.BlobKey
		fill_err = tp.Fill(&metadataKey)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarkBlobFailed(ctx, metadataKey)
	})
}

func Fuzz_Nosy_SharedBlobStore_MarkBlobFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadataKey disperser.BlobKey
		fill_err = tp.Fill(&metadataKey)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarkBlobFinalized(ctx, metadataKey)
	})
}

func Fuzz_Nosy_SharedBlobStore_MarkBlobInsufficientSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
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
		if s == nil || existingMetadata == nil || confirmationInfo == nil {
			return
		}

		s.MarkBlobInsufficientSignatures(ctx, existingMetadata, confirmationInfo)
	})
}

func Fuzz_Nosy_SharedBlobStore_MarkBlobProcessing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadataKey disperser.BlobKey
		fill_err = tp.Fill(&metadataKey)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarkBlobProcessing(ctx, metadataKey)
	})
}

func Fuzz_Nosy_SharedBlobStore_StoreBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SharedBlobStore
		fill_err = tp.Fill(&s)
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
		if s == nil || blob == nil {
			return
		}

		s.StoreBlob(ctx, blob, requestedAt)
	})
}

// skipping Fuzz_Nosy_SharedBlobStore_getBlobContentParallel__ because parameters include func, chan, or unsupported interface: chan<- github.com/Layr-Labs/eigenda/disperser/common/blobstore.blobResultOrError

func Fuzz_Nosy_MarshalBlobMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var metadata *disperser.BlobMetadata
		fill_err = tp.Fill(&metadata)
		if fill_err != nil {
			return
		}
		if metadata == nil {
			return
		}

		MarshalBlobMetadata(metadata)
	})
}

func Fuzz_Nosy_blobObjectKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blobHash string) {
		blobObjectKey(blobHash)
	})
}

func Fuzz_Nosy_convertToAttribMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blobStoreExclusiveStartKey *disperser.BlobStoreExclusiveStartKey
		fill_err = tp.Fill(&blobStoreExclusiveStartKey)
		if fill_err != nil {
			return
		}
		if blobStoreExclusiveStartKey == nil {
			return
		}

		convertToAttribMap(blobStoreExclusiveStartKey)
	})
}

func Fuzz_Nosy_getBlobHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blob *core.Blob
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		if blob == nil {
			return
		}

		getBlobHash(blob)
	})
}

func Fuzz_Nosy_getMetadataHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var requestedAt uint64
		fill_err = tp.Fill(&requestedAt)
		if fill_err != nil {
			return
		}
		var securityParams []*core.SecurityParam
		fill_err = tp.Fill(&securityParams)
		if fill_err != nil {
			return
		}

		getMetadataHash(requestedAt, securityParams)
	})
}
