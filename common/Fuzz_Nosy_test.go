package common

import (
	"context"
	"testing"

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

// skipping Fuzz_Nosy_EthClient_BalanceAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_BlockByHash__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_BlockNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_CallContract__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_CallContractAtHash__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_ChainID__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_CodeAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_EnsureAnyTransactionEvaled__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_EnsureTransactionEvaled__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_EstimateGas__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_EstimateGasPriceAndLimitAndSendTx__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_FeeHistory__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_FilterLogs__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_GetAccountAddress__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_GetLatestGasCaps__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_GetNoSendTransactOpts__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_HeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_NetworkID__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_NonceAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_PeerCount__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_PendingBalanceAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_PendingCallContract__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_PendingCodeAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_PendingNonceAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_PendingStorageAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_PendingTransactionCount__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_SendTransaction__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_StorageAt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_SuggestGasPrice__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_SuggestGasTipCap__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_SyncProgress__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_TransactionByHash__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_TransactionCount__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_TransactionInBlock__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_TransactionReceipt__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_TransactionSender__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_EthClient_UpdateGas__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.EthClient

// skipping Fuzz_Nosy_KVStore[T any]_GetItem__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.KVStore[T any]

// skipping Fuzz_Nosy_KVStore[T any]_UpdateItem__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.KVStore[T any]

// skipping Fuzz_Nosy_RPCEthClient_BatchCall__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.RPCEthClient

// skipping Fuzz_Nosy_RPCEthClient_BatchCallContext__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.RPCEthClient

// skipping Fuzz_Nosy_RPCEthClient_Call__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.RPCEthClient

// skipping Fuzz_Nosy_RPCEthClient_CallContext__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.RPCEthClient

// skipping Fuzz_Nosy_RateLimiter_AllowRequest__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.RateLimiter

// skipping Fuzz_Nosy_WorkerPool_Pause__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_WorkerPool_Size__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_WorkerPool_Stop__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_WorkerPool_StopWait__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_WorkerPool_Stopped__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_WorkerPool_Submit__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_WorkerPool_SubmitWait__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_WorkerPool_WaitingQueueSize__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

// skipping Fuzz_Nosy_EncodeToBytes__ because parameters include func, chan, or unsupported interface: T

func Fuzz_Nosy_FireblocksCLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string, flagPrefix string) {
		FireblocksCLIFlags(envPrefix, flagPrefix)
	})
}

func Fuzz_Nosy_GetClientAddress__(f *testing.F) {
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
		var header string
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var numProxies int
		fill_err = tp.Fill(&numProxies)
		if fill_err != nil {
			return
		}
		var allowDirectConnectionFallback bool
		fill_err = tp.Fill(&allowDirectConnectionFallback)
		if fill_err != nil {
			return
		}

		GetClientAddress(ctx, header, numProxies, allowDirectConnectionFallback)
	})
}

// skipping Fuzz_Nosy_Hash__ because parameters include func, chan, or unsupported interface: T

func Fuzz_Nosy_LoggerCLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string, flagPrefix string) {
		LoggerCLIFlags(envPrefix, flagPrefix)
	})
}

func Fuzz_Nosy_PrefixEnvVar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, prefix string, suffix string) {
		PrefixEnvVar(prefix, suffix)
	})
}

func Fuzz_Nosy_PrefixFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, prefix string, suffix string) {
		PrefixFlag(prefix, suffix)
	})
}

func Fuzz_Nosy_splitHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header []string
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}

		splitHeader(header)
	})
}
