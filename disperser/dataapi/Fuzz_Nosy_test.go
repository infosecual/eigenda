package dataapi

import (
	"context"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenda/core"
	"github.com/Layr-Labs/eigenda/disperser"
	"github.com/Layr-Labs/eigenda/disperser/dataapi/subgraph"
	"github.com/gin-gonic/gin"
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

// skipping Fuzz_Nosy_DynamoDBCollector_Collect__ because parameters include func, chan, or unsupported interface: chan<- github.com/prometheus/client_golang/prometheus.Metric

// skipping Fuzz_Nosy_DynamoDBCollector_Describe__ because parameters include func, chan, or unsupported interface: chan<- *github.com/prometheus/client_golang/prometheus.Desc

func Fuzz_Nosy_EigenDAServiceAvailabilityCheck_CheckHealth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sac *EigenDAServiceAvailabilityCheck
		fill_err = tp.Fill(&sac)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var serviceName string
		fill_err = tp.Fill(&serviceName)
		if fill_err != nil {
			return
		}
		if sac == nil {
			return
		}

		sac.CheckHealth(ctx, serviceName)
	})
}

func Fuzz_Nosy_EigenDAServiceAvailabilityCheck_CloseConnections__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sac *EigenDAServiceAvailabilityCheck
		fill_err = tp.Fill(&sac)
		if fill_err != nil {
			return
		}
		if sac == nil {
			return
		}

		sac.CloseConnections()
	})
}

// skipping Fuzz_Nosy_GRPCDialerSkipTLS_Dial__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.DialOption

func Fuzz_Nosy_Metrics_IncrementFailedRequestNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.IncrementFailedRequestNum(method)
	})
}

func Fuzz_Nosy_Metrics_IncrementSuccessfulRequestNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.IncrementSuccessfulRequestNum(method)
	})
}

func Fuzz_Nosy_Metrics_ObserveLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		var latencyMs float64
		fill_err = tp.Fill(&latencyMs)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.ObserveLatency(method, latencyMs)
	})
}

func Fuzz_Nosy_Metrics_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *Metrics
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Start(ctx)
	})
}

func Fuzz_Nosy_prometheusClient_QueryDisperserAvgThroughputBlobSizeBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *prometheusClient
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var start time.Time
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end time.Time
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		var windowSizeInSec uint8
		fill_err = tp.Fill(&windowSizeInSec)
		if fill_err != nil {
			return
		}
		if pc == nil {
			return
		}

		pc.QueryDisperserAvgThroughputBlobSizeBytes(ctx, start, end, windowSizeInSec)
	})
}

func Fuzz_Nosy_prometheusClient_QueryDisperserBlobSizeBytesPerSecond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *prometheusClient
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var start time.Time
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end time.Time
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if pc == nil {
			return
		}

		pc.QueryDisperserBlobSizeBytesPerSecond(ctx, start, end)
	})
}

func Fuzz_Nosy_prometheusClient_queryRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *prometheusClient
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var query string
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var start time.Time
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end time.Time
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if pc == nil {
			return
		}

		pc.queryRange(ctx, query, start, end)
	})
}

func Fuzz_Nosy_server_FetchBlobHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *gin.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.FetchBlobHandler(c)
	})
}

func Fuzz_Nosy_server_FetchBlobsHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *gin.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.FetchBlobsHandler(c)
	})
}

func Fuzz_Nosy_server_FetchDeregisteredOperators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *gin.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.FetchDeregisteredOperators(c)
	})
}

func Fuzz_Nosy_server_FetchMetricsHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *gin.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.FetchMetricsHandler(c)
	})
}

func Fuzz_Nosy_server_FetchMetricsTroughputHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *gin.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.FetchMetricsTroughputHandler(c)
	})
}

func Fuzz_Nosy_server_FetchNonSigners__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *gin.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.FetchNonSigners(c)
	})
}

func Fuzz_Nosy_server_FetchOperatorsNonsigningPercentageHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *gin.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.FetchOperatorsNonsigningPercentageHandler(c)
	})
}

func Fuzz_Nosy_server_GetEigenDAServiceAvailability__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *gin.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.GetEigenDAServiceAvailability(c)
	})
}

func Fuzz_Nosy_server_Shutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Shutdown()
	})
}

func Fuzz_Nosy_server_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start()
	})
}

func Fuzz_Nosy_server_calculateTotalCostGasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.calculateTotalCostGasUsed(ctx)
	})
}

func Fuzz_Nosy_server_convertBlobMetadatasToBlobMetadataResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var metadatas []*disperser.BlobMetadata
		fill_err = tp.Fill(&metadatas)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.convertBlobMetadatasToBlobMetadataResponse(ctx, metadatas)
	})
}

func Fuzz_Nosy_server_createOperatorQuorumIntervals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var nonsigners []core.OperatorID
		fill_err = tp.Fill(&nonsigners)
		if fill_err != nil {
			return
		}
		var nonsignerAddressToId map[string]core.OperatorID
		fill_err = tp.Fill(&nonsignerAddressToId)
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
		if s == nil {
			return
		}

		s.createOperatorQuorumIntervals(ctx, nonsigners, nonsignerAddressToId, startBlock, endBlock)
	})
}

