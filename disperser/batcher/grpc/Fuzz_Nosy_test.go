package dispatcher

import (
	"context"
	"testing"

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

func Fuzz_Nosy_dispatcher_DisperseBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *dispatcher
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var state *core.IndexedOperatorState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var blobs []core.EncodedBlob
		fill_err = tp.Fill(&blobs)
		if fill_err != nil {
			return
		}
		var batchHeader *core.BatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}
		if c == nil || state == nil || batchHeader == nil {
			return
		}

		c.DisperseBatch(ctx, state, blobs, batchHeader)
	})
}

// skipping Fuzz_Nosy_dispatcher_sendAllChunks__ because parameters include func, chan, or unsupported interface: chan github.com/Layr-Labs/eigenda/core.SignerMessage

func Fuzz_Nosy_dispatcher_sendChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *dispatcher
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blobs []*core.BlobMessage
		fill_err = tp.Fill(&blobs)
		if fill_err != nil {
			return
		}
		var batchHeader *core.BatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}
		var op *core.IndexedOperatorInfo
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		if c == nil || batchHeader == nil || op == nil {
			return
		}

		c.sendChunks(ctx, blobs, batchHeader, op)
	})
}

func Fuzz_Nosy_GetStoreChunksRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blobMessages []*core.BlobMessage
		fill_err = tp.Fill(&blobMessages)
		if fill_err != nil {
			return
		}
		var batchHeader *core.BatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}
		if batchHeader == nil {
			return
		}

		GetStoreChunksRequest(blobMessages, batchHeader)
	})
}
