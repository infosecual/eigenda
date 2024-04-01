package contractIndexRegistry

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

func Fuzz_Nosy_ContractIndexRegistryCaller_CurrentOperatorIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.CurrentOperatorIndex(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCaller_GetLatestOperatorUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		var operatorIndex uint32
		fill_err = tp.Fill(&operatorIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.GetLatestOperatorUpdate(opts, quorumNumber, operatorIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCaller_GetLatestQuorumUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.GetLatestQuorumUpdate(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCaller_GetOperatorListAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.GetOperatorListAtBlockNumber(opts, quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCaller_GetOperatorUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		var operatorIndex uint32
		fill_err = tp.Fill(&operatorIndex)
		if fill_err != nil {
			return
		}
		var arrayIndex uint32
		fill_err = tp.Fill(&arrayIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.GetOperatorUpdateAtIndex(opts, quorumNumber, operatorIndex, arrayIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCaller_GetQuorumUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		var quorumIndex uint32
		fill_err = tp.Fill(&quorumIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.GetQuorumUpdateAtIndex(opts, quorumNumber, quorumIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCaller_OPERATORDOESNOTEXISTID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.OPERATORDOESNOTEXISTID(opts)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCaller_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.RegistryCoordinator(opts)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCaller_TotalOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCaller
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.TotalOperatorsForQuorum(opts, quorumNumber)
	})
}

// skipping Fuzz_Nosy_ContractIndexRegistryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_CurrentOperatorIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.CurrentOperatorIndex(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_GetLatestOperatorUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorIndex uint32
		fill_err = tp.Fill(&operatorIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetLatestOperatorUpdate(quorumNumber, operatorIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_GetLatestQuorumUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetLatestQuorumUpdate(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_GetOperatorListAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetOperatorListAtBlockNumber(quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_GetOperatorUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorIndex uint32
		fill_err = tp.Fill(&operatorIndex)
		if fill_err != nil {
			return
		}
		var arrayIndex uint32
		fill_err = tp.Fill(&arrayIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetOperatorUpdateAtIndex(quorumNumber, operatorIndex, arrayIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_GetQuorumUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var quorumIndex uint32
		fill_err = tp.Fill(&quorumIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetQuorumUpdateAtIndex(quorumNumber, quorumIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_OPERATORDOESNOTEXISTID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.OPERATORDOESNOTEXISTID()
	})
}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractIndexRegistryCallerSession_TotalOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryCallerSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.TotalOperatorsForQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistryFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryFilterer
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_ContractIndexRegistryFilterer_FilterQuorumIndexUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryFilterer
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.FilterQuorumIndexUpdate(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractIndexRegistryFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryFilterer
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.ParseInitialized(log)
	})
}

func Fuzz_Nosy_ContractIndexRegistryFilterer_ParseQuorumIndexUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryFilterer
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.ParseQuorumIndexUpdate(log)
	})
}

// skipping Fuzz_Nosy_ContractIndexRegistryFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/IndexRegistry.ContractIndexRegistryInitialized

// skipping Fuzz_Nosy_ContractIndexRegistryFilterer_WatchQuorumIndexUpdate__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/IndexRegistry.ContractIndexRegistryQuorumIndexUpdate

func Fuzz_Nosy_ContractIndexRegistryInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIndexRegistryInitializedIterator
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

func Fuzz_Nosy_ContractIndexRegistryInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIndexRegistryInitializedIterator
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

func Fuzz_Nosy_ContractIndexRegistryInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIndexRegistryInitializedIterator
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

func Fuzz_Nosy_ContractIndexRegistryQuorumIndexUpdateIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIndexRegistryQuorumIndexUpdateIterator
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

func Fuzz_Nosy_ContractIndexRegistryQuorumIndexUpdateIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIndexRegistryQuorumIndexUpdateIterator
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

func Fuzz_Nosy_ContractIndexRegistryQuorumIndexUpdateIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIndexRegistryQuorumIndexUpdateIterator
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

// skipping Fuzz_Nosy_ContractIndexRegistryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractIndexRegistryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractIndexRegistryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryRaw
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_CurrentOperatorIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 [32]byte
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.CurrentOperatorIndex(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.DeregisterOperator(operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_GetLatestOperatorUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorIndex uint32
		fill_err = tp.Fill(&operatorIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetLatestOperatorUpdate(quorumNumber, operatorIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_GetLatestQuorumUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetLatestQuorumUpdate(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_GetOperatorListAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetOperatorListAtBlockNumber(quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_GetOperatorUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorIndex uint32
		fill_err = tp.Fill(&operatorIndex)
		if fill_err != nil {
			return
		}
		var arrayIndex uint32
		fill_err = tp.Fill(&arrayIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetOperatorUpdateAtIndex(quorumNumber, operatorIndex, arrayIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_GetQuorumUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var quorumIndex uint32
		fill_err = tp.Fill(&quorumIndex)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.GetQuorumUpdateAtIndex(quorumNumber, quorumIndex)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.InitializeQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_OPERATORDOESNOTEXISTID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.OPERATORDOESNOTEXISTID()
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.RegisterOperator(operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractIndexRegistrySession_TotalOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistrySession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.TotalOperatorsForQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistryTransactor_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryTransactor
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.DeregisterOperator(opts, operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractIndexRegistryTransactor_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryTransactor
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.InitializeQuorum(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistryTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryTransactor
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.RegisterOperator(opts, operatorId, quorumNumbers)
	})
}

// skipping Fuzz_Nosy_ContractIndexRegistryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractIndexRegistryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryTransactorRaw
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil || opts == nil {
			return
		}

		_ContractIndexRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractIndexRegistryTransactorSession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryTransactorSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.DeregisterOperator(operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractIndexRegistryTransactorSession_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryTransactorSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.InitializeQuorum(quorumNumber)
	})
}

func Fuzz_Nosy_ContractIndexRegistryTransactorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIndexRegistry *ContractIndexRegistryTransactorSession
		fill_err = tp.Fill(&_ContractIndexRegistry)
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
		if _ContractIndexRegistry == nil {
			return
		}

		_ContractIndexRegistry.RegisterOperator(operatorId, quorumNumbers)
	})
}

// skipping Fuzz_Nosy_DeployContractIndexRegistry__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