func Fuzz_Nosy_server_getBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getBlob(ctx, key)
	})
}

func Fuzz_Nosy_server_getBlobMetadataByBatchesWithLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getBlobMetadataByBatchesWithLimit(ctx, limit)
	})
}

func Fuzz_Nosy_server_getBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getBlobs(ctx, limit)
	})
}

func Fuzz_Nosy_server_getDeregisteredOperatorForDays__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var days int32
		fill_err = tp.Fill(&days)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getDeregisteredOperatorForDays(ctx, days)
	})
}

func Fuzz_Nosy_server_getMetric__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var startTime int64
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var endTime int64
		fill_err = tp.Fill(&endTime)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getMetric(ctx, startTime, endTime, limit)
	})
}

func Fuzz_Nosy_server_getNonSigners__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.getNonSigners(ctx, intervalSeconds)
	})
}

func Fuzz_Nosy_server_getOperatorNonsigningRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.getOperatorNonsigningRate(ctx, intervalSeconds)
	})
}

func Fuzz_Nosy_server_getOperatorQuorumEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
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
		var nonsignerAddressToId map[string]core.OperatorID
		fill_err = tp.Fill(&nonsignerAddressToId)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getOperatorQuorumEvents(ctx, startBlock, endBlock, nonsignerAddressToId)
	})
}

func Fuzz_Nosy_server_getServiceAvailability__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var services []string
		fill_err = tp.Fill(&services)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getServiceAvailability(ctx, services)
	})
}

func Fuzz_Nosy_server_getThroughput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var start int64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end int64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getThroughput(ctx, start, end)
	})
}

func Fuzz_Nosy_subgraphClient_QueryBatchNonSigningInfoInInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *subgraphClient
		fill_err = tp.Fill(&sc)
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
		if sc == nil {
			return
		}

		sc.QueryBatchNonSigningInfoInInterval(ctx, intervalSeconds)
	})
}

func Fuzz_Nosy_subgraphClient_QueryBatchNonSigningOperatorIdsInInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *subgraphClient
		fill_err = tp.Fill(&sc)
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
		if sc == nil {
			return
		}

		sc.QueryBatchNonSigningOperatorIdsInInterval(ctx, intervalSeconds)
	})
}

func Fuzz_Nosy_subgraphClient_QueryBatchesWithLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *subgraphClient
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var skip int
		fill_err = tp.Fill(&skip)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.QueryBatchesWithLimit(ctx, limit, skip)
	})
}

func Fuzz_Nosy_subgraphClient_QueryIndexedDeregisteredOperatorsForTimeWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *subgraphClient
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var days int32
		fill_err = tp.Fill(&days)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.QueryIndexedDeregisteredOperatorsForTimeWindow(ctx, days)
	})
}

func Fuzz_Nosy_subgraphClient_QueryOperatorQuorumEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *subgraphClient
		fill_err = tp.Fill(&sc)
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
		if sc == nil {
			return
		}

		sc.QueryOperatorQuorumEvent(ctx, startBlock, endBlock)
	})
}

func Fuzz_Nosy_subgraphClient_QueryOperatorsWithLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *subgraphClient
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.QueryOperatorsWithLimit(ctx, limit)
	})
}

// skipping Fuzz_Nosy_EigenDAServiceChecker_CheckHealth__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.EigenDAServiceChecker

// skipping Fuzz_Nosy_EigenDAServiceChecker_CloseConnections__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.EigenDAServiceChecker

// skipping Fuzz_Nosy_GRPCConn_Dial__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.GRPCConn

func Fuzz_Nosy_OperatorQuorumIntervals_GetQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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
		var operatorInitialQuorum map[string][]uint8
		fill_err = tp.Fill(&operatorInitialQuorum)
		if fill_err != nil {
			return
		}
		var addedToQuorum map[string][]*OperatorQuorum
		fill_err = tp.Fill(&addedToQuorum)
		if fill_err != nil {
			return
		}
		var removedFromQuorum map[string][]*OperatorQuorum
		fill_err = tp.Fill(&removedFromQuorum)
		if fill_err != nil {
			return
		}
		var operatorId string
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var blockNum uint32
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}

		oqi, err := CreateOperatorQuorumIntervals(startBlock, endBlock, operatorInitialQuorum, addedToQuorum, removedFromQuorum)
		if err != nil {
			return
		}
		oqi.GetQuorums(operatorId, blockNum)
	})
}

// skipping Fuzz_Nosy_PrometheusClient_QueryDisperserAvgThroughputBlobSizeBytes__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.PrometheusClient

// skipping Fuzz_Nosy_PrometheusClient_QueryDisperserBlobSizeBytesPerSecond__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.PrometheusClient

// skipping Fuzz_Nosy_SubgraphClient_QueryBatchNonSigningInfoInInterval__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.SubgraphClient

