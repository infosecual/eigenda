package contractMockRollup

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func Fuzz_Nosy_ContractMockRollupCaller_Commitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupCaller
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || opts == nil || arg0 == nil {
			return
		}

		_ContractMockRollup.Commitments(opts, arg0)
	})
}

func Fuzz_Nosy_ContractMockRollupCaller_EigenDAServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupCaller
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || opts == nil {
			return
		}

		_ContractMockRollup.EigenDAServiceManager(opts)
	})
}

func Fuzz_Nosy_ContractMockRollupCaller_Tau__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupCaller
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || opts == nil {
			return
		}

		_ContractMockRollup.Tau(opts)
	})
}

// skipping Fuzz_Nosy_ContractMockRollupCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractMockRollupCallerSession_Commitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupCallerSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || arg0 == nil {
			return
		}

		_ContractMockRollup.Commitments(arg0)
	})
}

func Fuzz_Nosy_ContractMockRollupCallerSession_EigenDAServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupCallerSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil {
			return
		}

		_ContractMockRollup.EigenDAServiceManager()
	})
}

func Fuzz_Nosy_ContractMockRollupCallerSession_Tau__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupCallerSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil {
			return
		}

		_ContractMockRollup.Tau()
	})
}

// skipping Fuzz_Nosy_ContractMockRollupRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractMockRollupRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractMockRollupRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupRaw
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || opts == nil {
			return
		}

		_ContractMockRollup.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractMockRollupSession_ChallengeCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var timestamp *big.Int
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		var point *big.Int
		fill_err = tp.Fill(&point)
		if fill_err != nil {
			return
		}
		var proof BN254G2Point
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var challengeValue *big.Int
		fill_err = tp.Fill(&challengeValue)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || timestamp == nil || point == nil || challengeValue == nil {
			return
		}

		_ContractMockRollup.ChallengeCommitment(timestamp, point, proof, challengeValue)
	})
}

func Fuzz_Nosy_ContractMockRollupSession_Commitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || arg0 == nil {
			return
		}

		_ContractMockRollup.Commitments(arg0)
	})
}

func Fuzz_Nosy_ContractMockRollupSession_EigenDAServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil {
			return
		}

		_ContractMockRollup.EigenDAServiceManager()
	})
}

func Fuzz_Nosy_ContractMockRollupSession_PostCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var blobHeader IEigenDAServiceManagerBlobHeader
		fill_err = tp.Fill(&blobHeader)
		if fill_err != nil {
			return
		}
		var blobVerificationProof EigenDARollupUtilsBlobVerificationProof
		fill_err = tp.Fill(&blobVerificationProof)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil {
			return
		}

		_ContractMockRollup.PostCommitment(blobHeader, blobVerificationProof)
	})
}

func Fuzz_Nosy_ContractMockRollupSession_Tau__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil {
			return
		}

		_ContractMockRollup.Tau()
	})
}

func Fuzz_Nosy_ContractMockRollupTransactor_ChallengeCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupTransactor
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var timestamp *big.Int
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		var point *big.Int
		fill_err = tp.Fill(&point)
		if fill_err != nil {
			return
		}
		var proof BN254G2Point
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var challengeValue *big.Int
		fill_err = tp.Fill(&challengeValue)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || opts == nil || timestamp == nil || point == nil || challengeValue == nil {
			return
		}

		_ContractMockRollup.ChallengeCommitment(opts, timestamp, point, proof, challengeValue)
	})
}

func Fuzz_Nosy_ContractMockRollupTransactor_PostCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupTransactor
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var blobHeader IEigenDAServiceManagerBlobHeader
		fill_err = tp.Fill(&blobHeader)
		if fill_err != nil {
			return
		}
		var blobVerificationProof EigenDARollupUtilsBlobVerificationProof
		fill_err = tp.Fill(&blobVerificationProof)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || opts == nil {
			return
		}

		_ContractMockRollup.PostCommitment(opts, blobHeader, blobVerificationProof)
	})
}

// skipping Fuzz_Nosy_ContractMockRollupTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractMockRollupTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupTransactorRaw
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || opts == nil {
			return
		}

		_ContractMockRollup.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractMockRollupTransactorSession_ChallengeCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupTransactorSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var timestamp *big.Int
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		var point *big.Int
		fill_err = tp.Fill(&point)
		if fill_err != nil {
			return
		}
		var proof BN254G2Point
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var challengeValue *big.Int
		fill_err = tp.Fill(&challengeValue)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil || timestamp == nil || point == nil || challengeValue == nil {
			return
		}

		_ContractMockRollup.ChallengeCommitment(timestamp, point, proof, challengeValue)
	})
}

func Fuzz_Nosy_ContractMockRollupTransactorSession_PostCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractMockRollup *ContractMockRollupTransactorSession
		fill_err = tp.Fill(&_ContractMockRollup)
		if fill_err != nil {
			return
		}
		var blobHeader IEigenDAServiceManagerBlobHeader
		fill_err = tp.Fill(&blobHeader)
		if fill_err != nil {
			return
		}
		var blobVerificationProof EigenDARollupUtilsBlobVerificationProof
		fill_err = tp.Fill(&blobVerificationProof)
		if fill_err != nil {
			return
		}
		if _ContractMockRollup == nil {
			return
		}

		_ContractMockRollup.PostCommitment(blobHeader, blobVerificationProof)
	})
}

// skipping Fuzz_Nosy_DeployContractMockRollup__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
