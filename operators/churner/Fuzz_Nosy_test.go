package churner

import (
	"context"
	"testing"
	"time"

	"github.com/Layr-Labs/eigenda/core"
	"github.com/Layr-Labs/eigenda/operators/churner"
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
		var reason FailReason
		fill_err = tp.Fill(&reason)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.IncrementFailedRequestNum(method, reason)
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

func Fuzz_Nosy_Server_Churn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *churner.ChurnRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.Churn(ctx, req)
	})
}

func Fuzz_Nosy_Server_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var metricsConfig MetricsConfig
		fill_err = tp.Fill(&metricsConfig)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start(metricsConfig)
	})
}

func Fuzz_Nosy_Server_checkShouldBeRateLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var request ChurnRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.checkShouldBeRateLimited(now, request)
	})
}

func Fuzz_Nosy_Server_validateChurnRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *churner.ChurnRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.validateChurnRequest(ctx, req)
	})
}

func Fuzz_Nosy_churner_ProcessChurnRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *churner
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorToRegisterAddress common.Address
		fill_err = tp.Fill(&operatorToRegisterAddress)
		if fill_err != nil {
			return
		}
		var churnRequest *ChurnRequest
		fill_err = tp.Fill(&churnRequest)
		if fill_err != nil {
			return
		}
		if c == nil || churnRequest == nil {
			return
		}

		c.ProcessChurnRequest(ctx, operatorToRegisterAddress, churnRequest)
	})
}

func Fuzz_Nosy_churner_UpdateQuorumCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *churner
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

		c.UpdateQuorumCount(ctx)
	})
}

func Fuzz_Nosy_churner_VerifyRequestSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *churner
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var churnRequest *ChurnRequest
		fill_err = tp.Fill(&churnRequest)
		if fill_err != nil {
			return
		}
		if c == nil || churnRequest == nil {
			return
		}

		c.VerifyRequestSignature(ctx, churnRequest)
	})
}

func Fuzz_Nosy_churner_createChurnResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *churner
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorToRegisterAddress common.Address
		fill_err = tp.Fill(&operatorToRegisterAddress)
		if fill_err != nil {
			return
		}
		var operatorToRegisterId core.OperatorID
		fill_err = tp.Fill(&operatorToRegisterId)
		if fill_err != nil {
			return
		}
		var quorumIDs []uint8
		fill_err = tp.Fill(&quorumIDs)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.createChurnResponse(ctx, operatorToRegisterAddress, operatorToRegisterId, quorumIDs)
	})
}

func Fuzz_Nosy_churner_getOperatorsToChurn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *churner
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var quorumIDs []uint8
		fill_err = tp.Fill(&quorumIDs)
		if fill_err != nil {
			return
		}
		var operatorStakes core.OperatorStakes
		fill_err = tp.Fill(&operatorStakes)
		if fill_err != nil {
			return
		}
		var operatorToRegisterAddress common.Address
		fill_err = tp.Fill(&operatorToRegisterAddress)
		if fill_err != nil {
			return
		}
		var currentBlockNumber uint32
		fill_err = tp.Fill(&currentBlockNumber)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.getOperatorsToChurn(ctx, quorumIDs, operatorStakes, operatorToRegisterAddress, currentBlockNumber)
	})
}

func Fuzz_Nosy_churner_sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *churner
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var operatorToRegisterAddress common.Address
		fill_err = tp.Fill(&operatorToRegisterAddress)
		if fill_err != nil {
			return
		}
		var operatorToRegisterId core.OperatorID
		fill_err = tp.Fill(&operatorToRegisterId)
		if fill_err != nil {
			return
		}
		var operatorsToChurn []core.OperatorToChurn
		fill_err = tp.Fill(&operatorsToChurn)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.sign(ctx, operatorToRegisterAddress, operatorToRegisterId, operatorsToChurn)
	})
}

func Fuzz_Nosy_CalculateRequestHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var churnRequest *ChurnRequest
		fill_err = tp.Fill(&churnRequest)
		if fill_err != nil {
			return
		}
		if churnRequest == nil {
			return
		}

		CalculateRequestHash(churnRequest)
	})
}

func Fuzz_Nosy_convertToOperatorsToChurnGrpc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var operatorsToChurn []core.OperatorToChurn
		fill_err = tp.Fill(&operatorsToChurn)
		if fill_err != nil {
			return
		}

		convertToOperatorsToChurnGrpc(operatorsToChurn)
	})
}
