package s3

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

func Fuzz_Nosy_client_DeleteObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *client
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bucket string
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.DeleteObject(ctx, bucket, key)
	})
}

func Fuzz_Nosy_client_DownloadObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *client
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bucket string
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.DownloadObject(ctx, bucket, key)
	})
}

func Fuzz_Nosy_client_ListObjects__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *client
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bucket string
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		var prefix string
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ListObjects(ctx, bucket, prefix)
	})
}

func Fuzz_Nosy_client_UploadObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *client
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bucket string
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UploadObject(ctx, bucket, key, d5)
	})
}

// skipping Fuzz_Nosy_Client_DeleteObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common/aws/s3.Client

// skipping Fuzz_Nosy_Client_DownloadObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common/aws/s3.Client

// skipping Fuzz_Nosy_Client_ListObjects__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common/aws/s3.Client

// skipping Fuzz_Nosy_Client_UploadObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common/aws/s3.Client
