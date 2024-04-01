package contractStakeRegistry

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

func Fuzz_Nosy_ContractStakeRegistryCaller_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.Delegation(opts)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetCurrentStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetCurrentStake(opts, operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetCurrentTotalStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetCurrentTotalStake(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetLatestStakeUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetLatestStakeUpdate(opts, operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetStakeAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetStakeAtBlockNumber(opts, operatorId, quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetStakeAtBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetStakeAtBlockNumberAndIndex(opts, quorumNumber, blockNumber, operatorId, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetStakeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetStakeHistory(opts, operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetStakeHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetStakeHistoryLength(opts, operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetStakeUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetStakeUpdateAtIndex(opts, quorumNumber, operatorId, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetStakeUpdateIndexAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetStakeUpdateIndexAtBlockNumber(opts, operatorId, quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetTotalStakeAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeAtBlockNumberFromIndex(opts, quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetTotalStakeHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeHistoryLength(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetTotalStakeIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeIndicesAtBlockNumber(opts, blockNumber, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_GetTotalStakeUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeUpdateAtIndex(opts, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_MAXWEIGHINGFUNCTIONLENGTH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.MAXWEIGHINGFUNCTIONLENGTH(opts)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_MinimumStakeForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.MinimumStakeForQuorum(opts, arg0)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.RegistryCoordinator(opts)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_StrategiesPerQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil || arg1 == nil {
			return
		}

		_ContractStakeRegistry.StrategiesPerQuorum(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_StrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil || arg1 == nil {
			return
		}

		_ContractStakeRegistry.StrategyParams(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_StrategyParamsByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil || index == nil {
			return
		}

		_ContractStakeRegistry.StrategyParamsByIndex(opts, quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_StrategyParamsLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.StrategyParamsLength(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_WEIGHTINGDIVISOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.WEIGHTINGDIVISOR(opts)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCaller_WeightOfOperatorForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCaller
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.WeightOfOperatorForQuorum(opts, quorumNumber, operator)
	})
}

// skipping Fuzz_Nosy_ContractStakeRegistryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.Delegation()
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetCurrentStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetCurrentStake(operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetCurrentTotalStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetCurrentTotalStake(quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetLatestStakeUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetLatestStakeUpdate(operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetStakeAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetStakeAtBlockNumber(operatorId, quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetStakeAtBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetStakeAtBlockNumberAndIndex(quorumNumber, blockNumber, operatorId, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetStakeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetStakeHistory(operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetStakeHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetStakeHistoryLength(operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetStakeUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetStakeUpdateAtIndex(quorumNumber, operatorId, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetStakeUpdateIndexAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetStakeUpdateIndexAtBlockNumber(operatorId, quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetTotalStakeAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeAtBlockNumberFromIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetTotalStakeHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeHistoryLength(quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetTotalStakeIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeIndicesAtBlockNumber(blockNumber, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_GetTotalStakeUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeUpdateAtIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_MAXWEIGHINGFUNCTIONLENGTH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.MAXWEIGHINGFUNCTIONLENGTH()
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_MinimumStakeForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.MinimumStakeForQuorum(arg0)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_StrategiesPerQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || arg1 == nil {
			return
		}

		_ContractStakeRegistry.StrategiesPerQuorum(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_StrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || arg1 == nil {
			return
		}

		_ContractStakeRegistry.StrategyParams(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_StrategyParamsByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.StrategyParamsByIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_StrategyParamsLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.StrategyParamsLength(quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_WEIGHTINGDIVISOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.WEIGHTINGDIVISOR()
	})
}

func Fuzz_Nosy_ContractStakeRegistryCallerSession_WeightOfOperatorForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryCallerSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.WeightOfOperatorForQuorum(quorumNumber, operator)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_FilterMinimumStakeForQuorumUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber []uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.FilterMinimumStakeForQuorumUpdated(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_FilterOperatorStakeUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.FilterOperatorStakeUpdate(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_FilterQuorumCreated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber []uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.FilterQuorumCreated(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_FilterStrategyAddedToQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber []uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.FilterStrategyAddedToQuorum(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_FilterStrategyMultiplierUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber []uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.FilterStrategyMultiplierUpdated(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_FilterStrategyRemovedFromQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumber []uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.FilterStrategyRemovedFromQuorum(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_ParseMinimumStakeForQuorumUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.ParseMinimumStakeForQuorumUpdated(log)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_ParseOperatorStakeUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.ParseOperatorStakeUpdate(log)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_ParseQuorumCreated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.ParseQuorumCreated(log)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_ParseStrategyAddedToQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.ParseStrategyAddedToQuorum(log)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_ParseStrategyMultiplierUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.ParseStrategyMultiplierUpdated(log)
	})
}

func Fuzz_Nosy_ContractStakeRegistryFilterer_ParseStrategyRemovedFromQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryFilterer
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.ParseStrategyRemovedFromQuorum(log)
	})
}

// skipping Fuzz_Nosy_ContractStakeRegistryFilterer_WatchMinimumStakeForQuorumUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/StakeRegistry.ContractStakeRegistryMinimumStakeForQuorumUpdated

// skipping Fuzz_Nosy_ContractStakeRegistryFilterer_WatchOperatorStakeUpdate__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/StakeRegistry.ContractStakeRegistryOperatorStakeUpdate

// skipping Fuzz_Nosy_ContractStakeRegistryFilterer_WatchQuorumCreated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/StakeRegistry.ContractStakeRegistryQuorumCreated

// skipping Fuzz_Nosy_ContractStakeRegistryFilterer_WatchStrategyAddedToQuorum__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/StakeRegistry.ContractStakeRegistryStrategyAddedToQuorum

// skipping Fuzz_Nosy_ContractStakeRegistryFilterer_WatchStrategyMultiplierUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/StakeRegistry.ContractStakeRegistryStrategyMultiplierUpdated

// skipping Fuzz_Nosy_ContractStakeRegistryFilterer_WatchStrategyRemovedFromQuorum__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/StakeRegistry.ContractStakeRegistryStrategyRemovedFromQuorum

func Fuzz_Nosy_ContractStakeRegistryMinimumStakeForQuorumUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryMinimumStakeForQuorumUpdatedIterator
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

func Fuzz_Nosy_ContractStakeRegistryMinimumStakeForQuorumUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryMinimumStakeForQuorumUpdatedIterator
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

func Fuzz_Nosy_ContractStakeRegistryMinimumStakeForQuorumUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryMinimumStakeForQuorumUpdatedIterator
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

func Fuzz_Nosy_ContractStakeRegistryOperatorStakeUpdateIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryOperatorStakeUpdateIterator
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

func Fuzz_Nosy_ContractStakeRegistryOperatorStakeUpdateIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryOperatorStakeUpdateIterator
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

func Fuzz_Nosy_ContractStakeRegistryOperatorStakeUpdateIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryOperatorStakeUpdateIterator
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

func Fuzz_Nosy_ContractStakeRegistryQuorumCreatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryQuorumCreatedIterator
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

func Fuzz_Nosy_ContractStakeRegistryQuorumCreatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryQuorumCreatedIterator
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

func Fuzz_Nosy_ContractStakeRegistryQuorumCreatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryQuorumCreatedIterator
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

// skipping Fuzz_Nosy_ContractStakeRegistryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractStakeRegistryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractStakeRegistryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryRaw
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_AddStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var _strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.AddStrategies(quorumNumber, _strategyParams)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.Delegation()
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.DeregisterOperator(operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetCurrentStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetCurrentStake(operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetCurrentTotalStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetCurrentTotalStake(quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetLatestStakeUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetLatestStakeUpdate(operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetStakeAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetStakeAtBlockNumber(operatorId, quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetStakeAtBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetStakeAtBlockNumberAndIndex(quorumNumber, blockNumber, operatorId, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetStakeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetStakeHistory(operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetStakeHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetStakeHistoryLength(operatorId, quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetStakeUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetStakeUpdateAtIndex(quorumNumber, operatorId, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetStakeUpdateIndexAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetStakeUpdateIndexAtBlockNumber(operatorId, quorumNumber, blockNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetTotalStakeAtBlockNumberFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeAtBlockNumberFromIndex(quorumNumber, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetTotalStakeHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeHistoryLength(quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetTotalStakeIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeIndicesAtBlockNumber(blockNumber, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_GetTotalStakeUpdateAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.GetTotalStakeUpdateAtIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		var _strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || minimumStake == nil {
			return
		}

		_ContractStakeRegistry.InitializeQuorum(quorumNumber, minimumStake, _strategyParams)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_MAXWEIGHINGFUNCTIONLENGTH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.MAXWEIGHINGFUNCTIONLENGTH()
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_MinimumStakeForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.MinimumStakeForQuorum(arg0)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_ModifyStrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var strategyIndices []*big.Int
		fill_err = tp.Fill(&strategyIndices)
		if fill_err != nil {
			return
		}
		var newMultipliers []*big.Int
		fill_err = tp.Fill(&newMultipliers)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.ModifyStrategyParams(quorumNumber, strategyIndices, newMultipliers)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.RegisterOperator(operator, operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_RemoveStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var indicesToRemove []*big.Int
		fill_err = tp.Fill(&indicesToRemove)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.RemoveStrategies(quorumNumber, indicesToRemove)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_SetMinimumStakeForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || minimumStake == nil {
			return
		}

		_ContractStakeRegistry.SetMinimumStakeForQuorum(quorumNumber, minimumStake)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_StrategiesPerQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || arg1 == nil {
			return
		}

		_ContractStakeRegistry.StrategiesPerQuorum(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_StrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || arg1 == nil {
			return
		}

		_ContractStakeRegistry.StrategyParams(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_StrategyParamsByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || index == nil {
			return
		}

		_ContractStakeRegistry.StrategyParamsByIndex(quorumNumber, index)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_StrategyParamsLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.StrategyParamsLength(quorumNumber)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_UpdateOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.UpdateOperatorStake(operator, operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_WEIGHTINGDIVISOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.WEIGHTINGDIVISOR()
	})
}

func Fuzz_Nosy_ContractStakeRegistrySession_WeightOfOperatorForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistrySession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.WeightOfOperatorForQuorum(quorumNumber, operator)
	})
}

func Fuzz_Nosy_ContractStakeRegistryStrategyAddedToQuorumIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyAddedToQuorumIterator
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

func Fuzz_Nosy_ContractStakeRegistryStrategyAddedToQuorumIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyAddedToQuorumIterator
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

func Fuzz_Nosy_ContractStakeRegistryStrategyAddedToQuorumIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyAddedToQuorumIterator
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

func Fuzz_Nosy_ContractStakeRegistryStrategyMultiplierUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyMultiplierUpdatedIterator
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

func Fuzz_Nosy_ContractStakeRegistryStrategyMultiplierUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyMultiplierUpdatedIterator
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

func Fuzz_Nosy_ContractStakeRegistryStrategyMultiplierUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyMultiplierUpdatedIterator
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

func Fuzz_Nosy_ContractStakeRegistryStrategyRemovedFromQuorumIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyRemovedFromQuorumIterator
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

func Fuzz_Nosy_ContractStakeRegistryStrategyRemovedFromQuorumIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyRemovedFromQuorumIterator
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

func Fuzz_Nosy_ContractStakeRegistryStrategyRemovedFromQuorumIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractStakeRegistryStrategyRemovedFromQuorumIterator
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

func Fuzz_Nosy_ContractStakeRegistryTransactor_AddStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactor
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var _strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.AddStrategies(opts, quorumNumber, _strategyParams)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactor_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactor
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.DeregisterOperator(opts, operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactor_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactor
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		var _strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil || minimumStake == nil {
			return
		}

		_ContractStakeRegistry.InitializeQuorum(opts, quorumNumber, minimumStake, _strategyParams)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactor_ModifyStrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactor
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var strategyIndices []*big.Int
		fill_err = tp.Fill(&strategyIndices)
		if fill_err != nil {
			return
		}
		var newMultipliers []*big.Int
		fill_err = tp.Fill(&newMultipliers)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.ModifyStrategyParams(opts, quorumNumber, strategyIndices, newMultipliers)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactor
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.RegisterOperator(opts, operator, operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactor_RemoveStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactor
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var indicesToRemove []*big.Int
		fill_err = tp.Fill(&indicesToRemove)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.RemoveStrategies(opts, quorumNumber, indicesToRemove)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactor_SetMinimumStakeForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactor
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil || minimumStake == nil {
			return
		}

		_ContractStakeRegistry.SetMinimumStakeForQuorum(opts, quorumNumber, minimumStake)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactor_UpdateOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactor
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.UpdateOperatorStake(opts, operator, operatorId, quorumNumbers)
	})
}

// skipping Fuzz_Nosy_ContractStakeRegistryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractStakeRegistryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorRaw
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || opts == nil {
			return
		}

		_ContractStakeRegistry.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactorSession_AddStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var _strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.AddStrategies(quorumNumber, _strategyParams)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactorSession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.DeregisterOperator(operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactorSession_InitializeQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		var _strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || minimumStake == nil {
			return
		}

		_ContractStakeRegistry.InitializeQuorum(quorumNumber, minimumStake, _strategyParams)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactorSession_ModifyStrategyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var strategyIndices []*big.Int
		fill_err = tp.Fill(&strategyIndices)
		if fill_err != nil {
			return
		}
		var newMultipliers []*big.Int
		fill_err = tp.Fill(&newMultipliers)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.ModifyStrategyParams(quorumNumber, strategyIndices, newMultipliers)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.RegisterOperator(operator, operatorId, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactorSession_RemoveStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var indicesToRemove []*big.Int
		fill_err = tp.Fill(&indicesToRemove)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.RemoveStrategies(quorumNumber, indicesToRemove)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactorSession_SetMinimumStakeForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		if _ContractStakeRegistry == nil || minimumStake == nil {
			return
		}

		_ContractStakeRegistry.SetMinimumStakeForQuorum(quorumNumber, minimumStake)
	})
}

func Fuzz_Nosy_ContractStakeRegistryTransactorSession_UpdateOperatorStake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractStakeRegistry *ContractStakeRegistryTransactorSession
		fill_err = tp.Fill(&_ContractStakeRegistry)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
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
		if _ContractStakeRegistry == nil {
			return
		}

		_ContractStakeRegistry.UpdateOperatorStake(operator, operatorId, quorumNumbers)
	})
}

// skipping Fuzz_Nosy_DeployContractStakeRegistry__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
