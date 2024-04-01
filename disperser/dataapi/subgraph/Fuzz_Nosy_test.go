package subgraph

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

func Fuzz_Nosy_api_QueryBatchNonSigningInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var intervalSeconds int64
		fill_err = tp.Fill(&intervalSeconds)
		if fill_err != nil {
			return
		}

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryBatchNonSigningInfo(ctx, intervalSeconds)
	})
}

func Fuzz_Nosy_api_QueryBatchNonSigningOperatorIdsInInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var intervalSeconds int64
		fill_err = tp.Fill(&intervalSeconds)
		if fill_err != nil {
			return
		}

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryBatchNonSigningOperatorIdsInInterval(ctx, intervalSeconds)
	})
}

func Fuzz_Nosy_api_QueryBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
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

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryBatches(ctx, descending, orderByField, first, skip)
	})
}

func Fuzz_Nosy_api_QueryBatchesByBlockTimestampRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
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

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryBatchesByBlockTimestampRange(ctx, start, end)
	})
}

func Fuzz_Nosy_api_QueryDeregisteredOperatorsGreaterThanBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
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

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryDeregisteredOperatorsGreaterThanBlockTimestamp(ctx, blockTimestamp)
	})
}

func Fuzz_Nosy_api_QueryOperatorAddedToQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
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

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryOperatorAddedToQuorum(ctx, startBlock, endBlock)
	})
}

func Fuzz_Nosy_api_QueryOperatorInfoByOperatorIdAtBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
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

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryOperatorInfoByOperatorIdAtBlockNumber(ctx, operatorId, blockNumber)
	})
}

func Fuzz_Nosy_api_QueryOperatorRemovedFromQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
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

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryOperatorRemovedFromQuorum(ctx, startBlock, endBlock)
	})
}

func Fuzz_Nosy_api_QueryOperators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
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

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryOperators(ctx, first)
	})
}

func Fuzz_Nosy_api_QueryRegisteredOperatorsGreaterThanBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uiMonitoringSocketAddr string
		fill_err = tp.Fill(&uiMonitoringSocketAddr)
		if fill_err != nil {
			return
		}
		var operatorStateSocketAddr string
		fill_err = tp.Fill(&operatorStateSocketAddr)
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

		a := NewApi(uiMonitoringSocketAddr, operatorStateSocketAddr)
		a.QueryRegisteredOperatorsGreaterThanBlockTimestamp(ctx, blockTimestamp)
	})
}

// skipping Fuzz_Nosy_Api_QueryBatchNonSigningInfo__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryBatchNonSigningOperatorIdsInInterval__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryBatches__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryBatchesByBlockTimestampRange__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryDeregisteredOperatorsGreaterThanBlockTimestamp__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryOperatorAddedToQuorum__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryOperatorInfoByOperatorIdAtBlockNumber__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryOperatorRemovedFromQuorum__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryOperators__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api

// skipping Fuzz_Nosy_Api_QueryRegisteredOperatorsGreaterThanBlockTimestamp__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph.Api
