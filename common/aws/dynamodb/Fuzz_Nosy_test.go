package dynamodb

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

// skipping Fuzz_Nosy_Client_DeleteItem__ because parameters include func, chan, or unsupported interface: map[string]github.com/aws/aws-sdk-go-v2/service/dynamodb/types.AttributeValue

func Fuzz_Nosy_Client_DeleteItems__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tableName string
		fill_err = tp.Fill(&tableName)
		if fill_err != nil {
			return
		}
		var keys []map[string]types.AttributeValue
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.DeleteItems(ctx, tableName, keys)
	})
}

func Fuzz_Nosy_Client_DeleteTable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tableName string
		fill_err = tp.Fill(&tableName)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.DeleteTable(ctx, tableName)
	})
}

// skipping Fuzz_Nosy_Client_GetItem__ because parameters include func, chan, or unsupported interface: map[string]github.com/aws/aws-sdk-go-v2/service/dynamodb/types.AttributeValue

// skipping Fuzz_Nosy_Client_PutItem__ because parameters include func, chan, or unsupported interface: map[string]github.com/aws/aws-sdk-go-v2/service/dynamodb/types.AttributeValue

func Fuzz_Nosy_Client_PutItems__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tableName string
		fill_err = tp.Fill(&tableName)
		if fill_err != nil {
			return
		}
		var items []map[string]types.AttributeValue
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.PutItems(ctx, tableName, items)
	})
}

// skipping Fuzz_Nosy_Client_QueryIndex__ because parameters include func, chan, or unsupported interface: map[string]github.com/aws/aws-sdk-go-v2/service/dynamodb/types.AttributeValue

// skipping Fuzz_Nosy_Client_QueryIndexCount__ because parameters include func, chan, or unsupported interface: map[string]github.com/aws/aws-sdk-go-v2/service/dynamodb/types.AttributeValue

// skipping Fuzz_Nosy_Client_QueryIndexWithPagination__ because parameters include func, chan, or unsupported interface: map[string]github.com/aws/aws-sdk-go-v2/service/dynamodb/types.AttributeValue

// skipping Fuzz_Nosy_Client_UpdateItem__ because parameters include func, chan, or unsupported interface: map[string]github.com/aws/aws-sdk-go-v2/service/dynamodb/types.AttributeValue

func Fuzz_Nosy_Client_writeItems__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tableName string
		fill_err = tp.Fill(&tableName)
		if fill_err != nil {
			return
		}
		var requestItems []map[string]types.AttributeValue
		fill_err = tp.Fill(&requestItems)
		if fill_err != nil {
			return
		}
		var operation batchOperation
		fill_err = tp.Fill(&operation)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.writeItems(ctx, tableName, requestItems, operation)
	})
}