// skipping Fuzz_Nosy_SubgraphClient_QueryBatchNonSigningOperatorIdsInInterval__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.SubgraphClient

// skipping Fuzz_Nosy_SubgraphClient_QueryBatchesWithLimit__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.SubgraphClient

// skipping Fuzz_Nosy_SubgraphClient_QueryIndexedDeregisteredOperatorsForTimeWindow__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.SubgraphClient

// skipping Fuzz_Nosy_SubgraphClient_QueryOperatorQuorumEvent__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.SubgraphClient

// skipping Fuzz_Nosy_SubgraphClient_QueryOperatorsWithLimit__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/dataapi.SubgraphClient

func Fuzz_Nosy_ComputeNumBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var quorumBatches *QuorumBatches
		fill_err = tp.Fill(&quorumBatches)
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
		if quorumBatches == nil {
			return
		}

		ComputeNumBatches(quorumBatches, startBlock, endBlock)
	})
}

func Fuzz_Nosy_ConvertHexadecimalToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, byteHash []byte) {
		ConvertHexadecimalToBytes(byteHash)
	})
}

func Fuzz_Nosy_ConvertNanosecondToSecond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, timestamp uint64) {
		ConvertNanosecondToSecond(timestamp)
	})
}

func Fuzz_Nosy_CreatQuorumBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batches []*BatchNonSigningInfo
		fill_err = tp.Fill(&batches)
		if fill_err != nil {
			return
		}

		CreatQuorumBatches(batches)
	})
}

func Fuzz_Nosy_addOperatorWithErrorDetail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var operators map[core.OperatorID]*DeregisteredOperatorInfo
		fill_err = tp.Fill(&operators)
		if fill_err != nil {
			return
		}
		var operator *Operator
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}
		var operatorId [32]byte
		fill_err = tp.Fill(&operatorId)
		if fill_err != nil {
			return
		}
		var errorMessage string
		fill_err = tp.Fill(&errorMessage)
		if fill_err != nil {
			return
		}
		if operator == nil {
			return
		}

		addOperatorWithErrorDetail(operators, operator, operatorId, errorMessage)
	})
}

// skipping Fuzz_Nosy_checkIsOnlineAndProcessOperator__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/disperser/dataapi.DeregisteredOperatorMetadata

func Fuzz_Nosy_checkIsOperatorOnline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, socket string) {
		checkIsOperatorOnline(socket)
	})
}

func Fuzz_Nosy_computeNumFailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batches []*BatchNonSigningInfo
		fill_err = tp.Fill(&batches)
		if fill_err != nil {
			return
		}
		var operatorQuorumIntervals OperatorQuorumIntervals
		fill_err = tp.Fill(&operatorQuorumIntervals)
		if fill_err != nil {
			return
		}

		computeNumFailed(batches, operatorQuorumIntervals)
	})
}

func Fuzz_Nosy_computeNumResponsible__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batches []*BatchNonSigningInfo
		fill_err = tp.Fill(&batches)
		if fill_err != nil {
			return
		}
		var operatorQuorumIntervals OperatorQuorumIntervals
		fill_err = tp.Fill(&operatorQuorumIntervals)
		if fill_err != nil {
			return
		}

		computeNumResponsible(batches, operatorQuorumIntervals)
	})
}

func Fuzz_Nosy_convertBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var subgraphBatches []*subgraph.Batches
		fill_err = tp.Fill(&subgraphBatches)
		if fill_err != nil {
			return
		}

		convertBatches(subgraphBatches)
	})
}

// skipping Fuzz_Nosy_errorResponse__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_getLowerBoundIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var intervals []*NumBatchesAtBlock
		fill_err = tp.Fill(&intervals)
		if fill_err != nil {
			return
		}
		var blockNum uint32
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}

		getLowerBoundIndex(intervals, blockNum)
	})
}

func Fuzz_Nosy_getNonSigners__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batches []*BatchNonSigningInfo
		fill_err = tp.Fill(&batches)
		if fill_err != nil {
			return
		}

		getNonSigners(batches)
	})
}

func Fuzz_Nosy_getUpperBoundIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var intervals []*NumBatchesAtBlock
		fill_err = tp.Fill(&intervals)
		if fill_err != nil {
			return
		}
		var blockNum uint32
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}

		getUpperBoundIndex(intervals, blockNum)
	})
}

func Fuzz_Nosy_parseOperatorQuorum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var operatorQuorum []*subgraph.OperatorQuorum
		fill_err = tp.Fill(&operatorQuorum)
		if fill_err != nil {
			return
		}

		parseOperatorQuorum(operatorQuorum)
	})
}

// skipping Fuzz_Nosy_processOperatorOnlineCheck__ because parameters include func, chan, or unsupported interface: chan<- *github.com/Layr-Labs/eigenda/disperser/dataapi.DeregisteredOperatorMetadata

// skipping Fuzz_Nosy_run__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigensdk-go/logging.Logger
