package geth

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

func Fuzz_Nosy_EthClient_EnsureAnyTransactionEvaled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthClient
		fill_err = tp.Fill(&c)
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
		if c == nil {
			return
		}

		c.EnsureAnyTransactionEvaled(ctx, txs, tag)
	})
}

func Fuzz_Nosy_EthClient_EnsureTransactionEvaled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthClient
		fill_err = tp.Fill(&c)
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
		if c == nil || tx == nil {
			return
		}

		c.EnsureTransactionEvaled(ctx, tx, tag)
	})
}

func Fuzz_Nosy_EthClient_EstimateGasPriceAndLimitAndSendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthClient
		fill_err = tp.Fill(&c)
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
		if c == nil || tx == nil || value == nil {
			return
		}

		c.EstimateGasPriceAndLimitAndSendTx(ctx, tx, tag, value)
	})
}

func Fuzz_Nosy_EthClient_GetAccountAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetAccountAddress()
	})
}

func Fuzz_Nosy_EthClient_GetLatestGasCaps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetLatestGasCaps(ctx)
	})
}

func Fuzz_Nosy_EthClient_GetNoSendTransactOpts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetNoSendTransactOpts()
	})
}

func Fuzz_Nosy_EthClient_UpdateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthClient
		fill_err = tp.Fill(&c)
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
		if c == nil || tx == nil || value == nil || gasTipCap == nil || gasFeeCap == nil {
			return
		}

		c.UpdateGas(ctx, tx, value, gasTipCap, gasFeeCap)
	})
}

func Fuzz_Nosy_EthClient_waitMined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthClient
		fill_err = tp.Fill(&c)
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
		if c == nil {
			return
		}

		c.waitMined(ctx, txs)
	})
}

func Fuzz_Nosy_FailoverController_GetTotalNumberRpcFault__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FailoverController
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetTotalNumberRpcFault()
	})
}

// skipping Fuzz_Nosy_FailoverController_ProcessError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_FailoverController_handleError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_FailoverController_handleHttpError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FailoverController
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var httpRespError rpc.HTTPError
		fill_err = tp.Fill(&httpRespError)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.handleHttpError(httpRespError)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil || blockNumber == nil {
			return
		}

		iec.BalanceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil || number == nil {
			return
		}

		iec.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.BlockNumber(ctx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if iec == nil || blockNumber == nil {
			return
		}

		iec.CallContract(ctx, call, blockNumber)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_CallContractAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.CallContractAtHash(ctx, msg, blockHash)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.ChainID(ctx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var contract common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if iec == nil || blockNumber == nil {
			return
		}

		iec.CodeAt(ctx, contract, blockNumber)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.EstimateGas(ctx, call)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_EstimateGasPriceAndLimitAndSendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *InstrumentedEthClient
		fill_err = tp.Fill(&c)
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
		if c == nil || tx == nil || value == nil {
			return
		}

		c.EstimateGasPriceAndLimitAndSendTx(ctx, tx, tag, value)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_FeeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil || lastBlock == nil {
			return
		}

		iec.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var query ethereum.FilterQuery
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.FilterLogs(ctx, query)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil || number == nil {
			return
		}

		iec.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_NetworkID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.NetworkID(ctx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil || blockNumber == nil {
			return
		}

		iec.NonceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.PeerCount(ctx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_PendingBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.PendingBalanceAt(ctx, account)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_PendingCallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.PendingCallContract(ctx, call)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.PendingCodeAt(ctx, account)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.PendingNonceAt(ctx, account)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_PendingStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.PendingStorageAt(ctx, account, key)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_PendingTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.PendingTransactionCount(ctx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil || tx == nil {
			return
		}

		iec.SendTransaction(ctx, tx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil || blockNumber == nil {
			return
		}

		iec.StorageAt(ctx, account, key, blockNumber)
	})
}

// skipping Fuzz_Nosy_InstrumentedEthClient_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_InstrumentedEthClient_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_InstrumentedEthClient_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.SuggestGasPrice(ctx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if iec == nil {
			return
		}

		iec.SyncProgress(ctx)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.TransactionByHash(ctx, hash)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.TransactionCount(ctx, blockHash)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.TransactionInBlock(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil {
			return
		}

		iec.TransactionReceipt(ctx, txHash)
	})
}

func Fuzz_Nosy_InstrumentedEthClient_TransactionSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iec *InstrumentedEthClient
		fill_err = tp.Fill(&iec)
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
		if iec == nil || tx == nil {
			return
		}

		iec.TransactionSender(ctx, tx, block, index)
	})
}

