package clients

import (
	"context"
	"testing"
	"time"

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

func Fuzz_Nosy_disperserClient_DisperseBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *disperserClient
		fill_err = tp.Fill(&c)
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
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.DisperseBlob(ctx, d3, quorums)
	})
}

func Fuzz_Nosy_disperserClient_DisperseBlobAuthenticated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *disperserClient
		fill_err = tp.Fill(&c)
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
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.DisperseBlobAuthenticated(ctx, d3, quorums)
	})
}

func Fuzz_Nosy_disperserClient_GetBlobStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *disperserClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var requestID []byte
		fill_err = tp.Fill(&requestID)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetBlobStatus(ctx, requestID)
	})
}

func Fuzz_Nosy_disperserClient_getDialOptions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *disperserClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.getDialOptions()
	})
}

func Fuzz_Nosy_retrievalClient_RetrieveBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *retrievalClient
		fill_err = tp.Fill(&r)
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
		if r == nil {
			return
		}

		r.RetrieveBlob(ctx, batchHeaderHash, blobIndex, referenceBlockNumber, batchRoot, quorumID)
	})
}

// skipping Fuzz_Nosy_DisperserClient_DisperseBlob__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/clients.DisperserClient

// skipping Fuzz_Nosy_DisperserClient_DisperseBlobAuthenticated__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/clients.DisperserClient

// skipping Fuzz_Nosy_DisperserClient_GetBlobStatus__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/clients.DisperserClient

func Fuzz_Nosy_NodeClient_GetBlobHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
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

		_x1 := NewNodeClient(timeout)
		_x1.GetBlobHeader(ctx, socket, batchHeaderHash, blobIndex)
	})
}

// skipping Fuzz_Nosy_NodeClient_GetChunks__ because parameters include func, chan, or unsupported interface: chan github.com/Layr-Labs/eigenda/clients.RetrievedChunks

// skipping Fuzz_Nosy_RetrievalClient_RetrieveBlob__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/clients.RetrievalClient

func Fuzz_Nosy_client_GetBlobHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c client
		fill_err = tp.Fill(&c)
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

		c.GetBlobHeader(ctx, socket, batchHeaderHash, blobIndex)
	})
}

// skipping Fuzz_Nosy_client_GetChunks__ because parameters include func, chan, or unsupported interface: chan github.com/Layr-Labs/eigenda/clients.RetrievedChunks
