package contractBLSPubkeyRegistry

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

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_GetApkForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkForQuorum(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_GetApkHashForQuorumAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkHashForQuorumAtBlockNumberFromIndex(opts, quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_GetApkIndicesForQuorumsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || opts == nil || blockNumber == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkIndicesForQuorumsAtBlockNumber(opts, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_GetApkUpdateForQuorumByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkUpdateForQuorumByIndex(opts, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetOperatorFromPubkeyHash(opts, pubkeyHash)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_GetQuorumApkHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetQuorumApkHistoryLength(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_PubkeyCompendium__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.PubkeyCompendium(opts)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_QuorumApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.QuorumApk(opts, arg0)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_QuorumApkUpdates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || opts == nil || arg1 == nil {
			return
		}

		_ContractBLSPubkeyRegistry.QuorumApkUpdates(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCaller_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.RegistryCoordinator(opts)
	})
}

// skipping Fuzz_Nosy_ContractBLSPubkeyRegistryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_GetApkForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkForQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_GetApkHashForQuorumAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || index == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkHashForQuorumAtBlockNumberFromIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_GetApkIndicesForQuorumsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || blockNumber == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkIndicesForQuorumsAtBlockNumber(quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_GetApkUpdateForQuorumByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || index == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkUpdateForQuorumByIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var pubkeyHash [32]byte
		fill_err = tp.Fill(&pubkeyHash)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetOperatorFromPubkeyHash(pubkeyHash)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_GetQuorumApkHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetQuorumApkHistoryLength(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_PubkeyCompendium__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.PubkeyCompendium()
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_QuorumApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.QuorumApk(arg0)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_QuorumApkUpdates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || arg1 == nil {
			return
		}

		_ContractBLSPubkeyRegistry.QuorumApkUpdates(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryCallerSession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_FilterOperatorAddedToQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.FilterOperatorAddedToQuorums(opts)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_FilterOperatorRemovedFromQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.FilterOperatorRemovedFromQuorums(opts)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.ParseInitialized(log)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_ParseOperatorAddedToQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.ParseOperatorAddedToQuorums(log)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_ParseOperatorRemovedFromQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.ParseOperatorRemovedFromQuorums(log)
	})
}

// skipping Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSPubkeyRegistry.ContractBLSPubkeyRegistryInitialized

// skipping Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_WatchOperatorAddedToQuorums__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSPubkeyRegistry.ContractBLSPubkeyRegistryOperatorAddedToQuorums

// skipping Fuzz_Nosy_ContractBLSPubkeyRegistryFilterer_WatchOperatorRemovedFromQuorums__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSPubkeyRegistry.ContractBLSPubkeyRegistryOperatorRemovedFromQuorums

func Fuzz_Nosy_ContractBLSPubkeyRegistryInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryInitializedIterator
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

func Fuzz_Nosy_ContractBLSPubkeyRegistryInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryInitializedIterator
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

func Fuzz_Nosy_ContractBLSPubkeyRegistryInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryInitializedIterator
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

func Fuzz_Nosy_ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator
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

func Fuzz_Nosy_ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator
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

func Fuzz_Nosy_ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator
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

func Fuzz_Nosy_ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator
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

func Fuzz_Nosy_ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator
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

func Fuzz_Nosy_ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator
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

// skipping Fuzz_Nosy_ContractBLSPubkeyRegistryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractBLSPubkeyRegistryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSPubkeyRegistryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryRaw
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		var pubkey BN254G1Point
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.DeregisterOperator(operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_GetApkForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkForQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_GetApkHashForQuorumAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || index == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkHashForQuorumAtBlockNumberFromIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_GetApkIndicesForQuorumsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || blockNumber == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkIndicesForQuorumsAtBlockNumber(quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_GetApkUpdateForQuorumByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || index == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetApkUpdateForQuorumByIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var pubkeyHash [32]byte
		fill_err = tp.Fill(&pubkeyHash)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetOperatorFromPubkeyHash(pubkeyHash)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_GetQuorumApkHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.GetQuorumApkHistoryLength(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_PubkeyCompendium__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.PubkeyCompendium()
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_QuorumApk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.QuorumApk(arg0)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_QuorumApkUpdates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		if _ContractBLSPubkeyRegistry == nil || arg1 == nil {
			return
		}

		_ContractBLSPubkeyRegistry.QuorumApkUpdates(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		var pubkey BN254G1Point
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.RegisterOperator(operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistrySession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryTransactor_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactor
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		var pubkey BN254G1Point
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.DeregisterOperator(opts, operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactor
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		var pubkey BN254G1Point
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.RegisterOperator(opts, operator, quorumNumbers, pubkey)
	})
}

// skipping Fuzz_Nosy_ContractBLSPubkeyRegistryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSPubkeyRegistryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactorRaw
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractBLSPubkeyRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryTransactorSession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactorSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		var pubkey BN254G1Point
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.DeregisterOperator(operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSPubkeyRegistryTransactorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactorSession
		fill_err = tp.Fill(&_ContractBLSPubkeyRegistry)
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
		var pubkey BN254G1Point
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if _ContractBLSPubkeyRegistry == nil {
			return
		}

		_ContractBLSPubkeyRegistry.RegisterOperator(operator, quorumNumbers, pubkey)
	})
}

// skipping Fuzz_Nosy_DeployContractBLSPubkeyRegistry__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
