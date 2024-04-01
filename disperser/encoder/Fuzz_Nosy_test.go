package encoder

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenda/disperser/api/grpc/encoder"
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

func Fuzz_Nosy_Metrics_IncrementCanceledBlobRequestNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.IncrementCanceledBlobRequestNum()
	})
}

func Fuzz_Nosy_Metrics_IncrementFailedBlobRequestNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.IncrementFailedBlobRequestNum()
	})
}

func Fuzz_Nosy_Metrics_IncrementRateLimitedBlobRequestNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.IncrementRateLimitedBlobRequestNum()
	})
}

func Fuzz_Nosy_Metrics_IncrementSuccessfulBlobRequestNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.IncrementSuccessfulBlobRequestNum()
	})
}

func Fuzz_Nosy_Metrics_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Start(ctx)
	})
}

func Fuzz_Nosy_Metrics_TakeLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var encoding time.Duration
		fill_err = tp.Fill(&encoding)
		if fill_err != nil {
			return
		}
		var total time.Duration
		fill_err = tp.Fill(&total)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.TakeLatency(encoding, total)
	})
}

func Fuzz_Nosy_Metrics_shutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var server *http.Server
		fill_err = tp.Fill(&server)
		if fill_err != nil {
			return
		}
		if m == nil || server == nil {
			return
		}

		m.shutdown(server)
	})
}

func Fuzz_Nosy_Server_Close__(f *testing.F) {
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

		s.Close()
	})
}

func Fuzz_Nosy_Server_EncodeBlob__(f *testing.F) {
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
		var req *encoder.EncodeBlobRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.EncodeBlob(ctx, req)
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

func Fuzz_Nosy_Server_handleEncoding__(f *testing.F) {
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
		var req *encoder.EncodeBlobRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.handleEncoding(ctx, req)
	})
}

func Fuzz_Nosy_Server_popRequest__(f *testing.F) {
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

		s.popRequest()
	})
}

func Fuzz_Nosy_client_EncodeBlob__(f *testing.F) {
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

		c.EncodeBlob(ctx, d3, encodingParams)
	})
}
