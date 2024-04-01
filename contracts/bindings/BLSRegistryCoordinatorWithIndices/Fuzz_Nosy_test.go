package contractBLSRegistryCoordinatorWithIndices

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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_BlsPubkeyRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.BlsPubkeyRegistry(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_CalculateOperatorChurnApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var registeringOperatorId [32]byte
		fill_err = tp.Fill(&registeringOperatorId)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil || expiry == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.CalculateOperatorChurnApprovalDigestHash(opts, registeringOperatorId, operatorKickParams, salt, expiry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_ChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ChurnApprover(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_Ejector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Ejector(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetCurrentQuorumBitmapByOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetCurrentQuorumBitmapByOperatorId(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperator(opts, operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetOperatorFromId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorFromId(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorId(opts, operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorSetParams(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorStatus(opts, operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetQuorumBitmapByOperatorIdAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil || index == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapByOperatorIdAtBlockNumberByIndex(opts, operatorId, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var operatorIds [][32]byte
		fill_err = tp.Fill(&operatorIds)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber(opts, blockNumber, operatorIds)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetQuorumBitmapUpdateByOperatorIdByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil || index == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapUpdateByOperatorIdByIndex(opts, operatorId, index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_GetQuorumBitmapUpdateByOperatorIdLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapUpdateByOperatorIdLength(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_IndexRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.IndexRegistry(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_IsChurnApproverSaltUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.IsChurnApproverSaltUsed(opts, arg0)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_NumRegistries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.NumRegistries(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_OPERATORCHURNAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.OPERATORCHURNAPPROVALTYPEHASH(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Paused(opts, index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Paused0(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.PauserRegistry(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_Registries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil || arg0 == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Registries(opts, arg0)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_ServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ServiceManager(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Slasher(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCaller_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCaller
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.StakeRegistry(opts)
	})
}

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_BlsPubkeyRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.BlsPubkeyRegistry()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_CalculateOperatorChurnApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var registeringOperatorId [32]byte
		fill_err = tp.Fill(&registeringOperatorId)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || expiry == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.CalculateOperatorChurnApprovalDigestHash(registeringOperatorId, operatorKickParams, salt, expiry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_ChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ChurnApprover()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_Ejector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Ejector()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetCurrentQuorumBitmapByOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetCurrentQuorumBitmapByOperatorId(operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperator(operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetOperatorFromId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorFromId(operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorId(operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorSetParams(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorStatus(operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetQuorumBitmapByOperatorIdAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || index == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapByOperatorIdAtBlockNumberByIndex(operatorId, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var operatorIds [][32]byte
		fill_err = tp.Fill(&operatorIds)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber(blockNumber, operatorIds)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetQuorumBitmapUpdateByOperatorIdByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || index == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapUpdateByOperatorIdByIndex(operatorId, index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_GetQuorumBitmapUpdateByOperatorIdLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapUpdateByOperatorIdLength(operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_IndexRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.IndexRegistry()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_IsChurnApproverSaltUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.IsChurnApproverSaltUsed(arg0)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_NumRegistries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.NumRegistries()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_OPERATORCHURNAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.OPERATORCHURNAPPROVALTYPEHASH()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Paused(index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Paused0()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_Registries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || arg0 == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Registries(arg0)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_ServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ServiceManager()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Slasher()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesCallerSession_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesCallerSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.StakeRegistry()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesChurnApproverUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesChurnApproverUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesChurnApproverUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesChurnApproverUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesChurnApproverUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesChurnApproverUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesEjectorUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesEjectorUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesEjectorUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesEjectorUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesEjectorUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesEjectorUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterChurnApproverUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterChurnApproverUpdated(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterEjectorUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterEjectorUpdated(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterOperatorDeregistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var operatorId [][32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterOperatorDeregistered(opts, operator, operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterOperatorRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var operatorId [][32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterOperatorRegistered(opts, operator, operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterOperatorSetParamsUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterOperatorSetParamsUpdated(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterOperatorSocketUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterOperatorSocketUpdate(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParseChurnApproverUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParseChurnApproverUpdated(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParseEjectorUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParseEjectorUpdated(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParseInitialized(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParseOperatorDeregistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParseOperatorDeregistered(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParseOperatorRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParseOperatorRegistered(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParseOperatorSetParamsUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParseOperatorSetParamsUpdated(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParseOperatorSocketUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParseOperatorSocketUpdate(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParsePaused(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesFilterer
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ParseUnpaused(log)
	})
}

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchChurnApproverUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesChurnApproverUpdated

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchEjectorUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesEjectorUpdated

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesInitialized

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchOperatorDeregistered__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesOperatorDeregistered

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchOperatorRegistered__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesOperatorRegistered

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchOperatorSetParamsUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchOperatorSocketUpdate__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesOperatorSocketUpdate

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesPaused

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesPauserRegistrySet

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSRegistryCoordinatorWithIndices.ContractBLSRegistryCoordinatorWithIndicesUnpaused

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesInitializedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesInitializedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesInitializedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesPausedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesPausedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesPausedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesPauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesPauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesPauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesPauserRegistrySetIterator
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

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesRaw
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_BlsPubkeyRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.BlsPubkeyRegistry()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_CalculateOperatorChurnApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var registeringOperatorId [32]byte
		fill_err = tp.Fill(&registeringOperatorId)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || expiry == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.CalculateOperatorChurnApprovalDigestHash(registeringOperatorId, operatorKickParams, salt, expiry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_ChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ChurnApprover()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_DeregisterOperatorWithCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.DeregisterOperatorWithCoordinator(quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_DeregisterOperatorWithCoordinator0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var deregistrationData []byte
		fill_err = tp.Fill(&deregistrationData)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.DeregisterOperatorWithCoordinator0(quorumNumbers, deregistrationData)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_EjectOperatorFromCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.EjectOperatorFromCoordinator(operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_Ejector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Ejector()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetCurrentQuorumBitmapByOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetCurrentQuorumBitmapByOperatorId(operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperator(operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetOperatorFromId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorFromId(operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorId(operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorSetParams(quorumNumber)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetOperatorStatus(operator)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetQuorumBitmapByOperatorIdAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var index *big.Int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || index == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapByOperatorIdAtBlockNumberByIndex(operatorId, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var operatorIds [][32]byte
		fill_err = tp.Fill(&operatorIds)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber(blockNumber, operatorIds)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetQuorumBitmapUpdateByOperatorIdByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || index == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapUpdateByOperatorIdByIndex(operatorId, index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_GetQuorumBitmapUpdateByOperatorIdLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.GetQuorumBitmapUpdateByOperatorIdLength(operatorId)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_IndexRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.IndexRegistry()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var _churnApprover common.Address
		fill_err = tp.Fill(&_churnApprover)
		if fill_err != nil {
			return
		}
		var _ejector common.Address
		fill_err = tp.Fill(&_ejector)
		if fill_err != nil {
			return
		}
		var _operatorSetParams []IBLSRegistryCoordinatorWithIndicesOperatorSetParam
		fill_err = tp.Fill(&_operatorSetParams)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var _initialPausedStatus *big.Int
		fill_err = tp.Fill(&_initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || _initialPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Initialize(_churnApprover, _ejector, _operatorSetParams, _pauserRegistry, _initialPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_IsChurnApproverSaltUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.IsChurnApproverSaltUsed(arg0)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_NumRegistries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.NumRegistries()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_OPERATORCHURNAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.OPERATORCHURNAPPROVALTYPEHASH()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || newPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.PauseAll()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Paused(index)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Paused0()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_RegisterOperatorWithCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
		if fill_err != nil {
			return
		}
		var signatureWithSaltAndExpiry ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&signatureWithSaltAndExpiry)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator(quorumNumbers, pubkey, socket, operatorKickParams, signatureWithSaltAndExpiry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_RegisterOperatorWithCoordinator0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var registrationData []byte
		fill_err = tp.Fill(&registrationData)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator0(quorumNumbers, registrationData)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_RegisterOperatorWithCoordinator1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator1(quorumNumbers, pubkey, socket)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_Registries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || arg0 == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Registries(arg0)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_ServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.ServiceManager()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_SetChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var _churnApprover common.Address
		fill_err = tp.Fill(&_churnApprover)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetChurnApprover(_churnApprover)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_SetEjector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var _ejector common.Address
		fill_err = tp.Fill(&_ejector)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetEjector(_ejector)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_SetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorSetParam IBLSRegistryCoordinatorWithIndicesOperatorSetParam
		fill_err = tp.Fill(&operatorSetParam)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetOperatorSetParams(quorumNumber, operatorSetParam)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Slasher()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.StakeRegistry()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || newPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesSession_UpdateSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.UpdateSocket(socket)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_DeregisterOperatorWithCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.DeregisterOperatorWithCoordinator(opts, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_DeregisterOperatorWithCoordinator0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var deregistrationData []byte
		fill_err = tp.Fill(&deregistrationData)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.DeregisterOperatorWithCoordinator0(opts, quorumNumbers, deregistrationData)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_EjectOperatorFromCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.EjectOperatorFromCoordinator(opts, operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _churnApprover common.Address
		fill_err = tp.Fill(&_churnApprover)
		if fill_err != nil {
			return
		}
		var _ejector common.Address
		fill_err = tp.Fill(&_ejector)
		if fill_err != nil {
			return
		}
		var _operatorSetParams []IBLSRegistryCoordinatorWithIndicesOperatorSetParam
		fill_err = tp.Fill(&_operatorSetParams)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var _initialPausedStatus *big.Int
		fill_err = tp.Fill(&_initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil || _initialPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Initialize(opts, _churnApprover, _ejector, _operatorSetParams, _pauserRegistry, _initialPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.PauseAll(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_RegisterOperatorWithCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
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
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
		if fill_err != nil {
			return
		}
		var signatureWithSaltAndExpiry ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&signatureWithSaltAndExpiry)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator(opts, quorumNumbers, pubkey, socket, operatorKickParams, signatureWithSaltAndExpiry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_RegisterOperatorWithCoordinator0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var registrationData []byte
		fill_err = tp.Fill(&registrationData)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator0(opts, quorumNumbers, registrationData)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_RegisterOperatorWithCoordinator1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
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
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator1(opts, quorumNumbers, pubkey, socket)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_SetChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _churnApprover common.Address
		fill_err = tp.Fill(&_churnApprover)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetChurnApprover(opts, _churnApprover)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_SetEjector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _ejector common.Address
		fill_err = tp.Fill(&_ejector)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetEjector(opts, _ejector)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_SetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var operatorSetParam IBLSRegistryCoordinatorWithIndicesOperatorSetParam
		fill_err = tp.Fill(&operatorSetParam)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetOperatorSetParams(opts, quorumNumber, operatorSetParam)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactor_UpdateSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactor
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.UpdateSocket(opts, socket)
	})
}

// skipping Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorRaw
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || opts == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_DeregisterOperatorWithCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.DeregisterOperatorWithCoordinator(quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_DeregisterOperatorWithCoordinator0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var deregistrationData []byte
		fill_err = tp.Fill(&deregistrationData)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.DeregisterOperatorWithCoordinator0(quorumNumbers, deregistrationData)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_EjectOperatorFromCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.EjectOperatorFromCoordinator(operator, quorumNumbers, pubkey)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var _churnApprover common.Address
		fill_err = tp.Fill(&_churnApprover)
		if fill_err != nil {
			return
		}
		var _ejector common.Address
		fill_err = tp.Fill(&_ejector)
		if fill_err != nil {
			return
		}
		var _operatorSetParams []IBLSRegistryCoordinatorWithIndicesOperatorSetParam
		fill_err = tp.Fill(&_operatorSetParams)
		if fill_err != nil {
			return
		}
		var _pauserRegistry common.Address
		fill_err = tp.Fill(&_pauserRegistry)
		if fill_err != nil {
			return
		}
		var _initialPausedStatus *big.Int
		fill_err = tp.Fill(&_initialPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || _initialPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Initialize(_churnApprover, _ejector, _operatorSetParams, _pauserRegistry, _initialPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || newPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.PauseAll()
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_RegisterOperatorWithCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
		if fill_err != nil {
			return
		}
		var signatureWithSaltAndExpiry ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&signatureWithSaltAndExpiry)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator(quorumNumbers, pubkey, socket, operatorKickParams, signatureWithSaltAndExpiry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_RegisterOperatorWithCoordinator0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var registrationData []byte
		fill_err = tp.Fill(&registrationData)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator0(quorumNumbers, registrationData)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_RegisterOperatorWithCoordinator1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
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
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.RegisterOperatorWithCoordinator1(quorumNumbers, pubkey, socket)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_SetChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var _churnApprover common.Address
		fill_err = tp.Fill(&_churnApprover)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetChurnApprover(_churnApprover)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_SetEjector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var _ejector common.Address
		fill_err = tp.Fill(&_ejector)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetEjector(_ejector)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_SetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorSetParam IBLSRegistryCoordinatorWithIndicesOperatorSetParam
		fill_err = tp.Fill(&operatorSetParam)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetOperatorSetParams(quorumNumber, operatorSetParam)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil || newPausedStatus == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesTransactorSession_UpdateSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSRegistryCoordinatorWithIndices *ContractBLSRegistryCoordinatorWithIndicesTransactorSession
		fill_err = tp.Fill(&_ContractBLSRegistryCoordinatorWithIndices)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if _ContractBLSRegistryCoordinatorWithIndices == nil {
			return
		}

		_ContractBLSRegistryCoordinatorWithIndices.UpdateSocket(socket)
	})
}

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesUnpausedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesUnpausedIterator
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

func Fuzz_Nosy_ContractBLSRegistryCoordinatorWithIndicesUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSRegistryCoordinatorWithIndicesUnpausedIterator
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

// skipping Fuzz_Nosy_DeployContractBLSRegistryCoordinatorWithIndices__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
