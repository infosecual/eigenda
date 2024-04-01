package contractAVSDirectory

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

func Fuzz_Nosy_ContractAVSDirectoryAVSMetadataURIUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryAVSMetadataURIUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryAVSMetadataURIUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryCaller_AvsOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
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
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.AvsOperatorStatus(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_CalculateOperatorAVSRegistrationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
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
		var avs common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil || expiry == nil {
			return
		}

		_ContractAVSDirectory.CalculateOperatorAVSRegistrationDigestHash(opts, operator, avs, salt, expiry)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.DOMAINTYPEHASH(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.Delegation(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.DomainSeparator(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_OPERATORAVSREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.OPERATORAVSREGISTRATIONTYPEHASH(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_OperatorSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
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
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.OperatorSaltIsSpent(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.Owner(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.Paused(opts, index)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.Paused0(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCaller
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.PauserRegistry(opts)
	})
}

// skipping Fuzz_Nosy_ContractAVSDirectoryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_AvsOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.AvsOperatorStatus(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_CalculateOperatorAVSRegistrationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var avs common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || expiry == nil {
			return
		}

		_ContractAVSDirectory.CalculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.Delegation()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.DomainSeparator()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_OPERATORAVSREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.OPERATORAVSREGISTRATIONTYPEHASH()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_OperatorSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.OperatorSaltIsSpent(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.Owner()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.Paused(index)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.Paused0()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryCallerSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_FilterAVSMetadataURIUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var avs []common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.FilterAVSMetadataURIUpdated(opts, avs)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_FilterOperatorAVSRegistrationStatusUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
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
		var avs []common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.FilterOperatorAVSRegistrationStatusUpdated(opts, operator, avs)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var account []common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_ParseAVSMetadataURIUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.ParseAVSMetadataURIUpdated(log)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.ParseInitialized(log)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_ParseOperatorAVSRegistrationStatusUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.ParseOperatorAVSRegistrationStatusUpdated(log)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.ParsePaused(log)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryFilterer
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.ParseUnpaused(log)
	})
}

// skipping Fuzz_Nosy_ContractAVSDirectoryFilterer_WatchAVSMetadataURIUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/AVSDirectory.ContractAVSDirectoryAVSMetadataURIUpdated

// skipping Fuzz_Nosy_ContractAVSDirectoryFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/AVSDirectory.ContractAVSDirectoryInitialized

// skipping Fuzz_Nosy_ContractAVSDirectoryFilterer_WatchOperatorAVSRegistrationStatusUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/AVSDirectory.ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated

// skipping Fuzz_Nosy_ContractAVSDirectoryFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/AVSDirectory.ContractAVSDirectoryOwnershipTransferred

// skipping Fuzz_Nosy_ContractAVSDirectoryFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/AVSDirectory.ContractAVSDirectoryPaused

// skipping Fuzz_Nosy_ContractAVSDirectoryFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/AVSDirectory.ContractAVSDirectoryPauserRegistrySet

// skipping Fuzz_Nosy_ContractAVSDirectoryFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/AVSDirectory.ContractAVSDirectoryUnpaused

func Fuzz_Nosy_ContractAVSDirectoryInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryInitializedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryInitializedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryInitializedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractAVSDirectoryOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractAVSDirectoryOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractAVSDirectoryPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryPausedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryPausedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryPausedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryPauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractAVSDirectoryPauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractAVSDirectoryPauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryPauserRegistrySetIterator
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

// skipping Fuzz_Nosy_ContractAVSDirectoryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractAVSDirectoryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractAVSDirectoryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryRaw
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_AvsOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.AvsOperatorStatus(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_CalculateOperatorAVSRegistrationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var avs common.Address
		fill_err = tp.Fill(&avs)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || expiry == nil {
			return
		}

		_ContractAVSDirectory.CalculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.Delegation()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.DeregisterOperatorFromAVS(operator)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.DomainSeparator()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || initialPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Initialize(initialOwner, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_OPERATORAVSREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.OPERATORAVSREGISTRATIONTYPEHASH()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_OperatorSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.OperatorSaltIsSpent(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.Owner()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || newPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.PauseAll()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.Paused(index)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.Paused0()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.RegisterOperatorToAVS(operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.RenounceOwnership()
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || newPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectorySession_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectorySession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.UpdateAVSMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
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
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.DeregisterOperatorFromAVS(opts, operator)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil || initialPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Initialize(opts, initialOwner, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.PauseAll(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
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
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.RegisterOperatorToAVS(opts, operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactor_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactor
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.UpdateAVSMetadataURI(opts, metadataURI)
	})
}

// skipping Fuzz_Nosy_ContractAVSDirectoryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractAVSDirectoryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorRaw
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || opts == nil {
			return
		}

		_ContractAVSDirectory.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.DeregisterOperatorFromAVS(operator)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var initialOwner common.Address
		fill_err = tp.Fill(&initialOwner)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var initialPausedStatus *big.Int
		fill_err = tp.Fill(&initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || initialPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Initialize(initialOwner, _pauserRegistry, initialPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || newPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.PauseAll()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.RegisterOperatorToAVS(operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.RenounceOwnership()
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil || newPausedStatus == nil {
			return
		}

		_ContractAVSDirectory.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryTransactorSession_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractAVSDirectory *ContractAVSDirectoryTransactorSession
		fill_err = tp.Fill(&_ContractAVSDirectory)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractAVSDirectory == nil {
			return
		}

		_ContractAVSDirectory.UpdateAVSMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_ContractAVSDirectoryUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryUnpausedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryUnpausedIterator
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

func Fuzz_Nosy_ContractAVSDirectoryUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractAVSDirectoryUnpausedIterator
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

// skipping Fuzz_Nosy_DeployContractAVSDirectory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
