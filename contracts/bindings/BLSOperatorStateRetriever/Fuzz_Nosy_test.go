package contractBLSOperatorStateRetriever

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

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverCaller_GetCheckSignaturesIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverCaller
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetCheckSignaturesIndices(opts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
	})
}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverCaller_GetOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverCaller
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetOperatorState(opts, registryCoordinator, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverCaller_GetOperatorState0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverCaller
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetOperatorState0(opts, registryCoordinator, operatorId, blockNumber)
	})
}

// skipping Fuzz_Nosy_ContractBLSOperatorStateRetrieverCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverCallerSession_GetCheckSignaturesIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverCallerSession
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetCheckSignaturesIndices(registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
	})
}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverCallerSession_GetOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverCallerSession
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetOperatorState(registryCoordinator, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverCallerSession_GetOperatorState0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverCallerSession
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetOperatorState0(registryCoordinator, operatorId, blockNumber)
	})
}

// skipping Fuzz_Nosy_ContractBLSOperatorStateRetrieverRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractBLSOperatorStateRetrieverRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverRaw
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverSession_GetCheckSignaturesIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverSession
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetCheckSignaturesIndices(registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
	})
}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverSession_GetOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverSession
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetOperatorState(registryCoordinator, quorumNumbers, blockNumber)
	})
}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverSession_GetOperatorState0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverSession
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
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
		if _ContractBLSOperatorStateRetriever == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.GetOperatorState0(registryCoordinator, operatorId, blockNumber)
	})
}

// skipping Fuzz_Nosy_ContractBLSOperatorStateRetrieverTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSOperatorStateRetrieverTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSOperatorStateRetriever *ContractBLSOperatorStateRetrieverTransactorRaw
		fill_err = tp.Fill(&_ContractBLSOperatorStateRetriever)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSOperatorStateRetriever == nil || opts == nil {
			return
		}

		_ContractBLSOperatorStateRetriever.Transfer(opts)
	})
}

// skipping Fuzz_Nosy_DeployContractBLSOperatorStateRetriever__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
