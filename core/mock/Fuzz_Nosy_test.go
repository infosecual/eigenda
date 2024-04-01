package mock

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigenda/api/grpc/churner"
	"github.com/Layr-Labs/eigenda/core"
	"github.com/Layr-Labs/eigenda/indexer"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_ChainDataMock_GetCurrentBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numOperators uint) {
		d, err := MakeChainDataMock(numOperators)
		if err != nil {
			return
		}
		d.GetCurrentBlockNumber()
	})
}

func Fuzz_Nosy_ChainDataMock_GetIndexedOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var numOperators uint
		fill_err = tp.Fill(&numOperators)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}

		d, err := MakeChainDataMock(numOperators)
		if err != nil {
			return
		}
		d.GetIndexedOperatorState(ctx, blockNumber, quorums)
	})
}

func Fuzz_Nosy_ChainDataMock_GetOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var numOperators uint
		fill_err = tp.Fill(&numOperators)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}

		d, err := MakeChainDataMock(numOperators)
		if err != nil {
			return
		}
		d.GetOperatorState(ctx, blockNumber, quorums)
	})
}

func Fuzz_Nosy_ChainDataMock_GetOperatorStateByOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var numOperators uint
		fill_err = tp.Fill(&numOperators)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var operator core.OperatorID
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}

		d, err := MakeChainDataMock(numOperators)
		if err != nil {
			return
		}
		d.GetOperatorStateByOperator(ctx, blockNumber, operator)
	})
}

func Fuzz_Nosy_ChainDataMock_GetTotalOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var numOperators uint
		fill_err = tp.Fill(&numOperators)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}

		d, err := MakeChainDataMock(numOperators)
		if err != nil {
			return
		}
		d.GetTotalOperatorState(ctx, blockNumber)
	})
}

func Fuzz_Nosy_ChainDataMock_GetTotalOperatorStateWithQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var numOperators uint
		fill_err = tp.Fill(&numOperators)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}

		d, err := MakeChainDataMock(numOperators)
		if err != nil {
			return
		}
		d.GetTotalOperatorStateWithQuorums(ctx, blockNumber, quorums)
	})
}

func Fuzz_Nosy_ChainDataMock_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var numOperators uint
		fill_err = tp.Fill(&numOperators)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		d, err := MakeChainDataMock(numOperators)
		if err != nil {
			return
		}
		d.Start(_x2)
	})
}

func Fuzz_Nosy_MockIndexedChainState_GetIndexedOperatorInfoByOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockIndexedChainState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorId core.OperatorID
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetIndexedOperatorInfoByOperatorId(ctx, operatorId, blockNumber)
	})
}

func Fuzz_Nosy_MockIndexedChainState_GetIndexedOperatorState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockIndexedChainState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetIndexedOperatorState(ctx, blockNumber, quorums)
	})
}

func Fuzz_Nosy_MockOperatorSocketsFilterer_FilterFastMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockOperatorSocketsFilterer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.FilterFastMode(headers)
	})
}

func Fuzz_Nosy_MockOperatorSocketsFilterer_FilterHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockOperatorSocketsFilterer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.FilterHeaders(headers)
	})
}

func Fuzz_Nosy_MockOperatorSocketsFilterer_GetSyncPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockOperatorSocketsFilterer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var latestHeader *indexer.Header
		fill_err = tp.Fill(&latestHeader)
		if fill_err != nil {
			return
		}
		if t1 == nil || latestHeader == nil {
			return
		}

		t1.GetSyncPoint(latestHeader)
	})
}

func Fuzz_Nosy_MockOperatorSocketsFilterer_SetSyncPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockOperatorSocketsFilterer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var latestHeader *indexer.Header
		fill_err = tp.Fill(&latestHeader)
		if fill_err != nil {
			return
		}
		if t1 == nil || latestHeader == nil {
			return
		}

		t1.SetSyncPoint(latestHeader)
	})
}

func Fuzz_Nosy_MockOperatorSocketsFilterer_WatchOperatorSocketUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockOperatorSocketsFilterer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorId core.OperatorID
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.WatchOperatorSocketUpdate(ctx, operatorId)
	})
}

