package mock

import (
	"context"
	"testing"

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

func Fuzz_Nosy_MockDisperserClient_DisperseBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}

		c := NewMockDisperserClient()
		c.DisperseBlob(ctx, d2, quorums)
	})
}

func Fuzz_Nosy_MockDisperserClient_DisperseBlobAuthenticated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}

		c := NewMockDisperserClient()
		c.DisperseBlobAuthenticated(ctx, d2, quorums)
	})
}

func Fuzz_Nosy_MockDisperserClient_GetBlobStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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

		c := NewMockDisperserClient()
		c.GetBlobStatus(ctx, key)
	})
}

func Fuzz_Nosy_MockNodeClient_GetBlobHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
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

		c := NewNodeClient()
		c.GetBlobHeader(ctx, socket, batchHeaderHash, blobIndex)
	})
}

// skipping Fuzz_Nosy_MockNodeClient_GetChunks__ because parameters include func, chan, or unsupported interface: chan github.com/Layr-Labs/eigenda/clients.RetrievedChunks

func Fuzz_Nosy_MockRetrievalClient_RetrieveBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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
		var referenceBlockNumber uint
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		var batchRoot [32]byte
		fill_err = tp.Fill(&batchRoot)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}

		c := NewRetrievalClient()
		c.RetrieveBlob(ctx, batchHeaderHash, blobIndex, referenceBlockNumber, batchRoot, quorumID)
	})
}

func Fuzz_Nosy_MockRetrievalClient_StartIndexingChainState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		c := NewRetrievalClient()
		c.StartIndexingChainState(ctx)
	})
}
