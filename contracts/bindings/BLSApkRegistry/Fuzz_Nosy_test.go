package contractBLSApkRegistry

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_ContractBLSApkRegistryCaller_ApkHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil || arg1 == nil {
			return
		}

		_ContractBLSApkRegistry.ApkHistory(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_CurrentApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.CurrentApk(opts, arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_GetApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.GetApk(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_GetApkHashAtBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkHashAtBlockNumberAndIndex(opts, quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_GetApkHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkHistoryLength(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_GetApkIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil || blockNumber == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkIndicesAtBlockNumber(opts, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_GetApkUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkUpdateAtIndex(opts, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var pubkeyHash [32]byte
		fill_err = tp.Fill(&pubkeyHash)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.GetOperatorFromPubkeyHash(opts, pubkeyHash)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.GetOperatorId(opts, operator)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_GetRegisteredPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.GetRegisteredPubkey(opts, operator)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_OperatorToPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.OperatorToPubkey(opts, arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_OperatorToPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.OperatorToPubkeyHash(opts, arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_PubkeyHashToOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.PubkeyHashToOperator(opts, arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCaller_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCaller
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.RegistryCoordinator(opts)
	})
}

// skipping Fuzz_Nosy_ContractBLSApkRegistryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_ApkHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || arg1 == nil {
			return
		}

		_ContractBLSApkRegistry.ApkHistory(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_CurrentApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.CurrentApk(arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_GetApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetApk(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_GetApkHashAtBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || index == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkHashAtBlockNumberAndIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_GetApkHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkHistoryLength(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_GetApkIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || blockNumber == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkIndicesAtBlockNumber(quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_GetApkUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || index == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkUpdateAtIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var pubkeyHash [32]byte
		fill_err = tp.Fill(&pubkeyHash)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetOperatorFromPubkeyHash(pubkeyHash)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetOperatorId(operator)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_GetRegisteredPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetRegisteredPubkey(operator)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_OperatorToPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.OperatorToPubkey(arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_OperatorToPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.OperatorToPubkeyHash(arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_PubkeyHashToOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.PubkeyHashToOperator(arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryCallerSession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryFilterer_FilterNewPubkeyRegistration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.FilterNewPubkeyRegistration(opts, operator)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryFilterer_FilterOperatorAddedToQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.FilterOperatorAddedToQuorums(opts)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryFilterer_FilterOperatorRemovedFromQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.FilterOperatorRemovedFromQuorums(opts)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.ParseInitialized(log)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryFilterer_ParseNewPubkeyRegistration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.ParseNewPubkeyRegistration(log)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryFilterer_ParseOperatorAddedToQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.ParseOperatorAddedToQuorums(log)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryFilterer_ParseOperatorRemovedFromQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.ParseOperatorRemovedFromQuorums(log)
	})
}

// skipping Fuzz_Nosy_ContractBLSApkRegistryFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSApkRegistry.ContractBLSApkRegistryInitialized

// skipping Fuzz_Nosy_ContractBLSApkRegistryFilterer_WatchNewPubkeyRegistration__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSApkRegistry.ContractBLSApkRegistryNewPubkeyRegistration

// skipping Fuzz_Nosy_ContractBLSApkRegistryFilterer_WatchOperatorAddedToQuorums__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSApkRegistry.ContractBLSApkRegistryOperatorAddedToQuorums

// skipping Fuzz_Nosy_ContractBLSApkRegistryFilterer_WatchOperatorRemovedFromQuorums__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSApkRegistry.ContractBLSApkRegistryOperatorRemovedFromQuorums

func Fuzz_Nosy_ContractBLSApkRegistryInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryInitializedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryNewPubkeyRegistrationIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryNewPubkeyRegistrationIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryNewPubkeyRegistrationIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryNewPubkeyRegistrationIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryNewPubkeyRegistrationIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryNewPubkeyRegistrationIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryOperatorAddedToQuorumsIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryOperatorAddedToQuorumsIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryOperatorAddedToQuorumsIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_ContractBLSApkRegistryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractBLSApkRegistryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSApkRegistryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryRaw
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_ApkHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || arg1 == nil {
			return
		}

		_ContractBLSApkRegistry.ApkHistory(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_CurrentApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.CurrentApk(arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.DeregisterOperator(operator, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_GetApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetApk(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_GetApkHashAtBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || index == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkHashAtBlockNumberAndIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_GetApkHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkHistoryLength(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_GetApkIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || blockNumber == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkIndicesAtBlockNumber(quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_GetApkUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || index == nil {
			return
		}

		_ContractBLSApkRegistry.GetApkUpdateAtIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var pubkeyHash [32]byte
		fill_err = tp.Fill(&pubkeyHash)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetOperatorFromPubkeyHash(pubkeyHash)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetOperatorId(operator)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_GetRegisteredPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.GetRegisteredPubkey(operator)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.InitializeQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_OperatorToPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.OperatorToPubkey(arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_OperatorToPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.OperatorToPubkeyHash(arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_PubkeyHashToOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.PubkeyHashToOperator(arg0)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_RegisterBLSPublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var pubkeyRegistrationMessageHash BN254G1Point
		fill_err = tp.Fill(&pubkeyRegistrationMessageHash)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.RegisterBLSPublicKey(operator, params, pubkeyRegistrationMessageHash)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.RegisterOperator(operator, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistrySession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistrySession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryTransactor_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactor
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.DeregisterOperator(opts, operator, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryTransactor_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactor
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.InitializeQuorum(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryTransactor_RegisterBLSPublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactor
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var pubkeyRegistrationMessageHash BN254G1Point
		fill_err = tp.Fill(&pubkeyRegistrationMessageHash)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.RegisterBLSPublicKey(opts, operator, params, pubkeyRegistrationMessageHash)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactor
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.RegisterOperator(opts, operator, quorumNumbers)
	})
}

// skipping Fuzz_Nosy_ContractBLSApkRegistryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSApkRegistryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactorRaw
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil || opts == nil {
			return
		}

		_ContractBLSApkRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryTransactorSession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.DeregisterOperator(operator, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryTransactorSession_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.InitializeQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryTransactorSession_RegisterBLSPublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var pubkeyRegistrationMessageHash BN254G1Point
		fill_err = tp.Fill(&pubkeyRegistrationMessageHash)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.RegisterBLSPublicKey(operator, params, pubkeyRegistrationMessageHash)
	})
}

func Fuzz_Nosy_ContractBLSApkRegistryTransactorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession
		fill_err = tp.Fill(&_ContractBLSApkRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractBLSApkRegistry == nil {
			return
		}

		_ContractBLSApkRegistry.RegisterOperator(operator, quorumNumbers)
	})
}

// skipping Fuzz_Nosy_DeployContractBLSApkRegistry__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
