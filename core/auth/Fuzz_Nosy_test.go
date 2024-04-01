package auth

import (
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

func Fuzz_Nosy_authenticator_AuthenticateBlobRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *authenticator
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var header core.BlobAuthHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.AuthenticateBlobRequest(header)
	})
}

func Fuzz_Nosy_signer_GetAccountID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *signer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetAccountID()
	})
}

func Fuzz_Nosy_signer_SignBlobRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *signer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var header core.BlobAuthHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SignBlobRequest(header)
	})
}
