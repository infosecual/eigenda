package contractIBLSPubkeyRegistry

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

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCaller_GetApkForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkForQuorum(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCaller_GetApkHashForQuorumAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkHashForQuorumAtBlockNumberFromIndex(opts, quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCaller_GetApkIndicesForQuorumsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || opts == nil || blockNumber == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkIndicesForQuorumsAtBlockNumber(opts, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCaller_GetApkUpdateForQuorumByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkUpdateForQuorumByIndex(opts, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCaller_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetOperatorFromPubkeyHash(opts, pubkeyHash)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCaller_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCaller
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.RegistryCoordinator(opts)
	})
}

// skipping Fuzz_Nosy_ContractIBLSPubkeyRegistryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCallerSession_GetApkForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkForQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCallerSession_GetApkHashForQuorumAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || index == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkHashForQuorumAtBlockNumberFromIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCallerSession_GetApkIndicesForQuorumsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || blockNumber == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkIndicesForQuorumsAtBlockNumber(quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCallerSession_GetApkUpdateForQuorumByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || index == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkUpdateForQuorumByIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCallerSession_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var pubkeyHash [32]byte
		fill_err = tp.Fill(&pubkeyHash)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetOperatorFromPubkeyHash(pubkeyHash)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryCallerSession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryCallerSession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryFilterer_FilterOperatorAddedToQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.FilterOperatorAddedToQuorums(opts)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryFilterer_FilterOperatorRemovedFromQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.FilterOperatorRemovedFromQuorums(opts)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryFilterer_ParseOperatorAddedToQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.ParseOperatorAddedToQuorums(log)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryFilterer_ParseOperatorRemovedFromQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryFilterer
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.ParseOperatorRemovedFromQuorums(log)
	})
}

// skipping Fuzz_Nosy_ContractIBLSPubkeyRegistryFilterer_WatchOperatorAddedToQuorums__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/IBLSPubkeyRegistry.ContractIBLSPubkeyRegistryOperatorAddedToQuorums

// skipping Fuzz_Nosy_ContractIBLSPubkeyRegistryFilterer_WatchOperatorRemovedFromQuorums__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/IBLSPubkeyRegistry.ContractIBLSPubkeyRegistryOperatorRemovedFromQuorums

func Fuzz_Nosy_ContractIBLSPubkeyRegistryOperatorAddedToQuorumsIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIBLSPubkeyRegistryOperatorAddedToQuorumsIterator
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

func Fuzz_Nosy_ContractIBLSPubkeyRegistryOperatorAddedToQuorumsIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIBLSPubkeyRegistryOperatorAddedToQuorumsIterator
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

func Fuzz_Nosy_ContractIBLSPubkeyRegistryOperatorAddedToQuorumsIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIBLSPubkeyRegistryOperatorAddedToQuorumsIterator
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

func Fuzz_Nosy_ContractIBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator
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

func Fuzz_Nosy_ContractIBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator
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

func Fuzz_Nosy_ContractIBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator
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

// skipping Fuzz_Nosy_ContractIBLSPubkeyRegistryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractIBLSPubkeyRegistryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryRaw
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistrySession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.DeregisterOperator(operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistrySession_GetApkForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkForQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistrySession_GetApkHashForQuorumAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || index == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkHashForQuorumAtBlockNumberFromIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistrySession_GetApkIndicesForQuorumsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || blockNumber == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkIndicesForQuorumsAtBlockNumber(quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistrySession_GetApkUpdateForQuorumByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || index == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetApkUpdateForQuorumByIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistrySession_GetOperatorFromPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var pubkeyHash [32]byte
		fill_err = tp.Fill(&pubkeyHash)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.GetOperatorFromPubkeyHash(pubkeyHash)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistrySession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.RegisterOperator(operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistrySession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistrySession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryTransactor_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryTransactor
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.DeregisterOperator(opts, operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryTransactor
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.RegisterOperator(opts, operator, quorumNumbers, pubkey)
	})
}

// skipping Fuzz_Nosy_ContractIBLSPubkeyRegistryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryTransactorRaw
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIBLSPubkeyRegistry == nil || opts == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryTransactorSession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryTransactorSession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.DeregisterOperator(operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractIBLSPubkeyRegistryTransactorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIBLSPubkeyRegistry *ContractIBLSPubkeyRegistryTransactorSession
		fill_err = tp.Fill(&_ContractIBLSPubkeyRegistry)
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
		if _ContractIBLSPubkeyRegistry == nil {
			return
		}

		_ContractIBLSPubkeyRegistry.RegisterOperator(operator, quorumNumbers, pubkey)
	})
}
