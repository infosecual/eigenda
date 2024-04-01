package contractRegistryCoordinator

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

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_BlsApkRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.BlsApkRegistry(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_CalculateOperatorChurnApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var registeringOperator common.Address
		fill_err = tp.Fill(&registeringOperator)
		if fill_err != nil {
			return
		}
		var registeringOperatorId [32]byte
		fill_err = tp.Fill(&registeringOperatorId)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IRegistryCoordinatorOperatorKickParam
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
		if _ContractRegistryCoordinator == nil || opts == nil || expiry == nil {
			return
		}

		_ContractRegistryCoordinator.CalculateOperatorChurnApprovalDigestHash(opts, registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_ChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.ChurnApprover(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_Ejector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.Ejector(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetCurrentQuorumBitmap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.GetCurrentQuorumBitmap(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperator(opts, operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetOperatorFromId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorFromId(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorId(opts, operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorSetParams(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorStatus(opts, operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetQuorumBitmapAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil || index == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapAtBlockNumberByIndex(opts, operatorId, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetQuorumBitmapHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapHistoryLength(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetQuorumBitmapIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapIndicesAtBlockNumber(opts, blockNumber, operatorIds)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_GetQuorumBitmapUpdateByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil || index == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapUpdateByIndex(opts, operatorId, index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_IndexRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.IndexRegistry(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_IsChurnApproverSaltUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.IsChurnApproverSaltUsed(opts, arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_NumRegistries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.NumRegistries(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_OPERATORCHURNAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.OPERATORCHURNAPPROVALTYPEHASH(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.Owner(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_PUBKEYREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.PUBKEYREGISTRATIONTYPEHASH(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.Paused(opts, index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.Paused0(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.PauserRegistry(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_PubkeyRegistrationMessageHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.PubkeyRegistrationMessageHash(opts, operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_QuorumCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.QuorumCount(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_QuorumUpdateBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.QuorumUpdateBlockNumber(opts, arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_Registries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil || arg0 == nil {
			return
		}

		_ContractRegistryCoordinator.Registries(opts, arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_ServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.ServiceManager(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCaller_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCaller
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.StakeRegistry(opts)
	})
}

// skipping Fuzz_Nosy_ContractRegistryCoordinatorCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_BlsApkRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.BlsApkRegistry()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_CalculateOperatorChurnApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var registeringOperator common.Address
		fill_err = tp.Fill(&registeringOperator)
		if fill_err != nil {
			return
		}
		var registeringOperatorId [32]byte
		fill_err = tp.Fill(&registeringOperatorId)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IRegistryCoordinatorOperatorKickParam
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
		if _ContractRegistryCoordinator == nil || expiry == nil {
			return
		}

		_ContractRegistryCoordinator.CalculateOperatorChurnApprovalDigestHash(registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_ChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ChurnApprover()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_Ejector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.Ejector()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetCurrentQuorumBitmap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetCurrentQuorumBitmap(operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperator(operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetOperatorFromId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorFromId(operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorId(operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorSetParams(quorumNumber)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorStatus(operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetQuorumBitmapAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || index == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapAtBlockNumberByIndex(operatorId, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetQuorumBitmapHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapHistoryLength(operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetQuorumBitmapIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapIndicesAtBlockNumber(blockNumber, operatorIds)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_GetQuorumBitmapUpdateByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || index == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapUpdateByIndex(operatorId, index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_IndexRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.IndexRegistry()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_IsChurnApproverSaltUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.IsChurnApproverSaltUsed(arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_NumRegistries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.NumRegistries()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_OPERATORCHURNAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.OPERATORCHURNAPPROVALTYPEHASH()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.Owner()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_PUBKEYREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.PUBKEYREGISTRATIONTYPEHASH()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.Paused(index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.Paused0()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_PubkeyRegistrationMessageHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.PubkeyRegistrationMessageHash(operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_QuorumCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.QuorumCount()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_QuorumUpdateBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.QuorumUpdateBlockNumber(arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_Registries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || arg0 == nil {
			return
		}

		_ContractRegistryCoordinator.Registries(arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_ServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ServiceManager()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorCallerSession_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorCallerSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.StakeRegistry()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorChurnApproverUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorChurnApproverUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorChurnApproverUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorChurnApproverUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorChurnApproverUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorChurnApproverUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorEjectorUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorEjectorUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorEjectorUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorEjectorUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorEjectorUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorEjectorUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterChurnApproverUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterChurnApproverUpdated(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterEjectorUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterEjectorUpdated(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterOperatorDeregistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterOperatorDeregistered(opts, operator, operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterOperatorRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterOperatorRegistered(opts, operator, operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterOperatorSetParamsUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterOperatorSetParamsUpdated(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterOperatorSocketUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterOperatorSocketUpdate(opts, operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterQuorumBlockNumberUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterQuorumBlockNumberUpdated(opts, quorumNumber)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseChurnApproverUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseChurnApproverUpdated(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseEjectorUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseEjectorUpdated(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseInitialized(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseOperatorDeregistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseOperatorDeregistered(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseOperatorRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseOperatorRegistered(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseOperatorSetParamsUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseOperatorSetParamsUpdated(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseOperatorSocketUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseOperatorSocketUpdate(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParsePaused(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseQuorumBlockNumberUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseQuorumBlockNumberUpdated(log)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorFilterer
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ParseUnpaused(log)
	})
}

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchChurnApproverUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorChurnApproverUpdated

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchEjectorUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorEjectorUpdated

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorInitialized

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchOperatorDeregistered__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorOperatorDeregistered

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchOperatorRegistered__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorOperatorRegistered

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchOperatorSetParamsUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorOperatorSetParamsUpdated

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchOperatorSocketUpdate__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorOperatorSocketUpdate

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorOwnershipTransferred

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorPaused

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorPauserRegistrySet

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchQuorumBlockNumberUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorQuorumBlockNumberUpdated

// skipping Fuzz_Nosy_ContractRegistryCoordinatorFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/RegistryCoordinator.ContractRegistryCoordinatorUnpaused

func Fuzz_Nosy_ContractRegistryCoordinatorInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorInitializedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorInitializedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorInitializedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorDeregisteredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorDeregisteredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorDeregisteredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorDeregisteredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorDeregisteredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorDeregisteredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorRegisteredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorRegisteredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorRegisteredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorSetParamsUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorSocketUpdateIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorSocketUpdateIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorSocketUpdateIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorSocketUpdateIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOperatorSocketUpdateIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOperatorSocketUpdateIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorPausedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorPausedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorPausedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorPauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorPauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorPauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorQuorumBlockNumberUpdatedIterator
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

// skipping Fuzz_Nosy_ContractRegistryCoordinatorRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractRegistryCoordinatorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractRegistryCoordinatorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorRaw
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_BlsApkRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.BlsApkRegistry()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_CalculateOperatorChurnApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var registeringOperator common.Address
		fill_err = tp.Fill(&registeringOperator)
		if fill_err != nil {
			return
		}
		var registeringOperatorId [32]byte
		fill_err = tp.Fill(&registeringOperatorId)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IRegistryCoordinatorOperatorKickParam
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
		if _ContractRegistryCoordinator == nil || expiry == nil {
			return
		}

		_ContractRegistryCoordinator.CalculateOperatorChurnApprovalDigestHash(registeringOperator, registeringOperatorId, operatorKickParams, salt, expiry)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_ChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ChurnApprover()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_CreateQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorSetParams IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&operatorSetParams)
		if fill_err != nil {
			return
		}
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		var strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || minimumStake == nil {
			return
		}

		_ContractRegistryCoordinator.CreateQuorum(operatorSetParams, minimumStake, strategyParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.DeregisterOperator(quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_EjectOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.EjectOperator(operator, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_Ejector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.Ejector()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetCurrentQuorumBitmap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetCurrentQuorumBitmap(operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperator(operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetOperatorFromId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorFromId(operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorId(operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorSetParams(quorumNumber)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetOperatorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetOperatorStatus(operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetQuorumBitmapAtBlockNumberByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || index == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapAtBlockNumberByIndex(operatorId, blockNumber, index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetQuorumBitmapHistoryLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapHistoryLength(operatorId)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetQuorumBitmapIndicesAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapIndicesAtBlockNumber(blockNumber, operatorIds)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_GetQuorumBitmapUpdateByIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || index == nil {
			return
		}

		_ContractRegistryCoordinator.GetQuorumBitmapUpdateByIndex(operatorId, index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_IndexRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.IndexRegistry()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var _initialOwner common.Address
		fill_err = tp.Fill(&_initialOwner)
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
		var _operatorSetParams []IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&_operatorSetParams)
		if fill_err != nil {
			return
		}
		var _minimumStakes []*big.Int
		fill_err = tp.Fill(&_minimumStakes)
		if fill_err != nil {
			return
		}
		var _strategyParams [][]IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || _initialPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Initialize(_initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_IsChurnApproverSaltUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.IsChurnApproverSaltUsed(arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_NumRegistries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.NumRegistries()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_OPERATORCHURNAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.OPERATORCHURNAPPROVALTYPEHASH()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.Owner()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_PUBKEYREGISTRATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.PUBKEYREGISTRATIONTYPEHASH()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || newPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.PauseAll()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.Paused(index)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.Paused0()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_PubkeyRegistrationMessageHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.PubkeyRegistrationMessageHash(operator)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_QuorumCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.QuorumCount()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_QuorumUpdateBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var arg0 uint8
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.QuorumUpdateBlockNumber(arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.RegisterOperator(quorumNumbers, socket, params, operatorSignature)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_RegisterOperatorWithChurn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IRegistryCoordinatorOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
		if fill_err != nil {
			return
		}
		var churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&churnApproverSignature)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.RegisterOperatorWithChurn(quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_Registries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var arg0 *big.Int
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || arg0 == nil {
			return
		}

		_ContractRegistryCoordinator.Registries(arg0)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.RenounceOwnership()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_ServiceManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.ServiceManager()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_SetChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var _churnApprover common.Address
		fill_err = tp.Fill(&_churnApprover)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.SetChurnApprover(_churnApprover)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_SetEjector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var _ejector common.Address
		fill_err = tp.Fill(&_ejector)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.SetEjector(_ejector)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_SetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorSetParams IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&operatorSetParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.SetOperatorSetParams(quorumNumber, operatorSetParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.StakeRegistry()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || newPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_UpdateOperators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operators []common.Address
		fill_err = tp.Fill(&operators)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateOperators(operators)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_UpdateOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorsPerQuorum [][]common.Address
		fill_err = tp.Fill(&operatorsPerQuorum)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateOperatorsForQuorum(operatorsPerQuorum, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorSession_UpdateSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateSocket(socket)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_CreateQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operatorSetParams IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&operatorSetParams)
		if fill_err != nil {
			return
		}
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		var strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil || minimumStake == nil {
			return
		}

		_ContractRegistryCoordinator.CreateQuorum(opts, operatorSetParams, minimumStake, strategyParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.DeregisterOperator(opts, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_EjectOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.EjectOperator(opts, operator, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _initialOwner common.Address
		fill_err = tp.Fill(&_initialOwner)
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
		var _operatorSetParams []IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&_operatorSetParams)
		if fill_err != nil {
			return
		}
		var _minimumStakes []*big.Int
		fill_err = tp.Fill(&_minimumStakes)
		if fill_err != nil {
			return
		}
		var _strategyParams [][]IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil || _initialPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Initialize(opts, _initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.PauseAll(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.RegisterOperator(opts, quorumNumbers, socket, params, operatorSignature)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_RegisterOperatorWithChurn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IRegistryCoordinatorOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
		if fill_err != nil {
			return
		}
		var churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&churnApproverSignature)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.RegisterOperatorWithChurn(opts, quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_SetChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.SetChurnApprover(opts, _churnApprover)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_SetEjector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.SetEjector(opts, _ejector)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_SetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		var operatorSetParams IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&operatorSetParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.SetOperatorSetParams(opts, quorumNumber, operatorSetParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_UpdateOperators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operators []common.Address
		fill_err = tp.Fill(&operators)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateOperators(opts, operators)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_UpdateOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var operatorsPerQuorum [][]common.Address
		fill_err = tp.Fill(&operatorsPerQuorum)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateOperatorsForQuorum(opts, operatorsPerQuorum, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactor_UpdateSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactor
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateSocket(opts, socket)
	})
}

// skipping Fuzz_Nosy_ContractRegistryCoordinatorTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorRaw
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || opts == nil {
			return
		}

		_ContractRegistryCoordinator.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_CreateQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorSetParams IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&operatorSetParams)
		if fill_err != nil {
			return
		}
		var minimumStake *big.Int
		fill_err = tp.Fill(&minimumStake)
		if fill_err != nil {
			return
		}
		var strategyParams []IStakeRegistryStrategyParams
		fill_err = tp.Fill(&strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || minimumStake == nil {
			return
		}

		_ContractRegistryCoordinator.CreateQuorum(operatorSetParams, minimumStake, strategyParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.DeregisterOperator(quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_EjectOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
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
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.EjectOperator(operator, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var _initialOwner common.Address
		fill_err = tp.Fill(&_initialOwner)
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
		var _operatorSetParams []IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&_operatorSetParams)
		if fill_err != nil {
			return
		}
		var _minimumStakes []*big.Int
		fill_err = tp.Fill(&_minimumStakes)
		if fill_err != nil {
			return
		}
		var _strategyParams [][]IStakeRegistryStrategyParams
		fill_err = tp.Fill(&_strategyParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || _initialPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Initialize(_initialOwner, _churnApprover, _ejector, _pauserRegistry, _initialPausedStatus, _operatorSetParams, _minimumStakes, _strategyParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || newPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.PauseAll()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.RegisterOperator(quorumNumbers, socket, params, operatorSignature)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_RegisterOperatorWithChurn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var params IBLSApkRegistryPubkeyRegistrationParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var operatorKickParams []IRegistryCoordinatorOperatorKickParam
		fill_err = tp.Fill(&operatorKickParams)
		if fill_err != nil {
			return
		}
		var churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&churnApproverSignature)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.RegisterOperatorWithChurn(quorumNumbers, socket, params, operatorKickParams, churnApproverSignature, operatorSignature)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.RenounceOwnership()
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_SetChurnApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var _churnApprover common.Address
		fill_err = tp.Fill(&_churnApprover)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.SetChurnApprover(_churnApprover)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_SetEjector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var _ejector common.Address
		fill_err = tp.Fill(&_ejector)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.SetEjector(_ejector)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_SetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var quorumNumber uint8
		fill_err = tp.Fill(&quorumNumber)
		if fill_err != nil {
			return
		}
		var operatorSetParams IRegistryCoordinatorOperatorSetParam
		fill_err = tp.Fill(&operatorSetParams)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.SetOperatorSetParams(quorumNumber, operatorSetParams)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil || newPausedStatus == nil {
			return
		}

		_ContractRegistryCoordinator.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_UpdateOperators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operators []common.Address
		fill_err = tp.Fill(&operators)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateOperators(operators)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_UpdateOperatorsForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var operatorsPerQuorum [][]common.Address
		fill_err = tp.Fill(&operatorsPerQuorum)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateOperatorsForQuorum(operatorsPerQuorum, quorumNumbers)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorTransactorSession_UpdateSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractRegistryCoordinator *ContractRegistryCoordinatorTransactorSession
		fill_err = tp.Fill(&_ContractRegistryCoordinator)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if _ContractRegistryCoordinator == nil {
			return
		}

		_ContractRegistryCoordinator.UpdateSocket(socket)
	})
}

func Fuzz_Nosy_ContractRegistryCoordinatorUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorUnpausedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorUnpausedIterator
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

func Fuzz_Nosy_ContractRegistryCoordinatorUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractRegistryCoordinatorUnpausedIterator
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

// skipping Fuzz_Nosy_DeployContractRegistryCoordinator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
