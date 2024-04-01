package mock

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigenda/core"
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

func Fuzz_Nosy_MockSubgraphApi_QueryBatchNonSigningInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var first int64
		fill_err = tp.Fill(&first)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryBatchNonSigningInfo(ctx, first)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryBatchNonSigningOperatorIdsInInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var first int64
		fill_err = tp.Fill(&first)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryBatchNonSigningOperatorIdsInInterval(ctx, first)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var descending bool
		fill_err = tp.Fill(&descending)
		if fill_err != nil {
			return
		}
		var orderByField string
		fill_err = tp.Fill(&orderByField)
		if fill_err != nil {
			return
		}
		var first int
		fill_err = tp.Fill(&first)
		if fill_err != nil {
			return
		}
		var skip int
		fill_err = tp.Fill(&skip)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryBatches(ctx, descending, orderByField, first, skip)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryBatchesByBlockTimestampRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var start uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end uint64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryBatchesByBlockTimestampRange(ctx, start, end)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryDeregisteredOperatorsGreaterThanBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockTimestamp uint64
		fill_err = tp.Fill(&blockTimestamp)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryDeregisteredOperatorsGreaterThanBlockTimestamp(ctx, blockTimestamp)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryOperatorAddedToQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var startBlock uint32
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}
		var endBlock uint32
		fill_err = tp.Fill(&endBlock)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryOperatorAddedToQuorum(ctx, startBlock, endBlock)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryOperatorInfoByOperatorIdAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
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

		m.QueryOperatorInfoByOperatorIdAtBlockNumber(ctx, operatorId, blockNumber)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryOperatorRemovedFromQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var startBlock uint32
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}
		var endBlock uint32
		fill_err = tp.Fill(&endBlock)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryOperatorRemovedFromQuorum(ctx, startBlock, endBlock)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryOperators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var first int
		fill_err = tp.Fill(&first)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryOperators(ctx, first)
	})
}

func Fuzz_Nosy_MockSubgraphApi_QueryRegisteredOperatorsGreaterThanBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockSubgraphApi
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockTimestamp uint64
		fill_err = tp.Fill(&blockTimestamp)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryRegisteredOperatorsGreaterThanBlockTimestamp(ctx, blockTimestamp)
	})
}
