package contracts

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

func Fuzz_Nosy_WethApprovalIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethApprovalIterator
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

func Fuzz_Nosy_WethApprovalIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethApprovalIterator
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

func Fuzz_Nosy_WethApprovalIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethApprovalIterator
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

func Fuzz_Nosy_WethCaller_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCaller
		fill_err = tp.Fill(&_Weth)
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
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Allowance(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_WethCaller_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCaller
		fill_err = tp.Fill(&_Weth)
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
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.BalanceOf(opts, arg0)
	})
}

func Fuzz_Nosy_WethCaller_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCaller
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Decimals(opts)
	})
}

func Fuzz_Nosy_WethCaller_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCaller
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Name(opts)
	})
}

func Fuzz_Nosy_WethCaller_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCaller
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Symbol(opts)
	})
}

func Fuzz_Nosy_WethCaller_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCaller
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.TotalSupply(opts)
	})
}

// skipping Fuzz_Nosy_WethCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_WethCallerSession_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCallerSession
		fill_err = tp.Fill(&_Weth)
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
		if _Weth == nil {
			return
		}

		_Weth.Allowance(arg0, arg1)
	})
}

func Fuzz_Nosy_WethCallerSession_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCallerSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.BalanceOf(arg0)
	})
}

func Fuzz_Nosy_WethCallerSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCallerSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Decimals()
	})
}

func Fuzz_Nosy_WethCallerSession_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCallerSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Name()
	})
}

func Fuzz_Nosy_WethCallerSession_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCallerSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Symbol()
	})
}

func Fuzz_Nosy_WethCallerSession_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethCallerSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.TotalSupply()
	})
}

func Fuzz_Nosy_WethDepositIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethDepositIterator
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

func Fuzz_Nosy_WethDepositIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethDepositIterator
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

func Fuzz_Nosy_WethDepositIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethDepositIterator
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

func Fuzz_Nosy_WethFilterer_FilterApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethFilterer
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var src []common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var guy []common.Address
		fill_err = tp.Fill(&guy)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.FilterApproval(opts, src, guy)
	})
}

func Fuzz_Nosy_WethFilterer_FilterDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethFilterer
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var dst []common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.FilterDeposit(opts, dst)
	})
}

func Fuzz_Nosy_WethFilterer_FilterTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethFilterer
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var src []common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst []common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.FilterTransfer(opts, src, dst)
	})
}

func Fuzz_Nosy_WethFilterer_FilterWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethFilterer
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var src []common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.FilterWithdrawal(opts, src)
	})
}

func Fuzz_Nosy_WethFilterer_ParseApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethFilterer
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.ParseApproval(log)
	})
}

func Fuzz_Nosy_WethFilterer_ParseDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethFilterer
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.ParseDeposit(log)
	})
}

func Fuzz_Nosy_WethFilterer_ParseTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethFilterer
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.ParseTransfer(log)
	})
}

func Fuzz_Nosy_WethFilterer_ParseWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethFilterer
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.ParseWithdrawal(log)
	})
}

// skipping Fuzz_Nosy_WethFilterer_WatchApproval__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/indexer/test/contracts.WethApproval

// skipping Fuzz_Nosy_WethFilterer_WatchDeposit__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/indexer/test/contracts.WethDeposit

// skipping Fuzz_Nosy_WethFilterer_WatchTransfer__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/indexer/test/contracts.WethTransfer

// skipping Fuzz_Nosy_WethFilterer_WatchWithdrawal__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/indexer/test/contracts.WethWithdrawal

// skipping Fuzz_Nosy_WethRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_WethRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_WethRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethRaw
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Transfer(opts)
	})
}

func Fuzz_Nosy_WethSession_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
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
		if _Weth == nil {
			return
		}

		_Weth.Allowance(arg0, arg1)
	})
}

func Fuzz_Nosy_WethSession_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var guy common.Address
		fill_err = tp.Fill(&guy)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || wad == nil {
			return
		}

		_Weth.Approve(guy, wad)
	})
}

