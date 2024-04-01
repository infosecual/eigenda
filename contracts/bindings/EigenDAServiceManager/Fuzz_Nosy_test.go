package contractEigenDAServiceManager

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

func Fuzz_Nosy_ContractEigenDAServiceManagerBatchConfirmedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerBatchConfirmedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerBatchConfirmedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerBatchConfirmedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerBatchConfirmedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerBatchConfirmedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerBatchConfirmerChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerBatchConfirmerChangedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerBatchConfirmerChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerBatchConfirmerChangedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerBatchConfirmerChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerBatchConfirmerChangedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.AvsDirectory(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_BLOCKSTALEMEASURE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.BLOCKSTALEMEASURE(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_BatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchConfirmer(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_BatchId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchId(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_BatchIdToBatchMetadataHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchIdToBatchMetadataHash(opts, arg0)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_BlsApkRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.BlsApkRegistry(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_CheckSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var msgHash [32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		var params IBLSSignatureCheckerNonSignerStakesAndSignature
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.CheckSignatures(opts, msgHash, quorumNumbers, referenceBlockNumber, params)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.Delegation(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.GetOperatorRestakedStrategies(opts, operator)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.GetRestakeableStrategies(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_LatestServeUntilBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.LatestServeUntilBlock(opts, referenceBlockNumber)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.Owner(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.Paused(opts, index)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.Paused0(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.PauserRegistry(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_QuorumAdversaryThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumAdversaryThresholdPercentages(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_QuorumConfirmationThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumConfirmationThresholdPercentages(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_QuorumNumbersRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumNumbersRequired(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.RegistryCoordinator(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_STOREDURATIONBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.STOREDURATIONBLOCKS(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.StakeRegistry(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_StaleStakesForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.StaleStakesForbidden(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_THRESHOLDDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.THRESHOLDDENOMINATOR(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_TaskNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.TaskNumber(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCaller_TrySignatureAndApkVerification__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCaller
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var msgHash [32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		var apk BN254G1Point
		fill_err = tp.Fill(&apk)
		if fill_err != nil {
			return
		}
		var apkG2 BN254G2Point
		fill_err = tp.Fill(&apkG2)
		if fill_err != nil {
			return
		}
		var sigma BN254G1Point
		fill_err = tp.Fill(&sigma)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.TrySignatureAndApkVerification(opts, msgHash, apk, apkG2, sigma)
	})
}

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.AvsDirectory()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_BLOCKSTALEMEASURE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BLOCKSTALEMEASURE()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_BatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchConfirmer()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_BatchId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchId()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_BatchIdToBatchMetadataHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchIdToBatchMetadataHash(arg0)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_BlsApkRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BlsApkRegistry()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_CheckSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var msgHash [32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		var params IBLSSignatureCheckerNonSignerStakesAndSignature
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.CheckSignatures(msgHash, quorumNumbers, referenceBlockNumber, params)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.Delegation()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.GetOperatorRestakedStrategies(operator)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.GetRestakeableStrategies()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_LatestServeUntilBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.LatestServeUntilBlock(referenceBlockNumber)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.Owner()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.Paused(index)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.Paused0()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_QuorumAdversaryThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumAdversaryThresholdPercentages()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_QuorumConfirmationThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumConfirmationThresholdPercentages()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_QuorumNumbersRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumNumbersRequired()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_STOREDURATIONBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.STOREDURATIONBLOCKS()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.StakeRegistry()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_StaleStakesForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.StaleStakesForbidden()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_THRESHOLDDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.THRESHOLDDENOMINATOR()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_TaskNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.TaskNumber()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerCallerSession_TrySignatureAndApkVerification__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerCallerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var msgHash [32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		var apk BN254G1Point
		fill_err = tp.Fill(&apk)
		if fill_err != nil {
			return
		}
		var apkG2 BN254G2Point
		fill_err = tp.Fill(&apkG2)
		if fill_err != nil {
			return
		}
		var sigma BN254G1Point
		fill_err = tp.Fill(&sigma)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.TrySignatureAndApkVerification(msgHash, apk, apkG2, sigma)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_FilterBatchConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.FilterBatchConfirmed(opts, batchHeaderHash)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_FilterBatchConfirmerChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.FilterBatchConfirmerChanged(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_FilterPaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.FilterPaused(opts, account)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_FilterPauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.FilterPauserRegistrySet(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_FilterStaleStakesForbiddenUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.FilterStaleStakesForbiddenUpdate(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_FilterUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.FilterUnpaused(opts, account)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_ParseBatchConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ParseBatchConfirmed(log)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_ParseBatchConfirmerChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ParseBatchConfirmerChanged(log)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ParseInitialized(log)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_ParsePaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ParsePaused(log)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_ParsePauserRegistrySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ParsePauserRegistrySet(log)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_ParseStaleStakesForbiddenUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ParseStaleStakesForbiddenUpdate(log)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_ParseUnpaused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerFilterer
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ParseUnpaused(log)
	})
}

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_WatchBatchConfirmed__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager.ContractEigenDAServiceManagerBatchConfirmed

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_WatchBatchConfirmerChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager.ContractEigenDAServiceManagerBatchConfirmerChanged

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager.ContractEigenDAServiceManagerInitialized

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager.ContractEigenDAServiceManagerOwnershipTransferred

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_WatchPaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager.ContractEigenDAServiceManagerPaused

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_WatchPauserRegistrySet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager.ContractEigenDAServiceManagerPauserRegistrySet

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_WatchStaleStakesForbiddenUpdate__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager.ContractEigenDAServiceManagerStaleStakesForbiddenUpdate

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerFilterer_WatchUnpaused__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/EigenDAServiceManager.ContractEigenDAServiceManagerUnpaused

func Fuzz_Nosy_ContractEigenDAServiceManagerInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerInitializedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerInitializedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerInitializedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerOwnershipTransferredIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerPausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerPausedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerPausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerPausedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerPausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerPausedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerPauserRegistrySetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerPauserRegistrySetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerPauserRegistrySetIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerPauserRegistrySetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerPauserRegistrySetIterator
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

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractEigenDAServiceManagerRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerRaw
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_AvsDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.AvsDirectory()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_BLOCKSTALEMEASURE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BLOCKSTALEMEASURE()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_BatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchConfirmer()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_BatchId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchId()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_BatchIdToBatchMetadataHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BatchIdToBatchMetadataHash(arg0)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_BlsApkRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.BlsApkRegistry()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_CheckSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var msgHash [32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		var quorumNumbers []byte
		fill_err = tp.Fill(&quorumNumbers)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		var params IBLSSignatureCheckerNonSignerStakesAndSignature
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.CheckSignatures(msgHash, quorumNumbers, referenceBlockNumber, params)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_ConfirmBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ConfirmBatch(batchHeader, nonSignerStakesAndSignature)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_Delegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.Delegation()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.DeregisterOperatorFromAVS(operator)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_GetOperatorRestakedStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.GetOperatorRestakedStrategies(operator)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_GetRestakeableStrategies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.GetRestakeableStrategies()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		var _initialOwner common.Address
		fill_err = tp.Fill(&_initialOwner)
		if fill_err != nil {
			return
		}
		var _batchConfirmer common.Address
		fill_err = tp.Fill(&_batchConfirmer)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || _initialPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Initialize(_pauserRegistry, _initialPausedStatus, _initialOwner, _batchConfirmer)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_LatestServeUntilBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var referenceBlockNumber uint32
		fill_err = tp.Fill(&referenceBlockNumber)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.LatestServeUntilBlock(referenceBlockNumber)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.Owner()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || newPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.PauseAll()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var index uint8
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.Paused(index)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_Paused0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.Paused0()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_PauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.PauserRegistry()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_QuorumAdversaryThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumAdversaryThresholdPercentages()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_QuorumConfirmationThresholdPercentages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumConfirmationThresholdPercentages()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_QuorumNumbersRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.QuorumNumbersRequired()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.RegisterOperatorToAVS(operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_RegistryCoordinator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.RegistryCoordinator()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.RenounceOwnership()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_STOREDURATIONBLOCKS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.STOREDURATIONBLOCKS()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_SetBatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var _batchConfirmer common.Address
		fill_err = tp.Fill(&_batchConfirmer)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.SetBatchConfirmer(_batchConfirmer)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_SetStaleStakesForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var value bool
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.SetStaleStakesForbidden(value)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.StakeRegistry()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_StaleStakesForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.StaleStakesForbidden()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_THRESHOLDDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.THRESHOLDDENOMINATOR()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_TaskNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.TaskNumber()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_TrySignatureAndApkVerification__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var msgHash [32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		var apk BN254G1Point
		fill_err = tp.Fill(&apk)
		if fill_err != nil {
			return
		}
		var apkG2 BN254G2Point
		fill_err = tp.Fill(&apkG2)
		if fill_err != nil {
			return
		}
		var sigma BN254G1Point
		fill_err = tp.Fill(&sigma)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.TrySignatureAndApkVerification(msgHash, apk, apkG2, sigma)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || newPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerSession_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var _metadataURI string
		fill_err = tp.Fill(&_metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.UpdateAVSMetadataURI(_metadataURI)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerStaleStakesForbiddenUpdateIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerStaleStakesForbiddenUpdateIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerStaleStakesForbiddenUpdateIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerStaleStakesForbiddenUpdateIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerStaleStakesForbiddenUpdateIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerStaleStakesForbiddenUpdateIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_ConfirmBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.ConfirmBatch(opts, batchHeader, nonSignerStakesAndSignature)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.DeregisterOperatorFromAVS(opts, operator)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
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
		var _initialOwner common.Address
		fill_err = tp.Fill(&_initialOwner)
		if fill_err != nil {
			return
		}
		var _batchConfirmer common.Address
		fill_err = tp.Fill(&_batchConfirmer)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil || _initialPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Initialize(opts, _pauserRegistry, _initialPausedStatus, _initialOwner, _batchConfirmer)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Pause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.PauseAll(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.RegisterOperatorToAVS(opts, operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_SetBatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.SetBatchConfirmer(opts, _batchConfirmer)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.SetPauserRegistry(opts, newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_SetStaleStakesForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var value bool
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.SetStaleStakesForbidden(opts, value)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil || newPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Unpause(opts, newPausedStatus)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactor_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactor
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.UpdateAVSMetadataURI(opts, _metadataURI)
	})
}

// skipping Fuzz_Nosy_ContractEigenDAServiceManagerTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorRaw
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || opts == nil {
			return
		}

		_ContractEigenDAServiceManager.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_ConfirmBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.ConfirmBatch(batchHeader, nonSignerStakesAndSignature)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_DeregisterOperatorFromAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.DeregisterOperatorFromAVS(operator)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		var _initialOwner common.Address
		fill_err = tp.Fill(&_initialOwner)
		if fill_err != nil {
			return
		}
		var _batchConfirmer common.Address
		fill_err = tp.Fill(&_batchConfirmer)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || _initialPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Initialize(_pauserRegistry, _initialPausedStatus, _initialOwner, _batchConfirmer)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || newPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Pause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_PauseAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.PauseAll()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_RegisterOperatorToAVS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
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
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.RegisterOperatorToAVS(operator, operatorSignature)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.RenounceOwnership()
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_SetBatchConfirmer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var _batchConfirmer common.Address
		fill_err = tp.Fill(&_batchConfirmer)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.SetBatchConfirmer(_batchConfirmer)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_SetPauserRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var newPauserRegistry common.Address
		fill_err = tp.Fill(&newPauserRegistry)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.SetPauserRegistry(newPauserRegistry)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_SetStaleStakesForbidden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var value bool
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.SetStaleStakesForbidden(value)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_Unpause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var newPausedStatus *big.Int
		fill_err = tp.Fill(&newPausedStatus)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil || newPausedStatus == nil {
			return
		}

		_ContractEigenDAServiceManager.Unpause(newPausedStatus)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerTransactorSession_UpdateAVSMetadataURI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractEigenDAServiceManager *ContractEigenDAServiceManagerTransactorSession
		fill_err = tp.Fill(&_ContractEigenDAServiceManager)
		if fill_err != nil {
			return
		}
		var _metadataURI string
		fill_err = tp.Fill(&_metadataURI)
		if fill_err != nil {
			return
		}
		if _ContractEigenDAServiceManager == nil {
			return
		}

		_ContractEigenDAServiceManager.UpdateAVSMetadataURI(_metadataURI)
	})
}

func Fuzz_Nosy_ContractEigenDAServiceManagerUnpausedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerUnpausedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerUnpausedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerUnpausedIterator
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

func Fuzz_Nosy_ContractEigenDAServiceManagerUnpausedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractEigenDAServiceManagerUnpausedIterator
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

// skipping Fuzz_Nosy_DeployContractEigenDAServiceManager__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
