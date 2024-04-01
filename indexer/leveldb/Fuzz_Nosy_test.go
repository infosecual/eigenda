package leveldb

import (
	"testing"

	"github.com/Layr-Labs/eigenda/indexer"
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

func Fuzz_Nosy_HeaderStore_AddHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *HeaderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddHeaders(headers)
	})
}

// skipping Fuzz_Nosy_HeaderStore_AttachObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.AccumulatorObject

func Fuzz_Nosy_HeaderStore_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *HeaderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Close()
	})
}

func Fuzz_Nosy_HeaderStore_FastForward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *HeaderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.FastForward()
	})
}

func Fuzz_Nosy_HeaderStore_GetLatestHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *HeaderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var finalized bool
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetLatestHeader(finalized)
	})
}

// skipping Fuzz_Nosy_HeaderStore_GetLatestObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

// skipping Fuzz_Nosy_HeaderStore_GetObject__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

// skipping Fuzz_Nosy_HeaderStore_getAccumulatorEntry__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

func Fuzz_Nosy_HeaderStore_updateHeaderEntry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *HeaderStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var header *indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var accKey []byte
		fill_err = tp.Fill(&accKey)
		if fill_err != nil {
			return
		}
		var tx *transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if s == nil || header == nil || tx == nil {
			return
		}

		s.updateHeaderEntry(header, accKey, tx)
	})
}

func Fuzz_Nosy_headerEntry_UpdateAccumulatorKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		e := newHeaderEntry(header)
		e.UpdateAccumulatorKeys(key)
	})
}

func Fuzz_Nosy_iter_First__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		i := Iter(key)
		i.First()
	})
}

func Fuzz_Nosy_iter_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		i := Iter(key)
		i.Next()
	})
}

func Fuzz_Nosy_iter_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		i := Iter(key)
		i.Release()
	})
}

// skipping Fuzz_Nosy_iter_Value__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_levelDB_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Close()
	})
}

// skipping Fuzz_Nosy_levelDB_Get__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_levelDB_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Has(key)
	})
}

func Fuzz_Nosy_levelDB_Iter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var prefix []byte
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Iter(prefix)
	})
}

// skipping Fuzz_Nosy_levelDB_Put__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_levelDB_Tx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Tx()
	})
}

func Fuzz_Nosy_transaction_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		t1, err := newTransaction(l)
		if err != nil {
			return
		}
		t1.Commit()
	})
}

func Fuzz_Nosy_transaction_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		t1, err := newTransaction(l)
		if err != nil {
			return
		}
		t1.Delete(key)
	})
}

func Fuzz_Nosy_transaction_Discard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		t1, err := newTransaction(l)
		if err != nil {
			return
		}
		t1.Discard()
	})
}

func Fuzz_Nosy_transaction_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		t1, err := newTransaction(l)
		if err != nil {
			return
		}
		t1.Empty()
	})
}

// skipping Fuzz_Nosy_transaction_Get__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_transaction_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		t1, err := newTransaction(l)
		if err != nil {
			return
		}
		t1.Has(key)
	})
}

func Fuzz_Nosy_transaction_Iter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *levelDB
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var prefix []byte
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		t1, err := newTransaction(l)
		if err != nil {
			return
		}
		t1.Iter(prefix)
	})
}

// skipping Fuzz_Nosy_transaction_Put__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_transaction_SetErr__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_headerEntryReader_GetHeaderEntry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r headerEntryReader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		r.GetHeaderEntry(key)
	})
}

func Fuzz_Nosy_headerEntryReader_GetLatestHeaderEntry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r headerEntryReader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.GetLatestHeaderEntry()
	})
}

func Fuzz_Nosy_headerEntryWriter_PutHeaderEntries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w headerEntryWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		w.PutHeaderEntries(headers)
	})
}

func Fuzz_Nosy_headerEntryWriter_filterNew__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w headerEntryWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		w.filterNew(headers)
	})
}

func Fuzz_Nosy_headerEntryWriter_putFinalizedHeaderEntry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w headerEntryWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var headers indexer.Headers
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		w.putFinalizedHeaderEntry(headers)
	})
}

func Fuzz_Nosy_headerEntryWriter_putHeaderEntry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w headerEntryWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var header *indexer.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		w.putHeaderEntry(header)
	})
}

// skipping Fuzz_Nosy_keyValueReader_get__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_keyValueWriter_put__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_encode__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_newAccumulatorKey__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

// skipping Fuzz_Nosy_newAccumulatorKeyPrefix__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/indexer.Accumulator

func Fuzz_Nosy_newHeaderKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		newHeaderKey(v)
	})
}

func Fuzz_Nosy_newHeaderKeySuffix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		newHeaderKeySuffix(v)
	})
}
