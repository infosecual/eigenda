package contractOperatorStateRetriever

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_ContractOperatorStateRetrieverCaller_GetCheckSignaturesIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var nonSignerOperatorIds [][32]byte
		fill_err = tp.Fill(&nonSignerOperatorIds)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractOperatorStateRetriever.GetCheckSignaturesIndices(opts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverCaller_GetOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractOperatorStateRetriever.GetOperatorState(opts, registryCoordinator, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverCaller_GetOperatorState0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractOperatorStateRetriever.GetOperatorState0(opts, registryCoordinator, operatorId, blockNumber)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverCaller_GetQuorumBitmapsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorIds [][32]byte
		fill_err = tp.Fill(&operatorIds)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractOperatorStateRetriever.GetQuorumBitmapsAtBlockNumber(opts, registryCoordinator, operatorIds, blockNumber)
	})
}

// skipping Fuzz_Nosy_ContractOperatorStateRetrieverCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractOperatorStateRetrieverCallerSession_GetCheckSignaturesIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var nonSignerOperatorIds [][32]byte
		fill_err = tp.Fill(&nonSignerOperatorIds)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil {
			return
		}

		_ContractOperatorStateRetriever.GetCheckSignaturesIndices(registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverCallerSession_GetOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil {
			return
		}

		_ContractOperatorStateRetriever.GetOperatorState(registryCoordinator, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverCallerSession_GetOperatorState0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil {
			return
		}

		_ContractOperatorStateRetriever.GetOperatorState0(registryCoordinator, operatorId, blockNumber)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverCallerSession_GetQuorumBitmapsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorIds [][32]byte
		fill_err = tp.Fill(&operatorIds)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil {
			return
		}

		_ContractOperatorStateRetriever.GetQuorumBitmapsAtBlockNumber(registryCoordinator, operatorIds, blockNumber)
	})
}

// skipping Fuzz_Nosy_ContractOperatorStateRetrieverRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractOperatorStateRetrieverRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractOperatorStateRetrieverRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractOperatorStateRetriever.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverSession_GetCheckSignaturesIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var nonSignerOperatorIds [][32]byte
		fill_err = tp.Fill(&nonSignerOperatorIds)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil {
			return
		}

		_ContractOperatorStateRetriever.GetCheckSignaturesIndices(registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverSession_GetOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil {
			return
		}

		_ContractOperatorStateRetriever.GetOperatorState(registryCoordinator, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverSession_GetOperatorState0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil {
			return
		}

		_ContractOperatorStateRetriever.GetOperatorState0(registryCoordinator, operatorId, blockNumber)
	})
}

func Fuzz_Nosy_ContractOperatorStateRetrieverSession_GetQuorumBitmapsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var registryCoordinator common.Address
		fill_err = tp.Fill(&registryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorIds [][32]byte
		fill_err = tp.Fill(&operatorIds)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil {
			return
		}

		_ContractOperatorStateRetriever.GetQuorumBitmapsAtBlockNumber(registryCoordinator, operatorIds, blockNumber)
	})
}

// skipping Fuzz_Nosy_ContractOperatorStateRetrieverTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractOperatorStateRetrieverTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw
		fill_err = tp.Fill(&_ContractOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractOperatorStateRetriever.Transfer(opts)
	})
}

// skipping Fuzz_Nosy_DeployContractOperatorStateRetriever__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
