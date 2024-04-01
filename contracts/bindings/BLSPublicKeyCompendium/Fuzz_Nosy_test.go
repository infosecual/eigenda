package contractBLSPublicKeyCompendium

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

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumCaller_GetMessageHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCaller
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
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
		if _ContractBLSPublicKeyCompendium == nil || opts == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.GetMessageHash(opts, operator)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumCaller_OperatorToPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCaller
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
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
		if _ContractBLSPublicKeyCompendium == nil || opts == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.OperatorToPubkeyHash(opts, arg0)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumCaller_PubkeyHashToOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCaller
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
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
		if _ContractBLSPublicKeyCompendium == nil || opts == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.PubkeyHashToOperator(opts, arg0)
	})
}

// skipping Fuzz_Nosy_ContractBLSPublicKeyCompendiumCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumCallerSession_GetMessageHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCallerSession
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.GetMessageHash(operator)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumCallerSession_OperatorToPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCallerSession
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.OperatorToPubkeyHash(arg0)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumCallerSession_PubkeyHashToOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCallerSession
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.PubkeyHashToOperator(arg0)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumFilterer_FilterNewPubkeyRegistration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumFilterer
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
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
		if _ContractBLSPublicKeyCompendium == nil || opts == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.FilterNewPubkeyRegistration(opts, operator)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumFilterer_ParseNewPubkeyRegistration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumFilterer
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.ParseNewPubkeyRegistration(log)
	})
}

// skipping Fuzz_Nosy_ContractBLSPublicKeyCompendiumFilterer_WatchNewPubkeyRegistration__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/contracts/bindings/BLSPublicKeyCompendium.ContractBLSPublicKeyCompendiumNewPubkeyRegistration

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator
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

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator
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

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator
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

// skipping Fuzz_Nosy_ContractBLSPublicKeyCompendiumRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ContractBLSPublicKeyCompendiumRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumRaw
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil || opts == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumSession_GetMessageHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumSession
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.GetMessageHash(operator)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumSession_OperatorToPubkeyHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumSession
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.OperatorToPubkeyHash(arg0)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumSession_PubkeyHashToOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumSession
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.PubkeyHashToOperator(arg0)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumSession_RegisterBLSPublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumSession
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var signedMessageHash BN254G1Point
		fill_err = tp.Fill(&signedMessageHash)
		if fill_err != nil {
			return
		}
		var pubkeyG1 BN254G1Point
		fill_err = tp.Fill(&pubkeyG1)
		if fill_err != nil {
			return
		}
		var pubkeyG2 BN254G2Point
		fill_err = tp.Fill(&pubkeyG2)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.RegisterBLSPublicKey(signedMessageHash, pubkeyG1, pubkeyG2)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumTransactor_RegisterBLSPublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumTransactor
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var signedMessageHash BN254G1Point
		fill_err = tp.Fill(&signedMessageHash)
		if fill_err != nil {
			return
		}
		var pubkeyG1 BN254G1Point
		fill_err = tp.Fill(&pubkeyG1)
		if fill_err != nil {
			return
		}
		var pubkeyG2 BN254G2Point
		fill_err = tp.Fill(&pubkeyG2)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil || opts == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.RegisterBLSPublicKey(opts, signedMessageHash, pubkeyG1, pubkeyG2)
	})
}

// skipping Fuzz_Nosy_ContractBLSPublicKeyCompendiumTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumTransactorRaw
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil || opts == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.Transfer(opts)
	})
}

func Fuzz_Nosy_ContractBLSPublicKeyCompendiumTransactorSession_RegisterBLSPublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumTransactorSession
		fill_err = tp.Fill(&_ContractBLSPublicKeyCompendium)
		if fill_err != nil {
			return
		}
		var signedMessageHash BN254G1Point
		fill_err = tp.Fill(&signedMessageHash)
		if fill_err != nil {
			return
		}
		var pubkeyG1 BN254G1Point
		fill_err = tp.Fill(&pubkeyG1)
		if fill_err != nil {
			return
		}
		var pubkeyG2 BN254G2Point
		fill_err = tp.Fill(&pubkeyG2)
		if fill_err != nil {
			return
		}
		if _ContractBLSPublicKeyCompendium == nil {
			return
		}

		_ContractBLSPublicKeyCompendium.RegisterBLSPublicKey(signedMessageHash, pubkeyG1, pubkeyG2)
	})
}

// skipping Fuzz_Nosy_DeployContractBLSPublicKeyCompendium__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