func Fuzz_Nosy_WethSession_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.BalanceOf(arg0)
	})
}

func Fuzz_Nosy_WethSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Decimals()
	})
}

func Fuzz_Nosy_WethSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Deposit()
	})
}

func Fuzz_Nosy_WethSession_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Fallback(calldata)
	})
}

func Fuzz_Nosy_WethSession_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Name()
	})
}

func Fuzz_Nosy_WethSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Receive()
	})
}

func Fuzz_Nosy_WethSession_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Symbol()
	})
}

func Fuzz_Nosy_WethSession_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.TotalSupply()
	})
}

func Fuzz_Nosy_WethSession_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || wad == nil {
			return
		}

		_Weth.Transfer(dst, wad)
	})
}

func Fuzz_Nosy_WethSession_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var src common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || wad == nil {
			return
		}

		_Weth.TransferFrom(src, dst, wad)
	})
}

func Fuzz_Nosy_WethSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || wad == nil {
			return
		}

		_Weth.Withdraw(wad)
	})
}

func Fuzz_Nosy_WethTransactor_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactor
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var guy common.Address
		fill_err = tp.Fill(&guy)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil || wad == nil {
			return
		}

		_Weth.Approve(opts, guy, wad)
	})
}

func Fuzz_Nosy_WethTransactor_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactor
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Deposit(opts)
	})
}

func Fuzz_Nosy_WethTransactor_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactor
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Fallback(opts, calldata)
	})
}

func Fuzz_Nosy_WethTransactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactor
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Receive(opts)
	})
}

func Fuzz_Nosy_WethTransactor_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactor
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil || wad == nil {
			return
		}

		_Weth.Transfer(opts, dst, wad)
	})
}

func Fuzz_Nosy_WethTransactor_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactor
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var src common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil || wad == nil {
			return
		}

		_Weth.TransferFrom(opts, src, dst, wad)
	})
}

func Fuzz_Nosy_WethTransactor_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactor
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil || wad == nil {
			return
		}

		_Weth.Withdraw(opts, wad)
	})
}

// skipping Fuzz_Nosy_WethTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_WethTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactorRaw
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Weth == nil || opts == nil {
			return
		}

		_Weth.Transfer(opts)
	})
}

func Fuzz_Nosy_WethTransactorSession_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactorSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var guy common.Address
		fill_err = tp.Fill(&guy)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || wad == nil {
			return
		}

		_Weth.Approve(guy, wad)
	})
}

func Fuzz_Nosy_WethTransactorSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactorSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Deposit()
	})
}

func Fuzz_Nosy_WethTransactorSession_Fallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactorSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Fallback(calldata)
	})
}

func Fuzz_Nosy_WethTransactorSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactorSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		if _Weth == nil {
			return
		}

		_Weth.Receive()
	})
}

func Fuzz_Nosy_WethTransactorSession_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactorSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || wad == nil {
			return
		}

		_Weth.Transfer(dst, wad)
	})
}

func Fuzz_Nosy_WethTransactorSession_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactorSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var src common.Address
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || wad == nil {
			return
		}

		_Weth.TransferFrom(src, dst, wad)
	})
}

func Fuzz_Nosy_WethTransactorSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Weth *WethTransactorSession
		fill_err = tp.Fill(&_Weth)
		if fill_err != nil {
			return
		}
		var wad *big.Int
		fill_err = tp.Fill(&wad)
		if fill_err != nil {
			return
		}
		if _Weth == nil || wad == nil {
			return
		}

		_Weth.Withdraw(wad)
	})
}

func Fuzz_Nosy_WethTransferIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethTransferIterator
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

func Fuzz_Nosy_WethTransferIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethTransferIterator
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

func Fuzz_Nosy_WethTransferIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethTransferIterator
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

func Fuzz_Nosy_WethWithdrawalIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethWithdrawalIterator
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

func Fuzz_Nosy_WethWithdrawalIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethWithdrawalIterator
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

func Fuzz_Nosy_WethWithdrawalIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *WethWithdrawalIterator
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

// skipping Fuzz_Nosy_DeployWeth__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
