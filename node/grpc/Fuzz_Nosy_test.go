package grpc

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/node"
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

func Fuzz_Nosy_Server_GetBlobHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in *node.GetBlobHeaderRequest
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if s == nil || in == nil {
			return
		}

		s.GetBlobHeader(ctx, in)
	})
}

func Fuzz_Nosy_Server_RetrieveChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in *node.RetrieveChunksRequest
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if s == nil || in == nil {
			return
		}

		s.RetrieveChunks(ctx, in)
	})
}

func Fuzz_Nosy_Server_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start()
	})
}

func Fuzz_Nosy_Server_StoreChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in *node.StoreChunksRequest
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if s == nil || in == nil {
			return
		}

		s.StoreChunks(ctx, in)
	})
}

func Fuzz_Nosy_Server_getBlobHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
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

		s.getBlobHeader(ctx, batchHeaderHash, blobIndex)
	})
}

func Fuzz_Nosy_Server_handleStoreChunksRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in *node.StoreChunksRequest
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if s == nil || in == nil {
			return
		}

		s.handleStoreChunksRequest(ctx, in)
	})
}

func Fuzz_Nosy_Server_rebuildMerkleTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
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

		s.rebuildMerkleTree(batchHeaderHash)
	})
}

func Fuzz_Nosy_Server_serveDispersal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.serveDispersal()
	})
}

func Fuzz_Nosy_Server_serveRetrieval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.serveRetrieval()
	})
}

func Fuzz_Nosy_Server_validateStoreChunkRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var in *node.StoreChunksRequest
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if s == nil || in == nil {
			return
		}

		s.validateStoreChunkRequest(in)
	})
}

func Fuzz_Nosy_GetBlobMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *node.StoreChunksRequest
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if in == nil {
			return
		}

		GetBlobMessages(in)
	})
}