func Fuzz_Nosy_MockShardValidator_UpdateOperatorID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var operatorID core.OperatorID
		fill_err = tp.Fill(&operatorID)
		if fill_err != nil {
			return
		}

		v := NewMockShardValidator()
		v.UpdateOperatorID(operatorID)
	})
}

// skipping Fuzz_Nosy_MockShardValidator_ValidateBatch__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/common.WorkerPool

func Fuzz_Nosy_MockTransactor_BatchOperatorIDToAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorIds []core.OperatorID
		fill_err = tp.Fill(&operatorIds)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.BatchOperatorIDToAddress(ctx, operatorIds)
	})
}

func Fuzz_Nosy_MockTransactor_BuildConfirmBatchTxn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batchHeader *core.BatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}
		var quorums map[uint8]*core.QuorumResult
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}
		var signatureAggregation *core.SignatureAggregation
		fill_err = tp.Fill(&signatureAggregation)
		if fill_err != nil {
			return
		}
		if t1 == nil || batchHeader == nil || signatureAggregation == nil {
			return
		}

		t1.BuildConfirmBatchTxn(ctx, batchHeader, quorums, signatureAggregation)
	})
}

func Fuzz_Nosy_MockTransactor_CalculateOperatorChurnApprovalDigestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorAddress common.Address
		fill_err = tp.Fill(&operatorAddress)
		if fill_err != nil {
			return
		}
		var operatorId core.OperatorID
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var operatorsToChurn []core.OperatorToChurn
		fill_err = tp.Fill(&operatorsToChurn)
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
		if t1 == nil || expiry == nil {
			return
		}

		t1.CalculateOperatorChurnApprovalDigestHash(ctx, operatorAddress, operatorId, operatorsToChurn, salt, expiry)
	})
}

func Fuzz_Nosy_MockTransactor_ConfirmBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batchHeader *core.BatchHeader
		fill_err = tp.Fill(&batchHeader)
		if fill_err != nil {
			return
		}
		var quorums map[uint8]*core.QuorumResult
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}
		var signatureAggregation *core.SignatureAggregation
		fill_err = tp.Fill(&signatureAggregation)
		if fill_err != nil {
			return
		}
		if t1 == nil || batchHeader == nil || signatureAggregation == nil {
			return
		}

		t1.ConfirmBatch(ctx, batchHeader, quorums, signatureAggregation)
	})
}

func Fuzz_Nosy_MockTransactor_DeregisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var pubkeyG1 *core.G1Point
		fill_err = tp.Fill(&pubkeyG1)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var quorumIds []uint8
		fill_err = tp.Fill(&quorumIds)
		if fill_err != nil {
			return
		}
		if t1 == nil || pubkeyG1 == nil {
			return
		}

		t1.DeregisterOperator(ctx, pubkeyG1, blockNumber, quorumIds)
	})
}

func Fuzz_Nosy_MockTransactor_GetBlockStaleMeasure__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetBlockStaleMeasure(ctx)
	})
}

func Fuzz_Nosy_MockTransactor_GetCurrentBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetCurrentBlockNumber(ctx)
	})
}

func Fuzz_Nosy_MockTransactor_GetCurrentQuorumBitmapByOperatorId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorId core.OperatorID
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetCurrentQuorumBitmapByOperatorId(ctx, operatorId)
	})
}

func Fuzz_Nosy_MockTransactor_GetNumberOfRegisteredOperatorForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetNumberOfRegisteredOperatorForQuorum(ctx, quorumID)
	})
}

func Fuzz_Nosy_MockTransactor_GetOperatorSetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetOperatorSetParams(ctx, quorumID)
	})
}

func Fuzz_Nosy_MockTransactor_GetOperatorStakes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorId core.OperatorID
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetOperatorStakes(ctx, operatorId, blockNumber)
	})
}

func Fuzz_Nosy_MockTransactor_GetOperatorStakesForQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var quorums []uint8
		fill_err = tp.Fill(&quorums)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetOperatorStakesForQuorums(ctx, quorums, blockNumber)
	})
}

func Fuzz_Nosy_MockTransactor_GetQuorumBitmapForOperatorsAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorIds []core.OperatorID
		fill_err = tp.Fill(&operatorIds)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetQuorumBitmapForOperatorsAtBlockNumber(ctx, operatorIds, blockNumber)
	})
}

