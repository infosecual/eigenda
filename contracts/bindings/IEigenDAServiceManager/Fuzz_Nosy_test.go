package contractIEigenDAServiceManager

import (
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

func Fuzz_Nosy_ContractIEigenDAServiceManagerBatchConfirmedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIEigenDAServiceManagerBatchConfirmedIterator
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

func Fuzz_Nosy_ContractIEigenDAServiceManagerBatchConfirmedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIEigenDAServiceManagerBatchConfirmedIterator
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

func Fuzz_Nosy_ContractIEigenDAServiceManagerBatchConfirmedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIEigenDAServiceManagerBatchConfirmedIterator
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

func Fuzz_Nosy_ContractIEigenDAServiceManagerBatchConfirmerChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIEigenDAServiceManagerBatchConfirmerChangedIterator
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

func Fuzz_Nosy_ContractIEigenDAServiceManagerBatchConfirmerChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIEigenDAServiceManagerBatchConfirmerChangedIterator
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

func Fuzz_Nosy_ContractIEigenDAServiceManagerBatchConfirmerChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractIEigenDAServiceManagerBatchConfirmerChangedIterator
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

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.AvsDirectory(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_BLOCKSTALEMEASURE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.BLOCKSTALEMEASURE(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_BatchIdToBatchMetadataHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var batchId uint32
		fill_err = tp.Fill(&batchId)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.BatchIdToBatchMetadataHash(opts, batchId)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
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
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.GetOperatorRestakedStrategies(opts, operator)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.GetRestakeableStrategies(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_LatestServeUntilBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.LatestServeUntilBlock(opts, referenceBlockNumber)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_QuorumAdversaryThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumAdversaryThresholdPercentages(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_QuorumConfirmationThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumConfirmationThresholdPercentages(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_QuorumNumbersRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumNumbersRequired(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCaller_TaskNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.TaskNumber(opts)
	})
}

// skipping Fuzz_Nosy_ContractIEigenDAServiceManagerCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.AvsDirectory()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_BLOCKSTALEMEASURE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.BLOCKSTALEMEASURE()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_BatchIdToBatchMetadataHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var batchId uint32
		fill_err = tp.Fill(&batchId)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.BatchIdToBatchMetadataHash(batchId)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.GetOperatorRestakedStrategies(operator)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.GetRestakeableStrategies()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_LatestServeUntilBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.LatestServeUntilBlock(referenceBlockNumber)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_QuorumAdversaryThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumAdversaryThresholdPercentages()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_QuorumConfirmationThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumConfirmationThresholdPercentages()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_QuorumNumbersRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumNumbersRequired()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerCallerSession_TaskNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.TaskNumber()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerFilterer_FilterBatchConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var batchHeaderHash [][32]byte
		fill_err = tp.Fill(&batchHeaderHash)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.FilterBatchConfirmed(opts, batchHeaderHash)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerFilterer_FilterBatchConfirmerChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.FilterBatchConfirmerChanged(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerFilterer_ParseBatchConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.ParseBatchConfirmed(log)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerFilterer_ParseBatchConfirmerChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.ParseBatchConfirmerChanged(log)
	})
}

// skipping Fuzz_Nosy_ContractIEigenDAServiceManagerFilterer_WatchBatchConfirmed__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/IEigenDAServiceManager.ContractIEigenDAServiceManagerBatchConfirmed

// skipping Fuzz_Nosy_ContractIEigenDAServiceManagerFilterer_WatchBatchConfirmerChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/IEigenDAServiceManager.ContractIEigenDAServiceManagerBatchConfirmerChanged

// skipping Fuzz_Nosy_ContractIEigenDAServiceManagerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractIEigenDAServiceManagerRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractIEigenDAServiceManagerRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerRaw
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.AvsDirectory()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_BLOCKSTALEMEASURE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.BLOCKSTALEMEASURE()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_BatchIdToBatchMetadataHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var batchId uint32
		fill_err = tp.Fill(&batchId)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.BatchIdToBatchMetadataHash(batchId)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_ConfirmBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var batchHeader IEigenDAServiceManagerBatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}
		var nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature
		fill_err = tp.Fill(&nonSignerStakesAndSignature)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.ConfirmBatch(batchHeader, nonSignerStakesAndSignature)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.DeregisterOperatorFromAVS(operator)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.GetOperatorRestakedStrategies(operator)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.GetRestakeableStrategies()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_LatestServeUntilBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.LatestServeUntilBlock(referenceBlockNumber)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_QuorumAdversaryThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumAdversaryThresholdPercentages()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_QuorumConfirmationThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumConfirmationThresholdPercentages()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_QuorumNumbersRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.QuorumNumbersRequired()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.RegisterOperatorToAVS(operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_SetBatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var _batchConfirmer common.Address
		fill_err = tp.Fill(&_batchConfirmer)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.SetBatchConfirmer(_batchConfirmer)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_TaskNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.TaskNumber()
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerSession_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var _metadataURI string
		fill_err = tp.Fill(&_metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.UpdateAVSMetadataURI(_metadataURI)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactor_ConfirmBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var batchHeader IEigenDAServiceManagerBatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}
		var nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature
		fill_err = tp.Fill(&nonSignerStakesAndSignature)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.ConfirmBatch(opts, batchHeader, nonSignerStakesAndSignature)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactor_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
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
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.DeregisterOperatorFromAVS(opts, operator)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactor_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
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
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.RegisterOperatorToAVS(opts, operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactor_SetBatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _batchConfirmer common.Address
		fill_err = tp.Fill(&_batchConfirmer)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.SetBatchConfirmer(opts, _batchConfirmer)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactor_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _metadataURI string
		fill_err = tp.Fill(&_metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.UpdateAVSMetadataURI(opts, _metadataURI)
	})
}

// skipping Fuzz_Nosy_ContractIEigenDAServiceManagerTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactorRaw
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractIEigenDAServiceManager.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactorSession_ConfirmBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var batchHeader IEigenDAServiceManagerBatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}
		var nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature
		fill_err = tp.Fill(&nonSignerStakesAndSignature)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.ConfirmBatch(batchHeader, nonSignerStakesAndSignature)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactorSession_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.DeregisterOperatorFromAVS(operator)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactorSession_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
		fill_err = tp.Fill(&operatorSignature)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.RegisterOperatorToAVS(operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactorSession_SetBatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var _batchConfirmer common.Address
		fill_err = tp.Fill(&_batchConfirmer)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.SetBatchConfirmer(_batchConfirmer)
	})
}

func Fuzz_Nosy_ContractIEigenDAServiceManagerTransactorSession_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractIEigenDAServiceManager *ContractIEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractIEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var _metadataURI string
		fill_err = tp.Fill(&_metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractIEigenDAServiceManager == nil {
			return
		}

		_ContractIEigenDAServiceManager.UpdateAVSMetadataURI(_metadataURI)
	})
}
