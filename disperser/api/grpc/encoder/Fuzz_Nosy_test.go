package encoder

import (
	context "context"
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

func Fuzz_Nosy_BlobCommitment_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlobCommitment
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BlobCommitment_GetCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *BlobCommitment
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetCommitment()
	})
}

func Fuzz_Nosy_BlobCommitment_GetLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *BlobCommitment
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetLength()
	})
}

func Fuzz_Nosy_BlobCommitment_GetLengthCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *BlobCommitment
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetLengthCommitment()
	})
}

func Fuzz_Nosy_BlobCommitment_GetLengthProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *BlobCommitment
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetLengthProof()
	})
}

func Fuzz_Nosy_BlobCommitment_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlobCommitment
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BlobCommitment_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *BlobCommitment
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.ProtoReflect()
	})
}

func Fuzz_Nosy_BlobCommitment_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *BlobCommitment
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Reset()
	})
}

func Fuzz_Nosy_BlobCommitment_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *BlobCommitment
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_EncodeBlobReply_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EncodeBlobReply
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EncodeBlobReply_GetChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobReply
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetChunks()
	})
}

func Fuzz_Nosy_EncodeBlobReply_GetCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobReply
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetCommitment()
	})
}

func Fuzz_Nosy_EncodeBlobReply_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EncodeBlobReply
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EncodeBlobReply_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobReply
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.ProtoReflect()
	})
}

func Fuzz_Nosy_EncodeBlobReply_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobReply
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Reset()
	})
}

func Fuzz_Nosy_EncodeBlobReply_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobReply
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_EncodeBlobRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EncodeBlobRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EncodeBlobRequest_GetData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobRequest
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetData()
	})
}

func Fuzz_Nosy_EncodeBlobRequest_GetEncodingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobRequest
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetEncodingParams()
	})
}

func Fuzz_Nosy_EncodeBlobRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EncodeBlobRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EncodeBlobRequest_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobRequest
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.ProtoReflect()
	})
}

func Fuzz_Nosy_EncodeBlobRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobRequest
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Reset()
	})
}

func Fuzz_Nosy_EncodeBlobRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodeBlobRequest
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_EncodingParams_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EncodingParams
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EncodingParams_GetChunkLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodingParams
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetChunkLength()
	})
}

func Fuzz_Nosy_EncodingParams_GetNumChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodingParams
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.GetNumChunks()
	})
}

func Fuzz_Nosy_EncodingParams_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EncodingParams
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EncodingParams_ProtoReflect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodingParams
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.ProtoReflect()
	})
}

func Fuzz_Nosy_EncodingParams_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodingParams
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Reset()
	})
}

func Fuzz_Nosy_EncodingParams_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EncodingParams
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.String()
	})
}

// skipping Fuzz_Nosy_encoderClient_EncodeBlob__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_EncoderClient_EncodeBlob__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/api/grpc/encoder.EncoderClient

// skipping Fuzz_Nosy_EncoderServer_EncodeBlob__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/api/grpc/encoder.EncoderServer

// skipping Fuzz_Nosy_EncoderServer_mustEmbedUnimplementedEncoderServer__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/api/grpc/encoder.EncoderServer

func Fuzz_Nosy_UnimplementedEncoderServer_EncodeBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 UnimplementedEncoderServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *EncodeBlobRequest
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x3 == nil {
			return
		}

		_x1.EncodeBlob(_x2, _x3)
	})
}

func Fuzz_Nosy_UnimplementedEncoderServer_mustEmbedUnimplementedEncoderServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 UnimplementedEncoderServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.mustEmbedUnimplementedEncoderServer()
	})
}

// skipping Fuzz_Nosy_UnsafeEncoderServer_mustEmbedUnimplementedEncoderServer__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/disperser/api/grpc/encoder.UnsafeEncoderServer

// skipping Fuzz_Nosy_RegisterEncoderServer__ because parameters include func, chan, or unsupported interface: google.golang.org/grpc.ServiceRegistrar

// skipping Fuzz_Nosy__Encoder_EncodeBlob_Handler__ because parameters include func, chan, or unsupported interface: interface{}