func Fuzz_Nosy_MultiHomingClient_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || blockNumber == nil {
			return
		}

		m.BalanceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_MultiHomingClient_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MultiHomingClient_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || number == nil {
			return
		}

		m.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MultiHomingClient_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.BlockNumber(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m == nil || blockNumber == nil {
			return
		}

		m.CallContract(ctx, call, blockNumber)
	})
}

func Fuzz_Nosy_MultiHomingClient_CallContractAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.CallContractAtHash(ctx, msg, blockHash)
	})
}

func Fuzz_Nosy_MultiHomingClient_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ChainID(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var contract common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m == nil || blockNumber == nil {
			return
		}

		m.CodeAt(ctx, contract, blockNumber)
	})
}

func Fuzz_Nosy_MultiHomingClient_EnsureAnyTransactionEvaled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.EnsureAnyTransactionEvaled(ctx, txs, tag)
	})
}

func Fuzz_Nosy_MultiHomingClient_EnsureTransactionEvaled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || tx == nil {
			return
		}

		m.EnsureTransactionEvaled(ctx, tx, tag)
	})
}

func Fuzz_Nosy_MultiHomingClient_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.EstimateGas(ctx, msg)
	})
}

func Fuzz_Nosy_MultiHomingClient_EstimateGasPriceAndLimitAndSendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || tx == nil || value == nil {
			return
		}

		m.EstimateGasPriceAndLimitAndSendTx(ctx, tx, tag, value)
	})
}

func Fuzz_Nosy_MultiHomingClient_FeeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || lastBlock == nil {
			return
		}

		m.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	})
}

func Fuzz_Nosy_MultiHomingClient_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.FilterLogs(ctx, q)
	})
}

func Fuzz_Nosy_MultiHomingClient_GetAccountAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAccountAddress()
	})
}

func Fuzz_Nosy_MultiHomingClient_GetLatestGasCaps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetLatestGasCaps(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_GetNoSendTransactOpts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetNoSendTransactOpts()
	})
}

func Fuzz_Nosy_MultiHomingClient_GetRPCInstance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRPCInstance()
	})
}

func Fuzz_Nosy_MultiHomingClient_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MultiHomingClient_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || number == nil {
			return
		}

		m.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MultiHomingClient_NetworkID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.NetworkID(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || blockNumber == nil {
			return
		}

		m.NonceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_MultiHomingClient_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PeerCount(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_PendingBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.PendingBalanceAt(ctx, account)
	})
}

func Fuzz_Nosy_MultiHomingClient_PendingCallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.PendingCallContract(ctx, msg)
	})
}

func Fuzz_Nosy_MultiHomingClient_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.PendingCodeAt(ctx, account)
	})
}

func Fuzz_Nosy_MultiHomingClient_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.PendingNonceAt(ctx, account)
	})
}

func Fuzz_Nosy_MultiHomingClient_PendingStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.PendingStorageAt(ctx, account, key)
	})
}

func Fuzz_Nosy_MultiHomingClient_PendingTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PendingTransactionCount(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || tx == nil {
			return
		}

		m.SendTransaction(ctx, tx)
	})
}

func Fuzz_Nosy_MultiHomingClient_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || blockNumber == nil {
			return
		}

		m.StorageAt(ctx, account, key, blockNumber)
	})
}

// skipping Fuzz_Nosy_MultiHomingClient_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_MultiHomingClient_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_MultiHomingClient_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SuggestGasPrice(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SyncProgress(ctx)
	})
}

func Fuzz_Nosy_MultiHomingClient_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.TransactionByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MultiHomingClient_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.TransactionCount(ctx, blockHash)
	})
}

func Fuzz_Nosy_MultiHomingClient_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.TransactionInBlock(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_MultiHomingClient_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.TransactionReceipt(ctx, txHash)
	})
}

func Fuzz_Nosy_MultiHomingClient_TransactionSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || tx == nil {
			return
		}

		m.TransactionSender(ctx, tx, block, index)
	})
}

func Fuzz_Nosy_MultiHomingClient_UpdateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiHomingClient
		fill_err = tp.Fill(&m)
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
		if m == nil || tx == nil || value == nil || gasTipCap == nil || gasFeeCap == nil {
			return
		}

		m.UpdateGas(ctx, tx, value, gasTipCap, gasFeeCap)
	})
}

func Fuzz_Nosy_EthClientFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
		EthClientFlags(envPrefix)
	})
}

func Fuzz_Nosy_addGasBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, gasLimit uint64) {
		addGasBuffer(gasLimit)
	})
}

func Fuzz_Nosy_getClientAndVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var client *EthClient
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		if client == nil {
			return
		}

		getClientAndVersion(client)
	})
}

// skipping Fuzz_Nosy_instrumentFunction__ because parameters include func, chan, or unsupported interface: func() (T, error)
