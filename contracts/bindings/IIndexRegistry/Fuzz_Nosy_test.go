package contractIIndexRegistry

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func Fuzz_Nosy_ContractIIndexRegistryCaller_GetOperatorIndexForQuorumAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
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
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorIndexForQuorumAtBlockNumberByIndex(opts, operatorId, quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCaller_GetOperatorIndexUpdateOfOperatorIdForQuorumAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(opts, operatorId, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCaller_GetOperatorListForQuorumAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIIndexRegistry)
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
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorListForQuorumAtBlockNumber(opts, quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCaller_GetTotalOperatorsForQuorumAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIIndexRegistry)
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
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.GetTotalOperatorsForQuorumAtBlockNumberByIndex(opts, quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCaller_GetTotalOperatorsUpdateForQuorumAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIIndexRegistry)
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
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.GetTotalOperatorsUpdateForQuorumAtIndex(opts, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCaller_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.RegistryCoordinator(opts)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCaller_TotalOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIIndexRegistry)
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
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.TotalOperatorsForQuorum(opts, quorumNumber)
	})
}

// skipping Fuzz_Nosy_ContractIIndexRegistryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractIIndexRegistryCallerSession_GetOperatorIndexForQuorumAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
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
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorIndexForQuorumAtBlockNumberByIndex(operatorId, quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCallerSession_GetOperatorIndexUpdateOfOperatorIdForQuorumAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCallerSession_GetOperatorListForQuorumAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
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
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorListForQuorumAtBlockNumber(quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCallerSession_GetTotalOperatorsForQuorumAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
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
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetTotalOperatorsForQuorumAtBlockNumberByIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCallerSession_GetTotalOperatorsUpdateForQuorumAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetTotalOperatorsUpdateForQuorumAtIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCallerSession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractIIndexRegistryCallerSession_TotalOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.TotalOperatorsForQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryFilterer_FilterGlobalIndexUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryFilterer
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operatorId [][32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.FilterGlobalIndexUpdate(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryFilterer_FilterQuorumIndexUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryFilterer
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operatorId [][32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.FilterQuorumIndexUpdate(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryFilterer_ParseGlobalIndexUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryFilterer
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.ParseGlobalIndexUpdate(log)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryFilterer_ParseQuorumIndexUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryFilterer
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.ParseQuorumIndexUpdate(log)
	})
}

// skipping Fuzz_Nosy_ContractIIndexRegistryFilterer_WatchGlobalIndexUpdate__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/IIndexRegistry.ContractIIndexRegistryGlobalIndexUpdate

// skipping Fuzz_Nosy_ContractIIndexRegistryFilterer_WatchQuorumIndexUpdate__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/IIndexRegistry.ContractIIndexRegistryQuorumIndexUpdate

func Fuzz_Nosy_ContractIIndexRegistryGlobalIndexUpdateIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIIndexRegistryGlobalIndexUpdateIterator
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

func Fuzz_Nosy_ContractIIndexRegistryGlobalIndexUpdateIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIIndexRegistryGlobalIndexUpdateIterator
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

func Fuzz_Nosy_ContractIIndexRegistryGlobalIndexUpdateIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIIndexRegistryGlobalIndexUpdateIterator
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

func Fuzz_Nosy_ContractIIndexRegistryQuorumIndexUpdateIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIIndexRegistryQuorumIndexUpdateIterator
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

func Fuzz_Nosy_ContractIIndexRegistryQuorumIndexUpdateIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIIndexRegistryQuorumIndexUpdateIterator
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

func Fuzz_Nosy_ContractIIndexRegistryQuorumIndexUpdateIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIIndexRegistryQuorumIndexUpdateIterator
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

// skipping Fuzz_Nosy_ContractIIndexRegistryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractIIndexRegistryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractIIndexRegistryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryRaw
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var operatorIdsToSwap [][32]byte
		fill_err = tp.Fill(&operatorIdsToSwap)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.DeregisterOperator(operatorId, quorumNumbers, operatorIdsToSwap)
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_GetOperatorIndexForQuorumAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
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
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorIndexForQuorumAtBlockNumberByIndex(operatorId, quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_GetOperatorIndexUpdateOfOperatorIdForQuorumAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorIndexUpdateOfOperatorIdForQuorumAtIndex(operatorId, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_GetOperatorListForQuorumAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
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
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetOperatorListForQuorumAtBlockNumber(quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_GetTotalOperatorsForQuorumAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
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
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetTotalOperatorsForQuorumAtBlockNumberByIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_GetTotalOperatorsUpdateForQuorumAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.GetTotalOperatorsUpdateForQuorumAtIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.RegisterOperator(operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractIIndexRegistrySession_TotalOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistrySession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.TotalOperatorsForQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryTransactor_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryTransactor
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var operatorIdsToSwap [][32]byte
		fill_err = tp.Fill(&operatorIdsToSwap)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.DeregisterOperator(opts, operatorId, quorumNumbers, operatorIdsToSwap)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryTransactor
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.RegisterOperator(opts, operatorId, quorumNumbers)
	})
}

// skipping Fuzz_Nosy_ContractIIndexRegistryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractIIndexRegistryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryTransactorRaw
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIIndexRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryTransactorSession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryTransactorSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var operatorIdsToSwap [][32]byte
		fill_err = tp.Fill(&operatorIdsToSwap)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.DeregisterOperator(operatorId, quorumNumbers, operatorIdsToSwap)
	})
}

func Fuzz_Nosy_ContractIIndexRegistryTransactorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIIndexRegistry *ContractIIndexRegistryTransactorSession
		fill_err = tp.Fill(&_ContractIIndexRegistry)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractIIndexRegistry == nil {
			return
		}

		_ContractIIndexRegistry.RegisterOperator(operatorId, quorumNumbers)
	})
}
