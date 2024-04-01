package store

import (
	"testing"
	"github.com/Layr-Labs/eigenda/common"
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
	
	func Fuzz_Nosy_dynamodbBucketStore[T any]_GetItem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *dynamodbBucketStore[T]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var requesterID string
		fill_err = tp.Fill(&requesterID)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.GetItem(ctx, requesterID)
	})
}

// skipping Fuzz_Nosy_dynamodbBucketStore[T any]_UpdateItem__ because parameters include func, chan, or unsupported interface: *T

func Fuzz_Nosy_localParamStore[T any]_GetItem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *localParamStore[T]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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

	s.GetItem(ctx, key)
	})
}

// skipping Fuzz_Nosy_localParamStore[T any]_UpdateItem__ because parameters include func, chan, or unsupported interface: *T

