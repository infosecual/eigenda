package disperser

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/api/grpc/disperser"
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

func Fuzz_Nosy_BlobMetadata_GetBlobKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlobMetadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlobKey()
	})
}

func Fuzz_Nosy_BlobMetadata_IsConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlobMetadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.IsConfirmed()
	})
}

func Fuzz_Nosy_LocalEncoderClient_EncodeBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *LocalEncoderClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var encodingParams encoding.EncodingParams
		fill_err = tp.Fill(&encodingParams)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.EncodeBlob(ctx, d3, encodingParams)
	})
}

func Fuzz_Nosy_Metrics_HandleAccountRateLimitedRequest__(f *testing.F) {
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
		var quorum string
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var blobBytes int
		fill_err = tp.Fill(&blobBytes)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.HandleAccountRateLimitedRequest(quorum, blobBytes, method)
	})
}

func Fuzz_Nosy_Metrics_HandleBlobStoreFailedRequest__(f *testing.F) {
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
		var quorum string
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var blobBytes int
		fill_err = tp.Fill(&blobBytes)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.HandleBlobStoreFailedRequest(quorum, blobBytes, method)
	})
}

func Fuzz_Nosy_Metrics_HandleFailedRequest__(f *testing.F) {
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
		var statusCode string
		fill_err = tp.Fill(&statusCode)
		if fill_err != nil {
			return
		}
		var quorum string
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var blobBytes int
		fill_err = tp.Fill(&blobBytes)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.HandleFailedRequest(statusCode, quorum, blobBytes, method)
	})
}

func Fuzz_Nosy_Metrics_HandleInvalidArgRequest__(f *testing.F) {
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
		if g == nil {
			return
		}

		g.HandleInvalidArgRequest(method)
	})
}

func Fuzz_Nosy_Metrics_HandleSuccessfulRequest__(f *testing.F) {
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
		var quorum string
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var blobBytes int
		fill_err = tp.Fill(&blobBytes)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.HandleSuccessfulRequest(quorum, blobBytes, method)
	})
}

func Fuzz_Nosy_Metrics_HandleSystemRateLimitedRequest__(f *testing.F) {
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
		var quorum string
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var blobBytes int
		fill_err = tp.Fill(&blobBytes)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.HandleSystemRateLimitedRequest(quorum, blobBytes, method)
	})
}

func Fuzz_Nosy_Metrics_IncrementFailedBlobRequestNum__(f *testing.F) {
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
		var statusCode string
		fill_err = tp.Fill(&statusCode)
		if fill_err != nil {
			return
		}
		var quorum string
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.IncrementFailedBlobRequestNum(statusCode, quorum, method)
	})
}

func Fuzz_Nosy_Metrics_IncrementSuccessfulBlobRequestNum__(f *testing.F) {
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
		var quorum string
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.IncrementSuccessfulBlobRequestNum(quorum, method)
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
		var latencyMs float64
		fill_err = tp.Fill(&latencyMs)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.ObserveLatency(method, latencyMs)
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

func Fuzz_Nosy_BlobKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key string) {
		mk, err := ParseBlobKey(key)
		if err != nil {
			return
		}
		mk.String()
	})
}

func Fuzz_Nosy_BlobStatus_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status disperser.BlobStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		bs, err := FromBlobStatusProto(status)
		if err != nil {
			return
		}
		bs.String()
	})
}

// skipping Fuzz_Nosy_BlobStore_GetAllBlobMetadataByBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_GetBlobContent__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_GetBlobMetadata__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_GetBlobMetadataByStatus__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_GetBlobMetadataByStatusWithPagination__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_GetBlobsByMetadata__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_GetMetadataInBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_HandleBlobFailure__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_IncrementBlobRetryCount__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_MarkBlobConfirmed__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_MarkBlobFailed__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_MarkBlobFinalized__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_MarkBlobInsufficientSignatures__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_MarkBlobProcessing__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_BlobStore_StoreBlob__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.BlobStore

// skipping Fuzz_Nosy_Dispatcher_DisperseBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.Dispatcher

// skipping Fuzz_Nosy_EncoderClient_EncodeBlob__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser.EncoderClient

func Fuzz_Nosy_GenerateReverseIndexKey__(f *testing.F) {
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
		var blobIndex uint32
		fill_err = tp.Fill(&blobIndex)
		if fill_err != nil {
			return
		}

		GenerateReverseIndexKey(batchHeaderHash, blobIndex)
	})
}