func Fuzz_Nosy_MockTransactor_GetQuorumCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetQuorumCount(ctx, blockNumber)
	})
}

func Fuzz_Nosy_MockTransactor_GetQuorumSecurityParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetQuorumSecurityParams(ctx, blockNumber)
	})
}

func Fuzz_Nosy_MockTransactor_GetRegisteredQuorumIdsForOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operator core.OperatorID
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetRegisteredQuorumIdsForOperator(ctx, operator)
	})
}

func Fuzz_Nosy_MockTransactor_GetRequiredQuorumNumbers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber uint32
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetRequiredQuorumNumbers(ctx, blockNumber)
	})
}

func Fuzz_Nosy_MockTransactor_GetStoreDurationBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetStoreDurationBlocks(ctx)
	})
}

func Fuzz_Nosy_MockTransactor_OperatorIDToAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorId core.OperatorID
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.OperatorIDToAddress(ctx, operatorId)
	})
}

func Fuzz_Nosy_MockTransactor_PubkeyHashToOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorId core.OperatorID
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.PubkeyHashToOperator(ctx, operatorId)
	})
}

func Fuzz_Nosy_MockTransactor_RegisterOperator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var keypair *core.KeyPair
		fill_err = tp.Fill(&keypair)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var quorumIds []uint8
		fill_err = tp.Fill(&quorumIds)
		if fill_err != nil {
			return
		}
		var operatorEcdsaPrivateKey *ecdsa.PrivateKey
		fill_err = tp.Fill(&operatorEcdsaPrivateKey)
		if fill_err != nil {
			return
		}
		var operatorToAvsRegistrationSigSalt [32]byte
		fill_err = tp.Fill(&operatorToAvsRegistrationSigSalt)
		if fill_err != nil {
			return
		}
		var operatorToAvsRegistrationSigExpiry *big.Int
		fill_err = tp.Fill(&operatorToAvsRegistrationSigExpiry)
		if fill_err != nil {
			return
		}
		if t1 == nil || keypair == nil || operatorEcdsaPrivateKey == nil || operatorToAvsRegistrationSigExpiry == nil {
			return
		}

		t1.RegisterOperator(ctx, keypair, socket, quorumIds, operatorEcdsaPrivateKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry)
	})
}

func Fuzz_Nosy_MockTransactor_RegisterOperatorWithChurn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var keypair *core.KeyPair
		fill_err = tp.Fill(&keypair)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		var quorumIds []uint8
		fill_err = tp.Fill(&quorumIds)
		if fill_err != nil {
			return
		}
		var operatorEcdsaPrivateKey *ecdsa.PrivateKey
		fill_err = tp.Fill(&operatorEcdsaPrivateKey)
		if fill_err != nil {
			return
		}
		var operatorToAvsRegistrationSigSalt [32]byte
		fill_err = tp.Fill(&operatorToAvsRegistrationSigSalt)
		if fill_err != nil {
			return
		}
		var operatorToAvsRegistrationSigExpiry *big.Int
		fill_err = tp.Fill(&operatorToAvsRegistrationSigExpiry)
		if fill_err != nil {
			return
		}
		var churnReply *churner.ChurnReply
		fill_err = tp.Fill(&churnReply)
		if fill_err != nil {
			return
		}
		if t1 == nil || keypair == nil || operatorEcdsaPrivateKey == nil || operatorToAvsRegistrationSigExpiry == nil || churnReply == nil {
			return
		}

		t1.RegisterOperatorWithChurn(ctx, keypair, socket, quorumIds, operatorEcdsaPrivateKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry, churnReply)
	})
}

func Fuzz_Nosy_MockTransactor_StakeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.StakeRegistry(ctx)
	})
}

func Fuzz_Nosy_MockTransactor_UpdateOperatorSocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var socket string
		fill_err = tp.Fill(&socket)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UpdateOperatorSocket(ctx, socket)
	})
}

func Fuzz_Nosy_MockTransactor_WeightOfOperatorForQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *MockTransactor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var quorumID uint8
		fill_err = tp.Fill(&quorumID)
		if fill_err != nil {
			return
		}
		var operator common.Address
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.WeightOfOperatorForQuorum(ctx, quorumID, operator)
	})
}
