package contractDelegationManager

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

func Fuzz_Nosy_ContractDelegationManagerCaller_BeaconChainETHStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.BeaconChainETHStrategy(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_CalculateCurrentStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateCurrentStakerDelegationDigestHash(opts, staker, operator, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_CalculateDelegationApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var _delegationApprover common.Address
		fill_err = tp.Fill(&_delegationApprover)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateDelegationApprovalDigestHash(opts, staker, operator, _delegationApprover, approverSalt, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_CalculateStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var _stakerNonce *big.Int
		fill_err = tp.Fill(&_stakerNonce)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil || _stakerNonce == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateStakerDelegationDigestHash(opts, staker, _stakerNonce, operator, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.CalculateWithdrawalRoot(opts, withdrawal)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_CumulativeWithdrawalsQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.CumulativeWithdrawalsQueued(opts, arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_DELEGATIONAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.DELEGATIONAPPROVALTYPEHASH(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.DOMAINTYPEHASH(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_DelegatedTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.DelegatedTo(opts, arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_DelegationApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.DelegationApprover(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_DelegationApproverSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.DelegationApproverSaltIsSpent(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.DomainSeparator(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_EarningsReceiver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.EarningsReceiver(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.EigenPodManager(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_GetDelegatableShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.GetDelegatableShares(opts, staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_GetOperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.GetOperatorShares(opts, operator, strategies)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_GetWithdrawalDelay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.GetWithdrawalDelay(opts, strategies)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_IsDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.IsDelegated(opts, staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_IsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.IsOperator(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_MAXSTAKEROPTOUTWINDOWBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.MAXSTAKEROPTOUTWINDOWBLOCKS(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_MAXWITHDRAWALDELAYBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.MAXWITHDRAWALDELAYBLOCKS(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_MinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.MinWithdrawalDelayBlocks(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_OperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.OperatorDetails(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_OperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.OperatorShares(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.Owner(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.Paused(opts, index)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.Paused0(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.PauserRegistry(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_PendingWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.PendingWithdrawals(opts, arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_STAKERDELEGATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.STAKERDELEGATIONTYPEHASH(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.Slasher(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_StakerNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.StakerNonce(opts, arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_StakerOptOutWindowBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.StakerOptOutWindowBlocks(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.StrategyManager(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCaller_StrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCaller
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.StrategyWithdrawalDelayBlocks(opts, arg0)
	})
}

// skipping Fuzz_Nosy_ContractDelegationManagerCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_BeaconChainETHStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.BeaconChainETHStrategy()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_CalculateCurrentStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateCurrentStakerDelegationDigestHash(staker, operator, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_CalculateDelegationApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var _delegationApprover common.Address
		fill_err = tp.Fill(&_delegationApprover)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateDelegationApprovalDigestHash(staker, operator, _delegationApprover, approverSalt, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_CalculateStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var _stakerNonce *big.Int
		fill_err = tp.Fill(&_stakerNonce)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || _stakerNonce == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateStakerDelegationDigestHash(staker, _stakerNonce, operator, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.CalculateWithdrawalRoot(withdrawal)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_CumulativeWithdrawalsQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.CumulativeWithdrawalsQueued(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_DELEGATIONAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DELEGATIONAPPROVALTYPEHASH()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_DelegatedTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegatedTo(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_DelegationApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegationApprover(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_DelegationApproverSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegationApproverSaltIsSpent(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DomainSeparator()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_EarningsReceiver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.EarningsReceiver(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.EigenPodManager()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_GetDelegatableShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.GetDelegatableShares(staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_GetOperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.GetOperatorShares(operator, strategies)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_GetWithdrawalDelay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.GetWithdrawalDelay(strategies)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_IsDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.IsDelegated(staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_IsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.IsOperator(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_MAXSTAKEROPTOUTWINDOWBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.MAXSTAKEROPTOUTWINDOWBLOCKS()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_MAXWITHDRAWALDELAYBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.MAXWITHDRAWALDELAYBLOCKS()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_MinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.MinWithdrawalDelayBlocks()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_OperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.OperatorDetails(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_OperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.OperatorShares(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Owner()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Paused(index)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Paused0()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_PendingWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.PendingWithdrawals(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_STAKERDELEGATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.STAKERDELEGATIONTYPEHASH()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Slasher()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_StakerNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.StakerNonce(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_StakerOptOutWindowBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.StakerOptOutWindowBlocks(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.StrategyManager()
	})
}

func Fuzz_Nosy_ContractDelegationManagerCallerSession_StrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerCallerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.StrategyWithdrawalDelayBlocks(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterMinWithdrawalDelayBlocksSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterMinWithdrawalDelayBlocksSet(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterOperatorDetailsModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterOperatorDetailsModified(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterOperatorMetadataURIUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterOperatorMetadataURIUpdated(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterOperatorRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterOperatorRegistered(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterOperatorSharesDecreased__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterOperatorSharesDecreased(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterOperatorSharesIncreased__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterOperatorSharesIncreased(opts, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterStakerDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker []common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterStakerDelegated(opts, staker, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterStakerForceUndelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker []common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterStakerForceUndelegated(opts, staker, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterStakerUndelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker []common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator []common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterStakerUndelegated(opts, staker, operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterStrategyWithdrawalDelayBlocksSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterStrategyWithdrawalDelayBlocksSet(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterWithdrawalCompleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterWithdrawalCompleted(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterWithdrawalMigrated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterWithdrawalMigrated(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_FilterWithdrawalQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.FilterWithdrawalQueued(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseInitialized(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseMinWithdrawalDelayBlocksSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseMinWithdrawalDelayBlocksSet(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseOperatorDetailsModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseOperatorDetailsModified(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseOperatorMetadataURIUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseOperatorMetadataURIUpdated(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseOperatorRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseOperatorRegistered(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseOperatorSharesDecreased__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseOperatorSharesDecreased(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseOperatorSharesIncreased__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseOperatorSharesIncreased(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParsePaused(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseStakerDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseStakerDelegated(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseStakerForceUndelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseStakerForceUndelegated(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseStakerUndelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseStakerUndelegated(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseStrategyWithdrawalDelayBlocksSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseStrategyWithdrawalDelayBlocksSet(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseUnpaused(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseWithdrawalCompleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseWithdrawalCompleted(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseWithdrawalMigrated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseWithdrawalMigrated(log)
	})
}

func Fuzz_Nosy_ContractDelegationManagerFilterer_ParseWithdrawalQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerFilterer
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ParseWithdrawalQueued(log)
	})
}

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerInitialized

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchMinWithdrawalDelayBlocksSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerMinWithdrawalDelayBlocksSet

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchOperatorDetailsModified__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerOperatorDetailsModified

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchOperatorMetadataURIUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerOperatorMetadataURIUpdated

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchOperatorRegistered__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerOperatorRegistered

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchOperatorSharesDecreased__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerOperatorSharesDecreased

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchOperatorSharesIncreased__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerOperatorSharesIncreased

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerOwnershipTransferred

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerPaused

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerPauserRegistrySet

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchStakerDelegated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerStakerDelegated

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchStakerForceUndelegated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerStakerForceUndelegated

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchStakerUndelegated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerStakerUndelegated

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchStrategyWithdrawalDelayBlocksSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerStrategyWithdrawalDelayBlocksSet

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerUnpaused

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchWithdrawalCompleted__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerWithdrawalCompleted

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchWithdrawalMigrated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerWithdrawalMigrated

// skipping Fuzz_Nosy_ContractDelegationManagerFilterer_WatchWithdrawalQueued__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/DelegationManager.ContractDelegationManagerWithdrawalQueued

func Fuzz_Nosy_ContractDelegationManagerInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerInitializedIterator
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

func Fuzz_Nosy_ContractDelegationManagerInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerInitializedIterator
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

func Fuzz_Nosy_ContractDelegationManagerInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerInitializedIterator
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

func Fuzz_Nosy_ContractDelegationManagerMinWithdrawalDelayBlocksSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerMinWithdrawalDelayBlocksSetIterator
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

func Fuzz_Nosy_ContractDelegationManagerMinWithdrawalDelayBlocksSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerMinWithdrawalDelayBlocksSetIterator
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

func Fuzz_Nosy_ContractDelegationManagerMinWithdrawalDelayBlocksSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerMinWithdrawalDelayBlocksSetIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorDetailsModifiedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorDetailsModifiedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorDetailsModifiedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorDetailsModifiedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorDetailsModifiedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorDetailsModifiedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorMetadataURIUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorMetadataURIUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorMetadataURIUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorRegisteredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorRegisteredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorRegisteredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorRegisteredIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorSharesDecreasedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorSharesDecreasedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorSharesDecreasedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorSharesDecreasedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorSharesDecreasedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorSharesDecreasedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorSharesIncreasedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorSharesIncreasedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorSharesIncreasedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorSharesIncreasedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOperatorSharesIncreasedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOperatorSharesIncreasedIterator
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

func Fuzz_Nosy_ContractDelegationManagerOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractDelegationManagerOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractDelegationManagerOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractDelegationManagerPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerPausedIterator
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

func Fuzz_Nosy_ContractDelegationManagerPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerPausedIterator
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

func Fuzz_Nosy_ContractDelegationManagerPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerPausedIterator
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

func Fuzz_Nosy_ContractDelegationManagerPauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractDelegationManagerPauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractDelegationManagerPauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerPauserRegistrySetIterator
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

// skipping Fuzz_Nosy_ContractDelegationManagerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractDelegationManagerRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractDelegationManagerRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerRaw
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_BeaconChainETHStrategy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.BeaconChainETHStrategy()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_CalculateCurrentStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateCurrentStakerDelegationDigestHash(staker, operator, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_CalculateDelegationApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var _delegationApprover common.Address
		fill_err = tp.Fill(&_delegationApprover)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateDelegationApprovalDigestHash(staker, operator, _delegationApprover, approverSalt, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_CalculateStakerDelegationDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var _stakerNonce *big.Int
		fill_err = tp.Fill(&_stakerNonce)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var expiry *big.Int
		fill_err = tp.Fill(&expiry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || _stakerNonce == nil || expiry == nil {
			return
		}

		_ContractDelegationManager.CalculateStakerDelegationDigestHash(staker, _stakerNonce, operator, expiry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_CalculateWithdrawalRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.CalculateWithdrawalRoot(withdrawal)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_CompleteQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		var tokens []common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndex *big.Int
		fill_err = tp.Fill(&middlewareTimesIndex)
		if fill_err != nil {
			return
		}
		var receiveAsTokens bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || middlewareTimesIndex == nil {
			return
		}

		_ContractDelegationManager.CompleteQueuedWithdrawal(withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_CompleteQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawals []IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}
		var tokens [][]common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndexes []*big.Int
		fill_err = tp.Fill(&middlewareTimesIndexes)
		if fill_err != nil {
			return
		}
		var receiveAsTokens []bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.CompleteQueuedWithdrawals(withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_CumulativeWithdrawalsQueued__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.CumulativeWithdrawalsQueued(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DELEGATIONAPPROVALTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DELEGATIONAPPROVALTYPEHASH()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DOMAINTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DOMAINTYPEHASH()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DecreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || shares == nil {
			return
		}

		_ContractDelegationManager.DecreaseDelegatedShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DelegateTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegateTo(operator, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DelegateToBySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&stakerSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegateToBySignature(staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DelegatedTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegatedTo(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DelegationApprover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegationApprover(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DelegationApproverSaltIsSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegationApproverSaltIsSpent(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_DomainSeparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DomainSeparator()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_EarningsReceiver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.EarningsReceiver(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_EigenPodManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.EigenPodManager()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_GetDelegatableShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.GetDelegatableShares(staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_GetOperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.GetOperatorShares(operator, strategies)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_GetWithdrawalDelay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.GetWithdrawalDelay(strategies)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_IncreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || shares == nil {
			return
		}

		_ContractDelegationManager.IncreaseDelegatedShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		var _minWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&_minWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		var _strategies []common.Address
		fill_err = tp.Fill(&_strategies)
		if fill_err != nil {
			return
		}
		var _withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&_withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || initialPausedStatus == nil || _minWithdrawalDelayBlocks == nil {
			return
		}

		_ContractDelegationManager.Initialize(initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_IsDelegated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.IsDelegated(staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_IsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.IsOperator(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_MAXSTAKEROPTOUTWINDOWBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.MAXSTAKEROPTOUTWINDOWBLOCKS()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_MAXWITHDRAWALDELAYBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.MAXWITHDRAWALDELAYBLOCKS()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_MigrateQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&withdrawalsToMigrate)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.MigrateQueuedWithdrawals(withdrawalsToMigrate)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_MinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.MinWithdrawalDelayBlocks()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_ModifyOperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&newOperatorDetails)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ModifyOperatorDetails(newOperatorDetails)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_OperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.OperatorDetails(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_OperatorShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.OperatorShares(arg0, arg1)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Owner()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || newPausedStatus == nil {
			return
		}

		_ContractDelegationManager.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.PauseAll()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Paused(index)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Paused0()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_PendingWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.PendingWithdrawals(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_QueueWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams
		fill_err = tp.Fill(&queuedWithdrawalParams)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.QueueWithdrawals(queuedWithdrawalParams)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_RegisterAsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var registeringOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&registeringOperatorDetails)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.RegisterAsOperator(registeringOperatorDetails, metadataURI)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.RenounceOwnership()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_STAKERDELEGATIONTYPEHASH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.STAKERDELEGATIONTYPEHASH()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_SetMinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newMinWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&newMinWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || newMinWithdrawalDelayBlocks == nil {
			return
		}

		_ContractDelegationManager.SetMinWithdrawalDelayBlocks(newMinWithdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_SetStrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		var withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.SetStrategyWithdrawalDelayBlocks(strategies, withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_Slasher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Slasher()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_StakerNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.StakerNonce(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_StakerOptOutWindowBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.StakerOptOutWindowBlocks(operator)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_StrategyManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.StrategyManager()
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_StrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.StrategyWithdrawalDelayBlocks(arg0)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_Undelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Undelegate(staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || newPausedStatus == nil {
			return
		}

		_ContractDelegationManager.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractDelegationManagerSession_UpdateOperatorMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.UpdateOperatorMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_ContractDelegationManagerStakerDelegatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerDelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStakerDelegatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerDelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStakerDelegatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerDelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStakerForceUndelegatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerForceUndelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStakerForceUndelegatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerForceUndelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStakerForceUndelegatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerForceUndelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStakerUndelegatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerUndelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStakerUndelegatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerUndelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStakerUndelegatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStakerUndelegatedIterator
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

func Fuzz_Nosy_ContractDelegationManagerStrategyWithdrawalDelayBlocksSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStrategyWithdrawalDelayBlocksSetIterator
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

func Fuzz_Nosy_ContractDelegationManagerStrategyWithdrawalDelayBlocksSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStrategyWithdrawalDelayBlocksSetIterator
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

func Fuzz_Nosy_ContractDelegationManagerStrategyWithdrawalDelayBlocksSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerStrategyWithdrawalDelayBlocksSetIterator
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

func Fuzz_Nosy_ContractDelegationManagerTransactor_CompleteQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		var tokens []common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndex *big.Int
		fill_err = tp.Fill(&middlewareTimesIndex)
		if fill_err != nil {
			return
		}
		var receiveAsTokens bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil || middlewareTimesIndex == nil {
			return
		}

		_ContractDelegationManager.CompleteQueuedWithdrawal(opts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_CompleteQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawals []IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}
		var tokens [][]common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndexes []*big.Int
		fill_err = tp.Fill(&middlewareTimesIndexes)
		if fill_err != nil {
			return
		}
		var receiveAsTokens []bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.CompleteQueuedWithdrawals(opts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_DecreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil || shares == nil {
			return
		}

		_ContractDelegationManager.DecreaseDelegatedShares(opts, staker, strategy, shares)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_DelegateTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.DelegateTo(opts, operator, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_DelegateToBySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&stakerSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.DelegateToBySignature(opts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_IncreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil || shares == nil {
			return
		}

		_ContractDelegationManager.IncreaseDelegatedShares(opts, staker, strategy, shares)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		var _minWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&_minWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		var _strategies []common.Address
		fill_err = tp.Fill(&_strategies)
		if fill_err != nil {
			return
		}
		var _withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&_withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil || initialPausedStatus == nil || _minWithdrawalDelayBlocks == nil {
			return
		}

		_ContractDelegationManager.Initialize(opts, initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_MigrateQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&withdrawalsToMigrate)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.MigrateQueuedWithdrawals(opts, withdrawalsToMigrate)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_ModifyOperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&newOperatorDetails)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.ModifyOperatorDetails(opts, newOperatorDetails)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractDelegationManager.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.PauseAll(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_QueueWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams
		fill_err = tp.Fill(&queuedWithdrawalParams)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.QueueWithdrawals(opts, queuedWithdrawalParams)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_RegisterAsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var registeringOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&registeringOperatorDetails)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.RegisterAsOperator(opts, registeringOperatorDetails, metadataURI)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_SetMinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newMinWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&newMinWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil || newMinWithdrawalDelayBlocks == nil {
			return
		}

		_ContractDelegationManager.SetMinWithdrawalDelayBlocks(opts, newMinWithdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_SetStrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		var withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.SetStrategyWithdrawalDelayBlocks(opts, strategies, withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_Undelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.Undelegate(opts, staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractDelegationManager.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactor_UpdateOperatorMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactor
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.UpdateOperatorMetadataURI(opts, metadataURI)
	})
}

// skipping Fuzz_Nosy_ContractDelegationManagerTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractDelegationManagerTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorRaw
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || opts == nil {
			return
		}

		_ContractDelegationManager.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_CompleteQueuedWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawal IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		var tokens []common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndex *big.Int
		fill_err = tp.Fill(&middlewareTimesIndex)
		if fill_err != nil {
			return
		}
		var receiveAsTokens bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || middlewareTimesIndex == nil {
			return
		}

		_ContractDelegationManager.CompleteQueuedWithdrawal(withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_CompleteQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawals []IDelegationManagerWithdrawal
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}
		var tokens [][]common.Address
		fill_err = tp.Fill(&tokens)
		if fill_err != nil {
			return
		}
		var middlewareTimesIndexes []*big.Int
		fill_err = tp.Fill(&middlewareTimesIndexes)
		if fill_err != nil {
			return
		}
		var receiveAsTokens []bool
		fill_err = tp.Fill(&receiveAsTokens)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.CompleteQueuedWithdrawals(withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_DecreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || shares == nil {
			return
		}

		_ContractDelegationManager.DecreaseDelegatedShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_DelegateTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegateTo(operator, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_DelegateToBySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&stakerSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry
		fill_err = tp.Fill(&approverSignatureAndExpiry)
		if fill_err != nil {
			return
		}
		var approverSalt [32]byte
		fill_err = tp.Fill(&approverSalt)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.DelegateToBySignature(staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_IncreaseDelegatedShares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var strategy common.Address
		fill_err = tp.Fill(&strategy)
		if fill_err != nil {
			return
		}
		var shares *big.Int
		fill_err = tp.Fill(&shares)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || shares == nil {
			return
		}

		_ContractDelegationManager.IncreaseDelegatedShares(staker, strategy, shares)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
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
		var _minWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&_minWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		var _strategies []common.Address
		fill_err = tp.Fill(&_strategies)
		if fill_err != nil {
			return
		}
		var _withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&_withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || initialPausedStatus == nil || _minWithdrawalDelayBlocks == nil {
			return
		}

		_ContractDelegationManager.Initialize(initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_MigrateQueuedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var withdrawalsToMigrate []IStrategyManagerDeprecatedStructQueuedWithdrawal
		fill_err = tp.Fill(&withdrawalsToMigrate)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.MigrateQueuedWithdrawals(withdrawalsToMigrate)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_ModifyOperatorDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&newOperatorDetails)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.ModifyOperatorDetails(newOperatorDetails)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || newPausedStatus == nil {
			return
		}

		_ContractDelegationManager.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.PauseAll()
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_QueueWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams
		fill_err = tp.Fill(&queuedWithdrawalParams)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.QueueWithdrawals(queuedWithdrawalParams)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_RegisterAsOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var registeringOperatorDetails IDelegationManagerOperatorDetails
		fill_err = tp.Fill(&registeringOperatorDetails)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.RegisterAsOperator(registeringOperatorDetails, metadataURI)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.RenounceOwnership()
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_SetMinWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newMinWithdrawalDelayBlocks *big.Int
		fill_err = tp.Fill(&newMinWithdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || newMinWithdrawalDelayBlocks == nil {
			return
		}

		_ContractDelegationManager.SetMinWithdrawalDelayBlocks(newMinWithdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_SetStrategyWithdrawalDelayBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var strategies []common.Address
		fill_err = tp.Fill(&strategies)
		if fill_err != nil {
			return
		}
		var withdrawalDelayBlocks []*big.Int
		fill_err = tp.Fill(&withdrawalDelayBlocks)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.SetStrategyWithdrawalDelayBlocks(strategies, withdrawalDelayBlocks)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_Undelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var staker common.Address
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.Undelegate(staker)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil || newPausedStatus == nil {
			return
		}

		_ContractDelegationManager.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractDelegationManagerTransactorSession_UpdateOperatorMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractDelegationManager *ContractDelegationManagerTransactorSession
		fill_err = tp.Fill(&_ContractDelegationManager)
		if fill_err != nil {
			return
		}
		var metadataURI string
		fill_err = tp.Fill(&metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractDelegationManager == nil {
			return
		}

		_ContractDelegationManager.UpdateOperatorMetadataURI(metadataURI)
	})
}

func Fuzz_Nosy_ContractDelegationManagerUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerUnpausedIterator
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

func Fuzz_Nosy_ContractDelegationManagerUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerUnpausedIterator
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

func Fuzz_Nosy_ContractDelegationManagerUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerUnpausedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalCompletedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalCompletedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalCompletedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalCompletedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalCompletedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalCompletedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalMigratedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalMigratedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalMigratedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalMigratedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalMigratedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalMigratedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalQueuedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalQueuedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalQueuedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalQueuedIterator
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

func Fuzz_Nosy_ContractDelegationManagerWithdrawalQueuedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractDelegationManagerWithdrawalQueuedIterator
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

// skipping Fuzz_Nosy_DeployContractDelegationManager__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
