package apiserver

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

func Fuzz_Nosy_DispersalServer_DisperseBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *disperser.DisperseBlobRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.DisperseBlob(ctx, req)
	})
}

// skipping Fuzz_Nosy_DispersalServer_DisperseBlobAuthenticated__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/api/grpc/disperser.Disperser_DisperseBlobAuthenticatedServer

func Fuzz_Nosy_DispersalServer_GetBlobStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *disperser.BlobStatusRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.GetBlobStatus(ctx, req)
	})
}

func Fuzz_Nosy_DispersalServer_RetrieveBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *disperser.RetrieveBlobRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.RetrieveBlob(ctx, req)
	})
}

func Fuzz_Nosy_DispersalServer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start(ctx)
	})
}

func Fuzz_Nosy_DispersalServer_checkRateLimitsAndAddRatesToHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
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
		var origin string
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		var authenticatedAddress string
		fill_err = tp.Fill(&authenticatedAddress)
		if fill_err != nil {
			return
		}
		if s == nil || blob == nil {
			return
		}

		s.checkRateLimitsAndAddRatesToHeader(ctx, blob, origin, authenticatedAddress)
	})
}

func Fuzz_Nosy_DispersalServer_disperseBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
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
		var authenticatedAddress string
		fill_err = tp.Fill(&authenticatedAddress)
		if fill_err != nil {
			return
		}
		var apiMethodName string
		fill_err = tp.Fill(&apiMethodName)
		if fill_err != nil {
			return
		}
		if s == nil || blob == nil {
			return
		}

		s.disperseBlob(ctx, blob, authenticatedAddress, apiMethodName)
	})
}

func Fuzz_Nosy_DispersalServer_getAccountRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var origin string
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		var authenticatedAddress string
		fill_err = tp.Fill(&authenticatedAddress)
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

		s.getAccountRate(origin, authenticatedAddress, quorumID)
	})
}

func Fuzz_Nosy_DispersalServer_updateQuorumConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.updateQuorumConfig(ctx)
	})
}

func Fuzz_Nosy_DispersalServer_validateRequestAndGetBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DispersalServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *disperser.DisperseBlobRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.validateRequestAndGetBlob(ctx, req)
	})
}

func Fuzz_Nosy_RateType_Plug__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RateType
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.Plug()
	})
}

func Fuzz_Nosy_RateType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RateType
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.String()
	})
}

func Fuzz_Nosy_CLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
		CLIFlags(envPrefix)
	})
}
