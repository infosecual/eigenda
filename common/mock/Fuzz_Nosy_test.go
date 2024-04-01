package mock

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigenda/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
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

func Fuzz_Nosy_MockEthClient_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m1 == nil || blockNumber == nil {
			return
		}

		m1.BalanceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_MockEthClient_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockEthClient_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if m1 == nil || number == nil {
			return
		}

		m1.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockEthClient_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.BlockNumber(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m1 == nil || blockNumber == nil {
			return
		}

		m1.CallContract(ctx, msg, blockNumber)
	})
}

func Fuzz_Nosy_MockEthClient_CallContractAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.CallContractAtHash(ctx, msg, blockHash)
	})
}

func Fuzz_Nosy_MockEthClient_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.ChainID(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m1 == nil || blockNumber == nil {
			return
		}

		m1.CodeAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_MockEthClient_EnsureAnyTransactionEvaled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var tag string
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.EnsureAnyTransactionEvaled(ctx, txs, tag)
	})
}

func Fuzz_Nosy_MockEthClient_EnsureTransactionEvaled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var tag string
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		if m1 == nil || tx == nil {
			return
		}

		m1.EnsureTransactionEvaled(ctx, tx, tag)
	})
}

func Fuzz_Nosy_MockEthClient_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.EstimateGas(ctx, msg)
	})
}

func Fuzz_Nosy_MockEthClient_EstimateGasPriceAndLimitAndSendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var tag string
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if m1 == nil || tx == nil || value == nil {
			return
		}

		m1.EstimateGasPriceAndLimitAndSendTx(ctx, tx, tag, value)
	})
}

func Fuzz_Nosy_MockEthClient_FeeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockCount uint64
		fill_err = tp.Fill(&blockCount)
		if fill_err != nil {
			return
		}
		var lastBlock *big.Int
		fill_err = tp.Fill(&lastBlock)
		if fill_err != nil {
			return
		}
		var rewardPercentiles []float64
		fill_err = tp.Fill(&rewardPercentiles)
		if fill_err != nil {
			return
		}
		if m1 == nil || lastBlock == nil {
			return
		}

		m1.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	})
}

func Fuzz_Nosy_MockEthClient_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var q ethereum.FilterQuery
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.FilterLogs(ctx, q)
	})
}

func Fuzz_Nosy_MockEthClient_GetAccountAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.GetAccountAddress()
	})
}

func Fuzz_Nosy_MockEthClient_GetLatestGasCaps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.GetLatestGasCaps(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_GetNoSendTransactOpts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.GetNoSendTransactOpts()
	})
}

func Fuzz_Nosy_MockEthClient_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockEthClient_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if m1 == nil || number == nil {
			return
		}

		m1.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockEthClient_NetworkID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.NetworkID(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m1 == nil || blockNumber == nil {
			return
		}

		m1.NonceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_MockEthClient_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PeerCount(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_PendingBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PendingBalanceAt(ctx, account)
	})
}

func Fuzz_Nosy_MockEthClient_PendingCallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PendingCallContract(ctx, msg)
	})
}

func Fuzz_Nosy_MockEthClient_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PendingCodeAt(ctx, account)
	})
}

func Fuzz_Nosy_MockEthClient_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PendingNonceAt(ctx, account)
	})
}

func Fuzz_Nosy_MockEthClient_PendingStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PendingStorageAt(ctx, account, key)
	})
}

func Fuzz_Nosy_MockEthClient_PendingTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PendingTransactionCount(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if m1 == nil || tx == nil {
			return
		}

		m1.SendTransaction(ctx, tx)
	})
}

func Fuzz_Nosy_MockEthClient_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m1 == nil || blockNumber == nil {
			return
		}

		m1.StorageAt(ctx, account, key, blockNumber)
	})
}

// skipping Fuzz_Nosy_MockEthClient_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_MockEthClient_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_MockEthClient_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.SuggestGasPrice(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.SyncProgress(ctx)
	})
}

func Fuzz_Nosy_MockEthClient_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.TransactionByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockEthClient_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.TransactionCount(ctx, blockHash)
	})
}

func Fuzz_Nosy_MockEthClient_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.TransactionInBlock(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_MockEthClient_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.TransactionReceipt(ctx, txHash)
	})
}

func Fuzz_Nosy_MockEthClient_TransactionSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var block common.Hash
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if m1 == nil || tx == nil {
			return
		}

		m1.TransactionSender(ctx, tx, block, index)
	})
}

func Fuzz_Nosy_MockEthClient_UpdateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasTipCap *big.Int
		fill_err = tp.Fill(&gasTipCap)
		if fill_err != nil {
			return
		}
		var gasFeeCap *big.Int
		fill_err = tp.Fill(&gasFeeCap)
		if fill_err != nil {
			return
		}
		if m1 == nil || tx == nil || value == nil || gasTipCap == nil || gasFeeCap == nil {
			return
		}

		m1.UpdateGas(ctx, tx, value, gasTipCap, gasFeeCap)
	})
}

func Fuzz_Nosy_MockRPCEthClient_BatchCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockRPCEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var b []rpc.BatchElem
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.BatchCall(b)
	})
}

func Fuzz_Nosy_MockRPCEthClient_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockRPCEthClient
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var b []rpc.BatchElem
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_MockRPCEthClient_Call__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockRPCEthClient_CallContext__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockWorkerpool_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockWorkerpool
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Pause(ctx)
	})
}

func Fuzz_Nosy_MockWorkerpool_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockWorkerpool
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Size()
	})
}

func Fuzz_Nosy_MockWorkerpool_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockWorkerpool
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Stop()
	})
}

func Fuzz_Nosy_MockWorkerpool_StopWait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockWorkerpool
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.StopWait()
	})
}

func Fuzz_Nosy_MockWorkerpool_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockWorkerpool
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Stopped()
	})
}

// skipping Fuzz_Nosy_MockWorkerpool_Submit__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_MockWorkerpool_SubmitWait__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_MockWorkerpool_WaitingQueueSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *MockWorkerpool
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.WaitingQueueSize()
	})
}

func Fuzz_Nosy_NoopRatelimiter_AllowRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *NoopRatelimiter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params []common.RequestParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.AllowRequest(ctx, params)
	})
}

func Fuzz_Nosy_S3Client_DeleteObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bucket string
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		s := NewS3Client()
		s.DeleteObject(ctx, bucket, key)
	})
}

func Fuzz_Nosy_S3Client_DownloadObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bucket string
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		s := NewS3Client()
		s.DownloadObject(ctx, bucket, key)
	})
}

func Fuzz_Nosy_S3Client_ListObjects__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bucket string
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		var prefix string
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}

		s := NewS3Client()
		s.ListObjects(ctx, bucket, prefix)
	})
}

func Fuzz_Nosy_S3Client_UploadObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bucket string
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}

		s := NewS3Client()
		s.UploadObject(ctx, bucket, key, d4)
	})
}
