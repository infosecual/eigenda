package contractBN254

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

// skipping Fuzz_Nosy_ContractBN254CallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractBN254Raw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractBN254Raw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBN254Raw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBN254 *ContractBN254Raw
		fill_err = tp.Fill(&_ContractBN254)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBN254 == nil || opts == nil {
			return
		}

		_ContractBN254.Transfer(opts)
	})
}

// skipping Fuzz_Nosy_ContractBN254TransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBN254TransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBN254 *ContractBN254TransactorRaw
		fill_err = tp.Fill(&_ContractBN254)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBN254 == nil || opts == nil {
			return
		}

		_ContractBN254.Transfer(opts)
	})
}

// skipping Fuzz_Nosy_DeployContractBN254__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
